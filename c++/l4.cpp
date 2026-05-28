#include <iostream>
#include <string>
int main() {
	std::string x;
	std::cout << "Кто сергей" << std::endl;
	std::getline(std::cin, x);
	std::cout << "сергей " << x << std::endl;
}
