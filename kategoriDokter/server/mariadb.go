package server

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addDokter          = `INSERT INTO dokter(kode_dokter,nama_dokter,jenis_kelamin,alamat,nomor_telepon,kode_kategori_dokter,biaya) VALUES (?,?,?,?,?,?,?)`
	selectDokterByKode = `SELECT nama_dokter,jenis_kelamin,alamat,nomor_telepon,kode_kategori_dokter,biaya FROM dokter WHERE kode_dokter = ?`
	selectDokter       = `SELECT kode_dokter, nama_dokter, jenis_kelamin, alamat, nomor_telepon, kode_kategori_dokter,biaya FROM dokter`
	updateDokter       = `UPDATE dokter SET nama_dokter =?, jenis_kelamin =?, alamat = ?, nomor_telepon = ?, kode_kategori_dokter =?, biaya =? WHERE kode_dokter =?`
)

// Langkah ke 4

type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) AddDokter(dokter Dokter) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addDokter, dokter.KodeDokter, dokter.NamaDokter, dokter.JenisKelamin, dokter.Alamat, dokter.NomorTelepon, dokter.KodeKategoriDokter, dokter.Biaya)
	fmt.Println(err)
	if err != nil {
		return err

	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadDokterByKode(kode string) (Dokter, error) {
	dokter := Dokter{KodeDokter: kode}
	err := rw.db.QueryRow(selectDokterByKode, kode).Scan(&dokter.NamaDokter, &dokter.JenisKelamin, &dokter.Alamat, &dokter.NomorTelepon, &dokter.KodeKategoriDokter, &dokter.Biaya)

	if err != nil {
		return Dokter{}, err
	}

	return dokter, nil
}

func (rw *dbReadWriter) ReadDokter() (Dokters, error) {
	dokter := Dokters{}
	rows, _ := rw.db.Query(selectDokter)
	defer rows.Close()
	for rows.Next() {
		var d Dokter
		err := rows.Scan(&d.KodeDokter, &d.NamaDokter, &d.JenisKelamin, &d.Alamat, &d.NomorTelepon, &d.KodeKategoriDokter, &d.Biaya)
		if err != nil {
			fmt.Println("error query:", err)
			return dokter, err
		}
		dokter = append(dokter, d)
	}
	//fmt.Println("db nya:", customer)
	return dokter, nil
}

func (rw *dbReadWriter) UpdateDokter(dok Dokter) error {
	//fmt.Println("update")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateDokter, dok.NamaDokter, dok.JenisKelamin, dok.Alamat, dok.NomorTelepon, dok.KodeKategoriDokter, dok.Biaya, dok.KodeDokter)

	//fmt.Println("name:", cus.Name, cus.CustomerId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return tx.Commit()
}
