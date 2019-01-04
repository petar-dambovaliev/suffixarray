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

	for i := 0; i < 500; i++ {
		bb.WriteRune(rune(rand.Intn(26) + 97))
	}
	str := bb.String()
	b.StartTimer()
	sa := NewArray(str)

	for i := 0; i < b.N; i++ {
		sa.DistinctSub()
	}
	b.StopTimer()
}
