package processfiles

import (
	"fmt"
	_ "time"

	"github.com/keenstart/keennodes/dirnfiles"
	"github.com/keenstart/keennodes/khash"

	_ "github.com/keenstart/keennodes/gopfile"
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

		- get blockFNV
		- start position
		- location - mapkey

		if ok := blah[blockhash]; !ok{
	 		if create {
				//create hashstruct to save to file
			}else{
				//open file to save to hashstruct
			}
		}
		//add to blah
	}*/

}
