#include <iostream>
int main() {
	int x;
	std::cout << "Число " << std::endl;
	std::cin >> x;
	int x2;
	std::cout << "Второе число " << std::endl;
	std::cin >> x2;
	int sum = x+x2;
	std::cout << x << "+" << x2 << "=" << sum << std::endl;
}
