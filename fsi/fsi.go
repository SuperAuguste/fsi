package fsi

/*

Format for Streamlet identifiers:

22_random_base/10_incremental_numbers
Hkh2bjNWQ4yxuMI++kTixA/0000000000
xxxxxxxxxxxxxxxxxxxxxx/nnnnnnnnnn

*/

import (
	"math/rand"
	"strconv"
	"time"
	"sync/atomic"
)

type FSI struct {
	Base  string
	Index uint64
}

const pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var randSrc = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateBase() string {

	bytes := make([]byte, 22)
	for i := 0; i < 22; i++ {

		bytes[i] = pool[randSrc.Intn(52)]

	}
	return string(bytes)

}

func New() FSI {

	return FSI{

		Base:  generateBase(),
		Index: 0,
	}

}

func (fsi *FSI) Generate() string {

	if fsi.Index == 9999999999 {

		fsi.Base = generateBase()
		fsi.Index = 0

	}

	index := strconv.FormatUint(fsi.Index, 10)
	l := len(index)

	if l == 1 {

		index = "000000000" + index

	} else if l == 2 {

		index = "00000000" + index

	} else if l == 3 {

		index = "0000000" + index

	} else if l == 4 {

		index = "000000" + index

	} else if l == 5 {

		index = "00000" + index

	} else if l == 6 {

		index = "0000" + index

	} else if l == 7 {

		index = "000" + index

	} else if l == 8 {

		index = "00" + index

	} else if l == 9 {

		index = "0" + index

	}

	result := fsi.Base + "/" + index
	atomic.AddUint64(&fsi.Index, 1)
	return result

}
