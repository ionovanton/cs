#include <iostream>
#include <vector>
#include <cmath>
#include <climits>
#include <algorithm>

using namespace std;

auto find_sum(vector<int> &a, const int n, const int x) -> bool {
	sort(a.begin(), a.end());
	// int l = 0, r = n - 1;
	// while (l < r - 1 && (a[l] + a[r]) != x) {
	// 	if (x - a[r] > a[l])
	// 		++l;
	// 	else
	// 		--r;
	// }
	int sum = 0;
	for (int l = 0, r = n - 1; l < r && sum != x; sum = a[l] + a[r]) {
		if (x > sum)
			++l;
		else
			--r;
	}
	return sum == x;
}

int main(int argc, char **argv) {
	vector<int> a { 1, 2, 4};
	auto answer = find_sum(a, a.size(), 3);
	cout << answer << endl;
}
