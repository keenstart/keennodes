package blah

import "fmt"
import "bytes"
import "crypto/sha512"
import "hash/crc32"
import "hash/crc64"
import "hash/fnv"

//For Global checks -- use a FIFO to avoid loading too much blah in memory
type HashBlahmap  map[[]byte][uint][unit]*Blah //Blahblock sha512 - this is the 1024bytes{}
	//unit collison/unit location

/* wk -- not using
type Blah struct {
	BlahBlockSha512 []byte
	//wk - blah       map[int]*Dirinfo
}*/

type Blah struct {
	Startposition int16
	BlockChecksum  uint64
}
