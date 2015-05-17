package blah

import (
	"bytes"
	"crypto/sha512"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"

	"github.com/keenstart/keennodes/gopfile"
)

//sha512 - this is the 1024bytes{}
//unit collison/unit location

//For Global checks -- use a FIFO to avoid loading too much blah
//n memory

/*

GlobalBlahBlocks
	- BlockHashSha512
		- BlockFNV64
			- Collisions 1
				- Location
					- BlockCheckSum
					- StartPosition
				- Location
					- BlockCheckSum
					- StartPosition
				- Location
					- BlockCheckSum
					- StartPosition
			- Collisions 2
				- Location
					- BlockCheckSum
					- StartPosition
		- BlockFNV64
			- Collisions 1
				- Location
					- BlockCheckSum
					- StartPosition

*/

type GlobalBlahBlock struct {
	BlockHashSha512 [64]uint8
	BlockFNV64      uint64
}

type Collisions uint64
type Locations uint64

type BlockStatus struct {
	BlockCheckSum uint32
	Startposition uint16
}

type HashBlahmap map[GlobalBlahBlocks][Collisions][Locations]BlockStatus

func NewHashBlahmap() *HashBlahmap {

	hashBlahmap = make(map[GlobalBlahBlock][Collisions][Locations]BlockStatus)

	return &hashBlahmap

}

func NewGlobalBlahBlock(blkHashSha512 [64]uint8, blkFNV64 uint64) GlobalBlahBlock {

	return GlobalBlahBlock{BlockHashSha512: blkHashSha512, BlockFNV64: blkFNV64}

}

func NewBlockStatus(blockCheckSum, startposition) BlockStatus {
	return BlockStatus{blockCheckSum: blockCheckSum, startposition: startposition}
}

func (h *HashBlahmap) AddHashBlahmap(globalBlahBlk GlobalBlahBlock, collision Collisions,
	location Locations, blockStatus BlockStatus) {

	*h[globalBlahBlk][collision][location] = blockStatus

}

// One unique Block can be assocaite with one file. It does not server any redundant purpose
// to use the same file for the unique block more than once. However for redundancy
// protection the unique block can  be represented by a another location incase a location
// get lose or corrupted
func (h *HashBlahmap) GetLocationBlkStatus(globalBlahBlk GlobalBlahBlock, collision Collisions,
	location Locations) {

	return *h[globalBlahBlk][collision][location]

}

func (h *HashBlahmap) GetFiles() error {
	err := gopfile.Load(BLOBFILE, h)

	if err != nil {
		return err
	}

	return nil

}

func (h *HashBlahmap) SetFiles() error {
	err := gopfile.Save(BLOBFILE, h)

	if err != nil {
		return err
	}

	return nil

}

/*--



Padding




--*/
