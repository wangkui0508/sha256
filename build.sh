# -march=armv8-a+crypto 
gcc -c -O3 -msse4.1 -msha *.c
ar -crs libsha256.a *.o

