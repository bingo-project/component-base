package util

import (
	cryptoRand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	"github.com/spf13/cast"
)

func Empty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

func RandomNumber(length int) string {
	table := [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, length)
	n, err := io.ReadAtLeast(cryptoRand.Reader, b, length)
	if n != length {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func RandomNumberRange(min float64, max float64, precision int) float64 {
	rand.Seed(time.Now().UnixNano())

	data := min + rand.Float64()*(max-min)

	value := strconv.FormatFloat(data, 'f', precision, 64)

	return cast.ToFloat64(value)
}

func RandomWeightedIndex(weights []float32) int {
	if len(weights) == 1 {
		return 0
	}
	var sum float32 = 0.0
	for _, w := range weights {
		sum += w
	}

	rand.Seed(time.Now().UnixNano())
	r := rand.Float32() * sum
	var t float32 = 0.0
	for i, w := range weights {
		t += w
		if t > r {
			return i
		}
	}

	return len(weights) - 1
}

func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}

	return ""
}

func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
