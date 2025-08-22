package endigit

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDigit(t *testing.T) {
	strLen := 5
	lenMask, spool := factor(strLen)

	digit, err := NewDigit(Config{
		StrLen:  strLen,
		LenMask: lenMask,
		Spool:   spool,
	})
	if err != nil {
		t.Fatal(err)
	}

	num := 1291
	encode, err := digit.Encode(num)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encode)

	decode, err := digit.Decode(encode)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, num, decode)
	t.Log(decode)
}

func factor(strLen int) (string, []string) {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})

	mask := string(letters[:strLen-1])

	var spool []string
	for i := 0; i < strLen; i++ {
		rand.Shuffle(len(letters), func(i, j int) {
			letters[i], letters[j] = letters[j], letters[i]
		})

		spool = append(spool, string(letters[:10]))
	}

	return mask, spool
}
