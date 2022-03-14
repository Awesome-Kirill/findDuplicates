package find

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicates(t *testing.T) {

	arr := [][]string{

		{"1014065", "1", "4708", "4709"},

		{"1014071", "1", "4708", "4697"},

		{"1014130", "2", "9416", "8004"},

		{"1014130", "2", "9416", "1412"},

		{"1014238", "1", "4708", "4694"},

		{"1014240", "3", "14124", "14090"},
	}

	good := [][2]int{{0, 2}, {0, 3}, {1, 0}, {1, 1}, {1, 4}, {1, 2}, {1, 3}, {2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4}}
	result := Duplicates(arr)

	//fmt.Println(good)
	if !assert.ElementsMatch(t, result, good) {
		t.Error("Test fail")
	}
}
