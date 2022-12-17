#include <iostream>
#include <vector>
#include <cmath>
#include <climits>
#include <algorithm>

using namespace std;

auto selection_sort(vector<int> &a) -> void {
	const int n = a.size();
	for (int i = 0, k = 0; i < n - 1; ++i) {
		for (int j = i; j < n; ++j) {
			if (a[k] > a[j])
				k = j;
		}
		swap(a[i], a[k]);
	}
}

int main() {
	vector<int> a { 1, 16, 4, 25, 9, 55, 2, 1, 0 };
	selection_sort(a);

	for (auto e : a)
		cout << e << ' ';
	cout << endl;
}
