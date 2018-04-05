package endpoint

import (
	"context"
	"fmt"

	sv "rumahsakit/dokter/dokter/server"
)

func (de DokterEndpoint) AddDokterService(ctx context.Context, dokter sv.Dokter) error {
	_, err := de.AddDokterEndpoint(ctx, dokter)
	return err
}

func (de DokterEndpoint) ReadDokterByKodeService(ctx context.Context, Kode string) (sv.Dokter, error) {
	req := sv.Dokter{KodeDokter: Kode}
	fmt.Println(req)
	resp, err := de.ReadDokterByKodeEndpoint(ctx, req)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	cus := resp.(sv.Dokter)
	return cus, err
}

func (de DokterEndpoint) ReadDokterService(ctx context.Context) (sv.Dokters, error) {
	resp, err := de.ReadDokterEndpoint(ctx, nil)
	fmt.Println("ce resp", resp)
	if err != nil {
		fmt.Println("error pada endpoint: ", err)
	}
	return resp.(sv.Dokters), err
}

func (de DokterEndpoint) UpdateDokterService(ctx context.Context, dok sv.Dokter) error {
	_, err := de.UpdateDokterEndpoint(ctx, dok)
	if err != nil {
		fmt.Println("error pada endpoint:", err)
	}
	return err
}
