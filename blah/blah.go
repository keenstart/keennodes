package blah

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
	//BlockHashSha512 [8]int64 //[]byte
	BlockFNV64 uint64
}

type Collisions uint64
type Locations uint64

type BlockStatus struct {
	BlockCheckSum uint32
	Startposition uint16
}

func NewGlobalBlahBlock(blkFNV64 uint64) GlobalBlahBlock {
	return GlobalBlahBlock{BlockFNV64: blkFNV64}

}

func NewBlockStatus(blockCheckSum uint32, startposition uint16) BlockStatus {
	return BlockStatus{BlockCheckSum: blockCheckSum, Startposition: startposition}
}

/*--



Padding




--*/
