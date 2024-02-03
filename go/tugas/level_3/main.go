package main

import (
	"fmt"

	"github.com/abu-abbas/level_3/entity"
	"github.com/abu-abbas/level_3/model"
	"github.com/abu-abbas/level_3/utils"
)

var itemModel 	model.Item
var detailModel	model.Detail

func main() {
	// migrateTable()
	 
	// insertItemAndGet()

	// insertManyItem()

	// insertItemWithDetails()

	// updateItemStatus()

	dropDetailWithId()
}

func migrateTable() {
	// do migrate
	res := itemModel.Migrate()

	// parse result
    if  _, err := res.RowsAffected(); err != nil {
        utils.ErrorCheck(err)
    }

	// do migrate
	res = detailModel.Migrate()

	// parse result
	if _, err := res.RowsAffected(); err != nil {
		utils.ErrorCheck(err)
	}

    fmt.Println("migrasi item dan item_detail table berhasil")
}

func insertItemAndGet() {
	// prepare entity
	handphone := entity.Item{Name: "Handphone", Status: "draft", Amount: 1000}

	// do insert entity
    res, err := itemModel.Create(handphone)
    if err != nil {
        utils.ErrorCheck(err)
    }
	
	// get feedback
    lastId, errLastId := res.LastInsertId()
    if errLastId != nil {
        utils.ErrorCheck(errLastId)
    }
    fmt.Println("insert item berhasil")
	
	// fetching data with lastId
    fetch, errFetch := itemModel.FindById(lastId)
    if errFetch != nil {
        fmt.Println("error: ", errFetch)
    } else {
        println(fetch.ToString())
    }
}

func insertManyItem() {
	items := []entity.Item{
		{Name: "Laptop", Status: "draft", Amount: 2000},
		{Name: "Headphone", Status: "publish", Amount: 500},
	}

	affectedRow := itemModel.CreateMany(items)
	fmt.Printf("affectedRow: %d\n", affectedRow)
}

func insertItemWithDetails() {
	details := [] entity.Detail{
		{Name: "Ban"},
		{Name: "Velg"},
		{Name: "Lampu"},
		{Name: "Spion"},
	}
	
	item := entity.Item{
		Name: "Motor", 
		Status: "sold", 
		Amount: 100000,
		Detail: details,
	}

	res := itemModel.CreateWithDetail(item)

	fetchItem, errItem := itemModel.FindById(res)
	if errItem != nil {
		utils.ErrorCheck(errItem)
	}

	fetchDetail, errDetail := detailModel.FindByItemId(res)
	if errDetail != nil {
		utils.ErrorCheck(errDetail)
	}

	stringDetail := ""
	for i, _ := range fetchDetail {
		stringDetail += fetchDetail[i].ToString() 
	}
	
	fmt.Printf("%s\n%s", fetchItem.ToString(), stringDetail)
}

func updateItemStatus() {
	item := entity.Item{
        Status: "ready",
        Id: 1,
    }
	res, err := itemModel.UpdateItemStatus(item)
	if err != nil {
		utils.ErrorCheck(err)
	}

	affectedRow, errAr := res.RowsAffected()
	if errAr != nil {
		utils.ErrorCheck(errAr)
	}

	fmt.Printf("Data yang berhasil diubah adalah %d baris\n", affectedRow)
}

func dropDetailWithId() {
	res := itemModel.DeleteItemById(3)
	affectedRow, errAr := res.RowsAffected()

    if errAr != nil {
        utils.ErrorCheck(errAr)
    }

	fmt.Printf("Data yang berhasil dihapus adalah %d baris\n", affectedRow)
}
