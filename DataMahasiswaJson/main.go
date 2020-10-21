package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Mahasiswa struct {
	MahasiswaID int            `json:"id_mahasiswa"`
	NoBp        string         `json:"no_bp"`
	Nama        string         `json:"nama"`
	Jurusan     string         `json:"jurusan"`
	Prodi       string         `json:"prodi"`
	AlamatDet   []AlamatDetail `json:"alamat_detail"`
}


type AlamatDetail struct {
	MahasiswaID int    `json:"id_mahasiswa"`
	Jalan       string `json:"alamat"`
	Kelurahan   string `json:"kelurahan"`
	Kecamatan   string `json:"kecamatan"`
	KabKota     string `json:"kota_kabupaten"`
	Provinsi    string `json:"provinsi"`
}

func main(){

	url := "http://localhost:8080/mahasiswa"

	spaceClient := http.Client{
		Timeout: time.Second * 2, //Timeout after 2 second

	}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil{
		log.Fatal(err)
	}

	req.Header.Set("User-Agent","spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	if res.Body != nil{
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

	for _, alm := range mhs.AlamatDet {
		fmt.Println("Mahasiswa ID",alm.MahasiswaID)
		fmt.Println("Jalan", alm.Jalan)
		fmt.Println("Kelurahan", alm.Kelurahan)
		fmt.Println("Kecamatan", alm.Kecamatan)
		fmt.Println("Kota/Kabupaten",alm.KabKota)
		fmt.Println("Provinsi",alm.Provinsi)

	}



}