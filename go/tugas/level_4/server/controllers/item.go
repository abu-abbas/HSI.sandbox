package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/abu-abbas/level_4/server/entity"
	"github.com/abu-abbas/level_4/server/handler"
	"github.com/abu-abbas/level_4/server/model"
	"github.com/go-chi/chi"

	log "github.com/sirupsen/logrus"
)

type Item struct {
	model model.Item
	res   handler.JsonBody
}

func (i *Item) Index(w http.ResponseWriter, r *http.Request) {
	items, err := i.model.Get()

	if err != nil {
		log.Errorf("gagal mendapatkan items, error: %v\n", err)
		i.res.HttpStatus = "error"
		i.res.HttpCode = http.StatusInternalServerError
		i.res.Message = "Terjadi kesalahan pada server. Silakan hubungi Admin"
	} else {
		i.res.HttpStatus = "success"
		i.res.HttpCode = http.StatusOK
		i.res.Payload = items
	}

	handler.JsonResponse(w, i.res)
}

func (i *Item) findById(id int64, w http.ResponseWriter, r *http.Request) {
	fetch, err := i.model.FindById(id)

	if err != nil {
		log.Errorf("gagal mendapatkan item, error: %v\n", err)
		i.res.HttpStatus = "error"
		i.res.HttpCode = http.StatusInternalServerError
		i.res.Message = "Terjadi kesalahan pada server. Silakan hubungi Admin"
	} else {
		i.res.HttpStatus = "success"
		i.res.HttpCode = http.StatusOK
		i.res.Payload = fetch
	}

	handler.JsonResponse(w, i.res)
}

func (i *Item) Create(w http.ResponseWriter, r *http.Request) {
	entity := entity.Item{}

	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		log.Errorf("gagal saat parsing item, error: %v\n", err)
		i.res.HttpStatus = "error"
		i.res.HttpCode = http.StatusExpectationFailed
		i.res.Message = "Terjadi kesalahan saat parsing data"
		handler.JsonResponse(w, i.res)
	} else {
		itemId, err := i.model.Create(entity)
		if err != nil {
			log.Errorf("gagal saat insert item, error: %v\n", err)
			i.res.HttpStatus = "error"
			i.res.HttpCode = http.StatusInternalServerError
			i.res.Message = "Terjadi kesalahan saat insert data"
			handler.JsonResponse(w, i.res)
		} else {
			i.findById(itemId, w, r)
		}
	}
}

func (i *Item) Edit(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Errorf("gagal saat parsing id, error: %v\n", err)
		i.res.HttpStatus = "error"
		i.res.HttpCode = http.StatusInternalServerError
		i.res.Message = "Terjadi kesalahan parsing id"
		handler.JsonResponse(w, i.res)

		return
	}

	entity := entity.Item{Id: int64(id)}
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		log.Errorf("gagal saat parsing item, error: %v\n", err)
		i.res.HttpStatus = "error"
		i.res.HttpCode = http.StatusExpectationFailed
		i.res.Message = "Terjadi kesalahan saat parsing data"
		handler.JsonResponse(w, i.res)

		return
	}

	_, err = i.model.UpdateItemStatus(entity)
	if err != nil {
		log.Errorf("gagal saat update item, error: %v\n", err)
		i.res.HttpStatus = "error"
		i.res.HttpCode = http.StatusInternalServerError
		i.res.Message = "Terjadi kesalahan saat update data"
		handler.JsonResponse(w, i.res)

		return
	}

	i.findById(int64(id), w, r)
}

func (i *Item) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Errorf("gagal saat parsing id, error: %v\n", err)
		i.res.HttpStatus = "error"
		i.res.HttpCode = http.StatusInternalServerError
		i.res.Message = "Terjadi kesalahan parsing id"
		handler.JsonResponse(w, i.res)
	} else {
		_, err = i.model.DeleteItemById(int64(id))

		if err != nil {
			log.Errorf("gagal saat update item, error: %v\n", err)
			i.res.HttpStatus = "error"
			i.res.HttpCode = http.StatusInternalServerError
			i.res.Message = "Terjadi kesalahan saat update data"
			handler.JsonResponse(w, i.res)
		} else {
			i.res.HttpStatus = "success"
			i.res.HttpCode = http.StatusOK
			i.res.Message = "Data berhasil dihapus"
			handler.JsonResponse(w, i.res)
		}
	}
}
