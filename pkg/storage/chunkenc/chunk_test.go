// Copyright 2021 The Parca Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Copyright 2017 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chunkenc

import (
	"fmt"
	"io"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChunk(t *testing.T) {
	for enc, nc := range map[Encoding]func() Chunk{
		EncXOR:   func() Chunk { return NewXORChunk() },
		EncDelta: func() Chunk { return NewDeltaChunk() },
		EncRLE:   func() Chunk { return NewRLEChunk() },
	} {
		t.Run(fmt.Sprintf("%v", enc), func(t *testing.T) {
			for range make([]struct{}, 1) {
				c := nc()
				testChunk(t, c)
			}
		})
	}
}

func testChunk(t *testing.T, c Chunk) {
	app, err := c.Appender()
	require.NoError(t, err)

	exp := make([]int64, 0, 300)
	v := int64(1243535)

	for i := 0; i < 300; i++ {
		if i%2 == 0 {
			v += int64(rand.Intn(1000000))
		} else {
			v -= int64(rand.Intn(1000000))
		}

		// Start with a new appender every 10th sample. This emulates starting
		// appending to a partially filled chunk.
		if i%10 == 0 {
			app, err = c.Appender()
			require.NoError(t, err)
		}

		app.Append(v)
		exp = append(exp, v)
	}

	// 1. Expand iterator in simple case.
	it1 := c.Iterator(nil)
	res1 := make([]int64, 0, len(exp))
	for it1.Next() {
		res1 = append(res1, it1.At())
	}
	require.NoError(t, it1.Err())
	require.Equal(t, exp, res1)

	// 2. Expand second iterator while reusing first one.
	it2 := c.Iterator(it1)
	res2 := make([]int64, 0, len(exp))
	for it2.Next() {
		res2 = append(res2, it2.At())
	}
	require.NoError(t, it2.Err())
	require.Equal(t, exp, res2)

	// 3. Test iterator Seek.
	mid := uint16(len(exp) / 2)

	it3 := c.Iterator(nil)
	var res3 []int64
	require.Equal(t, true, it3.Seek(mid))
	res3 = append(res3, it3.At())

	for it3.Next() {
		res3 = append(res3, it3.At())
	}
	require.NoError(t, it3.Err())
	require.Equal(t, exp[mid:], res3)
	require.Equal(t, false, it3.Seek(uint16(len(exp))))

	// 4. Append at a given index with 0 leading up to it
	app.AppendAt(310, 42)

	it4 := c.Iterator(nil)
	var res4 []int64
	require.Equal(t, true, it4.Seek(300)) // Seek to where zeros start
	for it4.Next() {
		res4 = append(res4, it4.At())
	}
	require.NoError(t, it4.Err())
	require.Equal(t, []int64{0, 0, 0, 0, 0, 0, 0, 0, 0, 42}, res4)

	require.False(t, it4.Next())
}

// for i in {1..10}; do go test -bench=BenchmarkIterators --benchtime=500000000x ./pkg/storage/chunkenc >> pkg/storage/benchmark/iterator-same.txt; done
// for i in {1..10}; do go test -bench=BenchmarkIterators --benchtime=500000000x ./pkg/storage/chunkenc >> pkg/storage/benchmark/iterator-increasing.txt; done
// for i in {1..10}; do go test -bench=BenchmarkIterators --benchtime=500000000x ./pkg/storage/chunkenc >> pkg/storage/benchmark/iterator-random.txt; done
// benchstat pkg/storage/benchmark/iterator-same.txt pkg/storage/benchmark/iterator-increasing.txt pkg/storage/benchmark/iterator-random.txt >> pkg/storage/benchmark/iterator-benchstat.txt

func BenchmarkIterators(b *testing.B) {
	for _, c := range []struct {
		name     string
		newChunk func() Chunk
	}{{
		name:     "delta",
		newChunk: func() Chunk { return NewDeltaChunk() },
	}, {
		name:     "rle",
		newChunk: func() Chunk { return NewRLEChunk() },
	}, {
		name:     "xor",
		newChunk: func() Chunk { return NewXORChunk() },
	}} {
		b.Run(c.name, func(b *testing.B) {
			benchmarkIterator(b, c.newChunk)
		})
	}
}

func benchmarkIterator(b *testing.B, newChunk func() Chunk) {
	var (
		chunks []Chunk
		c      Chunk
		a      Appender
		err    error

		v    = int64(123456)
		data = make([]int64, 0, b.N)
	)
	for i := 0; i < b.N; i++ {
		// v = rand.Int63n(1_000_000) // random
		v += int64(100) // increasing
		// v = 100 // same
		data = append(data, v)
	}

	for i := 0; i < b.N; {
		if i%250 == 0 {
			if i != 0 {
				chunks = append(chunks, c)
			}
			c = newChunk()
			a, err = c.Appender()
			if err != nil {
				b.Fatalf("get appender: %s", err)
			}
		}

		j := 0
		for _, v := range data {
			if j > 250 {
				break
			}
			a.Append(v)
			i++
			j++
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	res := make([]int64, 0, 1024)

	var it Iterator
	for i := 0; i < len(chunks); i++ {
		c := chunks[i]
		it := c.Iterator(it)

		for it.Next() {
			res = append(res, it.At())
		}
		if it.Err() != io.EOF {
			require.NoError(b, it.Err())
		}
		res = res[:0]
	}
}

// for i in {1..10}; do go test -bench=BenchmarkAppenders --benchtime=500000000x ./pkg/storage/chunkenc >> pkg/storage/benchmark/appender-same.txt; done
// for i in {1..10}; do go test -bench=BenchmarkAppenders --benchtime=500000000x ./pkg/storage/chunkenc >> pkg/storage/benchmark/appender-increasing.txt; done
// for i in {1..10}; do go test -bench=BenchmarkAppenders --benchtime=500000000x ./pkg/storage/chunkenc >> pkg/storage/benchmark/appender-random.txt; done
// benchstat ./pkg/storage/benchmark/appender-same.txt ./pkg/storage/benchmark/appender-increasing.txt ./pkg/storage/benchmark/appender-random.txt

func BenchmarkAppenders(b *testing.B) {
	for _, c := range []struct {
		name     string
		newChunk func() Chunk
	}{{
		name:     "delta",
		newChunk: func() Chunk { return NewDeltaChunk() },
	}, {
		name:     "rle",
		newChunk: func() Chunk { return NewRLEChunk() },
	}, {
		name:     "xor",
		newChunk: func() Chunk { return NewXORChunk() },
	}} {
		b.Run(c.name, func(b *testing.B) {
			benchmarkAppender(b, c.newChunk)
		})
	}
}

func benchmarkAppender(b *testing.B, newChunk func() Chunk) {
	b.ReportAllocs()

	var (
		chunks []Chunk
		c      = newChunk()
		a      Appender
		err    error

		v    = int64(123)
		data = make([]int64, 0, b.N)
	)
	for i := 0; i < b.N; i++ {
		// v = rand.Int63n(1_000_000) // random
		v += int64(100) // increasing
		// v = 100 // same
		data = append(data, v)
	}

	b.ResetTimer()

	for i := 0; i < b.N; {
		if i%250 == 0 {
			chunks = append(chunks, c)
			c = newChunk()
			a, err = c.Appender()
			if err != nil {
				b.Fatalf("get appender: %s", err)
			}
		}

		j := 0
		for _, v := range data {
			if j > 250 {
				break
			}
			a.Append(v)
			i++
			j++
		}
	}

	b.StopTimer()

	totalBytes := 0
	for _, c := range chunks {
		totalBytes += len(c.Bytes())
	}

	b.ReportMetric(float64(totalBytes)/float64(b.N), "disk/op")
}
