package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lachlanbird88/hash_table/hashtable"
)

func main() {

	var url string

	fmt.Print("Please enter a URL to count words for: ")
	fmt.Scan(&url)

	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	scanner.Split(bufio.ScanWords)

	buckets := make([]map[string]int, hashtable.NumBuckets)

	for i := 0; i < hashtable.NumBuckets; i++ {
		buckets[i] = make(map[string]int)
	}

	for scanner.Scan() {
		word := scanner.Text()
		hash := hashtable.Hash(word)

		_, ok := buckets[hash][word]

		if ok {
			buckets[hash][word]++
		} else {
			buckets[hash][word] = 1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	var search string

	fmt.Print("Please enter the word you would like to search for: ")
	fmt.Scan(&search)

	count, _ := hashtable.Count(buckets, search)

	fmt.Println("The word", search, "is mentioned ", count, "times")
}
