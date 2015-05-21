package khash

import (
	_ "bytes"
	"crypto/sha512"
	"encoding/binary"
	_ "fmt"
	"hash/crc32"
	"hash/crc64"
	"hash/fnv"
	"io/ioutil"
	"log"
)

const (
	SIZEOF_INT32 = 4
	SIZEOF_INT64 = 8
) // bytes

func ConverttoInt32(covert []byte) [16]int32 {

	//fmt.Printf("\nConvert32: bytes covert %d count %d\n", covert, len(covert))

	var data [16]int32 //make([]int32, len(covert)/SIZEOF_INT32)

	for i := range data {
		// assuming little endian
		data[i] = int32(binary.LittleEndian.Uint32(covert[i*SIZEOF_INT32 : (i+1)*SIZEOF_INT32]))
	}
	//fmt.Printf("data bytes covert %x\n", data)
	return data
}

func ConverttoInt64(covert []byte) [8]int64 {

	//fmt.Printf("\nConvert32: bytes covert %d count %d\n", covert, len(covert))

	var data [8]int64 //make([]int32, len(covert)/SIZEOF_INT32)

	for i := range data {
		// assuming little endian
		data[i] = int64(binary.LittleEndian.Uint64(covert[i*SIZEOF_INT64 : (i+1)*SIZEOF_INT64]))
	}
	//fmt.Printf("data bytes covert %x\n", data)
	return data
}

func HashFNV64(s []byte) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	_, err := h.Write(s)
	if err != nil {
		log.Fatal(err)
		//return nil
	}

	//fmt.Printf("read FNV64 %d bytes.", count) //Debug

	return h.Sum64()
}

func Hashcrc64(s []byte) uint64 {
	h := crc64.New(crc64.MakeTable(crc64.ECMA))
	_, err := h.Write(s)
	if err != nil {
		log.Fatal(err)
		//return nil
	}
	//fmt.Printf("read CRC64 %d bytes. Checksum Func ->  ", count) //Debug

	return h.Sum64()

}

func Hashcrc32(s []byte) uint32 {
	h := crc32.NewIEEE()
	_, err := h.Write(s)
	if err != nil {
		log.Fatal(err)
		//return nil
	}

	///fmt.Printf("read CRC32 %d bytes.  ->  ", count) //Debug

	return h.Sum32()
}

func Sha512fn(s []byte) []byte {
	h512 := sha512.New()

	_, err := h512.Write(s)
	if err != nil {
		log.Fatal(err)
		//return nil
	}

	//fmt.Printf("read SHA512 %d bytes.  2nd ->  ", count) //Debug
	//fmt.Printf("%x/n ", h512.Sum(nil))                   //Debug

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
