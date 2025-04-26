package main

import {
	"fmt"
	"github.com/hriday-13th/bloom-filter"
}

func main() {
	bf := BloomFilter.New(1000, 3)

	elements := []string{"apple", "banana", "cherry"}
	
	for _, e := range elements {
		bf.Add([]byte(e))
	}

	// Check for membership
	fmt.Println("Contains 'apple':", bf.Contains([]byte("apple")))
	fmt.Println("Contains 'grape':", bf.Contains([]byte("grape")))

	// Print statistics
	fmt.Printf("Number of elements: %d\n", bf.Count())
	fmt.Printf("Estimated false positive rate: %.4f\n", bf.EstimatedFalsePositiveRate())

	bf2 := bloomfilter.New(1000, 3)
	bf2.Add([]byte("date"))
	union := bf.Union(bf2)
	fmt.Println("Union contains 'apple':", union.Contains([]byte("apple")))
	fmt.Println("Union contains 'date':", union.Contains([]byte("date")))

	serialized := bf.Serialize()
	deserialized := bloomfilter.Deserialize(serialized)
	fmt.Println("Deserialized filter contains 'banana':", deserialized.Contains([]byte("banana")))
}