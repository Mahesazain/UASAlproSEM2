package main

import "fmt"

const NMAX int = 100

type book struct {
	namaBuku   string
	kodeBuku   int
	kategori   string
	peminjaman borrowed
}

type borrowed struct {
	namaPeminjam string
	pinjam       bool
	count        int
	tglPinjam    int
	tglKembali   int
}

type arrPerpus [NMAX]book

func manageBooks(T *arrPerpus) {
	var input int

	fmt.Println("= = = = =")
	fmt.Println("Manage Books")
	fmt.Println("1. Add Book")
	fmt.Println("2. Edit Book")
	fmt.Println("3. Delete Book")
	fmt.Println("= = = = =")

	fmt.Scanln(&input)

	if input == 1 {
		fmt.Println("= = = = =")
		addBooks(T)
	} else if input == 2 {
		fmt.Println("= = = = =")
		fmt.Print("Enter book ID: ")
		fmt.Scanln(&input)
		editBook(T, input)
	} else if input == 3 {
		fmt.Println("= = = = =")
		fmt.Print("Enter book ID: ")
		fmt.Scanln(&input)
		deleteBook(T, input)
	} else {
		fmt.Println("= = = = =")
		fmt.Println("Maaf, pilihan tidak valid. Silakan pilih nomor dari menu.")
	}
}

func addBooks(T *arrPerpus) {
	var tempCategory, tempTitle string
	var numBooks, emptySlot int     // Keep track of the number of books added
	var stop bool = false           // Use a flag to exit the loop when necessary and check if the array slot is empty
	var foundEmptySlot bool = false // for some reason i can't put these 2 bool into one even tho it has the same value which is false
	fmt.Println("Notes jika ingin memiliki spasi di nama buku dan kategori mohon memakai '_' ")
	fmt.Println("dan jika ingin berhenti menambahkan buku tulis 'STOP' ")

	for i := 0; i < NMAX && !foundEmptySlot; i++ {
		if T[i].namaBuku == "" {
			emptySlot = i
			foundEmptySlot = true
		}
	}

	fmt.Print("Masukan nama buku: ")
	fmt.Scanln(&tempTitle)
	if tempTitle == "" {
		fmt.Println("mohon untuk menamakan yang benar")
		fmt.Print("Masukan nama buku: ")
		fmt.Scanln(&tempTitle)
	}

	for emptySlot < NMAX && !stop {
		if tempTitle == "STOP" {
			stop = true
		} else {
			fmt.Print("Masukan kategori buku: ")
			fmt.Scanln(&tempCategory)

			T[emptySlot].namaBuku = tempTitle
			T[emptySlot].kategori = tempCategory
			T[emptySlot].kodeBuku = emptySlot + 1
			fmt.Println("Buku berhasil ditambahkan dengan ID:", T[emptySlot].kodeBuku)
			fmt.Println("= = = = =")

			numBooks++
			emptySlot++
			for emptySlot < NMAX && T[emptySlot].namaBuku != "" {
				emptySlot++
			}
		}
		if !stop {
			fmt.Print("Masukan nama buku: ")
			fmt.Scanln(&tempTitle)
			if tempTitle == "" {
				fmt.Println("mohon untuk menamakan yang benar")
				fmt.Print("Masukan nama buku: ")
				fmt.Scanln(&tempTitle)
			}
		}
	}
	// Print a message if no books were added
	if numBooks == 0 {
		fmt.Println("Tidak ada buku yang ditambahkan.")
	}
}

func editBook(T *arrPerpus, id int) {
	found := false
	var i int
	for i = 0; i < NMAX && T[i].namaBuku != ""; i++ {
		if T[i].kodeBuku == id {
			found = true
			fmt.Println("= = = = =")
			fmt.Println("Book Details")
			fmt.Println("Judul:", T[i].namaBuku)
			fmt.Println("Kategori:", T[i].kategori)
			fmt.Println("ID:", T[i].kodeBuku)
			fmt.Println("= = = = =")
			fmt.Println("1. Edit book details")
			fmt.Println("2. Cancel")
			fmt.Println("= = = = =")
			var input int
			fmt.Scanln(&input)

			if input == 1 {
				var tempTitle, tempCategory string
				fmt.Println("= = = = =")
				fmt.Print("Masukan nama buku baru: ")
				fmt.Scanln(&tempTitle)
				fmt.Print("Masukan kategori buku baru: ")
				fmt.Scanln(&tempCategory)
				T[i].namaBuku = tempTitle
				T[i].kategori = tempCategory
				fmt.Println("= = = = =")
				fmt.Println("Buku berhasil diperbarui.")
			} else {
				fmt.Println("= = = = =")
				fmt.Println("Buku tidak diperbarui.")
			}
		}
	}

	if !found {
		fmt.Println("= = = = =")
		fmt.Println("Buku tidak ditemukan.")
	}
}

func deleteBook(T *arrPerpus, id int) {
	found := false
	var i int
	for i = 0; i < NMAX && T[i].namaBuku != ""; i++ {
		if T[i].kodeBuku == id {
			found = true
			for j := i; j < NMAX-1 && T[j+1].namaBuku != ""; j++ {
				T[j] = T[j+1]
			}
			T[NMAX-1] = book{} // Clear the last book
			fmt.Println("= = = = =")
			fmt.Println("Buku berhasil dihapus.")
			for j := i; j < NMAX-1 && T[j].namaBuku != ""; j++ {
				T[j].kodeBuku = j + 1
			}
		}
	}

	if !found {
		fmt.Println("= = = = =")
		fmt.Println("Buku tidak ditemukan.")
	}
}

func cetakBook(T arrPerpus) {
	var i int = 0
	if T[i].kodeBuku == 0 {
		fmt.Println("Tidak ada buku, Mohon untuk menambahkan buku")
		return
	} else {
		for i := 0; i < NMAX && T[i].namaBuku != ""; i++ {
			fmt.Println("Judul:", T[i].namaBuku, "Kategori:", T[i].kategori, "ID:", T[i].kodeBuku)
		}
	}
}

func SearchCategory(T arrPerpus, kategori string) {
	var i int = 0
	var found bool = false
	if T[i].kodeBuku == 0 {
		fmt.Println("Maaf anda belum menambahkan buku")
		return
	} else {
		for i := 0; i < NMAX && T[i].namaBuku != ""; i++ {
			if kategori == T[i].kategori {
				fmt.Println("Judul:", T[i].namaBuku, "Kategori:", T[i].kategori, "ID:", T[i].kodeBuku)
				found = true
			}
		}
		if !found {
			fmt.Println("Maaf, kami tidak dapat menemukan kategori yang Anda cari.")
		}
	}
}

func borrowOrReturnBook(T *arrPerpus, id int) {
	var input int

	fmt.Println("= = = = =")
	fmt.Println("Borrow or Return Book")
	fmt.Println("1. Borrow Book")
	fmt.Println("2. Return Book")
	fmt.Println("= = = = =")

	fmt.Scanln(&input)

	if input == 1 {
		borrowBook(T, id)
	} else if input == 2 {
		returnBook(T)
	} else {
		fmt.Println("= = = = =")
		fmt.Println("Maaf, pilihan tidak valid. Silakan pilih nomor dari menu.")
	}
}

func borrowBook(T *arrPerpus, id int) {
	var found bool = false
	for i := 0; i < NMAX && T[i].namaBuku != ""; i++ {
		if T[i].kodeBuku == id {
			found = true
			if T[i].peminjaman.pinjam == false {
				T[i].peminjaman.pinjam = true
				T[i].peminjaman.count++
				fmt.Println("= = = = =")
				fmt.Print("Nama peminjam: ")
				fmt.Scanln(&T[i].peminjaman.namaPeminjam)
				fmt.Print("Jumlah hari pinjam: ")
				fmt.Scanln(&T[i].peminjaman.tglPinjam)
				fmt.Println("Buku berhasil dipinjam.")
			} else {
				fmt.Println("= = = = =")
				fmt.Println("Buku tidak dapat dipinjam karena sedang dipinjam.")
			}
			break
		}
	}
	if !found {
		fmt.Println("= = = = =")
		fmt.Println("ID buku tidak ditemukan.")
	}
}

func returnBook(T *arrPerpus) {
	var input int

	fmt.Println("= = = = =")
	fmt.Print("Masukan ID buku yang dikembalikan: ")
	fmt.Scanln(&input)

	var found bool = false
	for i := 0; i < NMAX && T[i].namaBuku != ""; i++ {
		if T[i].kodeBuku == input {
			found = true
			if T[i].peminjaman.pinjam == true {
				T[i].peminjaman.pinjam = false

				fmt.Println("= = = = =")
				fmt.Print("Jumlah hari keterlambatan: ")
				var daysLate int
				fmt.Scanln(&daysLate)

				T[i].peminjaman.tglKembali = T[i].peminjaman.tglPinjam + daysLate

				if T[i].peminjaman.tglKembali > T[i].peminjaman.tglPinjam+7 {
					// Calculate late return fine
					fine := (T[i].peminjaman.tglKembali - (T[i].peminjaman.tglPinjam + 7)) * 2000
					fmt.Println("Denda keterlambatan:", fine)
				}
				fmt.Println("Terima kasih telah mengembalikan bukunya.")
			} else {
				fmt.Println("= = = = =")
				fmt.Println("Buku tidak sedang dipinjam.")
			}
			break
		}
	}
	if !found {
		fmt.Println("= = = = =")
		fmt.Println("ID buku tidak ditemukan.")
	}
}

func listBorrowed(T arrPerpus) {
	for i := 0; i < NMAX && T[i].namaBuku != ""; i++ {
		if T[i].peminjaman.pinjam == true {
			fmt.Println("Judul:", T[i].namaBuku, "Dipinjam oleh:", T[i].peminjaman.namaPeminjam)
		}
	}
}

func popularBooks(T arrPerpus) {
	var popBooks [5]book
	var i, j int

	// Sort the books by count (most borrowed first)
	for i = 0; i < NMAX-1; i++ {
		for j = 0; j < NMAX-i-1; j++ {
			if T[j].peminjaman.count < T[j+1].peminjaman.count {
				T[j], T[j+1] = T[j+1], T[j]
			}
		}
	}

	// Get the top 5 most popular books
	for i = 0; i < 5; i++ {
		popBooks[i] = T[i]
	}

	// Print the popular books
	fmt.Println("5 buku paling populer:")
	for i = 0; i < 5; i++ {
		fmt.Println(i+1, ". Judul:", popBooks[i].namaBuku, "Jumlah peminjaman:", popBooks[i].peminjaman.count)
	}
}

func menu() {
	var T1 arrPerpus
	var input int
	var masukan string

	for {
		fmt.Println("= = = = =")
		fmt.Println("Library Menu")
		fmt.Println("1. Manage Books")
		fmt.Println("2. List Books")
		fmt.Println("3. Search using category")
		fmt.Println("4. Borrow or Return Books")
		fmt.Println("5. List Borrowed Books")
		fmt.Println("6. Popular Books")
		fmt.Println("7. Exit")
		fmt.Println("= = = = =")

		fmt.Scanln(&input)

		if input == 1 {
			fmt.Println("= = = = =")
			manageBooks(&T1)
		} else if input == 2 {
			fmt.Println("= = = = =")
			cetakBook(T1)
		} else if input == 3 {
			fmt.Println("= = = = =")
			fmt.Print("Search Category: ")
			fmt.Scanln(&masukan)
			SearchCategory(T1, masukan)
		} else if input == 4 {
			fmt.Println("= = = = =")
			fmt.Print("Masukan ID buku: ")
			fmt.Scanln(&input)
			borrowOrReturnBook(&T1, input)
		} else if input == 5 {
			fmt.Println("= = = = =")
			listBorrowed(T1)
		} else if input == 6 {
			fmt.Println("= = = = =")
			popularBooks(T1)
		} else if input == 7 {
			fmt.Println("= = = = =")
			fmt.Println("Terima kasih telah menggunakan layanan perpustakaan kami.")
			return
		} else {
			fmt.Println("= = = = =")
			fmt.Println("Maaf, pilihan tidak valid. Silakan pilih nomor dari menu.")
		}
	}
}

func main() {
	menu()
}

/* books list
The_Great_Gatsby
Literature
Harry_Potter_and_the_Philosopher's_Stone
Fantasy
The_Hitchhiker's_Guide_to_the_Galaxy
Fiction
To_Kill_a_Mockingbird
Fiction
The_Catcher_in_the_Rye
Fiction
The_Lord_of_the_Rings
Fantasy
Pride_and_Prejudice
Romance
1984
Fiction
The_Hunger_Games
Fiction
The_Da_Vinci_Code
Mystery
The_Shining
Horror
The_Chronicles_of_Narnia
Fantasy
The_Color_Purple
Fiction
The_Girl_with_the_Dragon_Tattoo
Mystery
A_Game_of_Thrones
Fantasy
*/
