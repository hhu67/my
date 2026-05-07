#include <stdio.h>

int main() {
    int pol;
    printf("Enter number ");
    scanf("%d", &pol);
    int x1 = pol % 10;
    int x2 = (pol / 10) % 10;
    int x3 = ((pol / 10) / 10) % 10;
    int sum = x1+x2+x3;
    printf("%d\n", sum);
}