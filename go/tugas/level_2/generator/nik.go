package generator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/abu-abbas/level_2/str"
)

const PREFIX = "AR"

func NIK(gender string, tahun int, jumlah int, index int, smt string) ([]string, error) {
	var result 		[]string
	var semester	string
	
	// parse gender
    cg, err := getGenderCode(gender)
    if err != nil {
        return result, err
    }
	
	// parse tahun
    strTahun := strconv.Itoa(tahun)[2:]
	
	// parse semester
	if smt == "" {
		semester = getSemester()
	} else {
		semester = smt
	}
	
	// generate nik
    for i := 1; i <= jumlah; i++ {
        c := strconv.Itoa(index + i)
        u := str.LeftPad(c, 5, "0")
        s := fmt.Sprintf("%s%s%s%s-%s", PREFIX, cg, strTahun, semester, u)

        result = append(result, s)
    }

    return result, nil
}

func ParseNIK(nik string) (gender string, tahun int, smt string, urutan int) {
	gender	 = nik[2:][:1]
	tahun,_	 = strconv.Atoi("20" + nik[3:][:2])
	smt 	 = nik[5:][:1]
	urutan,_ = strconv.Atoi(nik[7:])

	return
}

func getGenderCode(gender string) (string, error) {
    mapped := map[string]string{
        "ikhwan": "N",
        "i": "N",
        "n": "N",
        "akhwat": "T",
        "a": "T",
        "t": "T",
    }

    code, ok := mapped[strings.ToLower(gender)]
    if !ok {
        strError := fmt.Sprintf("jenis kelamin %s tidak sesuai", gender)
        return "", errors.New(strError)
    }

    return code, nil
}

func getSemester() string {
    var semester string
    _, month, _ := time.Now().Date()

    if int(month) <= 6 {
        semester = "1"
    } else {
        semester = "2"
    }

    return semester
}

