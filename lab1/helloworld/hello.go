package main

import (
	"fmt"

	"github.com/adpit/latihangolang/lab1/stringutil"
	// 'workspace' C:\Users\adpit\Desktop\data\atom-x64-windows\clone\latihangolang\lab\stringutil (before)
	// 'package' C:\Users\adpit\go\src\github.com\adpit\lab\stringutil (after)
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH")) // go run hasilnya berhasil Hello, Go! (public fungsi)
	// fmt.Println(stringutil.reverseTwo("!oG ,olleH")) // go run hasilnya error (private fungsi)
}
