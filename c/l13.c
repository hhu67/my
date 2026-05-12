#include <stdio.h>
#include <string.h>
#include <ctype.h>

int main() {
    char pol[100];
    printf("Enter word ");
    scanf("%s", pol);
    for(int i = 0; pol[i]; i++) {
        pol[i] = tolower(pol[i]);
    }
    char x[100];
    strcpy(x, pol);
    int len = strlen(x);
    for(int i = 0; i < len / 2; i++) {
    char temp = x[i];
    x[i] = x[len - 1 - i];
    x[len - 1 - i] = temp;
    }
    if(strcmp(pol, x) == 0) {
        printf("equal\n");
    }
    else {
        printf("not equal\n");
    }
}