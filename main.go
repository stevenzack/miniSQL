package main

import (
	"encoding/binary"
	"fmt"
	"log"
)

const (
	BNODE_NODE = 1
	BNODE_LEAF = 2
)

const (
	HEADER             = 4
	BTREE_PAGE_SIZE    = 4096
	BTREE_MAX_KEY_SIZE = 1000
	BTREE_MAX_VAL_SIZE = 3000
)

func init() {
	node1max := HEADER + 8 + 2 + 4 + BTREE_MAX_KEY_SIZE + BTREE_MAX_VAL_SIZE
	if node1max > BTREE_PAGE_SIZE {
		log.Fatal(`node1max is not suppose to be bigger than BTREE_PAGE_SIZE`)
	}
}

func main() {
	bs := []byte{0, 1}
	fmt.Println(binary.LittleEndian.Uint16(bs))
}
