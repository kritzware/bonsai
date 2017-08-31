package bonsai

import (
	"errors"
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"strconv"
	"strings"
	"unsafe"
)

type Row struct {
	key    string
	hash   uint32
	value  []byte
	Locked bool
}

type Store struct {
	keys uint32
	data map[uint32]*Row
}

func (s *Store) Insert(key string, val []byte) error {
	hash := s.createHash(key)
	row := Row{key, hash, val, false}
	row.Locked = true
	defer func() {
		row.Locked = false
	}()
	s.insertIntoMap(hash, &row)
	s.keys++
	return nil
}

func (s *Store) Get(key string) (string, *Row, error) {
	hash := s.createHash(key)
	if row, ok := s.data[hash]; ok {
		row.Locked = true
		defer func() {
			row.Locked = false
		}()
		return string(row.value[:]), row, nil
	} else {
		return "", nil, errors.New("Key not found")
	}
}

func (s *Store) Save() error {
	var data []string
	for key, row := range s.data {
		stringKey := fmt.Sprint(key)
		raw := []string{row.key, stringKey, string(row.value[:])}
		joined := strings.Join(raw, ":")
		data = append(data, joined)
	}
	joinedRows := strings.Join(data, " ")
	bytes := []byte(joinedRows)
	err := ioutil.WriteFile("bonsai.roots", bytes, 0644)
	return err
}

func (s *Store) LoadFromBackup(savedData string) {
	rows := strings.Split(savedData, " ")
	for i := 0; i < len(rows); i++ {
		split := strings.Split(rows[i], ":")
		rowHash, _ := strconv.ParseUint(split[2], 10, 32)
		rowHash32 := uint32(rowHash)
		rowData := []byte(split[3])
		row := Row{split[1], rowHash32, rowData, false}
		s.insertIntoMap(rowHash32, &row)
		s.keys++
	}
}

func (s *Store) GetKeyCount() uint32 {
	return s.keys
}

func (s *Store) Size() int {
	size := 0
	for _, value := range s.data {
		size += int(unsafe.Sizeof(value))
	}
	return size
}

func (s *Store) insertIntoMap(hash uint32, row *Row) error {
	s.data[hash] = row
	return nil
}

func (s *Store) createHash(key string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return hash.Sum32()
}
