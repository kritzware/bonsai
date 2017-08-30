package main

import (
	"bufio"
	"errors"
	"fmt"
	"hash/fnv"
	"os"
	"strings"
)

func main() {
	db := newDBInstance()
	readInput(db)
}

func store(db map[uint32]string, args []string) error {
	if len(args) < 3 {
		return errors.New("store expects two arguments: 'store [key] [value]'")
	}
	key := args[1]
	value := args[2]

	hash := hashKey(key)
	db[hash] = value

	return nil
}

func get(db map[uint32]string, args []string) (*string, error) {
	if len(args) < 2 {
		return nil, errors.New("get expects one argument: 'get [key]'")
	}
	key := args[1]
	hash := hashKey(key)
	value := db[hash]
	if value != "" {
		return &value, nil
	}
	return nil, nil
}

func hashKey(s string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(s))
	return hash.Sum32()
}

func newDBInstance() map[uint32]string {
	return make(map[uint32]string)
}

func readInput(db map[uint32]string) {
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
			err := store(db, splitLine)
			if err != nil {
				fmt.Println("| Error:", err)
				break
			}
		case "get":
			value, err := get(db, splitLine)
			if err != nil {
				fmt.Println("| Error:", err)
				break
			}
			if value != nil {
				fmt.Printf("[%p] %s\n", value, *value)
				break
			}
			fmt.Println(nil)
		case "exit":
			fmt.Println("Goodbye!")
			os.Exit(1)
		default:
			fmt.Println("Error: Unknown command")
		}

	}
}
