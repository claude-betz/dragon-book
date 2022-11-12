#include<stdio.h>

struct id {
    int x: 1;
    int y: 10;
    int z: 14;
    int w: 8;
}A;

int main() {
    printf("%lu", sizeof(A));
    return 0;
}