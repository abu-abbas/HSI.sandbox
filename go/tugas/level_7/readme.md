# HSI Sandbox [Go] - Tugas Level 7

## Task
1. Ganti perhitungan amount pada table tugas 3 dengan scheduler
2. Ganti perhitungan amount pada table tugas 3 dengan goroutine


## Dependensi

Dependensi modul go yang saya gunakan adalah sbb:
```
[database]
github.com/jmoiron/sqlx v1.3.5
github.com/mattn/go-sqlite3 v1.14.22

[cronjob]
github.com/robfig/cron/v3 v3.0.0
```

Untuk membantu penjadwalan pada scheduler saya menggunakan [crontab.guru](https://crontab.guru/)

## Struktur folder
```
.
├── database
│  ├── connection.go
│  ├── entity.go
│  └── model.go
├── go.mod
├── go.sum
└── main.go

```

* modul `database/connection.go` bertugas untuk mengatur koneksi database
* modul `database/entity.go` sebagai abstraksi dari tabel
* modul `database/model.go` bertugas untuk operasi database
* modul `main.go` sebagai modul utama

## Penjelasan Modul

Inisialisasi pembuatan penjadwalan job (scheduler). Menggunakan modul dari `github.com/robfig/cron/v3` untuk proses penjadwalan job. Berikut ini code-nya:
```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/abu-abbas/level_7/database"
	cron "github.com/robfig/cron/v3"
)

type SafeJob struct {
	mu sync.Mutex
	v  string
}

func main() {
	jakTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		panic(err)
	}

	scheduler := cron.New(cron.WithLocation(jakTime))
	defer scheduler.Stop()

	cid, err := scheduler.AddFunc("*/1 * * * *", startUpdateAmountJob)
	if err != nil {
		panic(err)
	}

	fmt.Printf(
		"scheduler id: %d started at: %s\n\n",
		cid,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	go scheduler.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

```

pada code diatas saya membuat penjadwalan job yang akan menjalankan proses `startUpdateAmountJob` dengan interval 1 menit sekali menggunakan cron schedule expressions `"*/1 * * * *"`.

Dan pada proses `startUpdateAmountJob` saya menggunakan `goroutines` untuk melakukan proses update amount pada database. Terlihat seperti code dibawah ini:
```go
func startUpdateAmountJob() {
	newJob := &SafeJob{}
	go newJob.UpdateAmount()
}

func (job *SafeJob) UpdateAmount() {
	job.mu.Lock()
	job.v = "update_amount"

	amount := 100000
	model := database.Item{}
	entity := database.ItemEntity{Amount: amount}

	res, err := model.UpdateAmount(entity)
	if err != nil {
		panic(err)
	}

	if res > 0 {
		fmt.Printf("row/s affected: %d\n", res)
		item, err := model.Get()
		if err != nil {
			panic(err)
		}

		for _, v := range item {
			fmt.Println(v.ToString())
		}

		fmt.Printf("update amount success on: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))
	}

	job.mu.Unlock()
}

```

Karena schedule yang saya buat cukup singkat yaitu permenit sekali, saya menggunakan fitur _mutual exclusion_ untuk memastikan hanya satu `goroutines` yang dijalankan saat nantinya ada proses yang membutuhkan waktu lebih panjang dari interval schedule yang saya buat.

Lanjut menjalankan program scheduler.
Berikut ini adalah ilustrasi dari program scheduler yang saya buat:
![Scheduler illustration.](https://raw.githubusercontent.com/abu-abbas/HSI.sandbox/39c710d6dd3631427d4e2348403dd0268a0cc40d/go/tugas/level_7/images/scheduler.png "Scheduler illustration.")

Point 1: Ilustrasi saat cronjob pertama kali dijalankan. Saat interval 1 menit dalam database ditemukan record yang amount-nya != 100000. Maka dilakukan operasi `update amout`

setelah itu, saya melakukan `proses insert` (menggunakan modul pada tugas level 5) yaitu item dengan HP dan amount 1000:
![Scheduler illustration.](https://raw.githubusercontent.com/abu-abbas/HSI.sandbox/39c710d6dd3631427d4e2348403dd0268a0cc40d/go/tugas/level_7/images/demo-insert.png "Scheduler illustration.")

Point 2: Ilustrasi saat cronjob menumukan ada record yang amount-nya != 100000. Data yang ditampilkan adalah seluruh record item pada database.
