// Final Commit Done

package chunk_based_processing

import (
	"bufio"
	"fmt"
	"os"
)

var (
	file   *os.File
	reader *bufio.Reader
	count  int
	chunk  []byte
)

func Start() {
	fmt.Println("\nProcessing Large Files in Chunks")
	fmt.Println()

	fmt.Println("Total Read Operations Happened =>", countIterations(512))
	fmt.Println()

	fmt.Println("Total Read Operations Happened =>", countIterations(1024))
	fmt.Println()

	fmt.Println("Total Read Operations Happened =>", countIterations(2048))
	fmt.Println()

	fmt.Println("Total Read Operations Happened =>", countIterations(4096))
	fmt.Println()
}

func countIterations(chunkSize int) int {
	file, _ = os.Open("22_Chunk_Based_File_Processing/data.txt")

	defer file.Close()

	reader = bufio.NewReader(file)
	count = 0
	chunk = make([]byte, chunkSize)

	for {
		n, err := reader.Read(chunk)
		if n > 0 {
			// fmt.Print(string(chunk[:n])) -->> Send to channel, the current chunk data
			// Here, I'm just trying to count iteration, chukSize based, how many
			count++
		}
		if err != nil {
			break
		}
	}
	return count
}
