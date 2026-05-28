#include <iostream>
int main() {
	int x;
	std::cout << "Число" << std::endl;
	std::cin >> x;
	if(x > 10) {
		std::cout << "Число больше 10" << std::endl;
	}
	else {
		std::cout << "Число меньше 10" << std::endl;
	}
}
