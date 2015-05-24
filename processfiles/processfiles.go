package processfiles

import (
	_ "bytes"
	_ "container/list"
	"encoding/hex"
	"fmt"
	"os"
	"sync"
	_ "time"

	"github.com/keenstart/keennodes/dirnfiles"
	"github.com/keenstart/keennodes/khash"

	_ "github.com/keenstart/keennodes/blah"
	"github.com/keenstart/keennodes/gopfile"
)

// Create a struct to store load blah( are hash name files)
// to memory instead of keep open  and close blab file.
// Make all go routine be able to access the struct
// lock

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

const (
	//PROCESSROOT   = "/Users/garethharris/"
	MAXPROCCESSES = 170 // To limit the amount of goroutine
	BLAHPATH      = "./blahs/"
)

type ProcesService struct {
	dspro *dirnfiles.Dirs

	// 'blahtracker' is to allow other goroutine the ability to continue to work. only blocking
	// if another goroutine is using the same blah file that it needs.
//w blahtracker map[[8]int64]bool
	blahtracker map[[8]int64]chan int//w
	sync.Mutex

	sync.WaitGroup

	//BlahMemoryList *list.List; //Type to add is (GlobalBlahBlock )

}

func NewProSerives() (*ProcesService, error) {

	p := &ProcesService{
		dspro:       dirnfiles.NewDirs(),
		blahtracker: make(map[[8]int64]chan),//
		//w blahtracker: make(map[[8]int64]bool),
		//HashBlahmap: make(map[blah.GlobalBlahBlock]map[blah.Collisions]map[blah.Locations]blah.BlockStatus),
		//hBlahmap: blah.NewHashBlahmap(),
	}

	err := p.dspro.GetFiles()

	if err != nil {
		return p, err
	}

	return p, nil
}

/*func (p *ProcesService) gethashBlahmap() *blah.HashBlahmap {
	return p.hBlahmap
}*/

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

	//sfile := make([]byte, files.Fsize, files.Fsize) //debug
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
		//blkHashSha512 [8]int64 //[]byte //[64]uint8
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

	blahfFileNameStr1 := "" //debug

	//loop BLOCKSIZE bytes at a time moving by 1 byte at a time
	for {

		hashBlahmap := make(map[GlobalBlahBlock]map[Collisions]map[Locations]BlockStatus)

		blahhashbyte := khash.Sha512fn(sfile[lo:hi])

		blahfFileNameStr := BLAHPATH + hex.EncodeToString(blahhashbyte)

		blahfFileNameStr1 = blahfFileNameStr //debug
		fmt.Println(blahfFileNameStr)        //debug
		fmt.Println("Start Map")             //debug

		// Check if files exist
		f, err := os.Open(blahfFileNameStr)
		f.Close()

		// This while loop will keep looping if a another goroutine is as load the
		// same blah file content. When the goroutine finish with the content it will
		// remove it from the 'blahtracker' map and the loop will exit
		//w for _, ok := p.blahtracker[khash.ConverttoInt64(blahhashbyte)]; ok; {
		//w }
		if _,ok:= p.blahtracker[khash.ConverttoInt64(blahhashbyte)];!ok;{ // w
			p.blahtracker[khash.ConverttoInt64(blahhashbyte)] = make(chan int) //w
		}
		p.blahtracker[khash.ConverttoInt64(blahhashbyte)] <- 1 //w
		
		//If file exist load it
		if !os.IsNotExist(err) {

			// Add the blah to the 'blahtracker' to start working withe the content of the blah
			//w p.blahtracker[khash.ConverttoInt64(blahhashbyte)] = true
			//w p.Mutex.Lock()
			fmt.Println("Load file = ", blahfFileNameStr) //debug
			gopfile.Load(blahfFileNameStr, &hashBlahmap)
			fmt.Println("Load hashBlahmap = ", hashBlahmap) //debug
			//w p.Mutex.Unlock()
		}

		fmt.Printf("Filenames : %x\n", blahhashbyte) //debug
		//blkHashSha512 = khash.ConverttoInt64(blahhashbyte) //	- get BlockHashSha512
		//fmt.Printf("blkHashSha512 : %d\n", blkHashSha512)  //debug

		blkFNV64 = khash.HashFNV64(sfile[lo:hi]) //	- get BlockFNV64
		fmt.Printf("blkFNV64 : %d\n", blkFNV64)  //debug

		fmt.Printf("Collisions : %d\n", collision) //debug

		fmt.Printf("Locations : %d\n", location) //debug

		startposition = uint16(lo)                        //	- start position
		fmt.Printf("startposition : %d\n", startposition) //debug

		blockCheckSum = khash.Hashcrc32(sfile[lo:hi])
		fmt.Printf("blockCheckSum : %d\n\n", blockCheckSum) //debug

		blockStatus = NewBlockStatus(blockCheckSum, startposition)
		globalBlahBlk = NewGlobalBlahBlock(blkFNV64)

		//_, ok := hashBlahmap[globalBlahBlk][collision][location]
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
			hashBlahmap[globalBlahBlk][collision][location] = blockStatus
		}

		// Save the new or updated
		gopfile.Save(blahfFileNameStr, hashBlahmap)
		fmt.Println("")
		fmt.Println("Save hashBlahmap = ", hashBlahmap)

		//if !ok {

		//p.hBlahmap.PutHashBlahmap(globalBlahBlk, collision, location, blockStatus)

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
		//fmt.Println("block = ", lo, " block = ", sfile[lo:hi]) //debug

		// Remove from 'blahtracker'  allow another goroutine that need that blah go ahead
		<-p.blahtracker[khash.ConverttoInt64(blahhashbyte)] //w
		delete(p.blahtracker, khash.ConverttoInt64(blahhashbyte))

		fmt.Printf("\nBlock# = %d, Block: %x \n\n",
			lo, sfile[lo:hi]) //Debug

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

	hashBlahmap := make(map[GlobalBlahBlock]map[Collisions]map[Locations]BlockStatus)
	gopfile.Load(blahfFileNameStr1, &hashBlahmap)
	fmt.Println("")
	fmt.Println("hashBlahmap = ", hashBlahmap)

	for _, g := range hashBlahmap {
		fmt.Println("globalBlahBlk = ", g)
		for _, c := range g {
			fmt.Println("Collisions = ", c)
			for _, l := range c {
				fmt.Println("location = ", l)
			}

		}
	}
	/*
		collision = 0
		location = 0
		//var b [8]int64
		fmt.Println("Get Map")
		blkHashSha512 = [...]int64{8542560806652977256, 3602876116890795078, 2414306596406457046, -4245526746901624870, -7073043113630984961, 2933880197123614138, 7460501996556034562, 253097422264543894}

		blkFNV64 = 12359158838599038801 //	- get BlockFNV64 2748473583

		globalBlahBlk = blah.NewGlobalBlahBlock(blkHashSha512, blkFNV64)

		blockStatus = hashBlahmap[globalBlahBlk][collision][location]

		fmt.Printf("startposition : %d\n", blockStatus.Startposition) //debug

		fmt.Printf("blockCheckSum : %d\n\n\n", blockStatus.BlockCheckSum)
	*/

	fmt.Println("files size", len(sfile))
	fmt.Printf("\nname = %s, Key: %d = %s with %d bytes. CRC \n\n",
		files.Name, files.Key, files.Path, files.Fsize)

	fmt.Printf("\n Files date %x\n\n",
		sfile)

}
