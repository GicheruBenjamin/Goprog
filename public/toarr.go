package public

import (
	"strings"
)

func Stringtoarray(s string) []string {
    return strings.Split(s, "")
}
func IntToArray(i int) []int {
    var result []int
    for i > 0 {
        result = append(result, i%10)
        i /= 10
    }
    return result
}

func FloatToArray(f float64) []float64 {
    var result []float64
    for f > 0 {
        result = append(result, f)
        f /= 10
    }
    return result
}