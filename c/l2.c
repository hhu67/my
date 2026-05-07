#include <stdio.h> 
int main() {
	char x[100];
	printf("Кто сергей? ");
	scanf("%99s", x);
	printf("Сергей %s\n", x);
}
