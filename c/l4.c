#include <stdio.h>
int main() {
    int x = 7;
    int s = 0;
    while(1) {
        int pol;
        printf("Number ");
        scanf("%d", &pol);
        if(pol == 7) {
            printf("the number is guessed\n");
            break;
        }
        else {
            printf("the number is not guessed\n");
            s++;
            if(s == 3) {
                printf("Game over\n");
                break;
            }
        }
    }
    return 0;
}