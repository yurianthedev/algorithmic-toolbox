#include <iostream>
#include <vector>

int minimumStops(int m, std::vector<int> &stops) {
	int numberOfStops = 0;
	int currentStop = 0;
	int lastFill = 0;

	while (currentStop < stops.size()) {
		if (stops.at(currentStop) - stops.at(lastFill) <= m) {
			currentStop++;
			continue;
		}
		if (currentStop == lastFill + 1) {
			return -1;
		}
		lastFill = currentStop - 1;
		numberOfStops++;
	}

	return numberOfStops;
}

int main() {
	int d, m, n;
	std::cin >> d >> m >> n;
	std::vector<int> stops;

	stops.push_back(0);
	for (int i = 0; i < n; i++) {
		int stop;
		std::cin >> stop;
		stops.push_back(stop);
	}
	stops.push_back(d);

	std::cout << minimumStops(m, stops);

	return 0;
}
