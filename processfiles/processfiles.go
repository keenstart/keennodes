package processfiles

import (
	"fmt"
	_ "time"
	_ "list"

	"github.com/keenstart/keennodes/dirnfiles"
	"github.com/keenstart/keennodes/khash"

	_ "github.com/keenstart/keennodes/gopfile"
	_ "github.com/keenstart/keennodes/blah"
)

// Create a struct to store load blah( are hash name files)
// to memory instead of keep open  and close blab file.
// Make all go routine be able to access the struct
// lock

const (
	PROCESSROOT   = "/Users/garethharris/"
	MAXPROCCESSES = 170 // To limit the amount of goroutine
)

type ProcesService struct {
	dspro *dirnfiles.Dirs
	

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
	maxprocessch := make(chan int, MAXPROCCESSES) // To limit the amount of goroutine

	for _, files := range p.dspro.Files {

		maxprocessch <- 1 // To limit the amount of goroutine
		go func(files *dirnfiles.Dirinfo) {
			process(files)
			<-maxprocessch // To limit the amount of goroutine
		}(files)

	}

}

func process(files *dirnfiles.Dirinfo) {

	x := khash.Sha512fn(khash.Filebytes(files.Path))
	
	fmt.Printf("\nKey: %d = %s with %d bytes. CRC %x \n\n",
		files.Key, files.Path, files.Fsize, x)

	//time.Sleep(1000 * time.Millisecond)
	
	
	/*for { //loop 1024 bytes at a time move 1 byte at a time

		break // when  slice less than 1024
		- get BlockHashSha512 
		- get BlockFNV64      
		- start position
		- location - files.Key
		
		if ok := HashBlahmap[GlobalBlahBlock]; !ok{
	 		if HashBlahmap fileexist (the filename = GlobalBlahBlock.BlockHashSha512)  {
	 			//sync.Rlock
				//open file to read and save to HashBlahmap 
				//sync.RUnlock	
				
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
		//sync.lock
		
		//Add to HashBlahmap with func AddHashBlahmap below
		//AddHashBlahmap(globalBlahBlk GlobalBlahBlock, collision Collisions,
					location Locations, blockStatus BlockStatus) 

		
		// 1. Save to file. 
		//2. remove element(GlobalBlahBlock) from the front of list BlahMemoryList if list is greater than MAXMEMORYBLAH
			//and also delete from map
			
		//sync.unlock	
		


	}*/	


}
