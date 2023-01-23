package src

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type Date struct {
	Year  uint
	Month uint
	Date  uint
}

func (self Date) String() string {
	return fmt.Sprintf("%d-%02d-%02d", self.Year, self.Month, self.Date)
}

func ParseDate(s string) (*Date, error) {
	reg := regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`)
	match := reg.FindAllStringSubmatch(s, -1)
	if match == nil {
		return nil, errors.New("Invalid Format")
	}

	if len(match) < 1 {
		return nil, errors.New("Invalid Format")
	}
	if len(match[0]) < 4 {
		return nil, errors.New("Invalid Format")
	}

	y, _ := strconv.ParseUint(match[0][1], 10, 32)
	m, _ := strconv.ParseUint(match[0][2], 10, 32)
	d, _ := strconv.ParseUint(match[0][3], 10, 32)

	date := Date{
		Year:  uint(y),
		Month: uint(m),
		Date:  uint(d),
	}
	return &date, nil
}
