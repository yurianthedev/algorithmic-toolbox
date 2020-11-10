#include <iostream>
#include <utility>
#include <vector>
#include <algorithm>

bool sorter(std::pair<double, double> i, std::pair<double, double> j) {
	return (i.first / i.second) > (j.first / j.second);
}

void sortItems(std::vector<std::pair<double, double>> &items) {
	std::sort(items.begin(), items.end(), sorter);
}

double maxValue(double capacity, std::vector<std::pair<double, double>> &items) {
	int position = 0;
	double value = 0;

	while (capacity > 0 && position < items.size()) {
		if (items.at(position).second <= capacity) {
			capacity -= items.at(position).second;
			value += items.at(position).first;
		} else {
			double take = capacity;
			capacity -= take;
			value += items.at(position).first / items.at(position).second * take;
		}

		position++;
	}

	return value;
}

int main() {
	double n, m;
	std::vector<std::pair<double, double>> items;
	std::cin >> n >> m;

	for (int i = 0; double(i) < n; i++) {
		double value, weight;
		std::cin >> value >> weight;

		items.emplace_back(value, weight);
	}

	sortItems(items);
	printf("%.4f", maxValue(m, items));

	return 0;
}
