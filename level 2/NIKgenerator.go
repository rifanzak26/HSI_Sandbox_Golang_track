package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

func generateNIK(gender string, tahun string, jumlahYangDigenerate int) []string {
	result := make([]string, jumlahYangDigenerate+1)
	var g string
	if gender == "ikhwan" {
		g = "N"
	} else {
		g = "T"
	}

	for i := 1; i < jumlahYangDigenerate+1; i++ {
		semester := rand.Intn(2) + 1
		nik := fmt.Sprintf("AR%s%s%d-%05d",
			g,
			tahun[len(tahun)-2:],
			semester,
			i)
		result[i] = nik
	}

	return result
}

func generateNIKLanjutan(nik string, jumlahYangDigenerate int) ([]string, error) {
	// Validasi NIK (bdsk length nik)
	if len(nik) != 12 {
		return nil, errors.New("NIK tidak valid")
	}

	// Ambil 4 digit terakhir dari NIK
	nomorUrut := nik[7:]
	noUrut, err := strconv.Atoi(nomorUrut)
	if err != nil {
		return nil, err
	}

	// Generate NIK lanjutan
	result := make([]string, jumlahYangDigenerate+1)
	for i := 1; i < jumlahYangDigenerate+1; i++ {
		nikLanjutan := fmt.Sprintf("%s-%05d", nik[:6], noUrut+i)
		result[i] = nikLanjutan
	}

	return result, nil
}

func main() {
	result := generateNIK("ikhwan", "2023", 10)
	for _, nik := range result {
		fmt.Println(nik)
	}

	var err error
	result, err = generateNIKLanjutan("ARN342-00001", 10)
	fmt.Println(err)
	for _, nik := range result {
		fmt.Println(nik)
	}
}

/* 1. Buatkan HSI NIK generator.


func generateNIK(gender, tahun, jumlahYangDigenerate) []string {
       return []string{}
}
generateNIK("ikhwan", 2019, 10) // Berarti generate 10 NIK.

dengan hanya menggunakan standard library Go.

Spesifikasi NIK adalah sebagai berikut:

AR<IKWAN_OR_AKHWAT><YEAR><SEMESTER>-<5_DIGITS_SEQUENCE>

Contoh: ARN192-00375

N: berarti ikhwan
19: Tahun pembuatan, 2019
2: Pembuatan semester akhir. Dalam satu bulan ada 2 semester. Sehingga bulan ke-1 sampai ke-5 adalah 1, sedangkan bulan ke-6 sampai ke-12 adalah 2.
00375: 5 digit urutan */
