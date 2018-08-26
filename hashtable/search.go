package hashtable

// Count the number of occurences for a specific word in the given hashtable
func Count(hashTable []map[string]int, word string) (count int, ok bool) {
	hash := Hash(word)

	count, ok = hashTable[hash][word]

	return
}
