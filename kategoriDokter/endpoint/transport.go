package endpoint

import (
	"context"

	scv "rumahsakit/dokter/dokter/server"

	pb "rumahsakit/dokter/dokter/grpc"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/tracing/opentracing"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
	stdopentracing "github.com/opentracing/opentracing-go"
	oldcontext "golang.org/x/net/context"
)

type grpcDokterServer struct {
	addDokter        grpctransport.Handler
	readDokterByKode grpctransport.Handler
	readDokter       grpctransport.Handler
	updateDokter     grpctransport.Handler
}

func NewGRPCDokterServer(endpoints DokterEndpoint, tracer stdopentracing.Tracer,
	logger log.Logger) pb.DokterServiceServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}
	return &grpcDokterServer{
		addDokter: grpctransport.NewServer(endpoints.AddDokterEndpoint,
			decodeAddDokterRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "AddDokter", logger)))...),
		readDokterByKode: grpctransport.NewServer(endpoints.ReadDokterByKodeEndpoint,
			decodeReadDokterByKodeRequest,
			encodeReadDokterByKodeResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDokterByKode", logger)))...),
		readDokter: grpctransport.NewServer(endpoints.ReadDokterEndpoint,
			decodeReadDokterRequest,
			encodeReadDokterResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "ReadDokter", logger)))...),
		updateDokter: grpctransport.NewServer(endpoints.UpdateDokterEndpoint,
			decodeUpdateDokterRequest,
			encodeEmptyResponse,
			append(options, grpctransport.ServerBefore(opentracing.GRPCToContext(tracer, "UpdateDokter", logger)))...),
	}
}

func decodeAddDokterRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.AddDokterReq)
	return scv.Dokter{KodeDokter: req.KodeDokter, NamaDokter: req.NamaDokter, JenisKelamin: req.JenisKelamin,
		Alamat: req.Alamat, NomorTelepon: req.NomorTelepon, KodeKategoriDokter: req.KodeKategoriDokter, Biaya: req.Biaya}, nil
}

func decodeReadDokterByKodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.ReadDokterByKodeReq)
	return scv.Dokter{KodeDokter: req.Kode}, nil
}

func decodeReadDokterRequest(_ context.Context, request interface{}) (interface{}, error) {
	return nil, nil
}

func decodeUpdateDokterRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UpdateDokterReq)
	return scv.Dokter{KodeDokter: req.KodeDokter, NamaDokter: req.NamaDokter, JenisKelamin: req.JenisKelamin,
		Alamat: req.Alamat, NomorTelepon: req.NomorTelepon, KodeKategoriDokter: req.KodeKategoriDokter, Biaya: req.Biaya}, nil
}

func encodeEmptyResponse(_ context.Context, response interface{}) (interface{}, error) {
	return &google_protobuf.Empty{}, nil
}

func encodeReadDokterByKodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Dokter)
	return &pb.ReadDokterByKodeResp{KodeDokter: resp.KodeDokter, NamaDokter: resp.NamaDokter, JenisKelamin: resp.JenisKelamin,
		Alamat: resp.Alamat, NomorTelepon: resp.NomorTelepon, KodeKategoriDokter: resp.KodeKategoriDokter, Biaya: resp.Biaya}, nil
}

func encodeReadDokterResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(scv.Dokters)

	rsp := &pb.ReadDokterResp{}

	for _, v := range resp {
		itm := &pb.ReadDokterByKodeResp{
			KodeDokter:         v.KodeDokter,
			NamaDokter:         v.NamaDokter,
			JenisKelamin:       v.JenisKelamin,
			Alamat:             v.Alamat,
			NomorTelepon:       v.NomorTelepon,
			KodeKategoriDokter: v.KodeKategoriDokter,
			Biaya:              v.Biaya,
		}
		rsp.Allkode = append(rsp.Allkode, itm)
	}
	return rsp, nil
}

func (s *grpcDokterServer) AddDokter(ctx oldcontext.Context, dokter *pb.AddDokterReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.addDokter.ServeGRPC(ctx, dokter)
	if err != nil {
		return nil, err
	}
	return resp.(*google_protobuf.Empty), nil
}

func (s *grpcDokterServer) ReadDokterByKode(ctx oldcontext.Context, kode *pb.ReadDokterByKodeReq) (*pb.ReadDokterByKodeResp, error) {
	_, resp, err := s.readDokterByKode.ServeGRPC(ctx, kode)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDokterByKodeResp), nil
}

func (s *grpcDokterServer) ReadDokter(ctx oldcontext.Context, e *google_protobuf.Empty) (*pb.ReadDokterResp, error) {
	_, resp, err := s.readDokter.ServeGRPC(ctx, e)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ReadDokterResp), nil
}

func (s *grpcDokterServer) UpdateDokter(ctx oldcontext.Context, cus *pb.UpdateDokterReq) (*google_protobuf.Empty, error) {
	_, resp, err := s.updateDokter.ServeGRPC(ctx, cus)
	if err != nil {
		return &google_protobuf.Empty{}, err
	}
	return resp.(*google_protobuf.Empty), nil
}
