#include <iostream>
#include <vector>
#include <cmath>
#include <climits>
#include <algorithm>

using namespace std;

auto merge(vector<int> &a, int l, int m, int r) -> void {
	const int nl = m - l + 1;
	const int nr = r - m;
	int x[nl], y[nr];

	for (int i = 0; i < nl; ++i)
		x[i] = a[i + l];
	for (int i = 0; i < nr; ++i)
		y[i] = a[m + 1 + i];

	int i = 0, j = 0, k = l;
	
	while (i < nl && j < nr) {
		if (x[i] <= y[j])
			a[k] = x[i++];
		else
			a[k] = y[j++];
		++k;
	}
	while (i < nl) {
		a[k] = x[i];
		++k, ++i;
	}
	while (j < nr) {
		a[k] = y[j];
		++k, ++j;
	}
}

auto merge_sort(vector<int> &a, int l, int r) -> void {
	if (l >= r)
		return ;
	int m = (r + l) / 2;
	merge_sort(a, l, m);
	merge_sort(a, m + 1, r);
	merge(a, l, m, r);
}

auto inversion(vector<int> &a, int l, int m, int r) -> int {
	const int nl = m - l + 1; // length of array to the left
	const int nr = r - m;
	int x[nl], y[nr], q = 0;

	for (int i = 0; i < nl; ++i)
		x[i] = a[i + l];
	for (int i = 0; i < nr; ++i)
		y[i] = a[m + 1 + i];

	int i = 0, j = 0, k = l;
	
	while (i < nl && j < nr) {
		if (x[i] <= y[j]) {
			a[k] = x[i++];
		} else {
			q += (nl - i);
			a[k] = y[j++];
		}
		++k;
	}
	while (i < nl) {
		a[k] = x[i];
		++k, ++i;
	}
	while (j < nr) {
		a[k] = y[j];
		++k, ++j;
	}
	return q;
}

auto inversion_count(vector<int> &a, int l, int r) -> int {
	if (l >= r)
		return 0;
	int m = (r + l) / 2, i = 0;
	i += inversion_count(a, l, m);
	i += inversion_count(a, m + 1, r);
	i += inversion(a, l, m, r);
	return i;
}

int main() {
	// vector<int> a { 1, 16, 4, 25, 9, 55, 2, 0, 7 };
	// vector<int> a { 1, 16, 4, 25, 9, 55, 2, 0, 7 };
	// vector<int> a { 1, 2, 3 };
	// vector<int> a { 3, 2, 1, 0 };
	// vector<int> a { 2, 3, 1 };
	vector<int> a { 7, 3, 1, 4 };
	int i = inversion_count(a, 0, a.size() - 1);

	for (auto e : a)
		cout << e << ' ';
	cout << endl << i << endl;
}
