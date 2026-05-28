#include <iostream>
#include <string>

int polz(void) {
	while(1) {
		std::string pol;
		std::cout << "Знак" << std::endl;
		std::cin >> pol;
		double x1;
    	std::cout << "Первое число" << std::endl;
		std::cin >> x1;
		double x2;
		std::cout << "Второе число" << std::endl;
		std::cin >> x2;
		if(pol == "+") {
			std::cout << x1 << pol << x2 << "=" << x1+x2 << std::endl;
		}
		else if(pol == "-") {
			std::cout << x1 << pol << x2 << "=" << x1-x2 << std::endl;
		}
		else if(pol == "exit") {
			break;
		}
		else if(pol == "*") {
			std::cout << x1 << pol << x2 << "=" << x1*x2 << std::endl;
		}
		else if(pol == "/") {
			std::cout << x1 << pol << x2 << "=" << x1/x2 << std::endl;
		}
		else {
			std::cout << "Неверный знак" << std::endl;
			break;
		}
	}
	return 1;	
}

int main() {
	while(1) {
		std::cout << "1. Начать\n2. Закончить" << std::endl;
		int pol1;
		std::cout << "Введите число" << std::endl;
		std::cin >> pol1;
		if(pol1 == 1) {
			polz();
		}
		else if(pol1 == 2) {
			return 0;
		}
		else {
			std::cout << "Введите верное число" << std::endl;
		}
	}
	return 0;
}
