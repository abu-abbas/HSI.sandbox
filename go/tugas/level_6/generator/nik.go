package generator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/abu-abbas/level_6/str"
)

const PREFIX = "AR"

type NikParser struct {
	Gender   string
	Tahun    int
	Semester string
	Urutan   int
	Jumlah   int
}

func NIK(n NikParser) ([]string, error) {
	var result []string
	var semester string

	// parse gender
	cg, err := getGenderCode(n.Gender)
	if err != nil {
		return result, err
	}

	// parse tahun
	strTahun := strconv.Itoa(n.Tahun)[2:]

	// parse semester
	if n.Semester == "" {
		semester = getSemester()
	} else {
		semester = n.Semester
	}

	// generate nik
	for i := 1; i <= n.Jumlah; i++ {
		c := strconv.Itoa(n.Urutan + i)
		u := str.LeftPad(c, 5, "0")
		s := fmt.Sprintf("%s%s%s%s-%s", PREFIX, cg, strTahun, semester, u)

		result = append(result, s)
	}

	return result, nil
}

func ParseNIK(n *NikParser, nik string) error {
	if !strings.HasPrefix(nik, PREFIX) {
		return errors.New("nik invalid")
	}

	gender := nik[2:3]
	_, err := getGenderCode(gender)
	if err != nil {
		return err
	}

	strTahun := nik[3:5]
	_, err = strconv.Atoi(strTahun)
	if err != nil {
		return errors.New("tahun tidak valid")
	}

	tahun, err := strconv.Atoi("20" + strTahun)
	if err != nil {
		return err
	}

	smt := nik[5:6]
	smtInt, err := strconv.Atoi(smt)
	if err != nil {
		return err
	}

	if smtInt < 0 || smtInt > 2 {
		return errors.New("semester tidak valid")
	}

	strUrutan := nik[7:]
	if len(strUrutan) > 5 || len(strUrutan) < 1 {
		return errors.New("urutan tidak valid")
	}

	urutan, err := strconv.Atoi(strUrutan)
	if err != nil {
		return err
	}

	n.Gender = gender
	n.Semester = smt
	n.Tahun = tahun
	n.Urutan = urutan

	return nil
}

func getGenderCode(gender string) (string, error) {
	mapped := map[string]string{
		"ikhwan": "N",
		"i":      "N",
		"n":      "N",
		"akhwat": "T",
		"a":      "T",
		"t":      "T",
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
