#include <iostream>
#include <string>

int main() {
	int s = 0;
	while(s <= 3) {
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
			s++;
		}
		else if(pol == "-") {
			std::cout << x1 << pol << x2 << "=" << x1-x2 << std::endl;
			s++;
		}
		else if(pol == "*") {
			std::cout << x1 << pol << x2 << "=" << x1*x2 << std::endl;
			s++;
		}
		else if(pol == "/") {
			std::cout << x1 << pol << x2 << "=" << x1/x2 << std::endl;
			s++;
		}
		else {
			std::cout << "Неверный знак" << std::endl;
			return 0;
		}
	}
	return 0;
}
