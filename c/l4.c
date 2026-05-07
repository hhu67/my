#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int main() {
    srand(time(NULL));
    int x = rand() % 11;
    int s = 0;
    while(1) {
        int pol;
        printf("Number ");
        scanf("%d", &pol);
        if(pol == x) {
            printf("the number is guessed\n");
            break;
        }
        else {
            printf("the number is not guessed\n");
            if(pol > x) {
                printf("less\n");
            }
            else {
                printf("more\n");
            }
            s++;
            if(s == 3) {
                printf("Game over\n");
                break;
            }
        }
    }
    return 0;
}