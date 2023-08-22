package endigit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigit(t *testing.T) {
	strLen := 16
	lenMask, spool := factor(strLen)

	digit, err := NewDigit(Config{
		StrLen:  16,
		LenMask: lenMask,
		Spool:   spool,
	})
	if err != nil {
		t.Fatal(err)
	}

	encode, err := digit.Encode(12)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encode)

	decode, err := digit.Decode(encode)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 12, decode)
	t.Log(decode)
}
