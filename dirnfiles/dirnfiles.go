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
	PROCESSROOT = "./"
	PROCESSEXT  = ".jpg,.JPG,.PNG,.png"
	BLOBFILE    = "/tmp/blob.bl"
	BLOCKSIZE   = 1024
)

// Hold the properties of the blahs
type Dirinfo struct {
	Key          int
	Path         string
	Fsize        int64
	Name         string //remove
	Modtime      string //remove
	Mode         string // remove
	FileChecksum uint64
}

// To keep a list of Blah (files that contain byte block to make other files)
// files location(address)
type Dirs struct {
	Files map[int]*Dirinfo
}

func NewDirs() *Dirs {
	return &Dirs{Files: make(map[int]*Dirinfo)}
}

func NewDirinfo(key int, path string, fsize int64, name string, modtime string, mode string) *Dirinfo {

	chksm := khash.Hashcrc64(khash.Filebytes(path))

	return &Dirinfo{
		Key:          key,
		Path:         path,
		Fsize:        fsize,
		Name:         name,
		Modtime:      modtime,
		Mode:         mode,
		FileChecksum: chksm,
	}

}

// Search system for qualified files to uses as blah files
func (d *Dirs) GetDirsfile() error {

	var key int

	err := filepath.Walk(PROCESSROOT, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() && f.Mode().IsRegular() && f.Size() > BLOCKSIZE {

			if strings.Contains(PROCESSEXT, filepath.Ext(path)) == true &&
				len(filepath.Ext(path)) > 1 {
				dd := NewDirinfo(key, path, f.Size(), f.Name(), f.ModTime().String(), f.Mode().String())
				d.Files[key] = dd
				key++

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

	return nil
}

// Display the blahs files and there properties
func (d *Dirs) DisplayPath() {
	fmt.Println("Display")
	for _, value := range d.Files {
		fmt.Printf("Key %d == %s with %d bytes. Name = %s, modify time = %s, file mode = %s FileChecksum %x\n",
			value.Key, value.Path, value.Fsize, value.Name, value.Modtime, value.Mode, value.FileChecksum)

	}

}

//Load  Blah files
func (d *Dirs) GetFiles() error {
	err := gopfile.Load(BLOBFILE, d)

	if err != nil {
		return err
	}

	return nil

}

// Save blah files
func (d *Dirs) SetFiles() error {
	err := gopfile.Save(BLOBFILE, d)

	if err != nil {
		return err
	}

	return nil

}
