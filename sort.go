	package main 

	import "fmt"

	func lessDate(tgl1, tgl2 string) bool{
		t1 := parseDate(tgl1)
		t2 := parseDate(tgl2)
		return t1.Before(t2)
	}

	func selectionSortByDate(data *daftarTagihan){
		n := len(*data)
		for i := 0; i < n-1; i++{
			minIdx := i
			for j := i + 1; j < n; j++{
				if lessDate((*data)[j].JatuhTempo, (*data)[minIdx].JatuhTempo){
					minIdx = j
				}
			}
			if minIdx != i {
				(*data)[i], (*data)[minIdx] = (*data)[minIdx], (*data)[i]
			}
		}
	}

	func insertionSortByDate(data *daftarTagihan){
		n := len(*data)
		for i := 1; i < n; i++{
			key := (*data)[i]
			j := i - 1
			for j >= 0 && lessDate(key.JatuhTempo, (*data)[j].JatuhTempo){
				(*data)[j+1] = (*data)[j]
				j-- 
			}
			(*data)[j+1] = key
		}
	}

	func urutkanTagihan(data *daftarTagihan){
		if len(*data) == 0{
			fmt.Println("\nTidak ada data tagihan untuk diurutkan.")
			return
		}

		fmt.Println("\n====== URUTKAN TAGIHAN ======")
		fmt.Println("Pilih algoritma yang akan digunakan: ")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		pilihAlgoritma := bacaInput("Pilih (1/2): ")

		switch pilihAlgoritma{
		case "1":
			selectionSortByDate(data)
			fmt.Println("\nPengurutan tagihan dengan Selection Sort berhasil")
		case "2":
			insertionSortByDate(data)
			fmt.Println("\nPengurutan tagihan dengan Insertion Sort berhasil")
		default:
			fmt.Println("Pilihan tidak valid")
			return
		}

		fmt.Println("\nHasil setelah diurutkan (tanggal terdekat terlebih): ")
		tampilSemua(*data)
	}