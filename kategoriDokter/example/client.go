package main

import (
	"context"
	"fmt"
	"time"

	cli "rumahsakit/dokter/dokter/endpoint"
	opt "rumahsakit/dokter/util/grpc"
	util "rumahsakit/dokter/util/microservice"

	tr "github.com/opentracing/opentracing-go"
)

func main() {
	logger := util.Logger()
	tracer := tr.GlobalTracer()
	// dicobanya 3x, timeoutnya, dan balancingnya
	option := opt.ClientOption{Retry: 3, RetryTimeout: 500 * time.Millisecond, Timeout: 30 * time.Second}

	client, err := cli.NewGRPCDokterClient([]string{"127.0.0.1:2181"}, nil, option, tracer, logger)
	if err != nil {
		logger.Log("error", err)
	}

	//Add Dokter
	//client.AddDokterService(context.Background(), svc.Dokter{KodeDokter: "DOK001", NamaDokter: "Olgi Tawakhal", JenisKelamin: "L", NomorTelepon: "087889967666", KodeKategoriDokter: "KKD001", Biaya: 2000000})

	//Get Dokter By Kode Kode
	dokKode, _ := client.ReadDokterByKodeService(context.Background(), "DOK001")
	fmt.Println("Dokter based on Kode:", dokKode)

	//List Customer
	// cuss, _ := client.ReadDokterService(context.Background())
	// fmt.Println("All Dokter:", cuss)

	//Update Customer
	// client.UpdateDokterService(context.Background(), svc.Dokter{NamaDokter: "Olgi Tawakhal", JenisKelamin: "L", Alamat: "oakdowkdoawkdoawkd", NomorTelepon: "01231", KodeKategoriDokter: "KKD001", Biaya: 5000, KodeDokter: "DOK001"})

	//List Dokter
	dokk, _ := client.ReadDokterService(context.Background())
	fmt.Println("All Dokter:", dokk)

}
