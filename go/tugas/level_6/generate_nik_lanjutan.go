package main

import (
	"errors"

	"github.com/abu-abbas/level_6/generator"
)

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
