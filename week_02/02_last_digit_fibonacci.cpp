#include <iostream>

short lastDigitOfNth(int n);

int main() {
	int n;

	std::cin >> n;
	std::cout << lastDigitOfNth(n);

	return 0;
}

short int lastDigitOfNth(int n) {
	if (n == 0) {
		return 0;
	} else if (n == 1) {
		return 1;
	}

	short int digits[n];

	digits[0] = 0;
	digits[1] = 1;

	for (int i = 2; i <= n; i++) {
		digits[i] = (short int) ((digits[i - 1] + digits[i - 2]) % 10);
	}

	return digits[n];
}
