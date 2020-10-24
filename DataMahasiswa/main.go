package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

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

//Get all mahasiswa

func getMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var mhs Mahasiswa
	var almdet AlamatDetail
	//var nilaidet NilaiDetail

	sql := `SELECT 
				 MahasiswaID,
				 IFNULL(NoBp,'') NoBp,
				 IFNULL(Nama,'') Nama,
				 IFNULL(Jurusan,'') Jurusan,
				 IFNULL(Prodi,'') Prodi
				 
				 FROM mahasiswa WHERE MahasiswaID IN (4)`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&mhs.MahasiswaID, &mhs.NoBp, &mhs.Nama, &mhs.Jurusan, &mhs.Prodi)

		if err != nil {
			panic(err.Error())

		}

		sqlDetail := `SELECT
						alamat_details.KodeAlamat,
						alamat_details.Jalan,
						alamat_details.Kelurahan,
						alamat_details.Kecamatan,
						alamat_details.KotaKabupaten,
						alamat_details.Provinsi
					FROM 
						mahasiswa
						INNER JOIN alamat_details
						ON (mahasiswa.MahasiswaID = alamat_details.MahasiswaID)

						WHERE mahasiswa.MahasiswaID = ?`

		mhsID := &mhs.MahasiswaID

		resultDetail, errDet := db.Query(sqlDetail, *mhsID)

		fmt.Println(*mhsID)
		defer resultDetail.Close()

		if errDet != nil {
			panic(err.Error())
		}

		for resultDetail.Next() {
			err := resultDetail.Scan(&almdet.KodeAlamat, &almdet.Jalan, &almdet.Kelurahan, &almdet.Kecamatan, &almdet.KotaKabupaten, &almdet.Provinsi)

			if err != nil {
				panic(err.Error())
			}

			mhs.AlamatDet = append(mhs.AlamatDet, almdet)

		}

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mhs)

}

func getNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var mhs Mahasiswa
	var nilaidet NilaiDetail

	sql := `SELECT 
				 MahasiswaID,
				 IFNULL(NoBp,'') NoBp,
				 IFNULL(Nama,'') Nama,
				 IFNULL(Jurusan,'') Jurusan,
				 IFNULL(Prodi,'') Prodi
		
				 FROM mahasiswa WHERE MahasiswaID IN (4)`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&mhs.MahasiswaID, &mhs.NoBp, &mhs.Nama, &mhs.Jurusan, &mhs.Prodi)

		if err != nil {
			panic(err.Error())

		}

		sqlDetail := `SELECT
						mata_kuliah.NamaMatkul,
						tabel_nilai.Nilai as Nilai,
						tabel_nilai.Semester
					FROM 
			
						tabel_nilai
						INNER JOIN mata_kuliah
						ON (tabel_nilai.KodeMatkul = mata_kuliah.KodeMatkul)
						INNER JOIN mahasiswa
						ON (mahasiswa.MahasiswaID = tabel_nilai.MahasiswaID)

						WHERE  mahasiswa.MahasiswaID = ?`

		mhsID := &mhs.MahasiswaID

		resultDetail, errDet := db.Query(sqlDetail, *mhsID)

		fmt.Println(*mhsID)
		defer resultDetail.Close()

		if errDet != nil {
			panic(err.Error())
		}

		for resultDetail.Next() {
			err := resultDetail.Scan(&nilaidet.NamaMatkul, &nilaidet.Nilai, &nilaidet.Semester)

			if err != nil {
				panic(err.Error())
			}

			mhs.NilaiDet = append(mhs.NilaiDet, nilaidet)

		}

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mhs)

}
func getAllData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var almdet AlamatDetail
	var mhs Mahasiswa
	var nilaidet NilaiDetail
	params := mux.Vars(r)

	sql := `SELECT 
				 MahasiswaID,
				 IFNULL(NoBp,'') NoBp,
				 IFNULL(Nama,'') Nama,
				 IFNULL(Jurusan,'') Jurusan,
				 IFNULL(Prodi,'') Prodi
		
				 FROM mahasiswa WHERE MahasiswaID = ?`

	result, err := db.Query(sql, params["id"])

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&mhs.MahasiswaID, &mhs.NoBp, &mhs.Nama, &mhs.Jurusan, &mhs.Prodi)

		if err != nil {
			panic(err.Error())

		}

		sqlDetail := `SELECT
						mata_kuliah.NamaMatkul,
						tabel_nilai.Nilai as Nilai,
						tabel_nilai.Semester
					FROM 
			
						tabel_nilai
						INNER JOIN mata_kuliah
						ON (tabel_nilai.KodeMatkul = mata_kuliah.KodeMatkul)
						INNER JOIN mahasiswa
						ON (mahasiswa.MahasiswaID = tabel_nilai.MahasiswaID)

						WHERE  mahasiswa.MahasiswaID = ?`

		sqlDetailAlamat := `SELECT
						alamat_details.KodeAlamat,
						alamat_details.Jalan,
						alamat_details.Kelurahan,
						alamat_details.Kecamatan,
						alamat_details.KotaKabupaten,
						alamat_details.Provinsi
					FROM 
						mahasiswa
						INNER JOIN alamat_details
						ON (mahasiswa.MahasiswaID = alamat_details.MahasiswaID)

						WHERE mahasiswa.MahasiswaID = ?`

		mhsID := &mhs.MahasiswaID

		resultDetail, errDet := db.Query(sqlDetail, *mhsID)
		resultDetailAlamat, errDet := db.Query(sqlDetailAlamat, *mhsID)

		fmt.Println(*mhsID)
		defer resultDetail.Close()
		defer resultDetailAlamat.Close()

		if errDet != nil {
			panic(err.Error())
		}

		for resultDetail.Next() {
			err := resultDetail.Scan(&nilaidet.NamaMatkul, &nilaidet.Nilai, &nilaidet.Semester)

			if err != nil {
				panic(err.Error())
			}

			mhs.NilaiDet = append(mhs.NilaiDet, nilaidet)

		}
		for resultDetailAlamat.Next() {
			err := resultDetailAlamat.Scan(&almdet.KodeAlamat, &almdet.Jalan, &almdet.Kelurahan, &almdet.Kecamatan, &almdet.KotaKabupaten, &almdet.Provinsi)

			if err != nil {
				panic(err.Error())
			}

			mhs.AlamatDet = append(mhs.AlamatDet, almdet)

		}

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mhs)

}

func createMahasiswa(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		Nama := r.FormValue("Nama")
		NoBp := r.FormValue("NoBp")
		Jurusan := r.FormValue("Jurusan")
		Prodi := r.FormValue("Prodi")

		stmt, err := db.Prepare("INSERT INTO mahasiswa (Nama,NoBp,Jurusan,Prodi) VALUES (?,?,?,?)")

		_, err = stmt.Exec(Nama, NoBp, Jurusan, Prodi)
		if err != nil {
			fmt.Fprint(w, "Data Duplicat")
		} else {
			fmt.Fprint(w, "Data Created")
		}

	}

}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/perkuliahan")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/mahasiswa", getMahasiswa).Methods("GET")
	r.HandleFunc("/nilaimhs", getNilaiMahasiswa).Methods("GET")
	r.HandleFunc("/mahasiswadata/{id}", getAllData).Methods("GET")
	r.HandleFunc("/mahasiswa", createMahasiswa).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
