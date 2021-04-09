package arrays_test

import (
	"fmt"
	"github.com/fdaines/arch-go/utils/arrays"
	"testing"
)

func TestContains(t *testing.T) {
	var tests = []struct {
		array []string
		searchFor string
		want bool
	}{
		{[]string{"foo", "foobar", "bar"}, "foobar", true},
		{[]string{"foo", "foobarX", "bar"}, "foobar", false},
		{[]string{}, "foobar", false},
	}

	for index, tt := range tests {
		t.Run(fmt.Sprintf("Search case %d", index+1), func(t *testing.T) {
			ans := arrays.Contains(tt.array, tt.searchFor)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}

}

