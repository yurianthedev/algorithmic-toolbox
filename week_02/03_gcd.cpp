#include <iostream>

unsigned long long int gcd(unsigned long long int a, unsigned long long int b);

int main() {
	unsigned long long int a, b;

	std::cin >> a >> b;
	std::cout << gcd(a, b);

	return 0;
}

unsigned long long int gcd(unsigned long long int a, unsigned long long int b) {
	if (b == 0) {
		return a;
	}
	unsigned long long int reminder = a % b;

	return gcd(b, reminder);
}