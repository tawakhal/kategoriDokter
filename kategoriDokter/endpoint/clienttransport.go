package endpoint

import (
	"context"
	"time"

	svc "rumahsakit/dokter/dokter/server"

	pb "rumahsakit/dokter/dokter/grpc"

	util "rumahsakit/dokter/util/grpc"
	disc "rumahsakit/dokter/util/microservice"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/lb"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	grpcName = "grpc.DokterService"
)

func NewGRPCDokterClient(nodes []string, creds credentials.TransportCredentials, option util.ClientOption,
	tracer stdopentracing.Tracer, logger log.Logger) (svc.DokterService, error) {

	instancer, err := disc.ServiceDiscovery(nodes, svc.ServiceID, logger)
	if err != nil {
		return nil, err
	}

	retryMax := option.Retry
	retryTimeout := option.RetryTimeout
	timeout := option.Timeout

	var addDokterEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientAddDokterEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		addDokterEp = retry
	}

	var readDokterByKodeEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDokterByKodeEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDokterByKodeEp = retry
	}

	var readDokterEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientReadDokterEndpoint, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		readDokterEp = retry
	}

	var updateDokterEp endpoint.Endpoint
	{
		factory := util.EndpointFactory(makeClientUpdateDokter, creds, timeout, tracer, logger)
		endpointer := sd.NewEndpointer(instancer, factory, logger)
		balancer := lb.NewRoundRobin(endpointer)
		retry := lb.Retry(retryMax, retryTimeout, balancer)
		updateDokterEp = retry
	}

	return DokterEndpoint{AddDokterEndpoint: addDokterEp, ReadDokterByKodeEndpoint: readDokterByKodeEp,
		ReadDokterEndpoint: readDokterEp, UpdateDokterEndpoint: updateDokterEp}, nil
}

func encodeAddDokterRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Dokter)
	return &pb.AddDokterReq{
		KodeDokter:         req.KodeDokter,
		NamaDokter:         req.NamaDokter,
		JenisKelamin:       req.JenisKelamin,
		Alamat:             req.Alamat,
		NomorTelepon:       req.NomorTelepon,
		KodeKategoriDokter: req.KodeKategoriDokter,
		Biaya:              req.Biaya,
	}, nil
}

func encodeReadDokterByKodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Dokter)
	return &pb.ReadDokterByKodeReq{Kode: req.KodeDokter}, nil
}

func encodeReadDokterRequest(_ context.Context, request interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeUpdateDokterRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(svc.Dokter)
	return &pb.UpdateDokterReq{
		KodeDokter:         req.KodeDokter,
		NamaDokter:         req.NamaDokter,
		JenisKelamin:       req.JenisKelamin,
		Alamat:             req.Alamat,
		NomorTelepon:       req.NomorTelepon,
		KodeKategoriDokter: req.KodeKategoriDokter,
		Biaya:              req.Biaya,
	}, nil
}

func decodeDokterResponse(_ context.Context, response interface{}) (interface{}, error) {
	return nil, nil
}

func decodeReadDokterByKodeRespones(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDokterByKodeResp)
	return svc.Dokter{
		KodeDokter:         resp.KodeDokter,
		NamaDokter:         resp.NamaDokter,
		JenisKelamin:       resp.JenisKelamin,
		Alamat:             resp.Alamat,
		NomorTelepon:       resp.NomorTelepon,
		KodeKategoriDokter: resp.KodeKategoriDokter,
		Biaya:              resp.Biaya,
	}, nil
}

func decodeReadDokterResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.ReadDokterResp)
	var rsp svc.Dokters

	for _, v := range resp.Allkode {
		itm := svc.Dokter{
			KodeDokter:         v.KodeDokter,
			NamaDokter:         v.NamaDokter,
			JenisKelamin:       v.JenisKelamin,
			Alamat:             v.Alamat,
			NomorTelepon:       v.NomorTelepon,
			KodeKategoriDokter: v.KodeKategoriDokter,
			Biaya:              v.Biaya}
		rsp = append(rsp, itm)
	}
	return rsp, nil
}

func makeClientAddDokterEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn,
		grpcName,
		"AddDokter",
		encodeAddDokterRequest,
		decodeDokterResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "AddDokter")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "AddDokter",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDokterByKodeEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDokterByKode",
		encodeReadDokterByKodeRequest,
		decodeReadDokterByKodeRespones,
		pb.ReadDokterByKodeResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDokterByKode")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDokterByKode",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientReadDokterEndpoint(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {

	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"ReadDokter",
		encodeReadDokterRequest,
		decodeReadDokterResponse,
		pb.ReadDokterResp{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "ReadDokter")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "ReadDokter",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}

func makeClientUpdateDokter(conn *grpc.ClientConn, timeout time.Duration, tracer stdopentracing.Tracer,
	logger log.Logger) endpoint.Endpoint {
	endpoint := grpctransport.NewClient(
		conn, grpcName,
		"UpdateDokter",
		encodeUpdateDokterRequest,
		decodeDokterResponse,
		google_protobuf.Empty{},
		grpctransport.ClientBefore(opentracing.ContextToGRPC(tracer, logger)),
	).Endpoint()

	endpoint = opentracing.TraceClient(tracer, "UpdateDokter")(endpoint)
	endpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "UpdateDokter",
		Timeout: timeout,
	}))(endpoint)

	return endpoint
}
