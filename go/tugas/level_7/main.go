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
