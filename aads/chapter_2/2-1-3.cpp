#include <iostream>
#include <vector>

using namespace std;

void insertion_sort_asc(vector<int> &a) {
	const int n = a.size();
	for (int i = 1, j; i < n; ++i) {
		auto key = a[i];
		j = i - 1;
		while (j > -1 && a[j] > key) {
			a[j + 1] = a[j];
			--j;
		}
		a[j + 1] = key;
	}
}

void insertion_sort_desc(vector<int> &a) {
	const int n = a.size();
	for (int i = 1, j; i < n; ++i) {
		auto key = a[i];
		j = i - 1;
		while (j > -1 && a[j] < key) {
			a[j + 1] = a[j];
			--j;
		}
		a[j + 1] = key;
	}
}

int main() {
	vector<int> a { 1, 16, 4, 25, 9, 55, 2, 1, 0 };
	vector<int> b { 1, 16, 4, 25, 9, 55, 2, 1, 0 };

	insertion_sort_asc(a);
	for (auto e : a)
		cout << e << ' ';
	cout << endl;
	insertion_sort_desc(b);
	for (auto e : b)
		cout << e << ' ';
	cout << endl;
	
}
