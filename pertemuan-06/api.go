package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// TODO: lihat SOAL.md untuk kontrak lengkap. Ganti setiap "panic" dengan
// implementasi yang benar; tambahkan import (encoding/json, errors, log,
// strconv, ...) sendiri kalau dibutuhkan.

// taskInput adalah bentuk JSON yang boleh dikirim klien untuk POST/PUT.
// Sengaja TIDAK memuat ID -- ID bukan urusan klien.
type taskInput struct {
	Judul   string `json:"judul"`
	Selesai bool   `json:"selesai"`
}

// errorBody adalah bentuk JSON semua respons error: {"error":"pesan"}.
type errorBody struct {
	Error string `json:"error"`
}

// API menyatukan routing dan handler HTTP di atas sebuah TaskStore.
type API struct {
	store TaskStore
}

// NewAPI membuat API di atas store (interface, bukan MemoryStore).
func NewAPI(store TaskStore) *API {
	return &API{store: store}
}

// Routes mendaftarkan semua endpoint memakai pola Go 1.22 ("METHOD /path").
func (a *API) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tasks", a.listTasks)
	mux.HandleFunc("POST /tasks", a.createTask)
	mux.HandleFunc("GET /tasks/{id}", a.getTask)
	mux.HandleFunc("PUT /tasks/{id}", a.updateTask)
	mux.HandleFunc("DELETE /tasks/{id}", a.deleteTask)
	return mux
}

// writeJSON menulis v sebagai JSON dengan status yang diminta.
// Urutan WAJIB: Header().Set(Content-Type) -> WriteHeader(status) -> Encode.
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// writeError menerjemahkan error aplikasi menjadi status + body errorBody.
func writeError(w http.ResponseWriter, err error) {
	if errors.Is(err, ErrTugasTidakDitemukan) {
		writeJSON(w, http.StatusNotFound, errorBody{Error: "tugas tidak ditemukan"})
		return
	}
	if errors.Is(err, ErrInputKosong) {
		writeJSON(w, http.StatusBadRequest, errorBody{Error: "input tidak boleh kosong"})
		return
	}

	// Error penanganan sintaks/tipe khusus kustom
	errStr := err.Error()
	if errStr == "id harus berupa angka" || errStr == "body bukan JSON yang valid" {
		writeJSON(w, http.StatusBadRequest, errorBody{Error: errStr})
		return
	}

	// Internal server error (Level 8: 500 generik, tanpa membocorkan detail sensitif)
	log.Printf("[SERVER ERROR] %v", err)
	writeJSON(w, http.StatusInternalServerError, errorBody{Error: "terjadi kesalahan internal"})
}

// GET /tasks
func (a *API) listTasks(w http.ResponseWriter, r *http.Request) {
	tasks := a.store.List()
	if tasks == nil {
		tasks = []Task{}
	}
	writeJSON(w, http.StatusOK, tasks)
}

// POST /tasks
func (a *API) createTask(w http.ResponseWriter, r *http.Request) {
	var in taskInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, fmt.Errorf("body bukan JSON yang valid"))
		return
	}

	if strings.TrimSpace(in.Judul) == "" {
		writeError(w, ErrInputKosong)
		return
	}

	t, err := a.store.Add(Task{
		Judul:   in.Judul,
		Selesai: in.Selesai,
	})
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, t)
}

// GET /tasks/{id}
func (a *API) getTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, fmt.Errorf("id harus berupa angka"))
		return
	}

	t, err := a.store.Get(id)
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

// PUT /tasks/{id}
func (a *API) updateTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, fmt.Errorf("id harus berupa angka"))
		return
	}

	var in taskInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		writeError(w, fmt.Errorf("body bukan JSON yang valid"))
		return
	}

	if strings.TrimSpace(in.Judul) == "" {
		writeError(w, ErrInputKosong)
		return
	}

	t, err := a.store.Update(Task{
		ID:      id,
		Judul:   in.Judul,
		Selesai: in.Selesai,
	})
	if err != nil {
		writeError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

// DELETE /tasks/{id}
func (a *API) deleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, fmt.Errorf("id harus berupa angka"))
		return
	}

	if err := a.store.Delete(id); err != nil {
		writeError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}