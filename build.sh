gcc -c -O3 -msse4.1 -msha *.c
ar -crs ../libsha256.a *.o

