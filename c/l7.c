#include <stdio.h>
#include <string.h>

int main() {
    int pol;
    printf("Enter number: ");
    if (scanf("%d", &pol) != 1);

    char st[100] = ""; 
    int x = 0;

    if (pol >= 100) pol = 99;

    while(x < pol) {
        int len = strlen(st);
        st[len] = '*';
        st[len + 1] = '\0';
        x++;
        printf("%s\n", st);
    }
}