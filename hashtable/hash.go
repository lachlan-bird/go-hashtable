package hashtable

// NumBuckets The number of buckets we would like to set in the hashtable. Currently
// hardcoded
const NumBuckets int = 12

// Hash Hashes a word into the number of buckets defined in NumBuckets
// Simple hashing algorithm of summing the UTF8 code point values of
// all the characters within the string
func Hash(word string) int {
	var length int

	for _, val := range word {
		length += int(val)
	}

	return length % NumBuckets
}
