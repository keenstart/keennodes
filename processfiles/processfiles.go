package processfiles

import (
	_ "container/list"
	"fmt"
	"sync"
	_ "time"

	"github.com/keenstart/keennodes/dirnfiles"
	"github.com/keenstart/keennodes/khash"

	//"github.com/keenstart/keennodes/blah"
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
	dspro *dirnfiles.Dirs
	sync.WaitGroup
	//sync.Mutex

	//BlahMemoryList *list.List; //Type to add is (GlobalBlahBlock )

}

func NewProSerives() (*ProcesService, error) {

	p := &ProcesService{
		dspro: dirnfiles.NewDirs(),
	}

	err := p.dspro.GetFiles()

	if err != nil {
		return p, err
	}

	return p, nil
}

func (p *ProcesService) ProFileSerives() {
	maxprocessch := make(chan int, 1) // MAXPROCCESSESTo limit the amount of goroutine

	wg := p.WaitGroup //debug purposes

	for _, files := range p.dspro.Files {
		wg.Add(1)         //debug purposes
		maxprocessch <- 1 // To limit the amount of goroutine
		go func(files *dirnfiles.Dirinfo) {
			fmt.Println("start sevrice go")
			process(files /*,p.sync.mutex*/)

			<-maxprocessch  // To limit the amount of goroutine
			defer wg.Done() //debug purposes
		}(files)
		break //debug purposes
	}
	wg.Wait() //debug purposes
}

func process(files *dirnfiles.Dirinfo /*, lock *sync.mutex*/) {

	// Get file bytes
	filesbytes := khash.Filebytes(files.Path)

	fmt.Printf("\nKey: %d = %s with %d bytes. CRC \n\n",
		files.Key, files.Path, files.Fsize) //debug purposes

	// Get BLOCKSIZE slice from file
	lo := 0
	hi := dirnfiles.BLOCKSIZE

	// Break thr for when the last valid BLOCKSIZE
	// can be process from the file
	brprocess := files.Fsize - dirnfiles.BLOCKSIZE

	sfile := filesbytes

	//loop BLOCKSIZE bytes at a time moving by 1 byte at a time
	for {

		fmt.Printf("blocks #%d , hi = %d,value = %x,len1024 = %d  cap = %d,file sizes = %d\n",
			lo, hi, sfile[lo:hi], len(sfile[lo:hi]), cap(sfile[lo:hi]), files.Fsize) //debug purposes

		/*
				- get BlockHashSha512
				- get BlockFNV64
				- start position
				- location - files.Key

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
		if int64(lo) >= brprocess {
			fmt.Println(" cap = ", cap(sfile[lo:hi]), " lo = ", lo, " hi = ", hi)
			break
		}

		// Move one byte at a time across the slice
		lo++
		hi++

	}

	fmt.Println("files size", len(filesbytes))
	fmt.Printf("\nname = %s, Key: %d = %s with %d bytes. CRC \n\n",
		files.Name, files.Key, files.Path, files.Fsize)

	fmt.Printf("\n Files date %x\n\n",
		filesbytes)

}
