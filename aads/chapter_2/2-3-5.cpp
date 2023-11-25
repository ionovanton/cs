#include <iostream>
#include <vector>
#include <cmath>
#include <climits>
#include <algorithm>

using namespace std;

auto recursive_insertion_sort(vector<int> &a, int n) -> void {
	if (n > 0) {
		recursive_insertion_sort(a, n - 1);
		int j = n - 1, k = a.size();
		for (int i = n; i < k; ++i) {
			if (a[i] < a[j])
				j = i;
		}
		swap(a[j], a[n - 1]);
	}
}

int main() {
	vector<int> a { 1, 16, 4, 25, 9, 55, 2, 1, 0 };
	recursive_insertion_sort(a, a.size());

	for (auto e : a)
		cout << e << ' ';
	cout << endl;
}
