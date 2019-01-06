package suffix

import (
	"bytes"
	"math/rand"
	"testing"
	"time"
)

func BenchmarkArray_DistinctSub(b *testing.B) {
	b.StopTimer()
	var bb bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 5000; i++ {
		bb.WriteRune(rune(rand.Intn(26) + 97))
	}
	str := bb.String()
	b.StartTimer()
	sa := NewArray([]byte(str))

	for i := 0; i < b.N; i++ {
		sa.DistinctSub()
	}
	b.StopTimer()
}

func TestArray_DistinctSub(t *testing.T) {
	sa := NewArray([]byte("azaza"))
	sub := sa.DistinctSub()

	if len(sub) != 9 {
		t.Errorf("Distinct substrings should be [a az aza azaz azaza z za zaz zaza]   %+v returned\n", sub)
	}

	r := [][]byte{
		{97},
		{97, 122},
		{97, 122, 97},
		{97, 122, 97, 122},
		{97, 122, 97, 122, 97},
		{122},
		{122, 97},
		{122, 97, 122},
		{122, 97, 122, 97},
	}

	for kk, vv := range sub {
		for k, v := range vv {
			if v != r[kk][k] {
				t.Errorf("substring %v should be %v %v given", kk, string(vv), string(r[kk]))
			}
		}
	}
}

func TestArray_DistinctSubCount(t *testing.T) {
	sa := NewArray([]byte("azaza"))
	c := sa.DistinctSubCount()

	if c != 9 {
		t.Errorf("expected 9 got %v\n", c)
	}
}

func TestArray_SubCount(t *testing.T) {
	sa := NewArray([]byte("azaza"))
	c := sa.SubCount()

	if c != 15 {
		t.Errorf("expected 15 got %v\n", c)
	}
}

func TestArray_LongestRepeatedSubs(t *testing.T) {
	sa := NewArray([]byte("abracadabra"))
	lrs := sa.LongestRepeatedSubs()

	if len(lrs) != 1 {
		t.Errorf("expected 1 longest repeated substring got %v\n", len(lrs))
	}

	if string(lrs[0]) != "abra" {
		t.Errorf("expected 'abra' got %v\n", string(lrs[0]))
	}

	sa = NewArray([]byte("ababbaabaa"))
	lrs = sa.LongestRepeatedSubs()

	if len(lrs) != 2 {
		t.Errorf("expected 2 longest repeated substring got %v\n", len(lrs))
	}

	if string(lrs[0]) != "aba" {
		t.Errorf("expected 'aba' got %v\n", string(lrs[0]))
	}

	if string(lrs[1]) != "baa" {
		t.Errorf("expected 'baa' got %v\n", string(lrs[0]))
	}
}
