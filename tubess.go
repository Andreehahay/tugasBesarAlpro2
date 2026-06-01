package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// Perangkat menyimpan data satu perangkat elektronik (Andre)
type Perangkat struct {
	ID        int
	Nama      string
	Watt      float64
	DurasiJam float64
	Ruangan   string
}

// ==============================
// KONSTANTA & VARIABEL GLOBAL
// ==============================

const MAX_PERANGKAT = 100 // kapasitas maksimum array statis (Erlan)

// Array global utama penyimpan data perangkat (Andre)
var daftarPerangkat [MAX_PERANGKAT]Perangkat
var jumlahPerangkat int = 0
var nextID int = 1

// ==============================
// SUBPROGRAM UTILITAS INPUT
// ==============================

// bacaString membaca satu baris input dari pengguna (Naufal)
func bacaString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	teks, _ := reader.ReadString('\n')
	return strings.TrimSpace(teks)
}

// bacaFloat membaca input angka desimal dari pengguna (Erlan)
func bacaFloat(prompt string) float64 {
	var angka float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanf("%f\n", &angka)
		if err == nil && angka >= 0 {
			return angka
		}
		fmt.Println("  [!] Input tidak valid. Masukkan angka positif.")
		// membersihkan buffer input yang tersisa (Naufal)
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
	}
}

// bacaInt membaca input bilangan bulat dari pengguna (Naufal)
func bacaInt(prompt string) int {
	var angka int
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanf("%d\n", &angka)
		if err == nil {
			return angka
		}
		fmt.Println("  [!] Input tidak valid. Masukkan bilangan bulat.")
		// membersihkan sisa input yang tidak terbaca (Erlan)
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')
	}
}

// ==============================
// SUBPROGRAM TAMPILAN
// ==============================

// cetakHeader menampilkan header aplikasi PowerLog (Andre)
func cetakHeader() {
	fmt.Println()
	fmt.Println("+++ PowerLog +++")
	fmt.Println("Aplikasi Pencatatan Konsumsi Listrik Perangkat")
	fmt.Println(strings.Repeat("-", 50))
}

// cetakGaris mencetak garis pemisah antar bagian tampilan (Erlan)
func cetakGaris() {
	fmt.Println(strings.Repeat("-", 50))
}

// cetakMenuUtama menampilkan daftar menu utama kepada pengguna (Andre)
func cetakMenuUtama() {
	fmt.Println()
	fmt.Println("  MENU UTAMA")
	cetakGaris()
	fmt.Println("  1. Tambah Perangkat")
	fmt.Println("  2. Lihat Semua Perangkat")
	fmt.Println("  3. Cari Perangkat")
	fmt.Println("  4. Ubah Data Perangkat")
	fmt.Println("  5. Hapus Perangkat")
	fmt.Println("  6. Urutkan Perangkat")
	fmt.Println("  7. Statistik Penggunaan Daya")
	fmt.Println("  0. Keluar")
	cetakGaris()
}

// cetakPerangkat menampilkan satu baris data perangkat ke layar (Naufal)
func cetakPerangkat(p Perangkat) {
	konsumsiHarian := hitungKonsumsiHarian(p.Watt, p.DurasiJam)
	fmt.Printf("  [%d] %-20s | %7.1f W | %5.1f jam/hari | %-12s | %.3f kWh/hari\n",
		p.ID, p.Nama, p.Watt, p.DurasiJam, p.Ruangan, konsumsiHarian)
}

// cetakHeaderTabel menampilkan baris judul kolom tabel perangkat (Erlan)
func cetakHeaderTabel() {
	cetakGaris()
	fmt.Printf("  %-4s %-20s | %7s | %14s | %-12s | %s\n",
		"ID", "Nama Perangkat", "Daya(W)", "Durasi", "Ruangan", "Konsumsi/Hari")
	cetakGaris()
}

// ==============================
// SUBPROGRAM KALKULASI
// ==============================

// hitungKonsumsiHarian menghitung konsumsi energi harian dalam kWh (Andre)
// Parameter: watt = daya perangkat, durasiJam = lama pemakaian per hari
// Return: konsumsi dalam kWh
func hitungKonsumsiHarian(watt float64, durasiJam float64) float64 {
	return (watt * durasiJam) / 1000.0
}

// hitungTotalKonsumsi menjumlahkan seluruh konsumsi harian semua perangkat (Naufal)
// Return: total kWh per hari dari semua perangkat yang tercatat
func hitungTotalKonsumsi() float64 {
	total := 0.0
	i := 0
	for i < jumlahPerangkat {
		total += hitungKonsumsiHarian(daftarPerangkat[i].Watt, daftarPerangkat[i].DurasiJam)
		i++
	}
	return total
}

// ==============================
// SUBPROGRAM CRUD
// ==============================

// tambahPerangkat menambahkan satu data perangkat baru ke array global (Andre)
func tambahPerangkat() {
	fmt.Println()
	fmt.Println("  [ TAMBAH PERANGKAT BARU ]")
	cetakGaris()

	// memastikan array belum penuh sebelum menambah data baru (Erlan)
	if jumlahPerangkat >= MAX_PERANGKAT {
		fmt.Println("  [!] Daftar perangkat sudah penuh (maks. 50 perangkat).")
		return
	}

	nama := bacaString("  Nama Perangkat  : ")
	if nama == "" {
		fmt.Println("  [!] Nama perangkat tidak boleh kosong.")
		return
	}
	watt := bacaFloat("  Daya (Watt)      : ")
	durasi := bacaFloat("  Durasi (jam/hari): ")
	ruangan := bacaString("  Lokasi Ruangan   : ")
	if ruangan == "" {
		fmt.Println("  [!] Lokasi ruangan tidak boleh kosong.")
		return
	}

	// membentuk struct perangkat baru dan menyimpannya ke array statis (Naufal)
	perangkatBaru := Perangkat{
		ID:        nextID,
		Nama:      nama,
		Watt:      watt,
		DurasiJam: durasi,
		Ruangan:   ruangan,
	}

	daftarPerangkat[jumlahPerangkat] = perangkatBaru
	jumlahPerangkat++
	nextID++

	fmt.Printf("  [✓] Perangkat \"%s\" berhasil ditambahkan.\n", nama)
}

// lihatSemuaPerangkat menampilkan seluruh isi array perangkat ke layar (Naufal)
func lihatSemuaPerangkat() {
	fmt.Println()
	fmt.Println("  [ DAFTAR SEMUA PERANGKAT ]")

	// memeriksa apakah ada data sebelum menampilkan tabel (Erlan)
	if jumlahPerangkat == 0 {
		fmt.Println("  [!] Belum ada perangkat yang tercatat.")
		return
	}

	cetakHeaderTabel()
	i := 0
	// menampilkan tiap perangkat satu per satu menggunakan indeks (Naufal)
	for i < jumlahPerangkat {
		cetakPerangkat(daftarPerangkat[i])
		i++
	}
	cetakGaris()
	fmt.Printf("  Total: %d perangkat | Konsumsi harian: %.3f kWh\n",
		jumlahPerangkat, hitungTotalKonsumsi())
}

// cariIndeksByID mencari indeks array berdasarkan ID perangkat (Andre)
// Menggunakan Sequential Search — memeriksa tiap elemen satu per satu
// Parameter: id = ID yang dicari
// Return: indeks array jika ditemukan, -1 jika tidak ditemukan
func cariIndeksByID(id int) int {
	indeks := -1
	i := 0
	ketemu := false
	// menelusuri array dari awal hingga ID ditemukan atau habis (Erlan)
	for i < jumlahPerangkat && !ketemu {
		if daftarPerangkat[i].ID == id {
			indeks = i
			ketemu = true
		}
		i++
	}
	return indeks
}

// ubahPerangkat memperbarui data perangkat yang sudah ada berdasarkan ID (Andre)
func ubahPerangkat() {
	fmt.Println()
	fmt.Println("  [ UBAH DATA PERANGKAT ]")
	cetakGaris()

	if jumlahPerangkat == 0 {
		fmt.Println("  [!] Belum ada perangkat yang bisa diubah.")
		return
	}

	id := bacaInt("  Masukkan ID perangkat yang akan diubah: ")
	// mencari posisi perangkat di array menggunakan sequential search (Naufal)
	indeks := cariIndeksByID(id)

	if indeks == -1 {
		fmt.Printf("  [!] Perangkat dengan ID %d tidak ditemukan.\n", id)
		return
	}

	fmt.Println()
	fmt.Println("  Data saat ini:")
	cetakHeaderTabel()
	cetakPerangkat(daftarPerangkat[indeks])
	cetakGaris()
	fmt.Println("  (Kosongkan input untuk mempertahankan nilai lama)")
	fmt.Println()

	// memperbarui nama jika pengguna mengisi input baru (Erlan)
	namaBaru := bacaString(fmt.Sprintf("  Nama Baru [%s]: ", daftarPerangkat[indeks].Nama))
	if namaBaru != "" {
		daftarPerangkat[indeks].Nama = namaBaru
	}

	// memperbarui watt jika nilai baru berbeda dari nilai saat ini (Naufal)
	wattBaru := bacaFloat(fmt.Sprintf("  Watt Baru [%.1f]: ", daftarPerangkat[indeks].Watt))
	if wattBaru != daftarPerangkat[indeks].Watt {
		daftarPerangkat[indeks].Watt = wattBaru
	}

	durasiBaru := bacaFloat(fmt.Sprintf("  Durasi Baru [%.1f jam]: ", daftarPerangkat[indeks].DurasiJam))
	if durasiBaru != daftarPerangkat[indeks].DurasiJam {
		daftarPerangkat[indeks].DurasiJam = durasiBaru
	}

	// memperbarui ruangan jika input baru tidak kosong (Erlan)
	ruanganBaru := bacaString(fmt.Sprintf("  Ruangan Baru [%s]: ", daftarPerangkat[indeks].Ruangan))
	if ruanganBaru != "" {
		daftarPerangkat[indeks].Ruangan = ruanganBaru
	}

	fmt.Printf("  [✓] Data perangkat ID %d berhasil diperbarui.\n", id)
}

// hapusPerangkat menghapus satu perangkat dari array berdasarkan ID (Andre)
func hapusPerangkat() {
	fmt.Println()
	fmt.Println("  [ HAPUS PERANGKAT ]")
	cetakGaris()

	if jumlahPerangkat == 0 {
		fmt.Println("  [!] Tidak ada perangkat untuk dihapus.")
		return
	}

	id := bacaInt("  Masukkan ID perangkat yang akan dihapus: ")
	// mencari posisi perangkat yang akan dihapus (Naufal)
	indeks := cariIndeksByID(id)

	if indeks == -1 {
		fmt.Printf("  [!] Perangkat dengan ID %d tidak ditemukan.\n", id)
		return
	}

	namaHapus := daftarPerangkat[indeks].Nama
	konfirmasi := bacaString(fmt.Sprintf("  Hapus \"%s\"? (y/n): ", namaHapus))

	if strings.ToLower(konfirmasi) == "y" {
		// menggeser semua elemen setelah posisi hapus ke kiri satu langkah (Erlan)
		i := indeks
		for i < jumlahPerangkat-1 {
			daftarPerangkat[i] = daftarPerangkat[i+1]
			i++
		}
		// mengosongkan slot terakhir setelah pergeseran selesai (Naufal)
		daftarPerangkat[jumlahPerangkat-1] = Perangkat{}
		jumlahPerangkat--
		fmt.Printf("  [✓] Perangkat \"%s\" berhasil dihapus.\n", namaHapus)
	} else {
		fmt.Println("  [!] Penghapusan dibatalkan.")
	}
}

// ==============================
// SUBPROGRAM PENCARIAN
// ==============================

// sequentialSearchNama mencari semua perangkat yang namanya mengandung kata kunci (Andre)
// Menggunakan Sequential Search — memeriksa setiap elemen array satu per satu
// Parameter: keyword = kata kunci yang dicari dalam nama perangkat
func sequentialSearchNama(keyword string) {
	keyword = strings.ToLower(keyword)
	ketemu := false
	fmt.Println()
	fmt.Printf("  Hasil Sequential Search nama \"%s\":\n", keyword)
	cetakHeaderTabel()

	i := 0
	// menelusuri seluruh array dan menampilkan elemen yang cocok (Erlan)
	for i < jumlahPerangkat {
		if strings.Contains(strings.ToLower(daftarPerangkat[i].Nama), keyword) {
			cetakPerangkat(daftarPerangkat[i])
			ketemu = true
		}
		i++
	}
	cetakGaris()
	if !ketemu {
		fmt.Printf("  [!] Tidak ditemukan perangkat dengan nama mengandung \"%s\".\n", keyword)
	}
}

// sequentialSearchRuangan mencari semua perangkat di ruangan tertentu (Naufal)
// Menggunakan Sequential Search — cocok untuk data yang belum terurut
// Parameter: ruangan = nama ruangan yang ingin dicari
func sequentialSearchRuangan(ruangan string) {
	ruangan = strings.ToLower(ruangan)
	ketemu := false
	fmt.Println()
	fmt.Printf("  Hasil Sequential Search ruangan \"%s\":\n", ruangan)
	cetakHeaderTabel()

	i := 0
	// membandingkan tiap ruangan perangkat dengan kata kunci yang dicari (Erlan)
	for i < jumlahPerangkat {
		if strings.Contains(strings.ToLower(daftarPerangkat[i].Ruangan), ruangan) {
			cetakPerangkat(daftarPerangkat[i])
			ketemu = true
		}
		i++
	}
	cetakGaris()
	if !ketemu {
		fmt.Printf("  [!] Tidak ditemukan perangkat di ruangan \"%s\".\n", ruangan)
	}
}

// salinArray menyalin isi array global ke array lokal sementara (Naufal)
// Parameter: temp = array tujuan salinan, n = jumlah elemen yang disalin
func salinArray(temp *[MAX_PERANGKAT]Perangkat, n int) {
	i := 0
	for i < n {
		temp[i] = daftarPerangkat[i]
		i++
	}
}

// insertionSortNama mengurutkan array perangkat berdasarkan nama menggunakan Insertion Sort (Andre)
// Insertion Sort bekerja dengan menyisipkan tiap elemen ke posisi yang tepat
// Parameter: arr = pointer array, n = jumlah elemen, asc = true untuk ascending
func insertionSortNama(arr *[MAX_PERANGKAT]Perangkat, n int, asc bool) {
	i := 1
	for i < n {
		kunci := arr[i]
		j := i - 1
		// menggeser elemen yang tidak sesuai urutan ke kanan untuk memberi ruang (Naufal)
		for j >= 0 && ((asc && strings.ToLower(arr[j].Nama) > strings.ToLower(kunci.Nama)) ||
			(!asc && strings.ToLower(arr[j].Nama) < strings.ToLower(kunci.Nama))) {
			arr[j+1] = arr[j]
			j--
		}
		// menyisipkan elemen kunci ke posisi yang sudah ditemukan (Andre)
		arr[j+1] = kunci
		i++
	}
}

// selectionSortKonsumsi mengurutkan array berdasarkan konsumsi energi menggunakan Selection Sort (Andre)
// Selection Sort memilih elemen minimum/maksimum lalu menukarnya ke posisi yang benar
// Parameter: arr = pointer array, n = jumlah elemen, asc = true untuk ascending
func selectionSortKonsumsi(arr *[MAX_PERANGKAT]Perangkat, n int, asc bool) {
	i := 0
	for i < n-1 {
		indeksEkstrem := i
		j := i + 1
		// mencari elemen paling kecil atau paling besar dari sisa array (Erlan)
		for j < n {
			konsumsiJ := hitungKonsumsiHarian(arr[j].Watt, arr[j].DurasiJam)
			konsumsiEkstrem := hitungKonsumsiHarian(arr[indeksEkstrem].Watt, arr[indeksEkstrem].DurasiJam)
			if (asc && konsumsiJ < konsumsiEkstrem) || (!asc && konsumsiJ > konsumsiEkstrem) {
				indeksEkstrem = j
			}
			j++
		}
		// menukar elemen terpilih dengan elemen di posisi i (Naufal)
		if indeksEkstrem != i {
			arr[i], arr[indeksEkstrem] = arr[indeksEkstrem], arr[i]
		}
		i++
	}
}

// binarySearchNama melakukan pencarian biner pada array terurut berdasarkan nama (Andre)
// Prasyarat: array harus sudah diurutkan ascending berdasarkan nama terlebih dahulu
// Parameter: arr = array terurut, n = jumlah elemen, keyword = nama yang dicari
func binarySearchNama(arr [MAX_PERANGKAT]Perangkat, n int, keyword string) {
	keyword = strings.ToLower(keyword)
	kiri := 0
	kanan := n - 1
	ditemukan := false
	indeksDitemukan := -1

	fmt.Println()
	fmt.Printf("  Hasil Binary Search nama \"%s\":\n", keyword)

	// membagi ruang pencarian menjadi dua di setiap iterasi (Andre)
	for kiri <= kanan && !ditemukan {
		tengah := (kiri + kanan) / 2
		namaTengah := strings.ToLower(arr[tengah].Nama)

		if namaTengah == keyword {
			indeksDitemukan = tengah
			ditemukan = true
		} else if namaTengah < keyword {
			// kata kunci berada di separuh kanan, geser batas kiri (Naufal)
			kiri = tengah + 1
		} else {
			// kata kunci berada di separuh kiri, geser batas kanan (Erlan)
			kanan = tengah - 1
		}
	}

	cetakHeaderTabel()
	if ditemukan {
		cetakPerangkat(arr[indeksDitemukan])
	} else {
		fmt.Printf("  [!] Perangkat dengan nama \"%s\" tidak ditemukan (binary search).\n", keyword)
	}
	cetakGaris()
}

// menuCariPerangkat menampilkan submenu dan mengarahkan ke metode pencarian yang dipilih (Andre)
func menuCariPerangkat() {
	fmt.Println()
	fmt.Println("  [ CARI PERANGKAT ]")
	cetakGaris()

	if jumlahPerangkat == 0 {
		fmt.Println("  [!] Belum ada perangkat untuk dicari.")
		return
	}

	fmt.Println("  Metode Pencarian:")
	fmt.Println("  1. Sequential Search (berdasarkan nama)")
	fmt.Println("  2. Sequential Search (berdasarkan ruangan)")
	fmt.Println("  3. Binary Search (berdasarkan nama persis - data akan diurutkan dahulu)")
	cetakGaris()

	pilihan := bacaInt("  Pilih metode (1-3): ")

	if pilihan == 1 {
		keyword := bacaString("  Masukkan kata kunci nama: ")
		sequentialSearchNama(keyword)
	} else if pilihan == 2 {
		ruangan := bacaString("  Masukkan nama ruangan: ")
		sequentialSearchRuangan(ruangan)
	} else if pilihan == 3 {
		// mengurutkan salinan array dulu sebelum binary search dijalankan (Naufal)
		var arrSortir [MAX_PERANGKAT]Perangkat
		salinArray(&arrSortir, jumlahPerangkat)
		insertionSortNama(&arrSortir, jumlahPerangkat, true)
		keyword := bacaString("  Masukkan nama perangkat (harus tepat): ")
		binarySearchNama(arrSortir, jumlahPerangkat, keyword)
	} else {
		fmt.Println("  [!] Pilihan tidak valid.")
	}
}

// ==============================
// SUBPROGRAM PENGURUTAN
// ==============================

// tampilkanHasilUrut menampilkan isi array yang sudah diurutkan beserta judulnya (Erlan)
// Parameter: arr = array terurut, n = jumlah elemen, judul = keterangan pengurutan
func tampilkanHasilUrut(arr [MAX_PERANGKAT]Perangkat, n int, judul string) {
	fmt.Println()
	fmt.Printf("  [ %s ]\n", judul)
	cetakHeaderTabel()
	i := 0
	for i < n {
		cetakPerangkat(arr[i])
		i++
	}
	cetakGaris()
}

// menuUrutkanPerangkat menampilkan submenu pengurutan dan menjalankan proses sorting (Andre)
func menuUrutkanPerangkat() {
	fmt.Println()
	fmt.Println("  [ URUTKAN PERANGKAT ]")
	cetakGaris()

	if jumlahPerangkat == 0 {
		fmt.Println("  [!] Belum ada perangkat untuk diurutkan.")
		return
	}

	fmt.Println("  Kategori Pengurutan:")
	fmt.Println("  1. Berdasarkan Nama (Insertion Sort)")
	fmt.Println("  2. Berdasarkan Konsumsi Energi (Selection Sort)")
	cetakGaris()

	pilihan := bacaInt("  Pilih kategori (1-2): ")

	// menanyakan arah pengurutan setelah kategori dipilih (Erlan)
	fmt.Println("  Urutan:")
	fmt.Println("  1. Ascending (A-Z / Terendah ke Tertinggi)")
	fmt.Println("  2. Descending (Z-A / Tertinggi ke Terendah)")
	arahInput := bacaInt("  Pilih urutan (1-2): ")
	asc := arahInput == 1

	// menyalin array global ke sementara agar urutan asli tidak berubah (Erlan)
	var arrSortir [MAX_PERANGKAT]Perangkat
	salinArray(&arrSortir, jumlahPerangkat)

	if pilihan == 1 {
		// mengurutkan nama menggunakan algoritma Insertion Sort (Andre)
		insertionSortNama(&arrSortir, jumlahPerangkat, asc)
		judulUrut := "DIURUTKAN BERDASARKAN NAMA"
		if asc {
			judulUrut += " (A → Z)"
		} else {
			judulUrut += " (Z → A)"
		}
		tampilkanHasilUrut(arrSortir, jumlahPerangkat, judulUrut)
	} else if pilihan == 2 {
		// mengurutkan konsumsi energi menggunakan algoritma Selection Sort (Andre)
		selectionSortKonsumsi(&arrSortir, jumlahPerangkat, asc)
		judulUrut := "DIURUTKAN BERDASARKAN KONSUMSI ENERGI"
		if asc {
			judulUrut += " (Terendah → Tertinggi)"
		} else {
			judulUrut += " (Tertinggi → Terendah)"
		}
		tampilkanHasilUrut(arrSortir, jumlahPerangkat, judulUrut)
	} else {
		fmt.Println("  [!] Pilihan tidak valid.")
	}
}

// ==============================
// SUBPROGRAM STATISTIK
// ==============================

// tampilkanStatistik menghitung dan menampilkan ringkasan statistik penggunaan daya (Andre)
// Menampilkan total konsumsi harian, estimasi bulanan, dan peringkat perangkat paling boros
func tampilkanStatistik() {
	fmt.Println()
	fmt.Println("  [ STATISTIK PENGGUNAAN DAYA ]")
	cetakGaris()

	if jumlahPerangkat == 0 {
		fmt.Println("  [!] Belum ada data perangkat.")
		return
	}

	totalKonsumsi := hitungTotalKonsumsi()
	totalWatt := 0.0
	i := 0
	// menjumlahkan total watt seluruh perangkat yang terdaftar (Naufal)
	for i < jumlahPerangkat {
		totalWatt += daftarPerangkat[i].Watt
		i++
	}

	fmt.Printf("  Jumlah Perangkat    : %d\n", jumlahPerangkat)
	fmt.Printf("  Total Daya Terpasang: %.1f Watt\n", totalWatt)
	fmt.Printf("  Total Konsumsi/Hari : %.3f kWh\n", totalKonsumsi)
	fmt.Printf("  Estimasi/Bulan      : %.2f kWh\n", totalKonsumsi*30)
	fmt.Println()

	// membuat salinan dan mengurutkan descending untuk menemukan yang paling boros (Erlan)
	var arrSortir [MAX_PERANGKAT]Perangkat
	salinArray(&arrSortir, jumlahPerangkat)
	selectionSortKonsumsi(&arrSortir, jumlahPerangkat, false)

	// menampilkan maksimal 5 perangkat paling boros di bagian atas (Naufal)
	tampil := jumlahPerangkat
	if tampil > 5 {
		tampil = 5
	}

	fmt.Println("  TOP PERANGKAT PALING BOROS:")
	cetakHeaderTabel()
	i = 0
	for i < tampil {
		cetakPerangkat(arrSortir[i])
		i++
	}
	cetakGaris()

	// menampilkan bar chart persentase konsumsi tiap perangkat (Naufal)
	if totalKonsumsi > 0 {
		fmt.Println()
		fmt.Println("  PERSENTASE KONSUMSI PER PERANGKAT:")
		cetakGaris()
		i = 0
		for i < jumlahPerangkat {
			k := hitungKonsumsiHarian(daftarPerangkat[i].Watt, daftarPerangkat[i].DurasiJam)
			persen := (k / totalKonsumsi) * 100
			// menghitung panjang bar sesuai persentase, tiap blok mewakili 5% (Erlan)
			barPanjang := int(math.Round(persen / 5))
			bar := strings.Repeat("█", barPanjang)
			fmt.Printf("  %-20s: %5.1f%% %s\n", daftarPerangkat[i].Nama, persen, bar)
			i++
		}
		cetakGaris()
	}
}

// ==============================
// MAIN PROGRAM
// ==============================

// main adalah titik masuk program, menampilkan menu dan memproses pilihan pengguna (Andre)
func main() {
	// eel @jebb_24
	cetakHeader()
	fmt.Println("  Selamat datang di PowerLog!")

	// program terus berjalan sampai pengguna memilih menu keluar (Erlan)
	berjalan := true
	for berjalan {
		cetakMenuUtama()
		pilihan := bacaInt("  Pilih menu: ")

		if pilihan == 1 {
			tambahPerangkat()
		} else if pilihan == 2 {
			lihatSemuaPerangkat()
		} else if pilihan == 3 {
			menuCariPerangkat()
		} else if pilihan == 4 {
			ubahPerangkat()
		} else if pilihan == 5 {
			hapusPerangkat()
		} else if pilihan == 6 {
			menuUrutkanPerangkat()
		} else if pilihan == 7 {
			tampilkanStatistik()
		} else if pilihan == 0 {
			// mengakhiri perulangan dan menampilkan pesan perpisahan (Andre)
			fmt.Println()
			fmt.Println("  Terima kasih telah menggunakan PowerLog!")
			fmt.Println("+++ PowerLog +++")
			berjalan = false
		} else {
			fmt.Println("  [!] Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}