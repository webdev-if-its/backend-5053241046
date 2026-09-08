package main

import "fmt"

// TODO(Level 1): 

func HitungSubtotal(qty int, hargaSatuan float64) float64 {

var HargaSatuan float64
var JumlahBarang float64

HargaSatuan = hargaSatuan
JumlahBarang = float64(qty)

return HargaSatuan * JumlahBarang 
}

// TODO(Level 2):

func HitungTotalPesanan(qty []int, hargaSatuan []float64) float64 {

	if len(qty) == len(hargaSatuan) {
		var total float64
		for i := 0; i < len(qty); i++ {
			total += HitungSubtotal(qty[i], hargaSatuan[i])
		}
		return total
	}
	return 0

}

// TODO(Level 3):

func TerapkanPajak(total float64, tarifPajak float64) float64 {
	// return total + (total * tarifPajak)
}

// TODO(Level 4):

func HitungDiskon(total float64) float64 {
	panic("belum diimplementasikan")
}

// TODO(Level 5):

func TotalSetelahDiskon(qty []int, hargaSatuan []float64, tarifPajak float64) float64 {
	panic("belum diimplementasikan")
}
// TODO(Level 6):

func ValidasiPesanan(qty []int, hargaSatuan []float64) (bool, string) {
	panic("belum diimplementasikan")
}

// TODO(Level 7):

func TentukanStatus(total float64) string {
	panic("belum diimplementasikan")
}

// TODO(Level 8):

func RingkasanPesanan(qty []int, hargaSatuan []float64, tarifPajak float64) string {
	panic("belum diimplementasikan")
}

// TODO(Level 9): signature ini SUDAH benar (cari tahu sendiri kenapa
// bentuknya begini - lihat SOAL.md) - tinggal implementasikan isinya.
func Total(harga ...float64) float64 {
	panic("belum diimplementasikan")
}

// TODO(Level 10, bonus): signature ini SUDAH benar (cari tahu sendiri
// kenapa ada dua nilai balik - lihat SOAL.md) - tinggal implementasikan isinya.
func HitungOngkosKirim(beratKg float64, jarakKm float64) (float64, error) {
	panic("belum diimplementasikan")
}

func main() {
	fmt.Println("Sales Order Processor - pertemuan 2")
}
