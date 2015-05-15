package dirnfiles

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/keenstart/keennodes/gopfile"
	"github.com/keenstart/keennodes/khash"
)

const (
	PROCESSROOT = "/Users/garethharris/"
	PROCESSEXT  = ".jpg,.JPG,.PNG,.png" //,.PNG,.png

	BLOBFILE = "/tmp/blob.bl"
)

type Dirinfo struct {
	Path         string
	Fsize        int64
	Name         string
	Modtime      string
	Mode         string
	FileChecksum uint64
}

type Dirs struct {
	Files map[int]*Dirinfo
}

func NewDirs() *Dirs {
	return &Dirs{Files: make(map[int]*Dirinfo)}
}

func NewDirinfo(path string, fsize int64, name string, modtime string, mode string) *Dirinfo {

	chksm := khash.Hashcrc64(khash.Filebytes(path))

	return &Dirinfo{Path: path,
		Fsize:        fsize,
		Name:         name,
		Modtime:      modtime,
		Mode:         mode,
		FileChecksum: chksm,
	}

}

func (d *Dirs) GetDirsfile() error {

	count := 0

	err := filepath.Walk(PROCESSROOT, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() != true && f.Mode().IsRegular() {

			if strings.Contains(PROCESSEXT, filepath.Ext(path)) == true &&
				len(filepath.Ext(path)) > 1 {

				fmt.Println("EXT = ", filepath.Ext(path)) //Debug
				//dd := &Dirinfo{path: path, fsize: f.Size(), name: f.Name(), modtime: f.ModTime().String()}
				dd := NewDirinfo(path, f.Size(), f.Name(), f.ModTime().String(), f.Mode().String())

				d.Files[count] = dd

				count++

				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("Total files = ", count) //Debug

	return nil

}

func (d *Dirs) DisplayPath() {
	fmt.Println("Display")
	for _, value := range d.Files {
		fmt.Printf("%s with %d bytes. Name = %s, modify time = %s, file mode = %s FileChecksum %x\n",
			value.Path, value.Fsize, value.Name, value.Modtime, value.Mode, value.FileChecksum)

	}

}

func (d *Dirs) GetFiles() error {
	err := gopfile.Load(BLOBFILE, d)

	if err != nil {
		return err
	}

	return nil

}

func (d *Dirs) SetFiles() error {
	err := gopfile.Save(BLOBFILE, d)

	if err != nil {
		return err
	}

	return nil

}
