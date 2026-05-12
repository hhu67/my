#include <stdio.h>

int main() {
    char x[100] = "False";
    while(1) {
        printf("%s\n", x);
        if(x == "True") {
            return 0;
        }
    }
}