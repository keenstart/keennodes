package blah

import "fmt"
import "bytes"
import "crypto/sha512"
import "hash/crc32"
import "hash/crc64"
import "hash/fnv"

// type Blahmap  map[[]byte][uint][unit]*Blah //sha512/unit collison/unit location
/* wk -- not using
type Blah struct {
	BlahSha512 []byte
	//wk - blah       map[int]*Dirinfo
}*/

type Blah struct {
	Startposition int16
	BlockChecksum  uint64
}
