package main

import (
	"math/rand"
	"sort"
)

type Ints []int

func (s Ints) Reduce(initial interface{}, fn func(value interface{}, elem int) interface{}) (ret interface{}) {
	ret = initial
	for _, elem := range s {
		ret = fn(ret, elem)
	}
	return
}

func (s Ints) Map(fn func(int) int) (ret Ints) {
	for _, elem := range s {
		ret = append(ret, fn(elem))
	}
	return
}

func (s Ints) Filter(filter func(int) bool) (ret Ints) {
	for _, elem := range s {
		if filter(elem) {
			ret = append(ret, elem)
		}
	}
	return
}

func (s Ints) All(predict func(int) bool) (ret bool) {
	ret = true
	for _, elem := range s {
		ret = predict(elem) && ret
	}
	return
}

func (s Ints) Any(predict func(int) bool) (ret bool) {
	for _, elem := range s {
		ret = predict(elem) || ret
	}
	return
}

func (s Ints) Each(fn func(e int)) {
	for _, elem := range s {
		fn(elem)
	}
}

func (s Ints) Shuffle() {
	for i := len(s) - 1; i >= 1; i-- {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func (s Ints) Sort(cmp func(a, b int) bool) {
	sorter := sliceSorter{
		l: len(s),
		less: func(i, j int) bool {
			return cmp(s[i], s[j])
		},
		swap: func(i, j int) {
			s[i], s[j] = s[j], s[i]
		},
	}
	_ = sorter.Len
	_ = sorter.Less
	_ = sorter.Swap
	sort.Sort(sorter)
}

type sliceSorter struct {
	l    int
	less func(i, j int) bool
	swap func(i, j int)
}

func (t sliceSorter) Len() int {
	return t.l
}

func (t sliceSorter) Less(i, j int) bool {
	return t.less(i, j)
}

func (t sliceSorter) Swap(i, j int) {
	t.swap(i, j)
}

func (s Ints) Clone() (ret Ints) {
	ret = make([]int, len(s))
	copy(ret, s)
	return
}
