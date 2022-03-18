package mycmp

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"golang.org/x/exp/constraints"
)

func Diff[T any](a, b T) string {
	return cmp.Diff(a,b)
}

func SortDiff[T constraints.Ordered](a, b []T) string {
	opt := cmpopts.SortSlices(func(i, j T) bool {
        return i < j
    })
	return cmp.Diff(a,b,opt)
}

