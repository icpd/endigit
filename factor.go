package endigit

import (
	"math/rand"
	"time"
)

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
