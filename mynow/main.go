package mynow

import "time"

var WeekStartDay = time.Sunday

var TimeFormats = []string{
	"2006", "2006-1", "2006-1-2", "2006-1-2 15", "2006-1-2 15:4", "2006-1-2 15:4:5", "1-2",
	"15:4:5", "15:4", "15",
	"15:4:5 Jan 2, 2006 MST", "2006-01-02 15:04:05.999999999 -0700 MST", "2006-01-02T15:04:05-07:00",
	"2006.1.2", "2006.1.2 15:04:05", "2006.01.02", "2006.01.02 15:04:05", "2006.01.02 15:04:05.999999999",
	"1/2/2006", "1/2/2006 15:4:5", "2006/01/02", "2006/01/02 15:04:05",
	time.ANSIC, time.UnixDate, time.RubyDate, time.RFC822, time.RFC822Z, time.RFC850,
	time.RFC1123, time.RFC1123Z, time.RFC3339, time.RFC3339Nano,
	time.Kitchen, time.Stamp, time.StampMilli, time.StampMicro, time.StampNano,
}

// Config configuration for now package
type Config struct {
	WeekStartDay time.Weekday
	TimeLocation *time.Location
	TimeFormats  []string
}

// DefaultConfig default config
var DefaultConfig *Config

func (config *Config) With(t time.Time) *Now {
	return &Now{
		Time:   t,
		Config: config,
	}
}

// Parse parse string to time based on configuration
func (config *Config) Parse(strs ...string) (time.Time, error) {
	if config.TimeLocation == nil {
		return config.With(time.Now()).Parse(strs...)
	} else {
		return config.With(time.Now().In(config.TimeLocation)).Parse(strs...)
	}
}

// MustParse must parse string to time or will panic
func (config *Config) MustParse(strs ...string) time.Time {
	if config.TimeLocation == nil {
		return config.With(time.Now()).MustParse(strs...)
	} else {
		return config.With(time.Now().In(config.TimeLocation)).MustParse(strs...)
	}
}

// Now now struct
type Now struct {
	time.Time
	*Config
}

// With initialize Now with time
func With(t time.Time) *Now {
	config := DefaultConfig
	if config == nil {
		config = &Config{
			WeekStartDay: WeekStartDay,
			TimeFormats:  TimeFormats,
		}
	}

	return &Now{
		Time:   t,
		Config: config,
	}
}

// New initialize Now with time
func New(t time.Time) *Now {
	return With(t)
}
