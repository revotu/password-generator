package pkg

import (
	"bytes"

	"github.com/revotu/password-generator/internal"
)

type PasswordGenerator struct{}

func New() *PasswordGenerator {
	return &PasswordGenerator{}
}

func (pg *PasswordGenerator) DigitsOnly(lenght int) string {
	var buf bytes.Buffer
	digits := "0123456789"

	for i := 0; i < lenght; i++ {
		buf.WriteByte(digits[internal.RandNumber(10)])
	}

	return buf.String()
}

func (pg * PasswordGenerator) CharsOnly(length int) string {
	var buf bytes.Buffer
	chars := "abcdefghijklmnopkrstuvwxyz"
	for i := 0; i < length; i++ {
		buf.WriteByte(chars[internal.RandNumber(26)])
	}
	return buf.String()
}