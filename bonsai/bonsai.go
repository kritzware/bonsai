package bonsai

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func CreateBonsaiInstance() *Store {
	db := Store{0, make(map[uint32]*Row)}
	existingData, err := ioutil.ReadFile("bonsai.roots")
	if err != nil {
		return &db
	}
	db.LoadFromBackup(string(existingData))
	return &db
}

func store(db *Store, splitLine []string) {
	key := splitLine[1]
	val := []byte(strings.Join(splitLine[2:], " "))
	err := db.Insert(key, val)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func get(db *Store, splitLine []string) {
	key := splitLine[1]
	result, row, err := db.Get(key)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("[%p] %s\n", &row, result)
}

func save(db *Store) {
	startTime := time.Now().UnixNano() / int64(time.Millisecond)
	err := db.Save()
	if err != nil {
		fmt.Println("Error saving db:", err)
		return
	}
	endTime := time.Now().UnixNano() / int64(time.Millisecond)
	diff := (endTime - startTime)
	fmt.Printf("Finished in %dms\n", diff)
}

func fill(db *Store, splitLine []string) {
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
}

func exit() {
	fmt.Println("Goodbye!")
	os.Exit(1)
}

func ReadInput(db *Store) {
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
			store(db, splitLine)
		case "get":
			get(db, splitLine)
		case "save":
			save(db)
		case "status":
			fmt.Printf("db: %T, %d bytes\n", db, db.Size())
			fmt.Printf("keys: %d\n", db.GetKeyCount())
		case "fill":
			fill(db, splitLine)
		case "exit":
			exit()
		default:
			fmt.Println("Error: Unknown command")
		}
	}
}
