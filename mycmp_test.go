package mycmp_test

import (
	"mycmp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)


type testError struct {
	Msg string
}

func (e testError) Error() string {
	return e.Msg
}

func TestCompareDiffSliceInt(t *testing.T) {
	opt := cmpopts.SortSlices(func(i, j int) bool {
        return i < j
    })
	testcases := map[string] struct {
		a []int
		b []int
	}{
		"compare int" : {[]int{1,2,3,4,5}, []int{5,4,3,2,1}},
	}
	for _, tc := range testcases {
		t.Run("compare another type", func(t *testing.T) {
			diff := mycmp.SortDiff(tc.a, tc.b)
			real := cmp.Diff(tc.a, tc.b, opt)
			if diff != real {
				t.Errorf("this func doesn't return correct diff (%v, %v)", diff, real)
			}
		})
	}
}



func TestCompareDiffSliceString(t *testing.T) {
	opt := cmpopts.SortSlices(func(i, j string) bool {
        return i < j
    })
	testcases := map[string] struct {
		a []string
		b []string
	}{
		"compare string" : {[]string{"a","b","c","d","e"}, []string{"e","d","c","b","a"}},
	}
	for _, tc := range testcases {
		t.Run("compare another type", func(t *testing.T) {
			diff := mycmp.SortDiff(tc.a, tc.b)
			real := cmp.Diff(tc.a, tc.b, opt)
			if diff != real {
				t.Errorf("this func doesn't return correct diff (%v, %v)", diff, real)
			}
		})
	}
}
