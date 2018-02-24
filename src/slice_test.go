package spec

import (
  "testing"
  "sort"
  "github.com/stretchr/testify/assert"
)

func TestSortStrings(t *testing.T) {
  strings := []string{"a", "c", "b"}
  sort.Strings(strings)
  assert.EqualValues(t, strings, []string{"a", "b", "c"})
  assert.True(t, sort.StringsAreSorted(strings))
}

func TestSortInts(t *testing.T) {
  ints := []int{3, 1, 2}
  sort.Ints(ints)
  assert.EqualValues(t, ints, []int{1, 2, 3})
  assert.True(t, sort.IntsAreSorted(ints))
}
