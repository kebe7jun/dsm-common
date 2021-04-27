package utils

import (
	"fmt"
	"testing"
)

func TestTimeFormat(t *testing.T) {

	fmt.Println(TimeFormat("1590464690690417"))

	for _, date := range GetDurationDateList("1590464690690417", "1591464690691418") {
		fmt.Println(date)
	}
}
