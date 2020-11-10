#include <iostream>

unsigned long long int greatestFibonacci(int n);

int main() {
	int n;

	std::cin >> n;
	std::cout << greatestFibonacci(n);

	return 0;
}

unsigned long long int greatestFibonacci(int n) {
	unsigned long long int previous = 0;
	unsigned long long int current = 1;
	unsigned long long int temp;

	if (n == 0) {
		return 0;
	}
	for (int i = 2; i <= n; i++) {
		temp = current + previous;
		previous = current;
		current = temp;
	}

	return current;
}
