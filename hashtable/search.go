package hashtable

func Search(hashTable []map[string]int, word string) (count int, ok bool) {
	hash := Hash(word)

	count, ok = hashTable[hash][word]

	return
}
