package server

import (
	"context"
)

type Status int32

const (
	//ServiceID is dispatch service ID

	//ini adalah konfigurasi sub domainnya
	ServiceID        = "rumahsakit"
	CreatedBy        = "Olgi"
	onAdd     Status = 1
)

type Dokter struct {
	KodeDokter         string
	NamaDokter         string
	JenisKelamin       string
	Alamat             string
	NomorTelepon       string
	KodeKategoriDokter string
	Biaya              int64
}
type Dokters []Dokter

/*type Location struct {
	customerID   int64
	label        []int32
	locationType []int32
	name         []string
	street       string
	village      string
	district     string
	city         string
	province     string
	latitude     float64
	longitude    float64
}*/

// ini interface untuk melakukan read
type ReadWriter interface {
	AddDokter(Dokter) error
	ReadDokterByKode(string) (Dokter, error)
	ReadDokter() (Dokters, error)
	UpdateDokter(Dokter) error
}

// ini interface yang mempunyai nilai return yang berupa interfase
type DokterService interface {
	AddDokterService(context.Context, Dokter) error
	ReadDokterByKodeService(context.Context, string) (Dokter, error)
	ReadDokterService(context.Context) (Dokters, error)
	UpdateDokterService(context.Context, Dokter) error
}
