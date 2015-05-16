package processfiles

import (
	"fmt"

	"github.com/keenstart/keennodes/dirnfiles"
	"github.com/keenstart/keennodes/khash"

	_ "github.com/keenstart/keennodes/gopfile"
)
// Create a struct to store load blah( are hash name files)
// to memory instead of keep open  and close blab file.
// Make all go routine be able to access the struct 
// lock 

const (
	PROCESSROOT = "/Users/garethharris/"
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

	for key, res := range p.dspro.Files {
		// go func TODO pass Path 'res'.
		// Move khash.Sha512fn(khash.Filebytes(res.Path)) to go func too
		x := khash.Sha512fn(khash.Filebytes(res.Path))

		fmt.Printf("Key: %d = %s with %d bytes. CRC %x \n\n",
			key, res.Path, res.Fsize, x)

	}

}
/*

go func ?{

	if ok := blah[blockhash]; ok{
 		used hash

	}else{
	
		//create or open blah file to used
		if create {
			create hashstruct to save to file
		}else{
			open file to save to hashstruct
			
		}

	}
	
	for { //loop 1024 bytes at a time move 1 byte at a time 
		- get blockhash
		- start position
		break / when  slice less than 1024
	}


}


*/
