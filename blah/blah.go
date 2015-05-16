package blah

import "fmt"
import "bytes"
import "crypto/sha512"
import "hash/crc32"
import "hash/crc64"
import "hash/fnv"

//For Global checks -- use a FIFO to avoid loading too much blah
//n memory
/*

GlobalBlahBlocks
	- BlockHashSha512
		- BlockFNV64
			- Collision 1
				- Location
					- BlockCheckSum
					- StartPosition
				- Location
					- BlockCheckSum
					- StartPosition
				- Location
					- BlockCheckSum
					- StartPosition
			- Collision 2
				- Location
					- BlockCheckSum
					- StartPosition
		- BlockFNV64
			- Collision 1
				- Location
					- BlockCheckSum
					- StartPosition

*/
type GlobalBlahBlocks struct {
	BlkHashSha512 map[[64]uint8]*BlockHashSha512
}

type BlockHashSha512 struct {
	BlkFNV64 map[uint64]*BlockFNV64
}

type BlockFNV64 struct {
	Collison map[uint64]*Collisons
}

type Collisons struct {
	Location map[uint64]*Locations
}

type Locations struct {
	BlockCheckSum uint32
	Startposition uint16
}

type HashBlahmap map[[]byte][uint][unit]*Blah //Blahblock
//sha512 - this is the 1024bytes{}
//unit collison/unit location
