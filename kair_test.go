package kair

import (
	"fmt"
	"testing"
	"time"
)

//TestPersonalDate - Check KDate func return
func TestPersonalDate(t *testing.T) {
	aR := Date(20, 05, 2018)
	eR := "2018-05-20 00:00:00 +0000 UTC"
	if aR.Time.String() != eR {
		t.Fatalf("Expected %s but go %s", eR, aR)
	} else {
		fmt.Println("TestPersonalDate - OK")
	}
}

//TestPersonalDateTime - Check KDateTime func return
func TestPersonalDateTime(t *testing.T) {
	aR := DateTime(20, 05, 2018, 14, 20, 30)
	eR := "2018-05-20 14:20:30 +0000 UTC"
	if aR.Time.String() != eR {
		t.Fatalf("Expected %s but go %s", eR, aR)
	} else {
		fmt.Println("TestPersonalDateTime - OK")
	}
}

//TestNowGetter - Check KFormat func return
func TestNowGetter(t *testing.T) {
	aR := Now().Format("LLL")
	eR := time.Now().Format("January 2, 2006 3:04 PM")
	if aR != eR {
		t.Fatalf("Expected %s but go %s", eR, aR)
	} else {
		fmt.Println("TestNowGetter - OK")
	}
}

//TestKFormat - Check all possible default format results
func TestKFormat(t *testing.T) {
	date := DateTime(20, 05, 2018, 10, 20, 0)
	testList := map[string]string{
		"LT":   "10:20 AM",
		"LTS":  "10:20:00 AM",
		"L":    "20/05/2018",
		"l":    "20/5/2018",
		"LL":   "May 20, 2018",
		"ll":   "May 20, 2018",
		"LLL":  "May 20, 2018 10:20 AM",
		"lll":  "May 20, 2018 10:20 AM",
		"LLLL": "Sunday, May 20, 2018 10:20 AM",
		"llll": "Sun, May 20, 2018 10:20 AM",
		"":     "2018-05-20 10:20:00 +0000 UTC",
	}

	for i, v := range testList {
		aR := date.Format(i)
		eR := v
		if aR != eR {
			t.Fatalf("Expected %s but go %s", eR, aR)
		}
	}
	fmt.Println("TestKFormat - OK")
}

//CustomFormat - Checks whether custom return is correct
func TestPersonalFormat(t *testing.T) {
	date := DateTime(20, 05, 2018, 10, 20, 0)
	aR := date.CustomFormat("MMM/dd/YY h:m:s")
	eR := "May/20/18 10:20:0"
	if aR != eR {
		t.Fatalf("Expected %s but go %s", eR, aR)
	} else {
		fmt.Println("TestPersonalFormat - OK")
	}
}
