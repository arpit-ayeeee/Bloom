package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"time"
)

type BloomFilter struct {
	bitArray []bool
	size     int
	hashFns  []func([]byte) int
}

// NewBloomFilter initializes a new Bloom filter with the given size and number of hash functions.
func NewBloomFilter(size int, numHashFns int) *BloomFilter {
	rand.Seed(time.Now().UnixNano())
	hashFns := make([]func([]byte) int, numHashFns)

	// Create simple hash functions with random seeds
	for i := 0; i < numHashFns; i++ {
		seed := rand.Int()
		hashFns[i] = func(data []byte) int {
			h := fnv.New32a()
			h.Write(data)
			h.Write([]byte{byte(seed)})
			return int(h.Sum32()) % size
		}
	}

	return &BloomFilter{
		bitArray: make([]bool, size),
		size:     size,
		hashFns:  hashFns,
	}
}

// Add inserts an element into the Bloom filter.
func (bf *BloomFilter) Add(data []byte) {
	for _, hashFn := range bf.hashFns {
		index := hashFn(data)
		bf.bitArray[index] = true
	}
}

// Check checks whether an element might exist in the Bloom filter.
// Returns true if it might exist, or false if it definitely does not exist.
func (bf *BloomFilter) Check(data []byte) bool {
	for _, hashFn := range bf.hashFns {
		index := hashFn(data)
		if !bf.bitArray[index] {
			return false
		}
	}
	return true
}

func main() {
	// Initialize a Bloom filter with 1000 bits and 3 hash functions.
	bloomFilter := NewBloomFilter(1000, 3)

	// Add elements to the Bloom filter.
	bloomFilter.Add([]byte("apple"))
	bloomFilter.Add([]byte("banana"))
	bloomFilter.Add([]byte("cherry"))

	// Check for membership.
	fmt.Println("apple:", bloomFilter.Check([]byte("apple")))   // Should print: true
	fmt.Println("banana:", bloomFilter.Check([]byte("banana"))) // Should print: true
	fmt.Println("cherry:", bloomFilter.Check([]byte("cherry"))) // Should print: true
	fmt.Println("grape:", bloomFilter.Check([]byte("grape")))   // Might print: false (or true, due to false positives)
	fmt.Println("mango:", bloomFilter.Check([]byte("mango")))   // Might print: false (or true, due to false positives)
}
