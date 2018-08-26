package hashtable

var NumBuckets int = 12

func Hash(word string) int {
	var length int

	for _, val := range word {
		length += int(val)
	}

	return length % NumBuckets
}
