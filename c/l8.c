#include <stdio.h>
#include <string.h>

int main() {
    char x[100] = "hello";
    x[0] = 'H';
    x[1] = 'i';
    strcpy(&x[2], &x[5]);
    strcat(x, "!");
    int len = strlen(x);
    printf("Symbol: %d\n%s", len, x);
}