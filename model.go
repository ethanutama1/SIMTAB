package main

type Tagihan struct{
	ID int
	Nama string
	Nominal float64
	JatuhTempo string
	Status string
	Kategori string
}

type daftarTagihan []Tagihan