#include <iostream>

int minimumExchange(int n) {
	int coin = 0;
	int counter = 0;
	int coins[] = {10, 5, 1};

	while (n > 0) {
		if (n >= coins[coin]) {
			n -= coins[coin];
			counter++;
		} else {
			coin++;
		}
	}

	return counter;
}

int main() {
	int n;
	std::cin >> n;

	std::cout << minimumExchange(n);

	return 0;
}
