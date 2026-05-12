#include <stdio.h>

int main() {
    int num[5];
    for (int i = 0; i < 5; i++) {
        printf("Enter number ");
        scanf("%d", &num[i]);
    }

    int max = num[0];
    for (int i = 1; i < 5; i++) {
        if (num[i] > max) {
            max = num[i];
        }
    }

    printf("Maximum number: %d\n", max);
    return 0;
}