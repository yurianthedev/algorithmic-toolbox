#include <iostream>
#include <vector>
#include <algorithm>

long long int bestValue(int n, std::vector<long long int> &profits, std::vector<long long int> &clicks) {
	long long int value = 0;
	for (int i = 0; i < n; i++) {
		value += profits.at(i) * clicks.at(i);
	}

	return value;
}

int main() {
	int n;
	std::vector<long long int> profits;
	std::vector<long long int> clicks;

	std::cin >> n;

	for (int i = 0; i < n; i++) {
		long long int profit;
		std::cin >> profit;
		profits.push_back(profit);
	}
	for (int i = 0; i < n; i++) {
		long long int click;
		std::cin >> click;
		clicks.push_back(click);
	}

	std::sort(profits.begin(), profits.end());
	std::sort(clicks.begin(), clicks.end());

	std::cout << bestValue(n, profits, clicks);

	return 0;
}
