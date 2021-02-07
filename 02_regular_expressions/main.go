package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`Hell.?`)
	fmt.Printf("%q\n", re.Find([]byte(`Hello World!`)))
	re = regexp.MustCompile(`Hell.?`)
	fmt.Println(re.Match([]byte(`Hello World!`)))
	re = regexp.MustCompile(`Hell.?`)
	fmt.Printf("%s\n", re.ReplaceAll([]byte("Hello World!"), []byte("Hi")))
	re = regexp.MustCompile(` `)
	subs := re.Split("Hello World!", -1)
	fmt.Println(subs)
	for _, s := range subs {
		fmt.Println(s)
	}
}
