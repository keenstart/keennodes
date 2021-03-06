package gopfile

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
)

//Utilities to load data structure to files
func Load(filepath string, i interface{}) error {
	n, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("File Read error:", err)
		return err
	}

	p := bytes.NewBuffer(n)
	dec := gob.NewDecoder(p)

	err = dec.Decode(i)
	if err != nil {
		log.Fatal("Decode error:", err)
		return err
	}
	return nil
}

//Utilities to save data structure to files
func Save(filepath string, i interface{}) error {
	m := new(bytes.Buffer)
	enc := gob.NewEncoder(m)

	err := enc.Encode(i)
	if err != nil {
		log.Fatal("Encode error:", err)
		return err
	}

	err = ioutil.WriteFile(filepath, m.Bytes(), 0600)
	if err != nil {
		log.Fatal("File Write error:", err)
		return err
	}
	return nil
}
