package processfiles

import (
	_ "bytes"
	_ "container/list"
	"fmt"
	"sync"
	_ "time"

	"github.com/keenstart/keennodes/dirnfiles"
	"github.com/keenstart/keennodes/khash"

	"github.com/keenstart/keennodes/blah"
	_ "github.com/keenstart/keennodes/gopfile"
)

// Create a struct to store load blah( are hash name files)
// to memory instead of keep open  and close blab file.
// Make all go routine be able to access the struct
// lock

const (
	//PROCESSROOT   = "/Users/garethharris/"
	MAXPROCCESSES = 170 // To limit the amount of goroutine
)

type ProcesService struct {
	dspro    *dirnfiles.Dirs
	hBlahmap *blah.HashBlahmap
	sync.WaitGroup
	//sync.Mutex

	//BlahMemoryList *list.List; //Type to add is (GlobalBlahBlock )

}

func NewProSerives() (*ProcesService, error) {

	p := &ProcesService{
		dspro:    dirnfiles.NewDirs(),
		hBlahmap: blah.NewHashBlahmap(),
	}

	err := p.dspro.GetFiles()

	if err != nil {
		return p, err
	}

	return p, nil
}

func (p *ProcesService) gethashBlahmap() *blah.HashBlahmap {
	return p.hBlahmap
}

func (p *ProcesService) ProFileSerives() {
	// MAXPROCCESSESTo limit the amount of goroutine
	maxprocessch := make(chan int, 1) // MAXPROCCESSESTo limit the amount of goroutine

	wg := p.WaitGroup //debug purposes

	for _, files := range p.dspro.Files {
		wg.Add(1) //debug purposes

		// To limit the amount of goroutine
		maxprocessch <- 1 // To limit the amount of goroutine
		go func(files *dirnfiles.Dirinfo) {
			fmt.Println("start sevrice go")
			process(files, p /*,p.sync.mutex*/)

			<-maxprocessch  // To limit the amount of goroutine
			defer wg.Done() //debug purposes
		}(files)
		break //debug purposes
	}
	wg.Wait() //debug purposes
}

func process(files *dirnfiles.Dirinfo, p *ProcesService /*, lock *sync.mutex*/) {

	fmt.Printf("\nKey: %d = %s with %d bytes. CRC \n\n",
		files.Key, files.Path, files.Fsize) //debug purposes

	// Get file bytes
	sfile := khash.Filebytes(files.Path)

	// Get BLOCKSIZE slice from file
	lo := 0
	hi := dirnfiles.BLOCKSIZE

	// Break the 'for loop' when the last valid BLOCKSIZE
	// can be process from the file
	brprocess := files.Fsize - dirnfiles.BLOCKSIZE

	var (
		startposition uint16
		blkHashSha512 []byte //[64]uint8
		blkFNV64      uint64
		blockCheckSum uint32
		globalBlahBlk blah.GlobalBlahBlock
		blockStatus   blah.BlockStatus
		location      blah.Locations
		collision     blah.Collisions
	)
	// File location represented by it's key to save on space.
	// To avoid repeating its location int the blah which
	// takes up more space than the key
	location = blah.Locations(files.Key)

	//loop BLOCKSIZE bytes at a time moving by 1 byte at a time
	for {

		fmt.Printf("blocks #%d , hi = %d,value = %x,len1024 = %d  cap = %d,file sizes = %d\n",
			lo, hi, sfile[lo:hi], len(sfile[lo:hi]), cap(sfile[lo:hi]), files.Fsize) //debug purposes

		blkHashSha512 = khash.Sha512fn(sfile[lo:hi]) //	- get BlockHashSha512
		blkFNV64 = khash.HashFNV64(sfile[lo:hi])     //	- get BlockFNV64

		startposition = uint16(lo) //	- start position
		blockCheckSum = khash.Hashcrc32(sfile[lo:hi])

		blah.NewBlockStatus(blockCheckSum, startposition)
		blah.NewGlobalBlahBlock(blkHashSha512, blkFNV64)
		p.hBlahmap.PutHashBlahmap(globalBlahBlk, collision, location, blockStatus)

		/*

				if ok := HashBlahmap[GlobalBlahBlock]; !ok{
			 		if HashBlahmap fileexist (the filename = GlobalBlahBlock.BlockHashSha512)  {
			 			// lock sync.Rlock
						//open file to read and save to HashBlahmap
						//lock sync.RUnlock

						 //check for collision
						-if collision{ //make this a func
							collision++ // increment collison
						}
					}else{

						//create HashBlahmap to save to file (the filename = GlobalBlahBlock.BlockHashSha512)
						collision := 1
					}

					//add to back of list BlahMemoryList


				}else{
					- check for collision
					-if collision{//make this a func
						collision++ // increment collison
					}
				}
				//lock sync.lock

				//Add to HashBlahmap with func AddHashBlahmap below
				//AddHashBlahmap(globalBlahBlk GlobalBlahBlock, collision Collisions,
							location Locations, blockStatus BlockStatus)
		*/

		// 1. Save to file.
		//2. remove element(GlobalBlahBlock) from the front of list BlahMemoryList if list is greater than MAXMEMORYBLAH
		//and also delete from map

		// lock sync.unlock

		// Break loop when last valid block is process
		//if cap(sfile[lo:hi]) == dirnfiles.BLOCKSIZE {
		if int64(lo) >= brprocess {
			fmt.Println(" cap = ", cap(sfile[lo:hi]), " lo = ", lo, " hi = ", hi)
			break
		}

		// Move one byte at a time across the slice
		lo++
		hi++

	}

	fmt.Println("files size", len(sfile))
	fmt.Printf("\nname = %s, Key: %d = %s with %d bytes. CRC \n\n",
		files.Name, files.Key, files.Path, files.Fsize)

	fmt.Printf("\n Files date %x\n\n",
		sfile)

}
