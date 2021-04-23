test `uname -p ` == aarch64 || gcc -c -O3 -msse4.1 -msha *.c
test `uname -p ` == aarch64 && gcc -c -O3 -march=armv8-a+crypto *.c
ar -crs libsha256.a *.o

