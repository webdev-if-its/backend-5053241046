## Identitas
Nama: Berlian Yafi Kania
NRP: 5053241046
Kelas: Backend M (RPL)

## Commit vs Push
git commit menyimpan perubahan file ke dalam riwayat repository lokal di komputer kita, sedangkan git push mengunggah commit dari repository lokal tersebut ke repository remote seperti GitHub/GitLab. Jika seseorang melakukan commit tetapi lupa melakukan push, perubahan kode miliknya tidak akan terlihat oleh anggota tim lain. Akibatnya, rekan tim akan bekerja dengan kode versi lama dan berpotensi mengalami bentrok (conflict) saat mengirimkan pembaruan di kemudian hari.

## Reproducibility
Perbedaan versi Go di antara anggota tim dapat memicu masalah jika kode memanfaatkan fitur sintaksis baru, pustaka standar baru, atau perubahan perilaku pengompilasi. Kapan itu jadi masalah nyata: ketika anggota tim menggunakan fitur baru Go (seperti perbaikan *range over int* di Go 1.22) sementara tim lain memakai Go versi lama, yang menyebabkan kegagalan kompilasi. Kapan tidak jadi masalah: saat program hanya menggunakan fitur dasar bahasa Go yang stabil dan tidak mengalami perubahan antar versi standar.

## Catatan Merge Conflict
Baris yang mengalami bentrok adalah baris nilai kembalian (return) pada fungsi `CetakInfo` di file `main.go`. Konflik terjadi karena cabang utama (*main*) dan cabang fitur (*fitur-sapaan*) sama-sama mengubah struktur string yang dikembalikan secara bersamaan. Keputusan hasil akhir dibuat dengan menggabungkan kedua perubahan tersebut secara manual, yaitu mempertahankan format informasi Nama, NRP, dan versi Go, sekaligus menyertakan hasil pemanggilan fungsi `Sapa(nama)`.

## Kenapa .gitignore Penting
File seperti biner hasil build (`bin/`, `*.exe`) atau konfigurasi editor (`.vscode/`, `.idea/`) berukuran dinamis dan spesifik untuk setiap lingkungan komputer pengembang. Jika file-file ini ter-commit ke repository, file biner dapat mengotori riwayat kompilasi dan memicu konflik permanen saat anggota tim lain mengompilasi ulang dengan pengaturan IDE atau OS yang berbeda.

## Refleksi
Bagian yang paling membingungkan awal pertemuan ini adalah menyelesaikan konflik merge git saat baris kode yang sama diubah di dua cabang berbeda. Akhirnya hal ini dipahami setelah memeriksa langsung penanda konflik (`<<<<<<<`, `=======`, `>>>>>>>`), memilih baris kode yang benar secara manual, lalu melakukan commit resolusi merge.

<!-- 
# backend-nrp

Repo tugas mata kuliah **Pengembangan Backend Dasar**, dibuat dari template [`webdev-if-its/backend-template`](https://github.com/webdev-if-its/backend-template). Ganti judul di atas jadi nama repo kalian sendiri (`backend-nrp`, contoh: `backend-5025201012`).

## Aturan Umum

- Tugas tiap pertemuan disimpan di folder `pertemuan-XX/` pada repo ini.
- Commit message wajib menyebut level yang dicapai: `pertemuan-XX: level N selesai`.
- Deadline push: sebelum pertemuan berikutnya dimulai.
- Semua level dicek otomatis lewat `go test` — baca `pertemuan-XX/SOAL.md` tiap minggu untuk detail levelnya.

## Mengambil Pertemuan Baru Tiap Minggu

Repo ini **tidak otomatis sinkron** dengan template dosen. Begitu ada pertemuan baru, jalankan (ganti `pertemuan-02` sesuai minggu berjalan):

```bash
git fetch https://github.com/webdev-if-its/backend-template.git main
git checkout FETCH_HEAD -- pertemuan-02
```

Perintah ini **aman dijalankan kapan pun** — tidak akan menimpa folder pertemuan lain yang sudah kalian kerjakan, karena hanya mengambil folder yang disebutkan. Setelah itu, commit folder barunya seperti biasa.

Kalau dosen memperbaiki sesuatu di pertemuan yang sudah dirilis (mis. ada bug di test), biasanya cukup ambil ulang file yang diperbaiki saja, bukan seluruh folder — akan diumumkan file mana yang berubah.

---

Bagian di bawah ini **isi bertahap** sesuai level yang sedang kalian kerjakan (lihat `pertemuan-01/SOAL.md`) — heading-nya dicek otomatis, jangan diganti namanya.

## Identitas
- Nama: (tulis di sini)
- NRP: (tulis di sini)
- Kelas: (tulis di sini)

## Commit vs Push
(tulis di sini)

## Reproducibility
(tulis di sini)

## Catatan Merge Conflict
(tulis di sini)

## Kenapa .gitignore Penting
(tulis di sini)

## Refleksi
(tulis di sini) -->
