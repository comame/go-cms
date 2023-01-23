package src

import (
	"testing"
)

func TestDateToString(t *testing.T) {
	date := Date{
		Year:  1970,
		Month: 1,
		Date:  10,
	}

	if date.String() != "1970-01-10" {
		t.Fail()
	}
}

func TestParseDate(t *testing.T) {
	date := "1970-01-10"

	got, err := ParseDate(date)

	if err != nil {
		t.FailNow()
	}

	if got.Year != 1970 {
		t.Fail()
	}
	if got.Month != 1 {
		t.Fail()
	}
	if got.Date != 10 {
		t.Fail()
	}
}

func TestParseDateFails(t *testing.T) {
	date := "1970-1-10"

	_, err := ParseDate(date)
	if err == nil {
		t.Fail()
	}
}
