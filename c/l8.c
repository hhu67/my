#include <stdio.h>
#include <string.h>

int main() {
    char x[100] = "hello";
    x[0] = 'H';
    x[1] = 'a';
    int index = 2;
    strcpy(&x[index], &x[index + 1]);
    printf("%s\n", x);
    int len = strlen(x);
    printf("%d\n", len);
}