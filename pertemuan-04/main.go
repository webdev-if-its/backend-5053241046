package main

import (
	"cmp"
	"errors"
	"fmt"
	"strings"
	"time"
)

// TODO: lihat SOAL.md untuk kontrak lengkap tiap fungsi/method di bawah.
// Ganti setiap "panic" dengan implementasi yang benar. Tambahkan import
// (mis. "strings") sendiri kalau memang dibutuhkan.

// ErrTugasTidakDitemukan dikembalikan ketika ID tugas tidak ada.
var ErrTugasTidakDitemukan = errors.New("tugas tidak ditemukan")

// ErrInputKosong dikembalikan ketika judul kosong (atau hanya spasi).
var ErrInputKosong = errors.New("input tidak boleh kosong")

// ErrDaftarKosong dikembalikan oleh Max ketika slice-nya kosong.
var ErrDaftarKosong = errors.New("daftar kosong")

// Timestamps dipakai lewat embedding (composition) di Task.
type Timestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Touch menyetel UpdatedAt ke waktu sekarang. (Level 9)
func (ts *Timestamps) Touch() {
	ts.UpdatedAt = time.Now()
}

// Task merepresentasikan satu tugas.
type Task struct {
	ID      int
	Judul   string
	Selesai bool
	Timestamps
}

// NewTask membuat Task baru dari judul. (Level 1)
func NewTask(judul string) (Task, error) {
	judul = strings.TrimSpace(judul)
	if judul == "" {
		return Task{}, ErrInputKosong
	}
	now := time.Now()
	t := Task{
		Judul: judul,
		Timestamps: Timestamps{
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
	return t, nil
}

// MarkDone menandai tugas selesai. (Level 2)
func (t *Task) MarkDone() {
	t.Selesai = true
	t.Touch()
}

// Rename mengganti judul tugas. (Level 3)
func (t *Task) Rename(judul string) error {
	judul = strings.TrimSpace(judul)
	if judul == "" {
		return ErrInputKosong
	}
	t.Judul = judul
	t.Touch()
	return nil
}

// String membuat Task memenuhi fmt.Stringer. (Level 10)
func (t Task) String() string {
	status := "belum selesai"
	if t.Selesai {
		status = "selesai"
	}
	return fmt.Sprintf("[%d] %s (%s)", t.ID, t.Judul, status)
}

// TaskStore adalah kontrak penyimpanan tugas. JANGAN diubah -- yang kalian
// tulis adalah implementasinya (MemoryStore).
type TaskStore interface {
	Add(t Task) (Task, error)
	Get(id int) (Task, error)
	List() []Task
	Delete(id int) error
}

// MemoryStore menyimpan tugas di memori.
type MemoryStore struct {
	// TODO: tambahkan field yang kalian butuhkan (mis. slice tugas dan penghitung ID)
	tugas []Task
	nextIDCounter int
}

// Pastikan *MemoryStore memenuhi TaskStore -- kalau tidak, kode gagal
// dikompilasi di sini, bukan baru ketahuan saat dijalankan.
var _ TaskStore = (*MemoryStore)(nil)

// NewMemoryStore membuat store kosong.
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		tugas: []Task{},
		nextIDCounter: 1,
	}
}

// Add menyimpan tugas dan memberinya ID baru. (Level 4)
func (m *MemoryStore) Add(t Task) (Task, error) {
	if strings.TrimSpace(t.Judul) == "" {
		return Task{}, ErrInputKosong
	}

	for {
		used := false
		for _, existing := range m.tugas {
			if existing.ID == m.nextIDCounter {
				used = true
				m.nextIDCounter++
				break
			}
		}
		if !used {
			break
		}
	}
	t.ID = m.nextIDCounter
	m.nextIDCounter++
	m.tugas = append(m.tugas, t)
	return t, nil
}

func (m *MemoryStore) Get(id int) (Task, error) {
	for _, t := range m.tugas {
		if t.ID == id {
			// Task dikembalikan berdasarkan value, sehingga perubahan pada hasil
			// pemanggilan tidak mengubah data yang tersimpan.
			return t, nil
		}
	}
	return Task{}, ErrTugasTidakDitemukan
}

// List mengembalikan seluruh tugas. (Level 5)
func (m *MemoryStore) List() []Task {
	var hasil = []Task{}
	for _, t := range m.tugas {
		hasil = append(hasil, t)
	}
	return hasil
}

// Delete menghapus tugas menurut ID. (Level 6)
func (m *MemoryStore) Delete(id int) error {
	for i, t := range m.tugas {
		if t.ID == id {
			m.tugas = append(m.tugas[:i], m.tugas[i+1:]...)
			return nil
		}
	}
	return ErrTugasTidakDitemukan
}

// Filter mengembalikan elemen xs yang lolos pred. (Level 7)
func Filter[T any](xs []T, pred func(T) bool) []T {
	hasil := make([]T, 0)
	for _, x := range xs {
		if pred(x) {
			hasil = append(hasil, x)
		}
	}
	return hasil
}

// Map mengubah tiap elemen xs dengan f. (Level 7)
func Map[T, U any](xs []T, f func(T) U) []U {
	hasil := make([]U, 0, len(xs))
	for _, x := range xs {
		hasil = append(hasil, f(x))
	}
	return hasil
}

// Contains melaporkan apakah v ada di xs. (Level 8)
func Contains[T comparable](xs []T, v T) bool {
	for _, x := range xs {
		if x == v {
			return true
		}
	}
	return false
}

// Max mengembalikan elemen terbesar di xs. (Level 8)
func Max[T cmp.Ordered](xs []T) (T, error) {
	var zero T
	if len(xs) == 0 {
		return zero, ErrDaftarKosong
	}
	max := xs[0]
	for _, x := range xs[1:] {
		if x > max {
			max = x
		}
	}
	return max, nil
}

// Gabung menyambung String() tiap elemen dengan pemisah sep. (Level 10)
func Gabung[T fmt.Stringer](xs []T, sep string) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = x.String()
	}
	return strings.Join(parts, sep)
}

func main() {
	fmt.Println("Task Manager v1 - pertemuan 4")

	var store TaskStore = NewMemoryStore()
	for _, judul := range []string{"Belajar struct", "Belajar interface"} {
		t, err := NewTask(judul)
		if err != nil {
			fmt.Println("gagal membuat tugas:", err)
			continue
		}
		if _, err := store.Add(t); err != nil {
			fmt.Println("gagal menyimpan tugas:", err)
		}
	}
	for _, t := range store.List() {
		fmt.Println(t)
	}
}
