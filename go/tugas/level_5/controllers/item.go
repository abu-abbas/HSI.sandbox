package controllers

import (
	"strconv"

	"github.com/abu-abbas/level_5/entity"
	"github.com/abu-abbas/level_5/model"
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	model model.Item
}

func internalServerErr(err error, c *fiber.Ctx) error {
	return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
		"status":  "error",
		"message": "Terjadi kesalahan pada server",
	})
}

func (i *Item) Index(c *fiber.Ctx) error {
	items, err := i.model.Get()
	if err != nil {
		return internalServerErr(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   items,
	})
}

func (i *Item) findById(id int64, c *fiber.Ctx) error {
	fetch, err := i.model.FindById(id)
	if err != nil {
		return internalServerErr(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   fetch,
	})
}

func (i *Item) Create(c *fiber.Ctx) error {
	entity := entity.Item{}
	err := c.BodyParser(&entity)
	if err != nil {
		return internalServerErr(err, c)
	}

	itemId, err := i.model.Create(entity)
	if err != nil {
		return internalServerErr(err, c)
	}

	return i.findById(itemId, c)
}

func (i *Item) Edit(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return internalServerErr(err, c)
	}

	entity := entity.Item{Id: int64(id)}
	err = c.BodyParser(&entity)
	if err != nil {
		return internalServerErr(err, c)
	}

	_, err = i.model.UpdateItemStatus(entity)
	if err != nil {
		return internalServerErr(err, c)
	}

	return i.findById(int64(id), c)
}

func (i *Item) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return internalServerErr(err, c)
	}

	_, err = i.model.DeleteItemById(int64(id))
	if err != nil {
		return internalServerErr(err, c)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Data berhasil dihapus",
	})
}
