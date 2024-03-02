package main

import "fmt"

const NIK_FROM_CONST string = "ARN171-06140"

func main() {
	coba, err := GenerateNIKLanjutan(getNikFromConst, 10)
	fmt.Println(coba, err)
}

func getNikFromConst() string {
	return NIK_FROM_CONST
}
