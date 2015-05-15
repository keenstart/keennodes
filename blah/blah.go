package blah

import "fmt"
import "bytes"
import "crypto/sha512"
import "hash/crc32"
import "hash/crc64"
import "hash/fnv"

// type Blahmap  map[[]byte][uint]*Blah //unit collison

type Blah struct {
	BlahSha512 []byte
	//wk - blah       map[int]*Dirinfo
}

type Blah struct {
	Collision uint64
	Location  uint64
}
