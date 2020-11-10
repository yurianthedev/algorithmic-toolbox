#include <iostream>

unsigned long long int gcd(unsigned long long int a, unsigned long long int b);

unsigned long long int lcm(unsigned long long int a, unsigned long long int b);

int main() {
	unsigned long long int a, b;

	std::cin >> a >> b;
	std::cout << lcm(a, b);

	return 0;
}

unsigned long long int gcd(unsigned long long int a, unsigned long long int b) {
	if (b == 0) {
		return a;
	}
	unsigned long long int reminder = a % b;

	return gcd(b, reminder);
}

unsigned long long int lcm(unsigned long long int a, unsigned long long int b) {
	return (a * b) / gcd(a, b);
}
