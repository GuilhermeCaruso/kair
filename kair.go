//
package Kair

import (
	"time"
)

var (
	units  = []string{"LT", "LTS", "L", "l", "LL", "ll", "LLL", "lll", "LLLL", "llll"}
	months = map[int]time.Month{
		1:  time.January,
		2:  time.February,
		3:  time.March,
		4:  time.April,
		5:  time.May,
		6:  time.June,
		7:  time.July,
		8:  time.August,
		9:  time.September,
		10: time.October,
		11: time.November,
		12: time.December,
	}
)

type KairStruct struct {
	Time time.Time
}

func (k KairStruct) KFormat(format string) string {
	timeLint := k.Time
	var timeriz string
	switch format {
	case units[0]:
		timeriz = timeLint.Format("3:04 PM")
	case units[1]:
		timeriz = timeLint.Format("3:04:05 PM")
	case units[2]:
		timeriz = timeLint.Format("02/01/2006")
	case units[3]:
		timeriz = timeLint.Format("02/1/2006")
	case units[4]:
		timeriz = timeLint.Format("January 2, 2006")
	case units[5]:
		timeriz = timeLint.Format("Jan 2, 2006")
	case units[6]:
		timeriz = timeLint.Format("January 2, 2006 3:04 PM")
	case units[7]:
		timeriz = timeLint.Format("Jan 2, 2006 3:04 PM")
	case units[8]:
		timeriz = timeLint.Format("Monday, January 2, 2006 3:04 PM")
	case units[9]:
		timeriz = timeLint.Format("Mon, Jan 2, 2006 3:04 PM")
	default:
		timeriz = timeLint.String()
	}
	return timeriz
}

func Now() KairStruct {
	var k KairStruct
	k.Time = time.Now()
	return k
}

func KDate(day int, month int, year int) KairStruct {
	var k KairStruct
	k.Time = time.Date(year, months[month], day, 0, 0, 0, 0, time.UTC)
	return k
}

func KDateTime(day int, month int, year int, hour int, min int, sec int) KairStruct {
	var k KairStruct
	k.Time = time.Date(year, months[month], day, hour, min, sec, 0, time.UTC)
	return k
}
