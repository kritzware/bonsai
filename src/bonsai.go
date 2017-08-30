package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	db := Store{0, make(map[uint32]*Row)}
	readInput(&db)
}

func readInput(db *Store) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("bonsai> ")
		line, err := reader.ReadString('\n')
		line = strings.ToLower(strings.Trim(line, "\r\n"))

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		splitLine := strings.Split(line, " ")
		var command = splitLine[0]

		switch command {
		case "store":
			key := splitLine[1]
			val := []byte(splitLine[2])
			err := db.Insert(key, val)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
		case "get":
			key := splitLine[1]
			result, _, err := db.Get(key)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			fmt.Println(result)
		case "exit":
			fmt.Println("Goodbye!")
			os.Exit(1)
		default:
			fmt.Println("Error: Unknown command")
		}

	}
}
