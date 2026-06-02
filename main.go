package main

import (
	"fmt"
	"strings"
)

func tampilkanHeader(){
	fmt.Println()
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("=	SIMTAB - Sistem Manajemen Tagihan	 =")
	fmt.Println(strings.Repeat("=", 50))
}

func tampilkanMenu(){
	fmt.Println("\n ===== MENU UTAMA =====")
	fmt.Println(" 1. Tampilkan Semua Tagihan")
	fmt.Println(" 2. Tambah Tagihan Baru")
	fmt.Println(" 3. Edit Tagihan")
	fmt.Println(" 4. Hapus Tagihan")
	fmt.Println(" 5. Cari Tagihan")
	fmt.Println("----------------------------")
	fmt.Println(" 6. Urutkan Tagihan (Coming Soon)")
	fmt.Println(" 7. Statistik Tagihan (Coming Soon)")
	fmt.Println("----------------------------")
	fmt.Println(" 0. Keluar")
	fmt.Print("\n Pilih menu: ")
}

func main(){

	var dataTagihan daftarTagihan
	nextID := 1

	tampilkanHeader()

	for{
		tampilkanMenu()

		inputPilihan := bacaInput("")
		var pilihan int
		_, err := fmt.Sscan(inputPilihan, &pilihan)
		if err != nil{
			scanner.Scan()
			fmt.Println("\n Input tidak valid. Silahkan masukkan angka 0-7")
			continue
		}

		switch pilihan{
		case 1:
			fmt.Println("\nDAFTAR TAGIHAN ")
			tampilSemua(dataTagihan)

		case 2:
			tambahTagihan(&dataTagihan, &nextID)
		
		case 3: 
			ubahTagihan(&dataTagihan)

		case 4:
			hapusTagihan(&dataTagihan)

		
		case 5:
			cariTagihan(dataTagihan)

		case 6, 7:
			fmt.Println("Fitur ini akan tersedia dalam waktu dekat!")

		case 0:
			fmt.Println("\n Terima kasih sudah menggunakan SIMTAB!")
			fmt.Println(strings.Repeat("=", 50))
			return	

		default:
			fmt.Println("\n Pilihan yang dimasukkan tidak tersedia. Silahkan masukkan angka 0-7!")
		}
	}
}