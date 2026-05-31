#include <iostream>
#include <random>
#include <string>

int all() {
	std::random_device rd;
	std::mt19937 gen(rd());
	std::uniform_int_distribution<int> distrib(1, 100);
	int i = 0;
	while(i <= 3) {
		int x1 = distrib(gen);
		int x2 = distrib(gen);
		int x3 = x1+x2;
		int pol;
		std::cout << x1 << "+" << x2 << "=?" << std::endl;
		std::cin >> pol;
		if(pol == x3) {
			std::cout << "Верно!" << std::endl;
			std::cout << x1 << "+" << x2 << "=" << x3 << std::endl;
		}
		else {
			std::cout << "Не верно!" << std::endl;
			std::cout << x1 << "+" << x2 << "=" << x3 << std::endl;
			return 0;
		}
	}
	return 1;
}

int main() {
	while(1) {
		std::string polz;
		std::cout << "1. Начать\n2. Закончить" << std::endl;
		std::cin >> polz;
		if(polz == "1") {
			all();
		}
		else if(polz == "2") {
			return 0;
		}
	}
	return 0;
}
