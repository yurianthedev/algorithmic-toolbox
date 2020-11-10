#include <iostream>
#include <vector>

int pisanoLength(int m) {
	int i = 2;
	unsigned long long int previous = 0;
	unsigned long long int current = 1;
	unsigned long long int temp;

	if (m % 10 == 0) {
		return m * 6;
	}
	while (true) {
		temp = (current + previous) % m;
		previous = current;
		current = temp;

		i++;
		if (current == 1 && previous == 0) {
			return i - 2;
		}
	}
}

unsigned long long int pisanoAt(unsigned long long int reminder, int m) {
	unsigned long long int previous = 0;
	unsigned long long int current = 1;
	unsigned long long int temp;

	if (reminder == 1) {
		return current;
	} else if (reminder == 0) {
		return previous;
	}
	for (int i = 2; i <= reminder; i++) {
		temp = (current + previous) % m;
		previous = current;
		current = temp;
	}

	return current;
}

int main() {
	unsigned long long int n;
	int m;
	std::cin >> n >> m;

	int pLength = pisanoLength(m);
	std::cout << pLength << std::endl;
	int reminder = (int) (n % pLength);

	std::cout << pisanoAt(reminder, m);

	return 0;
}
