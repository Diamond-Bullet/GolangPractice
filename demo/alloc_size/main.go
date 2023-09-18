package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
	"unsafe"
)

func main() {
	hs := []*map[Key]Value{}
	n := 1
	before := Alloc()
	for i := 0; i < n; i++ {
		h := map[Key]Value{}
		hs = append(hs, &h)
	}
	after := Alloc()
	emptyPerMap := float64(after-before) / float64(n)
	fmt.Printf("Bytes used for %d empty maps: %d, bytes/map %.1f\n", n, after-before, emptyPerMap)
	hs = nil

	k := 400000
	for p := 1; p < 2; p++ {
		before = Alloc()
		for i := 0; i < n; i++ {
			h := map[Key]Value{}
			for j := 0; j < k; j++ {
				h[Key{
					Exchange: randomString(8),
					Symbol:   randomString(8),
				}] = Value{
					Tpe:      randomString(8),
					Property: randomString(8),
					Name:     randomString(20),
				}
			}
			hs = append(hs, &h)
		}
		after = Alloc()
		fullPerMap := float64(after-before) / float64(n)
		fmt.Printf("Bytes used for %d maps with %d entries: %d, bytes/map %.1f\n", n, k, after-before, fullPerMap)
		fmt.Printf("Bytes per entry %.1f\n", (fullPerMap-emptyPerMap)/float64(k))
		k *= 2
	}
}

func PrintSingle[K comparable](v K) {
	fmt.Println(v)
}

var hs = []*map[Key]Value{}

func Alloc() uint64 {
	var stats runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats)
	return stats.Alloc - uint64(unsafe.Sizeof(hs[0]))*uint64(len(hs))
}

type Key struct {
	Exchange, Symbol string
}

type Value struct {
	Tpe, Property, Name string
}

func randomString(n int) string {
	source := rand.New(rand.NewSource(time.Now().UnixNano()))
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = byte(source.Intn(128))
	}
	return string(res)
}
