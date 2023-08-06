package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	ID string
	Nama string
	Alamat string
	Pekerjaan string
	Alasan string

}


var biodataList = []Biodata{
	{ID: "1", Nama: "Fao1", Alamat: "Bintaro", Pekerjaan: "Karyawan", Alasan: "Dibutuhkan dalam pekerjaan"},
	{ID: "2", Nama: "Fao2", Alamat: "Jakarta", Pekerjaan: "Mahasiswa", Alasan: "Dibutuhkan dalam kuliah"},
	{ID: "3", Nama: "Fao3", Alamat: "Bandung", Pekerjaan: "Wiraswasta", Alasan: "Mengisi waktu luang"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Tolong masukan nama atau nomor absen")
		fmt.Println("Contoh: 'go run main.go Fao2' atau 'go run main.go 2'")
		os.Exit(1)
	}

	biodataName := os.Args[1]

	biodataList := findBiodata(biodataName)

	if biodataList != nil {
		fmt.Println("ID:", biodataList.ID)
		fmt.Println("Nama:", biodataList.Nama)
		fmt.Println("Alamat:", biodataList.Alamat)
		fmt.Println("Pekerjaan:", biodataList.Pekerjaan)
		fmt.Println("Alasan:", biodataList.Alasan)
	} else {
		fmt.Println("Biodata tidak ditemukan.")
	}
}

func findBiodata(name string) *Biodata {
	for i := range biodataList {
		if biodataList[i].Nama == name || biodataList[i].ID == name {
			return &biodataList[i]
		}
	}
	return nil
}
