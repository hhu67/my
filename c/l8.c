#include <stdio.h>
#include <string.h>

int main() {
    char pol[100];
    printf("Enter word ");
    scanf("%s", pol);
    char x[100] = "hello";
    x[0] = 'H';
    x[1] = 'i';
    strcpy(&x[2], &x[5]);
    strcat(x, "!");
    if(strcmp(x, pol) == 0) {
        printf("equal\n");
    }
    else {
        printf("not equal\n");
    }
    int len = strlen(x), len2 = strlen(pol);
    printf("Symbol: %d\n%s\nSymbol your word: %d", len, x, len2);
}