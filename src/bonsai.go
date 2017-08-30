package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
			result, row, err := db.Get(key)
			if err != nil {
				fmt.Println("Error:", err)
				break
			}
			fmt.Printf("[%p] \"%s\"\n", &row, result)
		case "status":
			fmt.Printf("db: %T, %d bytes\n", db, db.Size())
			fmt.Printf("keys: %d\n", db.GetKeyCount())
		case "fill":
			startTime := time.Now().UnixNano() / int64(time.Millisecond)
			iterations := splitLine[1]
			max, _ := strconv.Atoi(iterations)
			for i := 0; i < max; i++ {
				val := []byte(strconv.Itoa(i))
				key := strconv.Itoa(i)
				err := db.Insert("key:"+key, val)
				if err != nil {
					fmt.Println("Erorr:", err)
				}
			}
			endTime := time.Now().UnixNano() / int64(time.Millisecond)
			diff := (endTime - startTime)
			fmt.Printf("Finished in %dms\n", diff)
		case "exit":
			fmt.Println("Goodbye!")
			os.Exit(1)
		default:
			fmt.Println("Error: Unknown command")
		}

	}
}
