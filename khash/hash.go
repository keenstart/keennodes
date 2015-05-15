package khash

import (
	_ "bytes"
	"crypto/sha512"
	"fmt"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"io/ioutil"
	"log"
)

func HashFNV64(s []byte) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	count, err := h.Write(s)
	if err != nil {
		log.Fatal(err)
		//return nil
	}

	fmt.Printf("read FNV64 %d bytes.", count) //Debug

	return h.Sum64()
}

func Hashcrc64(s []byte) uint64 {
	h := crc64.New(crc64.MakeTable(crc64.ECMA))
	count, err := h.Write(s)
	if err != nil {
		log.Fatal(err)
		//return nil
	}
	fmt.Printf("read CRC64 %d bytes. Checksum Func ->  ", count) //Debug

	return h.Sum64()

}

func Hashcrc32(s []byte) uint32 {
	h := crc32.NewIEEE()
	count, err := h.Write(s)
	if err != nil {
		log.Fatal(err)
		//return nil
	}

	fmt.Printf("read CRC32 %d bytes.  ->  ", count) //Debug

	return h.Sum32()
}

func Sha512fn(s []byte) []byte {
	h512 := sha512.New()

	count, err := h512.Write(s)
	if err != nil {
		log.Fatal(err)
		//return nil
	}

	fmt.Printf("read SHA512 %d bytes.  2nd ->  ", count) //Debug
	fmt.Printf("%x/n ", h512.Sum(nil))                   //Debug

	return h512.Sum(nil)
}

func Filebytes(path string) []byte {
	n, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal("File Read error:", err)
		panic("File Read error:Filebytes") //debug
		//return _, err
	}

	return n

}
