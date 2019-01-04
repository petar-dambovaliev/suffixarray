package suffix

import (
	"sort"
)

type Suffix interface {
	DistinctSubCount() int
	DistinctSub() []string
	SubCount() int
}

type array struct {
	txt  string
	sa   []int
	lcp  []int
	subs []string
}

func NewArray(txt string) Suffix {
	a := &array{txt: txt}
	a.newArray()
	a.newLcp()
	a.genAllSubs()

	return a
}

func (a *array) genAllSubs() {
	cnt := a.SubCount()
	allSubs := make([]string, cnt)

	k := 0

	for i := 0; i < len(a.txt); i++ {
		for j := 0; j < len(a.txt)-i; j++ {
			allSubs[k] = a.txt[j : j+i+1]
			k++
		}
	}

	a.subs = allSubs
}

func (a *array) DistinctSub() []string {
	n := a.DistinctSubCount()
	dups := make([]string, a.SubCount()-n)
	dist := make([]string, n)

	k := 0

	for i := 1; i < len(a.lcp); i++ {
		for j := 0; j < a.lcp[i]; j++ {
			dups[k] = a.txt[a.sa[i] : a.sa[i]+j+1]
			k++
		}
	}

	var m, ind int
	for _, v := range a.subs {
		m = -1
		for k, vv := range dups {
			if v == vv {
				m = k
			}
		}

		if m > -1 {
			dups[m] = ""
			continue
		}

		dist[ind] = v
		ind++
	}

	return dist
}

func (a *array) SubCount() int {
	return (len(a.txt) * (len(a.txt) + 1)) / 2
}

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

func (a *array) newArray() {
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

	a.sa = suffixArr
}

func (a *array) newLcp() {
	n := len(a.sa)
	lcp := make([]int, n)

	var common, s1, s2 int

	for i := 1; i < n; i++ {
		s1 = a.sa[i-1]
		s2 = a.sa[i]

		for s1 < len(a.txt) && s2 < len(a.txt) && (a.txt)[s1] == (a.txt)[s2] {
			common++
			s1++
			s2++
		}
		lcp[i] = common
		common = 0
	}
	a.lcp = lcp
}