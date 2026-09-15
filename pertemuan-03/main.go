package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var ErrTugasTidakDitemukan = errors.New("tugas tidak ditemukan")
var ErrInputKosong = errors.New("input tidak boleh kosong")

type Task struct {
	ID      int
	Judul   string
	Selesai bool
}

type TokoTugas struct {
	Daftar []Task
	NextID int
	Log    []string
}

// TODO (LEVEL 1 && 6)
func TambahTugas(toko *TokoTugas, judul string) (Task, error) {
	if strings.TrimSpace(judul) == "" {
		return Task{}, ErrInputKosong
	}

	if toko.NextID == 0 {
		toko.NextID = 1
	}

	tugasBaru := Task{
		ID:      toko.NextID,
		Judul:   judul,
		Selesai: false,
	}

	toko.Daftar = append(toko.Daftar, tugasBaru)
	toko.NextID++

	return tugasBaru, nil
}

// TODO (LEVEL 2 && 4)
func LihatTugas(toko *TokoTugas, id int) (Task, error) {
	for _, tugas := range toko.Daftar {
		if tugas.ID == id {
			return tugas, nil
		}
	}
	return Task{}, ErrTugasTidakDitemukan
}

// TODO (LEVEL 3 && 5)
func HapusTugas(toko *TokoTugas, id int) error {
	for i, tugas := range toko.Daftar {
		if tugas.ID == id {
			toko.Daftar = append(toko.Daftar[:i], toko.Daftar[i+1:]...)
			return nil
		}
	}
	return ErrTugasTidakDitemukan
}

// TODO (LEVEL 7)
func HapusTugasTercatat(toko *TokoTugas, id int) (err error) {
	defer func() {
		if err == nil {
			toko.Log = append(toko.Log, fmt.Sprintf("hapus id=%d: berhasil", id))
		} else {
			toko.Log = append(toko.Log, fmt.Sprintf("hapus id=%d: gagal (%v)", id, err))
		}
	}()

	err = HapusTugas(toko, id) 
	return err
}

// TODO (LEVEL 8)
func AmankanPanggilan(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("terjadi panic: %v", r)
		}
	}()
	err = fn()
	return err
}

// TODO (LEVEL 9)
func AmankanHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := AmankanPanggilan(func() error {
			next(w, r)
			return nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// TODO (LEVEL 10 - Part 1)
func RekapStatus(toko *TokoTugas) map[string]int {
	status := map[string]int{
		"selesai":       0,
		"belum selesai": 0,
	}

	for _, tugas := range toko.Daftar {
		if tugas.Selesai {
			status["selesai"]++
		} else {
			status["belum selesai"]++
		}
	}
	return status
}

// TODO (LEVEL 10 - Part 2)
func BuatHandlerTugas(toko *TokoTugas) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, t := range toko.Daftar {
			statusStr := "belum selesai"
			if t.Selesai {
				statusStr = "selesai"
			}
			fmt.Fprintf(w, "%d. %s [%s]\n", t.ID, t.Judul, statusStr)
		}

		rekap := RekapStatus(toko)
		fmt.Fprintf(w, "\nRingkasan: %d selesai, %d belum selesai\n", rekap["selesai"], rekap["belum selesai"])
	}
}

func main() {
	fmt.Println("Task Manager v0 - pertemuan 3")

	toko := &TokoTugas{}
	http.HandleFunc("/tasks", AmankanHandler(BuatHandlerTugas(toko)))

	fmt.Println("Server jalan di :8080")
	http.ListenAndServe(":8080", nil)
}