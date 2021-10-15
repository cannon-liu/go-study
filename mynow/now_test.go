package mynow

import (
	"testing"
	"time"
)

var (
	format          = "2006-01-02 15:04:05.999999999"
	locationCaracas *time.Location
	locationBerlin  *time.Location
	timeCaracas     time.Time
)

func init() {
	var err error
	if locationCaracas, err = time.LoadLocation("America/Caracas"); err != nil {
		panic(err)
	}

	if locationBerlin, err = time.LoadLocation("Europe/Berlin"); err != nil {
		panic(err)
	}

	timeCaracas = time.Date(2016, 1, 1, 12, 10, 0, 0, locationCaracas)
}

func assertT(t *testing.T) func(time.Time, string, string) {
	return func(actual time.Time, expected string, msg string) {
		actualStr := actual.Format(format)
		if actualStr != expected {
			t.Errorf("Failed %s: actial: %v,expected: %v", msg, actualStr, expected)
		}
	}
}

func TestBeginningOf(t *testing.T) {
	assert := assertT(t)

	n := time.Date(2013, 11, 18, 17, 51, 49, 123456789, time.UTC)

	assert(With(n).BeginningOfMinute(), "2013-11-18 17:51:00", "BeginningOfMinute")

}
