package server

import (
	"context"
)

// langka ke lima
type dokter struct {
	writer ReadWriter
}

func NewDokter(writer ReadWriter) DokterService {
	return &dokter{writer: writer}
}

//Methode pada interface CustomerService di service.go
func (d *dokter) AddDokterService(ctx context.Context, dokter Dokter) error {
	//fmt.Println("customer")
	err := d.writer.AddDokter(dokter)
	if err != nil {
		return err
	}

	return nil
}

func (d *dokter) ReadDokterByKodeService(ctx context.Context, mob string) (Dokter, error) {
	cus, err := d.writer.ReadDokterByKode(mob)
	//fmt.Println(cus)
	if err != nil {
		return cus, err
	}
	return cus, nil
}

func (d *dokter) ReadDokterService(ctx context.Context) (Dokters, error) {
	dok, err := d.writer.ReadDokter()
	//fmt.Println("customer", cus)
	if err != nil {
		return dok, err
	}
	return dok, nil
}

func (d *dokter) UpdateDokterService(ctx context.Context, dok Dokter) error {
	err := d.writer.UpdateDokter(dok)
	if err != nil {
		return err
	}
	return nil
}
