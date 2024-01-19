# Tugas HSI.Sandbox - Level 1

**Pertanyaan 1**
_Bagaimana dependency management dalam golang?_

**Jawab:**
Dependency management dalam golang menggunakan Go Modules. Perintah yang digunakan adalah `go get nama-dependensi`. Contohnya penggunaannya adalah sebagai berikut:

 1. Menambahakan dependensi: `go get rsc.io/quote`
 2. Menandai versi dependensi: `go mod tidy`
 3. File `go.mod` akan diperbaharui dengan informasi dependensi dan versinya

**Pertanyaan 2**
_Jelaskan kegunaan function `fmt.Sprintln` apa bedanya dengan `fmt.Println` beri contoh code copas outputnya?_

**Jawab:**
fmt.Sprintln untuk membuat string dengan nilai-nilai tertentu tergantung inputan dan tidak mencetaknya kedalam konsol.

contoh kodingan:
```go
result := fmt.Sprintln(quote.Hello())
fmt.Print(result)
```
output:
```shell
Hello, world.
```

Sedangkan `fmt.Println` digunakan untuk mencetak nilai-nilai kedalam konsol. `fmt.Println` juga selalu menambahkan spasi antara operand dan selalu menambahkan newline

contoh kodingan:
```go
result := fmt.Sprintln(quote.Hello())
fmt.Println("output:", result)
```

output:
```shell
output: Hello, world.

```

**Pertanyaan 3**
_Jelaskan kegunaan function fmt.Errorf. Apa bedanya dengan errors.New? Beri contoh code, copas outputnya_

**Jawab:**
Fungsi `fmt.Errorf` adalah untuk membuat nilai error dengan menggunakan format string secara **dinamis**.
contoh kodingan:
```go
code :=  504
err := fmt.Errorf("code http untuk gateway timeout adalah %v", code)
fmt.Println(err)
```

output:
```shell
code http untuk gateway timeout adalah 504

```

Sedangkan fungsi `errors.New` adalah untuk membuat nilai error dengan menggunakan format string secara **statis**.

contoh kodingan:
```go
err = errors.New("terjadi galat pada server")
fmt.Println(err)
```

output:
```shell
terjadi galat pada server

```
