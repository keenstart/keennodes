package dirnfiles

import (
	_ "bufio"
	_ "encoding/gob"
	"fmt"
	_ "io/ioutil"
	_ "log"
	"os"
	"path/filepath"
	//"io"
	_ "bytes"
)

type Dirinfo struct {
	Path    string
	Fsize   int64
	Name    string
	Modtime string
	Mode    string
}

type Dirs struct {
	Files map[int]*Dirinfo
}

func NewDirs() *Dirs {
	return &Dirs{Files: make(map[int]*Dirinfo)}
}

func NewDirinfo(path string, fsize int64, name string, modtime string, mode string) *Dirinfo {

	return &Dirinfo{Path: path,
		Fsize:   fsize,
		Name:    name,
		Modtime: modtime,
		Mode:    mode,
	}

}

func (d *Dirs) GetDirs(path string) error {

	count := 0

	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {

		if f.IsDir() != true {
			//dd := &Dirinfo{path: path, fsize: f.Size(), name: f.Name(), modtime: f.ModTime().String()}
			dd := NewDirinfo(path, f.Size(), f.Name(), f.ModTime().String(), f.Mode().String())

			d.Files[count] = dd

			count++

			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil

}

func (d *Dirs) DisplayPath() {
	fmt.Println("Display")
	for _, value := range d.Files {
		fmt.Printf("%s with %d bytes. Name = %s, modify time = %s, file mode = %s\n",
			value.Path, value.Fsize, value.Name, value.Modtime, value.Mode)
	}

}
