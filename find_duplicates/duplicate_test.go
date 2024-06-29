package findduplicates

import (
	"testing"
)

func TestFindDuplicatesWithMissing(t *testing.T) {
	input := []int{1, 2, 2, 3, 4, 5}
	computed := findDuplicateWithMissing(input)
	t.Logf("%+v", computed)
	if computed != 2 {
		panic("didn't find 2")
	}
	input = []int{1, 1, 2, 3, 4, 5}
	computed = findDuplicateWithMissing(input)
	t.Logf("%+v", computed)
	if computed != 1 {
		panic("didn't find 1")
	}

	input = []int{1, 2, 3, 5, 5, 6}
	computed = findDuplicateWithMissing(input)
	t.Logf("%+v", computed)
	if computed != 5 {
		panic("didn't find 5")
	}
}

func TestFindDuplicates(t *testing.T) {
	input := []int{1, 2, 2, 3, 4, 5}
	computed := findDuplicate(input)
	t.Logf("%+v", computed)
	if computed != 2 {
		panic("didn't find 2")
	}
	input = []int{1, 1, 2, 3, 4, 5}
	computed = findDuplicate(input)
	t.Logf("%+v", computed)
	if computed != 1 {
		panic("didn't find 1")
	}

	input = []int{1, 2, 3, 5, 5, 4}
	computed = findDuplicate(input)
	t.Logf("%+v", computed)
	if computed != 5 {
		panic("didn't find 5")
	}

	input = []int{1, 3, 4, 2, 2}
	computed = findDuplicate(input)
	t.Logf("%+v", computed)
	if computed != 2 {
		panic("didn't find 2")
	}
}
