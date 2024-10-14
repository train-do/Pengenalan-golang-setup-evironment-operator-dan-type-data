package main

import "fmt"

func main() {
	var hargaJual = 150000.0;
	var hargaBeli = 100000.0;
	var biayaOperasional = 1000.0;
	var diskon = 15.0;
	var jumlahTerjual = 100;
	hargaJualDiskon := hargaJual-hargaJual*(diskon/100);
	totalPendapatan := hargaJualDiskon*float64(jumlahTerjual);
	totalBiaya := hargaBeli*float64(jumlahTerjual);
	totalKeuntungan := totalPendapatan-(totalBiaya+(biayaOperasional*float64(jumlahTerjual)));
	fmt.Println("Harga Jual Setelah Diskon: ", hargaJualDiskon);
	fmt.Println("Total Pendapatan : ", totalPendapatan);
	fmt.Println("Total Biaya : ", totalBiaya);
	fmt.Println("Total Keuntungan : ", totalKeuntungan);
}