package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"time"
)

func main() {
	// 時間関係
	t := time.Now()
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month())

	// 正規表現
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	// ソート
	i := []int{5, 3, 2, 1, 4}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)

	// iota
	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3)

	// ioutil
	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
