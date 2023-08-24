package endigit

import (
	"bytes"
	"errors"
	"strconv"
	"strings"

	"github.com/spf13/cast"
)

type Config struct {
	StrLen  int      // encode string length
	LenMask string   // in encode string mark number length
	Spool   []string // encode string string pool
}

type Digit struct {
	strLen  int
	lenMask string
	spool   []string
}

// Encode encodes a given number into a string.
//
// It takes an integer number as input and returns a string representation of the encoded number.
// An error is returned if the number is negative or too large to be encoded.
func (d *Digit) Encode(num int) (string, error) {
	if num < 0 {
		return "", errors.New("num must be positive")
	}

	numLen := digitLen(num)
	if numLen+1 > d.strLen {
		return "", errors.New("num too large")
	}

	var (
		numStr  = strconv.Itoa(num)
		builder struct {
			hits int
			buf  bytes.Buffer
		}
	)

	for ; builder.hits < numLen; builder.hits++ {
		pos := byteToInt(numStr[builder.hits])
		char := d.spool[builder.hits][pos]
		builder.buf.WriteByte(char)
	}

	lastChar := byteToInt(numStr[builder.hits-1])
	for ; builder.hits < d.strLen-1; builder.hits++ {
		char := d.spool[lastChar][builder.hits%10]
		builder.buf.WriteByte(char)
	}

	builder.buf.WriteByte(d.lenMask[numLen-1])
	return builder.buf.String(), nil
}

// Decode decodes a string into an integer.
func (d *Digit) Decode(str string) (int, error) {
	if len(str) != d.strLen {
		return 0, errors.New("invalid string length")
	}

	var (
		sb     strings.Builder
		numLen = strings.Index(d.lenMask, str[len(str)-1:]) + 1
	)
	for i := 0; i < numLen; i++ {
		pos := strings.Index(d.spool[i], str[i:i+1])
		sb.WriteString(strconv.Itoa(pos))
	}

	return strconv.Atoi(sb.String())
}

func byteToInt(b byte) int {
	return cast.ToInt(string(b))
}

func digitLen(n int) int {
	if n == 0 {
		return 1
	}

	var count int
	for n > 0 {
		n /= 10
		count++
	}

	return count
}

// NewDigit creates a new Digit instance based on the provided configuration.
//
// It takes a Config struct as a parameter, which contains the following fields:
//   - StrLen: the length of the encoded string (required)
//   - LenMask: the mask used to generate the integer of length (required)
//   - Spool: the set of characters that can be used to construct the integer (required)
//
// It returns a pointer to a Digit instance and an error. The error is non-nil
// if any of the required fields are missing or if the length of LenMask or Spool
// does not match the value of StrLen.
func NewDigit(config Config) (*Digit, error) {

	if config.StrLen == 0 {
		return nil, errors.New("config.StrLen is required")
	}

	if config.LenMask == "" {
		return nil, errors.New("config.LenMask is required")
	}

	m := make(map[int32]struct{})
	for _, b := range config.LenMask {
		if _, ok := m[b]; ok {
			return nil, errors.New("config.LenMask must not contain duplicate characters")
		}

		m[b] = struct{}{}
	}

	if config.Spool == nil {
		return nil, errors.New("config.Spool is required")
	}

	if len(config.LenMask) != config.StrLen-1 {
		return nil, errors.New("config.LenMask length must be equal to config.StrLen-1")
	}

	if len(config.Spool) != config.StrLen {
		return nil, errors.New("config.Spool length must be equal to config.StrLen")
	}

	for _, s := range config.Spool {
		if len(s) != 10 {
			return nil, errors.New("each element of config.Spool must have a length of 10")
		}

		m := make(map[int32]struct{})
		for _, b := range s {
			if _, ok := m[b]; ok {
				return nil, errors.New("each element of config.Spool must not contain duplicate characters")
			}

			m[b] = struct{}{}
		}
	}

	return &Digit{
		strLen:  config.StrLen,
		lenMask: config.LenMask,
		spool:   config.Spool,
	}, nil
}
