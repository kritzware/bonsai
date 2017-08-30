package main

import (
	"errors"
	"hash/fnv"
	"unsafe"
)

type Row struct {
	key   string
	hash  uint32
	value []byte
}

type Store struct {
	keys uint32
	data map[uint32]*Row
}

func (s *Store) Insert(key string, val []byte) error {
	hash := s.createHash(key)
	row := &Row{key, hash, val}
	s.insertIntoMap(hash, row)
	s.keys++
	return nil
}

func (s *Store) Get(key string) (string, *Row, error) {
	hash := s.createHash(key)
	if row, ok := s.data[hash]; ok {
		return string(row.value[:]), row, nil
	} else {
		return "", nil, errors.New("Key not found")
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
