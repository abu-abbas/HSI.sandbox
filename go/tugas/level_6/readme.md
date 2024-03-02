# HSI Sandbox [Go] - Tugas Level 6

1. Buatkan unit test dari fungsi sebelumnya NIK generator

    Dependensi yang saya gunakan
    ```
    [testing-unit]
    require github.com/stretchr/testify v1.9.0
    ```

    Struktur folder
    ```
        .
        ├── generate_nik_lanjutan.go
        ├── generate_nik_lanjutan_test.go
        ├── generator
        │  └── nik.go
        ├── go.mod
        ├── go.sum
        ├── main.go
        └── str
           └── leftPad.go

    ```

    Berikut ini adalah hasil unit testing pada fungsi `GenerateNIKLanjutan`
    ```shell
    go test -v

    === RUN   TestGenerateNIKLanjutan
    === RUN   TestGenerateNIKLanjutan/empty_nik
    === RUN   TestGenerateNIKLanjutan/invalid_prefix_nik
    === RUN   TestGenerateNIKLanjutan/valid_test
    --- PASS: TestGenerateNIKLanjutan (0.00s)
        --- PASS: TestGenerateNIKLanjutan/empty_nik (0.00s)
        --- PASS: TestGenerateNIKLanjutan/invalid_prefix_nik (0.00s)
        --- PASS: TestGenerateNIKLanjutan/valid_test (0.00s)
    PASS
    ok  	github.com/abu-abbas/level_6	0.005s


    ```


2. Modifikasi `generateNIKLanjutan` agar dapat mengambil NIK awal dari function lain
    `func generateNIKLanjutan(f func() string, 10) []string`

    2.1. Perubahan pada fungsi `generateNIKLanjutan`
    ```go
    func GenerateNIKLanjutan(fn func() string, jumlah int) ([]string, error) {
    	currentNik := fn()

    	if currentNik == "" {
    		return nil, errors.New("current nik tidak boleh kosong")
    	}

    	nikParser := generator.NikParser{Jumlah: jumlah}
    	err := generator.ParseNIK(&nikParser, currentNik)
    	if err != nil {
    		return nil, err
    	}

    	result, err := generator.NIK(nikParser)
    	if err != nil {
    		return nil, errors.New("nik tidak boleh nil")
    	}

    	return result, nil
    }
    ```

    2.2. Perubahan pada fungsi `ParseNIK` untuk menambah validasi inputan
    ```go
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
    ```
    2.3. Perubahan parameter pada fungsi `NIK`
    ```go
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
    ```


3. Buatkan stub dari tugas no. 2, dan simulasikan beberapa test case

    Berikut ini test case yang saya buat untuk melakukan unit test pada fungsi `generateNIKLanjutan`

    ```go
    package main

    import (
    	"testing"

    	"github.com/stretchr/testify/assert"
    	"github.com/stretchr/testify/require"
    )

    func TestGenerateNIKLanjutan(t *testing.T) {
    	for scenario, fn := range map[string]func(t *testing.T){
    		"valid test":         testValidNik,
    		"empty nik":          testEmptyNik,
    		"invalid prefix nik": testInvalidPrefixNik,
    	} {
    		t.Run(scenario, func(t *testing.T) {
    			fn(t)
    		})
    	}
    }

    func testValidNik(t *testing.T) {
    	validMockFn := func() string {
    		return "ARN171-06140"
    	}

    	want := []string{"ARN171-06141", "ARN171-06142"}
    	got, err := GenerateNIKLanjutan(validMockFn, 2)
    	require.Nil(t, err)
    	assert.Equal(t, want, got)
    }

    func testEmptyNik(t *testing.T) {
    	emptyMockFn := func() string {
    		return ""
    	}

    	_, err := GenerateNIKLanjutan(emptyMockFn, 2)
    	require.Error(t, err)
    	assert.Equal(t, "current nik tidak boleh kosong", err.Error())
    }

    func testInvalidPrefixNik(t *testing.T) {
    	invalidPrefixNikMockFn := func() string {
    		return "RAN171-06140"
    	}

    	_, err := GenerateNIKLanjutan(invalidPrefixNikMockFn, 2)
    	require.Error(t, err)
    	assert.Equal(t, "nik invalid", err.Error())
    }

    ```
