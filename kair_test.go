package Kair

import (
	"fmt"
	"testing"
)

// func TestKairFormatType(t *testing.T) {
// 	aR := Now("LTS")
// 	if aR != "LT" {
// 		t.Fatalf("Expected %s but go %s", "LT", aR)
// 	}
// }

func TestMain(t *testing.T) {
	fmt.Println(Now().KFormat("LT"))
	fmt.Println(KDate(20, 10, 2018).KFormat("L"))
	fmt.Println(KDateTime(20, 10, 2018, 12, 20, 30).KFormat("LLL"))
}
