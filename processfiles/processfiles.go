package processfiles

import (
	"fmt"

	"github.com/keenstart/keennodes/dirnfiles"
	"github.com/keenstart/keennodes/khash"

	_ "github.com/keenstart/keennodes/gopfile"
)

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

		x := khash.Sha512fn(khash.Filebytes(res.Path))

		fmt.Printf("Key: %d = %s with %d bytes. CRC %x \n\n",
			key, res.Path, res.Fsize, x)

	}

}
