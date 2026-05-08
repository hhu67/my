#include <stdio.h>
#include <string.h>

int main() {
    char pol[100];
    printf("Enter word ");
    scanf("%s", pol);
    int len = strlen(pol);
    if(len > 50) {
        printf("Error: too long\n");
        return 0;
    }
    strrev(pol);
    printf("Reversed: %s\n", pol);
    return 0;
}