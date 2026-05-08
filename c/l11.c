#include <stdio.h>
#include <string.h>
#include <unistd.h>

int main() {
    int x = 0;
    char str[100];
    printf("Enter word ");
    scanf("%s", str);
    int len = strlen(str);
    for(int i = 0; i < len; i++) {
        if(str[i] == 'a') {
            x++;
        }
        else if(str[i] == 'e') {
            x++;
        }
        else if(str[i] == 'i') {
            x++;
        }
        else if(str[i] == 'o') {
            x++;
        }
        else if(str[i] == 'u') {
            x++;
        }
        else if(str[i] == 'y') {
            x++;
        }
        else if(str[i] == 'A') {
            x++;
        }
        else if(str[i] == 'E') {
            x++;
        }
        else if(str[i] == 'I') {
            x++;
        }
        else if(str[i] == 'O') {
            x++;
        }
        else if(str[i] == 'U') {
            x++;
        }
        else if(str[i] == 'Y') {
            x++;
        }
    }
    printf("Vowels: %d\n", x);
    sleep(2);
}