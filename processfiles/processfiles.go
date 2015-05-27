package processfiles

import (
	"encoding/hex"
	"fmt"
	"os"
	"sync"
	_ "time"

	"github.com/keenstart/keennodes/dirnfiles"
	"github.com/keenstart/keennodes/gopfile"
	"github.com/keenstart/keennodes/khash"
)

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

const (
	MAXPROCCESSES = 170 // To limit the amount of goroutine
	BLAHPATH      = "./blahs/"
	MAXLOCATION   = 12 // Limit the amount of location per blah
)

type ProcesService struct {
	dspro *dirnfiles.Dirs
	// 'blahtracker' is to allow other goroutine the ability to continue to work. only blocking
	// if another goroutine is using the same blah file that it needs.
	blahtracker map[[8]int64]bool
	sync.Mutex
	sync.WaitGroup
}

func NewProSerives() (*ProcesService, error) {
	p := &ProcesService{
		dspro:       dirnfiles.NewDirs(),
		blahtracker: make(map[[8]int64]bool),
	}

	err := p.dspro.GetFiles()
	if err != nil {
		return p, err
	}

	return p, nil
}

func (p *ProcesService) ProFileSerives() {
	maxprocessch := make(chan int, 1) // MAXPROCCESSESTo limit the amount of goroutine
	wg := p.WaitGroup

	for _, files := range p.dspro.Files {
		wg.Add(1)

		// To limit the amount of goroutine
		maxprocessch <- 1

		go func(files *dirnfiles.Dirinfo) {
			process(files, p)
			<-maxprocessch
			defer wg.Done()
		}(files)
		break //debug purposes
	}
	wg.Wait()
}

func process(files *dirnfiles.Dirinfo, p *ProcesService /*, lock *sync.mutex*/) {
	sfile := khash.Filebytes(files.Path)
	lo := 0
	hi := dirnfiles.BLOCKSIZE

	// Break the 'for loop' when the last valid BLOCKSIZE
	// can be process from the file
	brprocess := files.Fsize - dirnfiles.BLOCKSIZE

	var (
		startposition uint16
		blkFNV64      uint64
		blockCheckSum uint32
		globalBlahBlk GlobalBlahBlock
		blockStatus   BlockStatus
		location      Locations
		collision     Collisions
	)

	// File location represented by it's key to save on space.
	// To avoid repeating its location int the blah which
	// takes up more space than the key
	location = Locations(files.Key)
	collision = 1

	//loop BLOCKSIZE bytes at a time moving by 1 byte at a time
	for {

		hashBlahmap := make(map[GlobalBlahBlock]map[Collisions]map[Locations]BlockStatus)
		blahhashbyte := khash.Sha512fn(sfile[lo:hi])
		blahfFileNameStr := BLAHPATH + hex.EncodeToString(blahhashbyte)

		// Check if files exist
		f, err := os.Open(blahfFileNameStr)
		f.Close()

		// This while loop will keep looping if a another goroutine as  the
		// same blah file content. When the goroutine finish with the content it will
		// remove it from the 'blahtracker' map and the loop will exit

		for {
			p.Mutex.Lock()
			if _, ok := p.blahtracker[khash.ConverttoInt64(blahhashbyte)]; !ok {
				p.blahtracker[khash.ConverttoInt64(blahhashbyte)] = true
				p.Mutex.Unlock()
				break
			}
			p.Mutex.Unlock()
			fmt.Println("WAITING:  ", blahhashbyte) //debug
		}

		//If file exist load it
		if !os.IsNotExist(err) {
			gopfile.Load(blahfFileNameStr, &hashBlahmap)
		}

		blkFNV64 = khash.HashFNV64(sfile[lo:hi]) //	- get BlockFNV64
		startposition = uint16(lo)               //	- start position
		blockCheckSum = khash.Hashcrc32(sfile[lo:hi])
		blockStatus = NewBlockStatus(blockCheckSum, startposition)
		globalBlahBlk = NewGlobalBlahBlock(blkFNV64)

		cc, ok := hashBlahmap[globalBlahBlk]
		if !ok {
			cc = make(map[Collisions]map[Locations]BlockStatus)
			hashBlahmap[globalBlahBlk] = cc
		}

		ll, ok := hashBlahmap[globalBlahBlk][collision]
		if !ok {
			ll = make(map[Locations]BlockStatus)
			hashBlahmap[globalBlahBlk][collision] = ll
		}

		_, ok = hashBlahmap[globalBlahBlk][collision][location]
		if !ok {
			// Limit the number of location with the same copy of Blah
			// It is necessary to have more than one location in case one is loss
			if len(hashBlahmap[globalBlahBlk][collision]) < MAXLOCATION {

				// TODO check for collision and incement collision count
				// numbercollision := len(hashBlahmap[globalBlahBlk])
				// use a utility function and return true if there is a collison

				hashBlahmap[globalBlahBlk][collision][location] = blockStatus
			}
		}

		// Save the new or updated
		gopfile.Save(blahfFileNameStr, hashBlahmap)

		p.Mutex.Lock()
		delete(p.blahtracker, khash.ConverttoInt64(blahhashbyte))
		p.Mutex.Unlock()

		// Break loop when last valid block is process
		if int64(lo) >= brprocess {
			break
		}

		// Move one byte at a time across the slice
		lo++
		hi++
	}

}
