#include <stdio.h>
#include "../sha256.h"

int main(void) {
	BYTE buf[SHA256_BLOCK_SIZE*2];
	for(int i=0; i<SHA256_BLOCK_SIZE*2; i++) {
		buf[i] = i;
	}
	SHA256_CTX ctx;
	for(int count=0; count<100*10000; count++) {
		sha256_init(&ctx);
		sha256_update(&ctx, buf, SHA256_BLOCK_SIZE*2);
		sha256_final(&ctx, buf);
		for(int i=0; i<SHA256_BLOCK_SIZE; i++) {
			buf[i+SHA256_BLOCK_SIZE] = ~buf[i];
		}
	}
	return 0;
}
