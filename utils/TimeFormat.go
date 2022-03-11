package utils

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type LocalTime time.Time
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = LocalTime(t1)
	return err
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *LocalTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		*t = LocalTime(vt)
	default:
		return nil
	}
	return nil
}

func (t *LocalTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}