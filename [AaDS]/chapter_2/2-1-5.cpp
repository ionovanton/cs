#include <iostream>
#include <vector>
#include <cmath>
#include <algorithm>

using namespace std;

auto add_binary_integers(vector<int> &a, vector<int> &b) -> vector<int> {
	const int n = a.size();
	vector<int> c(n + 1);

	auto carry = 0;
	for (int i = 0; i < n; ++i) {
		carry = a[i] + b[i] + carry;
		c[i] = carry % 2;
		carry /= 2;		
	}
	c[n] = carry;
	return c;
}

auto bin_to_int(const vector<int> &a) -> void {
	int r = 0; const int n = a.size();
	for (int i = 0; i < n; ++i) {
		r += a[i] * pow(2, i);
	}
	auto t(a);
	reverse(t.begin(), t.end());
	for (auto e : t)
		cout << e;
	cout << ": " << r << endl;
}

int main() {
	vector<int> a { 1, 0, 0, 1, 1, 1, 1, 1, 1 };
	vector<int> b { 1, 1, 0, 1, 0, 0, 0, 1, 0 };
	auto c = add_binary_integers(a, b);

	bin_to_int(a); bin_to_int(b); bin_to_int(c);


		
}
