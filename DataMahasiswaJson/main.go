package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Mahasiswa struct (Model) ...
type Mahasiswa struct {
	MahasiswaID int            `json:"MahasiswaID"`
	NoBp        string         `json:"NoBp"`
	Nama        string         `json:"Nama"`
	Jurusan     string         `json:"Jurusan"`
	Prodi       string         `json:"Prodi"`
	AlamatDet   []AlamatDetail `json:"AlamatDet"`
	NilaiDet    []NilaiDetail  `json:"NilaiDet"`
}

//AlamatDetail struct (Model)
type AlamatDetail struct {
	KodeAlamat    string `json:"KodeAlamat"`
	Jalan         string `json:"Jalan"`
	Kelurahan     string `json:"Kelurahan"`
	Kecamatan     string `json:"Kecamatan"`
	KotaKabupaten string `json:"KotaKabupaten"`
	Provinsi      string `json:"Provinsi"`
}

// NilaiMahasiswa struct
type NilaiDetail struct {
	NamaMatkul string `json:"NamaMatkul"`
	Nilai      string `json:"Nilai"`
	Semester   string `json:"Semester"`
}

func main() {

	url := "http://localhost:8080/mahasiswa"

	spaceClient := http.Client{
		Timeout: time.Second * 2, //Timeout after 2 second

	}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	mhs := Mahasiswa{}
	jsonErr := json.Unmarshal(body, &mhs)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(mhs.MahasiswaID)
	fmt.Println(mhs.Nama)
	fmt.Println(mhs.NoBp)

	for _, almt := range mhs.AlamatDet {
		fmt.Println("Jalan", almt.Jalan)
		fmt.Println("Kelurahan", almt.Kelurahan)
		fmt.Println("Kecamatan", almt.Kecamatan)
		fmt.Println("Kota/Kabupaten", almt.KotaKabupaten)
		fmt.Println("Provinsi", almt.Provinsi)

	}
	for _, nilai := range mhs.NilaiDet {
		fmt.Println("NamaMatkul", nilai.NamaMatkul)
		fmt.Println("Nilai", nilai.Nilai)
		fmt.Println("Semester", nilai.Semester)

	}

}
