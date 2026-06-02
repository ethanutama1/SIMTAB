package main

import (
	"fmt"
	"sort"
	"strings"
)

func getField(t Tagihan, field string)string{
	switch field{
	case "nama":
		return t.Nama
	case "kategori" :
		return t.Kategori
	default:
		return ""
	}
}

func sequentialSearch(data daftarTagihan, field, keyword string)daftarTagihan{

	var hasil daftarTagihan
	keyword = strings.ToLower(keyword)

	for _, t := range data{
		if strings.ToLower(getField(t, field)) == keyword{
			hasil = append(hasil, t)
		}
	}

	return hasil
}

func binarySearch(data daftarTagihan, field, keyword string)daftarTagihan{
	sorted := make(daftarTagihan, len(data))
	copy(sorted, data)

	sort.Slice(sorted, func(i, j int) bool {
		return strings.ToLower(getField(sorted[i], field)) < strings.ToLower(getField(sorted[i], field))
	})

	keyword = strings.ToLower(keyword)
	var hasil daftarTagihan

	low, high := 0, len(sorted)-1
	foundIdx := -1
	for low <= high{
		mid := (low + high) / 2
		midValue := strings.ToLower(getField(sorted[mid], field))
		if midValue == keyword{
			foundIdx = mid
			break
		}else if midValue < keyword{
			low = mid + 1
		}else{
			high = mid - 1
		}
	}

	if foundIdx != -1{
		for i := foundIdx; i >= 0 && strings.ToLower(getField(sorted[i], field)) == keyword; i--{
			hasil = append(hasil, sorted[i])
		}

		for i := foundIdx + 1; i < len(sorted) && strings.ToLower(getField(sorted[i], field)) == keyword; i++{
			hasil = append(hasil, sorted[i])
		}
	}
	return hasil
}

func tampilHasilPencarian(hasil	daftarTagihan){
	if len(hasil) == 0{
		fmt.Println("\nTidak ada data yang cocok!")
		return
	}
	fmt.Printf("\nDitemukan %d tagihan:\n", len(hasil))
	tampilSemua(hasil)
}

func cariTagihan(data daftarTagihan){
	fmt.Println("\n====== CARI TAGIHAN ======")
	fmt.Println("Pilih algoritma pencarian")
	fmt.Println("1. Sequential Search")
	fmt.Println("2. Binary Search")
	pilihAlgoritma := bacaInput("Pilih (1/2): ")

	if pilihAlgoritma != "1" && pilihAlgoritma != "2"{
		fmt.Println("Pilihan yang dimasukkan tidak valid.")
		return	
	}

	fmt.Println("\nPilih kriteria pencarian:")
	fmt.Println("1. Nama Tagihan")
	fmt.Println("2. Kategori")
	pilihKategori := bacaInput("Pilih (1/2): ")

	var field string
	switch pilihKategori {
	case "1":
		field = "nama"
	case "2":
		field = "kategori"
	default:
		fmt.Println("Kategori yang dimasukkan tidak valid.")
		return
	}

	keyword := bacaInput("Masukkan kata kunci: ")
	if keyword == " "{
		fmt.Println("Kata kunci tidak boleh kosong.")
		return
	}

	var hasil daftarTagihan
	if pilihAlgoritma == "1"{
		hasil = sequentialSearch(data, field, keyword)
	}else{
		hasil = binarySearch(data, field, keyword)
	}

	tampilHasilPencarian(hasil)

}