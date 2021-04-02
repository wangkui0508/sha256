package main

import (
	"fmt"

	"crypto/sha256"
)

/*
#include "sha256lib/sha256.h"
#include <stdint.h>

//struct bytes32 sha256go(_GoString_ gostr) {
//    size_t length = _GoStringLen(gostr);
//    const char * data = _GoStringPtr(gostr);
//    return sha256((const uint8_t *)data, length);
//}

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
// #cgo LDFLAGS: -lsha256 -L.
import "C"

func Sha256(s string) (result [32]byte){
	hash := C.sha256go(s)
	for i:=0; i<32; i++ {
		result[i] = byte(hash.bytes[i]);
	}
	return
}

func main() {
	//ok := C.supports_sha_ni();
	//fmt.Println("Here %d", ok)
	hello := "Hello! World 12122Hello! World 12122Hello! World 121222Hello! World 121222Hello! World 121222Hello! World 12122222Hello! World 121222"
	fmt.Printf("length %d\n", len(hello))
	fmt.Printf("hash.c  %#v\n", Sha256(hello))
	fmt.Printf("hash.go %#v\n", sha256.Sum256([]byte(hello)))
}


