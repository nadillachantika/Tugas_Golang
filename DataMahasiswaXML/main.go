package main

import(
	"database/sql"
	"encoding/xml"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

)

var db *sql.DB
var err error

// Order struct (Model) ...
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
func getMahasiswa(w http.ResponseWriter, r *http.Request) {
	
	var mhs Mahasiswa
	var almdet AlamatDetail

	sql := `SELECT 
				 id_mahasiswa,
				 IFNULL(no_bp,'') no_bp,
				 IFNULL(nama,'') nama,
				 IFNULL(jurusan,'') jurusan,
				 IFNULL(prodi,'') prodi,
				 FROM mahasiswa WHERE id_mahasiswa IN (7)`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&mhs.MahasiswaID, &mhs.NoBp, &mhs.Prodi)

		if err != nil {
			panic(err.Error())

		}

		sqlDetail := `SELECT
						alamat_detail.id_mahasiswa,
						mahasiswa.no_bp,
						mahasiswa.jurusan,
						mahasiswa.prodi,
						alamat_detail.jalan,
						alamat_detail.kelurahan,
						alamat_detail.kecamatan,
						alamat_detail.kota_kabupaten,
						alamat_detail.provinsi
					FROM 
						alamat_detail
						INNER JOIN mahasiswa
						ON (alamat_detail.id_mahasiswa = mahasiswa.id_mahasiswa)

						WHERE alamat_detail.id_mahasiswa = ?`

		mhsID := &mhs.MahasiswaID

		resultDetail, errDet := db.Query(sqlDetail, *mhsID)

		defer resultDetail.Close()

		if errDet != nil {
			panic(err.Error())
		}

		for resultDetail.Next() {
			err := resultDetail.Scan(&almdet.MahasiswaID, &almdet.Jalan, &almdet.Kelurahan, &almdet.Kecamatan, &almdet.KabKota, &almdet.Provinsi)

			if err != nil {
				panic(err.Error())
			}

			mhs.AlamatDet = append(mhs.AlamatDet, almdet)

		}

	}
	
	w.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"))
	xml.NewEncoder(w).Encode(mhs)

}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/akademik")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	//init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/mahasiswa", getMahasiswa).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}