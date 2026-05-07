#include <stdio.h>

int main() {
    int pol;
    printf("Enter number ");
    scanf("%d", &pol);
    if(pol > 9) {
        printf("Error: number must be from 1 to 9\n");
        return 0;
    }
    else if(pol < 1) {
        printf("Error: number must be from 1 to 9\n");
        return 0;
    }
    int x = 1;
    int s = 0;
    while(1) {
        printf("%d * %d = %d\n", pol, x, pol*x);
        x++;
        s++;
        if(s < 10) {
            continue;
        }
        else {
            break;
        }
    }
}