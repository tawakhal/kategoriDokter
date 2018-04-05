package endpoint

import (
	"context"

	svc "rumahsakit/dokter/dokter/server"

	kit "github.com/go-kit/kit/endpoint"
)

type DokterEndpoint struct {
	AddDokterEndpoint        kit.Endpoint
	ReadDokterByKodeEndpoint kit.Endpoint
	ReadDokterEndpoint       kit.Endpoint
	UpdateDokterEndpoint     kit.Endpoint
}

func NewDokterEndpoint(service svc.DokterService) DokterEndpoint {
	addDokterEp := makeAddDokterEndpoint(service)
	readDokterByKodeEp := makeReadDokterByKodeEndpoint(service)
	readDokterEp := makeReadDokterEndpoint(service)
	updateDokterEp := makeUpdateDokterEndpoint(service)
	return DokterEndpoint{AddDokterEndpoint: addDokterEp,
		ReadDokterByKodeEndpoint: readDokterByKodeEp,
		ReadDokterEndpoint:       readDokterEp,
		UpdateDokterEndpoint:     updateDokterEp,
	}
}

func makeAddDokterEndpoint(service svc.DokterService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Dokter)
		err := service.AddDokterService(ctx, req)
		return nil, err
	}
}

func makeReadDokterByKodeEndpoint(service svc.DokterService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Dokter)
		result, err := service.ReadDokterByKodeService(ctx, req.KodeDokter)
		/*return svc.Customer{CustomerId: result.CustomerId, Name: result.Name,
		CustomerType: result.CustomerType, Mobile: result.Mobile, Email: result.Email,
		Gender: result.Gender, CallbackPhone: result.CallbackPhone, Status: result.Status}, err*/
		return result, err
	}
}

func makeReadDokterEndpoint(service svc.DokterService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		result, err := service.ReadDokterService(ctx)
		return result, err
	}
}

func makeUpdateDokterEndpoint(service svc.DokterService) kit.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(svc.Dokter)
		err := service.UpdateDokterService(ctx, req)
		return nil, err
	}
}
