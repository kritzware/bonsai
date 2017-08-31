package main

import (
	"github.com/kritzware/bonsai/bonsai"
)

func main() {
	db := bonsai.CreateBonsaiInstance()
	bonsai.ReadInput(db)
}
