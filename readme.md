# Bloom Filter in Go

This is a simple implementation of a **Bloom Filter** in Go, a probabilistic data structure used to test whether an element is a member of a set. It is highly memory-efficient and allows for false positives but never false negatives.

---

## Features
- **Efficient Storage**: Compact representation of a set of elements.
- **Probabilistic Membership Testing**: Allows for quick existence checks with minimal memory usage.
- **Customizable**: Specify filter size and hash functions for fine-tuned performance.

---

## How It Works
1. **Insert an Element**: Hash the element using multiple hash functions and set the corresponding bits in a bit array.
2. **Check Membership**: Hash the element again and verify that the required bits are set. If any bit is not set, the element is definitely not in the set.

---

## Code Structure
- **`main.go`**: Demonstrates usage of the Bloom filter.
- **`bloom.go`**: Core implementation of the Bloom filter, including:
  - Setting and checking bits.
  - Hash function integration using **Murmur3**.

---

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/arpit-ayeee/bloom.git
   cd bloomfilter
   ```

2. **Initialize the module** (if not already done):
   ```bash
   go mod init github.com/yourusername/bloomfilter
   ```

3. **Download dependencies**:
   ```bash
   go mod tidy
   ```

---

## Usage

### Running the Example Code

1. Compile and run the program:
   ```bash
   go run main.go
   ```

2. The output will show Bloom filter operations, including element insertion and membership checks.

---

## Limitations
- **False Positives**: The Bloom filter may indicate that an element exists when it doesn’t.
- **No Deletions**: Removing an element is not supported without risking false negatives.

---
