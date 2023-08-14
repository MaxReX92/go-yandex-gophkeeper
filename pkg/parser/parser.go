package parser

import (
	"fmt"
	"strconv"
	"time"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

// ToFloat64 parse strings to float64.
func ToFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// ToInt32 parse string to int32.
func ToInt32(str string) (int32, error) {
	result, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}

	return int32(result), nil
}

// ToInt64 parse string to int64.
func ToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// FloatToString convert float64 to string.
func FloatToString(num float64) string {
	return fmt.Sprintf("%g", num)
}

// IntToString convert int64 to string.
func Int32ToString(num int32) string {
	return strconv.Itoa(int(num))
}

// IntToString convert int64 to string.
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func ToTime(str string) (time.Time, error) {
	result, err := time.Parse("01/06", str)
	if err != nil {
		return result, logger.WrapError("parse time value", ErrInvalidFormat)
	}

	return result, nil
}
