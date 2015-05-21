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

	"github.com/keenstart/keennodes/blah"
	"github.com/keenstart/keennodes/gopfile"
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
	//hBlahmap *blah.HashBlahmap

	//HashBlahmap map[blah.GlobalBlahBlock]map[blah.Collisions]map[blah.Locations]blah.BlockStatus

	sync.WaitGroup
	//sync.Mutex

	//BlahMemoryList *list.List; //Type to add is (GlobalBlahBlock )

}

func NewProSerives() (*ProcesService, error) {

	p := &ProcesService{
		dspro: dirnfiles.NewDirs(),
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
		blkHashSha512 [8]int64 //[]byte //[64]uint8
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
	location = 4 //blah.Locations(files.Key)
	collision = 2
	// test
	blahfFileNameStr1 := ""

	path := "./blahs/"
	//loop BLOCKSIZE bytes at a time moving by 1 byte at a time
	for {

		hashBlahmap := make(map[blah.GlobalBlahBlock]map[blah.Collisions]map[blah.Locations]blah.BlockStatus)

		blahhashbyte := khash.Sha512fn(sfile[lo:hi])

		blahfFileNameStr := hex.EncodeToString(blahhashbyte)

		blahfFileNameStr1 = blahfFileNameStr //debug
		fmt.Println(blahfFileNameStr)        //debug
		fmt.Println("Start Map")             //debug

		f, err := os.Open(path + blahfFileNameStr)

		//If file exist load it
		if !os.IsNotExist(err) {
			fmt.Println("Load file = ", path+blahfFileNameStr) //debug
			gopfile.Load(path+blahfFileNameStr, &hashBlahmap)
			fmt.Println("Load hashBlahmap = ", hashBlahmap)

		}
		f.Close()

		blkHashSha512 = khash.Convert64(blahhashbyte) //	- get BlockHashSha512
		//fmt.Printf("blkHashSha512 : %d\n", blkHashSha512) //debug

		blkFNV64 = khash.HashFNV64(sfile[lo:hi]) //	- get BlockFNV64
		fmt.Printf("blkFNV64 : %d\n", blkFNV64)  //debug

		fmt.Printf("Collisions : %d\n", collision) //debug

		fmt.Printf("Locations : %d\n", location) //debug

		startposition = uint16(lo) //	- start position
		//fmt.Printf("startposition : %d\n", startposition) //debug

		blockCheckSum = khash.Hashcrc32(sfile[lo:hi])
		//fmt.Printf("blockCheckSum : %d\n\n", blockCheckSum) //debug

		blockStatus = blah.NewBlockStatus(blockCheckSum, startposition)
		globalBlahBlk = blah.NewGlobalBlahBlock(blkHashSha512, blkFNV64)

		//_, ok := hashBlahmap[globalBlahBlk][collision][location]
		cc, ok := hashBlahmap[globalBlahBlk]
		if !ok {
			cc = make(map[blah.Collisions]map[blah.Locations]blah.BlockStatus)
			hashBlahmap[globalBlahBlk] = cc
		}

		ll, ok := hashBlahmap[globalBlahBlk][collision]
		if !ok {
			ll = make(map[blah.Locations]blah.BlockStatus)
			hashBlahmap[globalBlahBlk][collision] = ll
		}

		_, ok = hashBlahmap[globalBlahBlk][collision][location]
		if !ok {
			hashBlahmap[globalBlahBlk][collision][location] = blockStatus
		}

		// Save the new or updated
		gopfile.Save(path+blahfFileNameStr, hashBlahmap)
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

	hashBlahmap := make(map[blah.GlobalBlahBlock]map[blah.Collisions]map[blah.Locations]blah.BlockStatus)
	gopfile.Load(path+blahfFileNameStr1, &hashBlahmap)
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
	/*
		for _, value := range test[globalBlahBlk][collision][location] {

										}

		fmt.Println("files size", len(sfile))
		fmt.Printf("\nname = %s, Key: %d = %s with %d bytes. CRC \n\n",
			files.Name, files.Key, files.Path, files.Fsize)

		fmt.Printf("\n Files date %x\n\n",
			sfile)
	*/

}
