package suffix

import (
	"sort"
)

type Suffix interface {
	DistinctSubCount() int
	DistinctSub() [][]byte
	SubCount() int
	LongestRepeatedSubs() [][]byte
}

type array struct {
	txt []byte
	sa  []int
	lcp []int
}

func NewArray(txt []byte) Suffix {
	a := &array{txt: txt}
	a.sa = a.newArray()
	a.lcp = a.newLcp()

	return a
}

// LongestRepeatedSubs returns the longest
// repeated substring from a.txt
func (a *array) LongestRepeatedSubs() [][]byte {
	var max, key int
	lrs := make([][]byte, 0)

	for k, v := range a.lcp {
		if v > max {
			max = v
			key = k
		}
	}

	lrs = append(lrs, a.txt[a.sa[key]:a.sa[key]+max])

	for i := key + 1; i < len(a.lcp); i++ {
		if a.lcp[i] == max {
			lrs = append(lrs, a.txt[a.sa[i]:a.sa[i]+max])
		}
	}
	return lrs
}

// DistinctSub returns all the distinct substrings of a.txt
func (a *array) DistinctSub() [][]byte {
	n := a.DistinctSubCount()
	dist := make([][]byte, n)

	var x, k int
	for i, n := range a.sa {
		x = a.lcp[i] + n

		for x < len(a.sa) {
			dist[k] = a.txt[n : x+1]
			k++
			x += 1
		}
	}

	return dist
}

// SubCount returns the total substring count in a.txt
func (a *array) SubCount() int {
	return (len(a.txt) * (len(a.txt) + 1)) / 2
}

// DistinctSubCount returns the total distinct substrings in a.txt
func (a *array) DistinctSubCount() int {
	var dup int

	for _, v := range a.lcp {
		dup += v
	}

	return a.SubCount() - dup
}

type suffix struct {
	Index int
	Rank  [2]int
}

func (a *array) newArray() []int {
	n := len(a.txt)
	suffixes := make([]suffix, n)

	for i := range a.txt {
		suffixes[i].Index = i
		suffixes[i].Rank[0] = int(a.txt[i]) - 'a'

		if i+1 < n {
			suffixes[i].Rank[1] = int(a.txt[i+1]) - 'a'
		} else {
			suffixes[i].Rank[1] = -1
		}
	}

	sortFunc := func(i, j int) bool {
		if suffixes[i].Rank[0] == suffixes[j].Rank[0] {
			return suffixes[i].Rank[1] < suffixes[j].Rank[1]
		}

		return suffixes[i].Rank[0] < suffixes[j].Rank[0]
	}

	sort.Slice(suffixes, sortFunc)

	ind := make([]int, n)

	for k := 4; k < 2*n; k *= 2 {
		var rank int
		prev_rank := suffixes[0].Rank[0]
		suffixes[0].Rank[0] = rank
		ind[suffixes[0].Index] = 0

		for i := 1; i < n; i++ {
			if suffixes[i].Rank[0] == prev_rank &&
				suffixes[i].Rank[1] == suffixes[i-1].Rank[1] {
				prev_rank = suffixes[i].Rank[0]
				suffixes[i].Rank[0] = rank
			} else {
				prev_rank = suffixes[i].Rank[0]
				rank += 1
				suffixes[i].Rank[0] = rank
			}
			ind[suffixes[i].Index] = i
		}

		for i := 0; i < n; i++ {
			nextindex := suffixes[i].Index + k/2
			if nextindex < n {
				suffixes[i].Rank[1] = suffixes[ind[nextindex]].Rank[0]
			} else {
				suffixes[i].Rank[1] = -1
			}
		}

		sort.Slice(suffixes, sortFunc)
	}

	suffixArr := make([]int, n)
	for i := 0; i < n; i++ {
		suffixArr[i] = suffixes[i].Index
	}

	return suffixArr
}

func (a *array) newLcp() []int {
	n := len(a.sa)
	lcp := make([]int, n)

	var common, s1, s2 int

	for i := 1; i < n; i++ {
		s1 = a.sa[i-1]
		s2 = a.sa[i]

		for s1 < len(a.txt) && s2 < len(a.txt) && a.txt[s1] == a.txt[s2] {
			common++
			s1++
			s2++
		}
		lcp[i] = common
		common = 0
	}
	return lcp
}
