package utils

func ErrorCheck(e error) {
	if e != nil {
		panic(e)
	}
}
