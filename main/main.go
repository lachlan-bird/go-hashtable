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

	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	scanner.Split(bufio.ScanWords)

	buckets := make([]map[string]int, hashtable.NumBuckets)

	for i := 0; i < 12; i++ {
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

	count, _ := hashtable.Search(buckets, "Moby")

	fmt.Println(count)
}
