package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

var (
	readString func() string
	readBytes  func() []byte
	stdout     *bufio.Writer
)

func init() {
	readString, readBytes = newReadString(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
}
func newReadString(ior io.Reader) (func() string, func() []byte) {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)
	f1 := func() string {
		if !r.Scan() {
			panic("panic")
		}
		return r.Text()
	}
	f2 := func() []byte {
		if !r.Scan() {
			panic("panic")
		}
		return r.Bytes()
	}
	return f1, f2
}
func readInt() int {
	return int(readInt64())
}
func readInt64() int64 {
	i, err := strconv.ParseInt(readString(), 0, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}
func readFloat64() float64 {
	f, err := strconv.ParseFloat(readString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

type Restaurat struct {
	index int
	City  string
	Point int
}

type Restaurats []Restaurat

func (r Restaurats) Len() int {
	return len(r)
}

func (r Restaurats) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type ByCity struct {
	Restaurats
}

func (b ByCity) Less(i, j int) bool {
	// fmt.Println(b.Restaurats[i].City == b.Restaurats[j].City)
	return b.Restaurats[i].City < b.Restaurats[j].City
}

type ByPoint struct {
	Restaurats
}

func (b ByPoint) Less(i, j int) bool {
	return b.Restaurats[i].Point > b.Restaurats[j].Point
}

type ByCityByPoint struct {
	Restaurats
}

func (b ByCityByPoint) Less(i, j int) bool {
	if b.Restaurats[i].Point > b.Restaurats[j].Point {
		return b.Restaurats[i].City <= b.Restaurats[j].City
	}
	return b.Restaurats[i].City < b.Restaurats[j].City
}

func main() {
	n := readInt()
	restaurants := make(Restaurats, n)

	for i := 0; i < n; i++ {
		restaurants[i].index = i + 1
		restaurants[i].City = readString()
		restaurants[i].Point = readInt()
	}

	// sort.Sort(ByPoint{restaurants})
	// sort.Sort(ByCity{restaurants})
	sort.Sort(ByCityByPoint{restaurants})

	// sort.Slice(restaurants, func(i, j int) bool {
	// 	p, q := restaurants[i], restaurants[j]
	// 	if p.Point > q.Point {
	// 		return p.City <= q.City
	// 	}
	// 	return false
	// })

	for i := 0; i < n; i++ {
		// fmt.Println(restaurants[i].index)
		fmt.Println(restaurants[i].index)
	}
}
