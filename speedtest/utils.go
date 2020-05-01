package speedtest

import (
	"math/rand"
	"sort"
)

// http://stackoverflow.com/a/12810288
type randomDataMaker struct {
	src rand.Source
}

func (r *randomDataMaker) Read(p []byte) (n int, err error) {
	todo := len(p)
	offset := 0
	for {
		val := int64(r.src.Int63())
		for i := 0; i < 8; i++ {
			p[offset] = byte(val)
			todo--
			if todo == 0 {
				return len(p), nil
			}
			offset++
			val >>= 8
		}
	}
}

/*
SortServerBy sorts servers by
*/
type SortServerBy func(s1, s2 *TestServer) bool

/*
SortServer sorts server
*/
func (by SortServerBy) SortServer(servers []TestServer) {
	ss := &serverSorter{
		servers: servers,
		by:      by, // The Sort method's receiver is the function (closure) that defines the sort order.
	}
	sort.Sort(ss)
}

// serverSorter joins a By function and a slice of servers to be sorted.
type serverSorter struct {
	servers []TestServer
	by      func(s1, s2 *TestServer) bool // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (s *serverSorter) Len() int {
	return len(s.servers)
}

// Swap is part of sort.Interface.
func (s *serverSorter) Swap(i, j int) {
	s.servers[i], s.servers[j] = s.servers[j], s.servers[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *serverSorter) Less(i, j int) bool {
	return s.by(&s.servers[i], &s.servers[j])
}
