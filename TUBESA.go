package main

import "fmt"

const NMAX int = 100

type tugas struct {
	nama, kategori, status string
	prioritas              int
}

type arrTugas [NMAX]tugas

func main() {
	var pilihan int
	menuUtama(&pilihan)
}

func menuUtama(pilihan *int) {
	var data arrTugas
	var jumlahTugas int

	for *pilihan != 7 {
		fmt.Println()
		fmt.Println("===== APLIKASI MANAJEMEN TUGAS RUMAH TANGGA =====")
		fmt.Println("1. Tambah Tugas")
		fmt.Println("2. Ubah Status Tugas")
		fmt.Println("3. Hapus Tugas")
		fmt.Println("4. Tampilkan Tugas")
		fmt.Println("5. Cari Tugas")
		fmt.Println("6. Urutkan Tugas")
		fmt.Println("7. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(pilihan)

		if *pilihan == 1 {
			tambahTugas(&data, &jumlahTugas)
		} else if *pilihan == 2 {
			ubahStatus(&data, jumlahTugas)
		} else if *pilihan == 3 {
			hapusTugas(&data, &jumlahTugas)
		} else if *pilihan == 4 {
			tampilkanTugas(data, jumlahTugas)
		} else if *pilihan == 5 {
			menuCari(data, jumlahTugas)
		} else if *pilihan == 6 {
			menuSort(&data, jumlahTugas)
		}
	}
}

func tambahTugas(data *arrTugas, jumlah *int) {
	fmt.Println()

	fmt.Print("Nama Tugas     : ")
	fmt.Scan(&data[*jumlah].nama)

	fmt.Print("Kategori       : ")
	fmt.Scan(&data[*jumlah].kategori)

	fmt.Print("Prioritas(1-5) : ")
	fmt.Scan(&data[*jumlah].prioritas)

	data[*jumlah].status = "Belum"

	*jumlah++

	fmt.Println("Tugas berhasil ditambahkan!")
}

func tampilkanTugas(data arrTugas, jumlah int) {
	var selesai int

	fmt.Println()
	fmt.Println("----------------------------------------------------------------")
	fmt.Printf("| %-3s | %-20s | %-10s | %-10s | %-3s |\n",
		"No", "Nama", "Kategori", "Status", "P")
	fmt.Println("----------------------------------------------------------------")

	for i := 0; i < jumlah; i++ {
		fmt.Printf("| %-3d | %-20s | %-10s | %-10s | %-3d |\n",
			i+1,
			data[i].nama,
			data[i].kategori,
			data[i].status,
			data[i].prioritas)

		if data[i].status == "Selesai" {
			selesai++
		}
	}

	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Jumlah Tugas          :", jumlah)
	fmt.Println("Jumlah Selesai        :", selesai)
	fmt.Println("Jumlah Belum Selesai  :", jumlah-selesai)

	if jumlah > 0 {
		fmt.Printf("Progress              : %.2f%%\n",
			float64(selesai)/float64(jumlah)*100)
	}
}

func ubahStatus(data *arrTugas, jumlah int) {
	var nama string
	var idx int

	fmt.Print("Masukkan nama tugas: ")
	fmt.Scan(&nama)

	idx = binarySearch(*data, jumlah, nama)

	if idx == -1 {
		fmt.Println("Tugas tidak ditemukan!")
		return
	}

	data[idx].status = "Selesai"

	fmt.Println("Status berhasil diubah!")
}

func hapusTugas(data *arrTugas, jumlah *int) {
	var nama string
	var idx int

	fmt.Print("Masukkan nama tugas: ")
	fmt.Scan(&nama)

	idx = binarySearch(*data, *jumlah, nama)

	if idx == -1 {
		fmt.Println("Tugas tidak ditemukan!")
		return
	}

	for i := idx; i < *jumlah-1; i++ {
		data[i] = data[i+1]
	}

	*jumlah--

	fmt.Println("Tugas berhasil dihapus!")
}

func menuCari(data arrTugas, jumlah int) {
	var pilih int

	fmt.Println()
	fmt.Println("1. Cari Berdasarkan Nama")
	fmt.Println("2. Cari Berdasarkan Kategori")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		cariNama(data, jumlah)
	} else if pilih == 2 {
		cariKategori(data, jumlah)
	}
}

func cariNama(data arrTugas, jumlah int) {
	var nama string
	var idx int

	selectionSortNama(&data, jumlah)

	fmt.Print("Masukkan nama tugas: ")
	fmt.Scan(&nama)

	idx = binarySearch(data, jumlah, nama)

	if idx == -1 {
		fmt.Println("Tugas tidak ditemukan!")
	} else {
		fmt.Println("Tugas ditemukan!")
		fmt.Println("Nama      :", data[idx].nama)
		fmt.Println("Kategori  :", data[idx].kategori)
		fmt.Println("Status    :", data[idx].status)
		fmt.Println("Prioritas :", data[idx].prioritas)
	}
}

func cariKategori(data arrTugas, jumlah int) {
	var kategori string
	var ketemu bool

	fmt.Print("Masukkan kategori: ")
	fmt.Scan(&kategori)

	for i := 0; i < jumlah; i++ {
		if data[i].kategori == kategori {
			fmt.Println(data[i].nama)
			ketemu = true
		}
	}

	if !ketemu {
		fmt.Println("Tidak ada tugas ditemukan")
	}
}

func menuSort(data *arrTugas, jumlah int) {
	var pilih int

	fmt.Println()
	fmt.Println("1. Prioritas Terkecil")
	fmt.Println("2. Prioritas Terbesar")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		selectionSortPrioritas(data, jumlah)
		tampilkanTugas(*data, jumlah)
	} else if pilih == 2 {
		insertionSortPrioritas(data, jumlah)
		tampilkanTugas(*data, jumlah)
	}
}

func binarySearch(data arrTugas, n int, x string) int {
	selectionSortNama(&data, n)

	var left, right, mid int

	left = 0
	right = n - 1

	for left <= right {
		mid = (left + right) / 2

		if data[mid].nama == x {
			return mid
		} else if x < data[mid].nama {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func selectionSortNama(data *arrTugas, n int) {
	var minIdx int

	for i := 0; i < n-1; i++ {
		minIdx = i

		for j := i + 1; j < n; j++ {
			if data[j].nama < data[minIdx].nama {
				minIdx = j
			}
		}

		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func selectionSortPrioritas(data *arrTugas, n int) {
	var minIdx int

	for i := 0; i < n-1; i++ {
		minIdx = i

		for j := i + 1; j < n; j++ {
			if data[j].prioritas < data[minIdx].prioritas {
				minIdx = j
			}
		}

		data[i], data[minIdx] = data[minIdx], data[i]
	}
}

func insertionSortPrioritas(data *arrTugas, n int) {
	var pass, i int
	var temp tugas

	for pass = 1; pass < n; pass++ {
		temp = data[pass]
		i = pass

		for i > 0 && temp.prioritas > data[i-1].prioritas {
			data[i] = data[i-1]
			i--
		}

		data[i] = temp
	}
}