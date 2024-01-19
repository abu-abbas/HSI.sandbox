package main

import (
	"errors"
	"fmt"

	"rsc.io/quote"
)

/*
Soal no. 2
Jelaskan kegunaan function fmt.Sprintln apa bedanya dengan
fmt.Println, beri contoh code, copas outputnya
*/
func DiffPrint() {
	/*
		Jawab:
		fmt.Sprintln untuk membuat string dengan nilai-nilai
		tertentu tergantung inputan dan tidak mencetaknya kedalam konsol
	*/
	result := fmt.Sprintln(quote.Hello())
	fmt.Print("output fmt.Sprintln:", result)

	/*
		Jawab:
		fmt.Println digunakan untuk mencetak nilai-nilai kedalam konsol
		fmt.Println juga selalu menambahkan spasi antara operand dan
		selalu menambahkan newline
	*/
	fmt.Println("output fmt.Println:", result)
}

/*
Soal no. 3
Jelaskan kegunaan function fmt.Errorf. Apa bedanya dengan errors.New?
beri contoh code, copas outputnya
*/
func DiffError() {
	/*
		Jawab:
		fungsi fmt.Errorf adalah untuk membuat nilai error dengan menggunakan
		format string secara dinamis
	*/
	code := 504
	err := fmt.Errorf("output fmt.Errorf: code http untuk gateway timeout adalah %v", code)
	fmt.Println(err)

	/*
		Jawab:
		fungsi errors.New adalah untuk membuat nilai error dengan menggunakan
		format string secara statis
	*/
	err = errors.New("output errors.New")
	fmt.Println(err)
}

func main() {
	DiffPrint()

	DiffError()
}
