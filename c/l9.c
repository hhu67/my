#include <stdio.h>
#include <string.h>

int main() {
    char pol[100];
    printf("Enter word ");
    scanf("%s", pol);
    int len = strlen(pol);
    for(int i = 0; i < len / 2; i++) {
    char temp = pol[i];
    pol[i] = pol[len - 1 - i];
    pol[len - 1 - i] = temp;
    }
    printf("%s\n", pol);
    return 0;
}