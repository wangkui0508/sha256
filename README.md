# sha256
Fast sha256 implementation in C, based on:
1. https://github.com/B-Con/crypto-algorithms (https://github.com/B-Con/crypto-algorithms/commit/cfbde48414baacf51fc7c74f275190881f037d32)
2. https://github.com/noloader/SHA-Intrinsics

It utilizes [SHA-NI](https://en.wikipedia.org/wiki/Intel_SHA_extensions) on x86-64 machines.
