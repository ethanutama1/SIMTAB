package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)


func bacaInput(prompt string) string{
	fmt.Print(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func tampilSemua(data daftarTagihan){
	if len(data) == 0{
		fmt.Println("Belum ada data tagihan!")
		return	
	}

	fmt.Println()
	fmt.Println(strings.Repeat("=", 85))
	fmt.Printf(" %-4s %-20s %-15s %-14s %-15s %-12s\n", 
		"ID", "Nama Tagihan", "Nominal (Rp)", "Jatuh Tempo", "Status", "Kategori")
	fmt.Println(strings.Repeat("=", 85))

	for _, t := range data{
		statusText := t.Status
		if t.Status == "Lunas"{
			statusText = "Lunas"
		}else{
			statusText = "Belum Lunas"
		}
		fmt.Printf(" %-4d %-20s %-15.0f %-14s %-15s %-12s\n",
			t.ID, t.Nama, t.Nominal, t.JatuhTempo, statusText, t.Kategori)
	}
	fmt.Println(strings.Repeat("=", 85))
}


func tambahTagihan(data *daftarTagihan, nextID *int){
	fmt.Println("\nTAMBAH TAGIHAN BARU")
	fmt.Println(strings.Repeat("-", 35))

	//	----Masukkan nama tagihan
	nama := bacaInput("Nama Tagihan: ")
	if nama == ""{
		fmt.Println("Nama tagihan tidak boleh kosong!")
		return
	}

	var nominal float64

	//	----Masukkan nomimal tagihan
	fmt.Print("Nominal (Rp): ")
	_, err := fmt.Fscan(os.Stdin, &nominal)
	scanner.Scan()
	if err != nil || nominal <= 0{
		fmt.Println("Nomimal tidak valid!")
		return
	}

	//	----Masukkan waktu jatuh tempo
	jatuhTempo := bacaInput("Jatuh Tempo: ")
	if jatuhTempo == ""{
		fmt.Println("Waktu jatuh tempo tidak boleh kosong!")
		return
	}

	//	----Masukkan status tagihan
	fmt.Println("Status : 1. Belum lunas 2. Lunas")
	pilihStatus := bacaInput("Pilih (1/2): ")
	status := "Belum lunas"
	if pilihStatus == "2"{
		status = "Lunas"
	}

	//	----Masukkan kategori tagihan
	fmt.Println("Kategori: 1. Listrik  2. Air  3. Internet  4. Gas  5. Lainnya")
	pilihKategori := bacaInput("Pilih (1-5): ")
	kategoriMap := map[string]string{
		"1": "Listrik", "2": "Air", "3": "Internet", "4": "Gas", "5": "Lainnya",
	}
	kategori := kategoriMap[pilihKategori]
	if kategori == ""{
		kategori = "Lainnya"
	}

	tagihanBaru := Tagihan{
		ID: *nextID,
		Nama: nama,
		Nominal: nominal,
		JatuhTempo: jatuhTempo,
		Status: status,
		Kategori: kategori,
	}
	*data = append(*data, tagihanBaru)
	*nextID++

	fmt.Printf("\nTagihan '%s' berhasil ditambahkan!\n", nama)
}


func cariIndexByID(data daftarTagihan, id int)int{
	for i, t := range data {
		if t.ID == id{
			return i
		}
	}
	return -1
}


func ubahTagihan(data *daftarTagihan){
	fmt.Print("\nUBAH DATA TAGIHAN\n")
	fmt.Println(strings.Repeat("-", 35))

	//	----Menampilkan terlebih dahulu data
	tampilSemua(*data)

	var id int

	//	----Membaca ID yang akan diubah
	fmt.Print("\nMasukkan ID yang akan diubah: ")
	_, err := fmt.Fscan(os.Stdin, &id)
	scanner.Scan()
	if err != nil {
		fmt.Println("ID yang dimasukkan tidak valid!")
		return
	}

	//	----Validasi ID yang dicari
	idx := cariIndexByID(*data, id)
	if idx == -1{
		fmt.Printf("Tagihan dengan ID %d tidak ditemukan\n", id)
		return
	}

	t := &(*data)[idx]
	fmt.Printf("\nData lama -> Nama: %s | Nominal: %.0f | Jatuh Tempo: %s | Status: %s | Kategori: %s\n",
		t.Nama, t.Nominal, t.JatuhTempo, t.Status, t.Kategori)
	fmt.Println("(Tekan Enter untuk melewati field yang tidak ingin diubah)")
	fmt.Println(strings.Repeat("-", 40))

	//	----Masukkan nama tagihan yang baru
	namaBaru := bacaInput("Nama baru: ")
	if namaBaru != ""{
		t.Nama = namaBaru
	}

	//	----Masukkan nominal tagihan yang baru
	nominalStr := bacaInput("Nominal baru: ")
	if nominalStr != ""{
		var nominalBaru float64
		_, err := fmt.Sscanf(nominalStr, "%f", &nominalBaru)
		if err == nil && nominalBaru > 0{
			t.Nominal = nominalBaru
		}else{
			fmt.Println("Nominal tidak valid!")
		}
	}

	//  ----Masukkan waktu jatuh tempo yang baru
	jatuhTempoBaru := bacaInput("Jatuh tempo baru: ")
	if jatuhTempoBaru != ""{
		t.JatuhTempo = jatuhTempoBaru
	}

	//	----Masukkan status tagihan yang baru
	fmt.Println("Status : 1. Belum lunas 2. Lunas (Enter = skip)")
	pilihStatus := bacaInput("Pilih (1/2): ")
	if pilihStatus == "1"{
		t.Status = "Belum Lunas"
	}else{
		t.Status = "Lunas"
	}

	//	----Masukkan kategori tagihan yang baru
	fmt.Println("Pilih kategori baru: 1. Listrik  2. Air  3. Internet  4. Gas  5. Lainnya (Enter = skip)")
	pilihKategori := bacaInput("Pilih (1-5): ")
	kategoriMap := map[string]string{
		"1": "Listrik", "2": "Air", "3": "Internet", "4": "Gas", "5": "Lainnya",
	}
	if kat, ok := kategoriMap[pilihKategori]; ok{
		t.Kategori = kat
	}
	fmt.Printf("\n Tagihan ID '%d' berhasil diperbarui!\n", id)
}


func hapusTagihan(data *daftarTagihan){
	fmt.Println("\nHAPUS DATA TAGIHAN")
	fmt.Println(strings.Repeat("-", 35))

	//	----Menampilkan terlebih dahulu data
	tampilSemua(*data)

	var id int

	//	----Membaca ID yang akan dihapus
	fmt.Print("\nMasukkan ID tagihan yang akan dihapus: ")
	_, err := fmt.Fscan(os.Stdin, &id)
	scanner.Scan()
	if err != nil{
		fmt.Println("ID yang dimasukkan tidak valid!")
		return
	}

	//	----Mencari index tagihan berdasarkan ID
	idx := cariIndexByID(*data, id)
	if idx == -1{
		fmt.Printf("Tagihan dengan ID %d tidak ditemukan\n", id)
		return
	}

	//	----Konfirmasi penghapusan
	namaHapus := (*data)[idx].Nama
	konfirm := bacaInput(fmt.Sprintf("Yakin ingin menghapus '%s'? (y/n): ", namaHapus))
	if strings.ToLower(konfirm) != "y"{
		fmt.Println("Penghapusan data dibatalkan")
		return	
	}

	*data = append((*data)[:idx], (*data)[idx+1:]...)
	fmt.Printf("\nTagihan '%s' berhasil dihapus!\n", namaHapus)
}