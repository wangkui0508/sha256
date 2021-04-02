package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"

	"github.com/coinexchain/randsrc"
)

/*
#include "../sha256.h"
#include <stdint.h>

struct bytes32 {
    uint8_t bytes[32];
};

struct bytes32 sha256go(_GoString_ gostr) {
    struct bytes32 result;
    size_t length = _GoStringLen(gostr);
    const char * data = _GoStringPtr(gostr);
    SHA256_CTX ctx;
    sha256_init(&ctx);
    sha256_update(&ctx, (const BYTE *)data, length);
    sha256_final(&ctx, (uint8_t*)&result.bytes[0]);
    return result;
}

*/
// #cgo LDFLAGS: -lsha256 -L..
import "C"

func Sha256(s string) (result [32]byte){
	hash := C.sha256go(s)
	for i:=0; i<32; i++ {
		result[i] = byte(hash.bytes[i]);
	}
	return
}

func fuzzTest() {
	randFilename := os.Getenv("RANDFILE")
	if len(randFilename) == 0 {
		fmt.Printf("No RANDFILE specified. Exiting...")
		return
	}
	roundCount, err := strconv.Atoi(os.Getenv("RANDCOUNT"))
	if err != nil {
		fmt.Printf("Fuzz test not run. Error when Parsing RANDCOUNT: %#v\n", err)
		return
	}

	rs := randsrc.NewRandSrcFromFile(randFilename)

	for i := 0; i< roundCount; i++ {
		length := 1 + int(rs.GetUint32()) % 300
		s := rs.GetString(length)
		ref := sha256.Sum256([]byte(s))
		imp := Sha256(s)
		if ref != imp {
			panic("Not Equal")
		}
	}
	fmt.Printf("%d test finished\n", roundCount)
}

func main() {
	hello := "Hello! World 12122Hello! World 12122Hello! World 121222Hello! World 121222Hello! World 121222Hello! World 12122222Hello! World 121222"
	fmt.Printf("length %d\n", len(hello))
	fmt.Printf("hash.c  %#v\n", Sha256(hello))
	fmt.Printf("hash.go %#v\n", sha256.Sum256([]byte(hello)))
	fuzzTest()
}


