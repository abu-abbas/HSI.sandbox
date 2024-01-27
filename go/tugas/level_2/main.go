package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/abu-abbas/level_2/generator"
)

func main() {
	// generate NIK
	nikIkhwan, err 	:= generateNIK("ikhwan", 2024, 10)
	nikAkhwat, _ 	:= generateNIK("akhwat", 2024, 10)
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generate NIK Ikhwan:", nikIkhwan, "\n")
	fmt.Println("Generate NIK Akhwat:", nikAkhwat, "\n")

	// generate NIK Lanjutan
	adminIkhwan1, _ := generateNIKLanjutan("ARN171-07770", 2)
	adminIkhwan2, _ := generateNIKLanjutan("ARN181-08880", 2)
	adminAkhwat1, _ := generateNIKLanjutan("ART191-06660", 2)
	adminAkhwat2, _ := generateNIKLanjutan("ART201-05550", 2)

	fmt.Println("Generate NIK Lanjutan (Ikhwan):", adminIkhwan1, adminIkhwan2, "\n")
	fmt.Println("Generate NIK Lanjutan (Akhwat):", adminAkhwat1, adminAkhwat2, "\n")

	generateKelompok(
		nikIkhwan, 
		nikAkhwat, 
		adminIkhwan1, 
		adminIkhwan2, 
		adminAkhwat1,
		adminAkhwat2,
	)
}

func generateNIK(gender string, tahun int, jumlah int) ([]string, error) {
	result, err := generator.NIK(gender, tahun, jumlah, 0, "")
	return result, err 
}

func generateNIKLanjutan(nik string, jumlah int) ([]string, error) {
	gender, tahun, smt, urutan := generator.ParseNIK(nik)
	result, err := generator.NIK(gender, tahun, jumlah, urutan, smt)
	return result, err
}

func generateKelompok(ids ...[]string) {
	var nikIkhwan, nikAkhwat []string
	for _, id := range ids {
		for _, nik := range id {

			if strings.Contains(nik, "ARN") {
				nikIkhwan = append(nikIkhwan, nik)
			}        

			if strings.Contains(nik, "ART") {
				nikAkhwat = append(nikAkhwat, nik)
			} 
		}
	}

	sort.Strings(nikIkhwan)
	sort.Strings(nikAkhwat)

	fmt.Println("Kelompok Ikhwan:", nikIkhwan, "\n")
	fmt.Println("Kelompok Akhwat:", nikAkhwat, "\n")
}
