package main

// Setiap file pergi dimulai dengan nama paket. Nama paketnya harus sama dengan nama folder
// kecuali package main , package main adalah entry point untuk program anda.

import "fmt"

// Jika Anda menggunakan kode dari paket lain, Anda akan mencantumkan paket yang ingin Anda impor.
// Ini memungkinkan Anda untuk menggunakan kode dalam program Anda yang telah ditulis orang lain. Paket juga kadang disebut sebagai library.

// Paket "fmt" sedang diimpor.

func main() {

	fmt.Println("Hello World!")

}
