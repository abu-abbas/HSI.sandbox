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
