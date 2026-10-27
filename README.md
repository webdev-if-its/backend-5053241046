# backend-5053241046

Repo tugas mata kuliah Pengembangan Backend Dasar.

## Identitas
- Nama: Berlian Yafi Kania
- NRP: 5053241046
- Kelas: Backend M (RPL)

## Commit vs Push
git commit menyimpan perubahan file ke dalam riwayat repository lokal di komputer kita, sedangkan git push mengunggah commit dari repository lokal tersebut ke repository remote seperti GitHub/GitLab. Jika seseorang melakukan commit tetapi lupa melakukan push, perubahan kode miliknya tidak akan terlihat oleh anggota tim lain. Akibatnya, rekan tim akan bekerja dengan kode versi lama dan berpotensi mengalami bentrok (conflict) saat mengirimkan pembaruan di kemudian hari.

## Reproducibility
Perbedaan versi Go di antara anggota tim dapat memicu masalah jika kode memanfaatkan fitur sintaksis baru, pustaka standar baru, atau perubahan perilaku pengompilasi. Kapan itu jadi masalah nyata: ketika anggota tim menggunakan fitur baru Go (seperti perbaikan range over int di Go 1.22) sementara tim lain memakai Go versi lama, yang menyebabkan kegagalan kompilasi. Kapan tidak jadi masalah: saat program hanya menggunakan fitur dasar bahasa Go yang stabil dan tidak mengalami perubahan antar versi standar.

## Catatan Merge Conflict
Baris yang mengalami bentrok adalah baris nilai kembalian (return) pada fungsi CetakInfo di file main.go. Konflik terjadi karena cabang utama (main) dan cabang fitur (fitur-sapaan) sama-sama mengubah struktur string yang dikembalikan secara bersamaan. Keputusan hasil akhir dibuat dengan menggabungkan kedua perubahan tersebut secara manual, yaitu mempertahankan format informasi Nama, NRP, dan versi Go, sekaligus menyertakan hasil pemanggilan fungsi Sapa(nama).

## Kenapa .gitignore Penting
File seperti biner hasil build (bin/, *.exe) atau konfigurasi editor (.vscode/, .idea/) berukuran dinamis dan spesifik untuk setiap lingkungan komputer pengembang. Jika file-file ini ter-commit ke repository, file biner dapat mengotori riwayat kompilasi dan memicu konflik permanen saat anggota tim lain mengompilasi ulang dengan pengaturan IDE atau OS yang berbeda.

## Refleksi
Bagian yang paling membingungkan awal pertemuan ini adalah menyelesaikan konflik merge git saat baris kode yang sama diubah di dua cabang berbeda. Akhirnya hal ini dipahami setelah memeriksa langsung penanda konflik (<<<<<<<, =======, >>>>>>>), memilih baris kode yang benar secara manual, lalu melakukan commit resolusi merge.

## Analisis Pertemuan 4

1. Pointer vs Value Receiver
(a) Value Receiver (t T): Method bekerja pada salinan (copy) dari data asli. Perubahan di dalam method tidak mengubah data asli.
(b) Pointer Receiver (t *T): Method bekerja menggunakan pointer (alamat memori) data asli. Perubahan di dalam method akan mengubah data asli, dan menghindari proses copying data yang besar.

2. Method Set (Himpunan Method)
(a) Untuk tipe Value T: Method set hanya berisi method yang menggunakan Value Receiver.
(b) Untuk tipe Pointer *T: Method set berisi semua method, baik yang menggunakan Value maupun Pointer Receiver.

3. Hubungannya dengan Kata Interface
Sebuah tipe data dikatakan mengimplementasikan sebuah interface jika tipe data tersebut memiliki semua method yang diminta oleh interface tersebut. Karena interface hanya memegang salinan nilai, Go tidak mengizinkan interface tersebut memanggil method dengan Pointer Receiver.

## Analisis Pertemuan 5

1. Operasi hasil[i] = ... pada UnduhSemua aman tanpa Mutex karena setiap goroutine menulis ke elemen slice dengan lokasi indeks memory yang berbeda secara independen. Tidak ada dua goroutine yang mengakses elemen indeks yang sama. Sebaliknya, selesai++ pada Progres mengalami race condition jika dipanggil oleh banyak goroutine secara bersamaan tanpa Mutex atau lock. Hal ini terjadi karena operasi selesai++ bukanlah satu langkah atomic, melainkan terdiri dari tiga langkah: membaca nilai saat ini, menambah nilai sebesar 1, dan menyimpan kembali nilai baru. Tanpa penguncian mutex, dua goroutine dapat membaca nilai lama yang sama secara bersamaan sehingga salah satu pembaruan nilai akan hilang.

2. Jika wg.Add(1) ditulis di dalam goroutine (bukan sebelum perintah go), akan terjadi race condition antara goroutine utama dan goroutine anak. Goroutine utama bisa saja mengeksekusi wg.Wait() lebih cepat sebelum goroutine anak sempat berjalan dan memanggil wg.Add(1). Akibatnya, wg.Wait() berisiko langsung lewat atau selesai prematur tanpa menunggu unduhan yang sebenarnya masih diproses oleh goroutine anak.

## Dokumentasi API Task Manager

Daftar endpoint REST API:
- GET /tasks (status 200)
- POST /tasks (status 201)
- GET /tasks/{id} (status 200)
- PUT /tasks/{id} (status 200)
- DELETE /tasks/{id} (status 204)

Kode status response yang digunakan: 200, 201, 204, 400, 404, 500.

Contoh body request JSON:
{"judul": "Tugas Baru", "selesai": false}

Contoh respons error JSON:
{"error": "tugas tidak ditemukan"}

Analisis dan Penjelasan:

1. Kenapa ID bukan angka mengembalikan 400 Bad Request dan bukan 404 Not Found?
Permintaan seperti GET /tasks/abc merupakan kesalahan format sintaks dari masukan klien (malformed request). Parameter ID wajib berupa integer angka, sehingga kegagalan parsing ini dikategorikan sebagai kesalahan permintaan klien (400 Bad Request). Kode status 404 Not Found khusus digunakan jika format URL dan sintaks ID sudah valid berupa angka, namun data tugasnya memang tidak ditemukan di store.

2. Kenapa detail error 500 tidak ditampilkan ke klien?
Pesan error mentah dari kesalahan internal atau database berisiko membocorkan informasi sensitif seperti skema tabel, kredensial, maupun IP server. Oleh karena itu, detail error asli hanya dicatat pada log internal server (menggunakan log.Printf), sedangkan respons HTTP yang dikirim ke klien disembunyikan menggunakan pesan generik {"error": "terjadi kesalahan internal"} demi menjaga keamanan sistem internal.
