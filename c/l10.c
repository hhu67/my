#include <stdio.h>
#include <string.h>

int main() {
    char str[100];
    printf("Enter phrase ");
    scanf("%s", str);
    int len = strlen(str);
    if(len > 50) {
        printf("too long");
        return 0;
    }
    for(int i = 0; i < len; i++) {
        printf("%c\n", str[i]);
    }
}