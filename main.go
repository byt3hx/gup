package main

import(
	"fmt"
	"net/url"
	"bufio"
	"os"
	"sync"
	"flag"
)


func main() {

	var s bool
	var multi bool
	flag.BoolVar(&s, "s" , false , "For a single parameter")
	flag.BoolVar(&multi, "m" , false, "For mutliple parameters")
	flag.Parse()


	var urls []string

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		urls = append(urls, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n")
	}

	results := make(chan string)

	var wg sync.WaitGroup
	for _, link := range urls {
		wg.Add(1)
		go func(link string) {
			u , err := url.Parse(link)
			if err != nil {
				fmt.Println(err)
			}
			if multi == true {
				fmt.Println(u.RawQuery)
			}
			if s == true {
				m, err := url.ParseQuery(u.RawQuery)
				if err == nil {
					for   k := range m {
						fmt.Println(k)	
					}
				}
			}
			defer wg.Done()
		}(link)	
	}
	wg.Wait()
	close(results)	
}
