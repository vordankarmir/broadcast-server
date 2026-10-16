package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Read() {
	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)
	// Allocate a larger buffer for inputs with very long lines.
	// The first argument is the initial buffer, the second is the max capacity.
	scanner.Buffer(make([]byte, 1024), 1024*1024)

	// var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		fmt.Println(line)
		// socket.EmitEvent(client, line)
		// lines = append(lines, line)
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("output:")
	// for _, l := range lines {
	// 	fmt.Println(l)
	// }
}
