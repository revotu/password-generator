package pkg

import (
	"fmt"
	"testing"
)

func Test_PasswordGenerator_DigitsOnly(t *testing.T) {
	pwdGen := New()
	fmt.Println(pwdGen.DigitsOnly(6))
}

func Test_PasswordGenerator_CharsOnly(t *testing.T) {
	pwdGen := New()
	fmt.Println(pwdGen.CharsOnly(6))
}