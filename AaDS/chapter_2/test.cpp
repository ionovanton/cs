
#include <iostream>
#include <vector>
using namespace std;
 
// returns value of poly[0]x(n-1) + poly[1]x(n-2) + .. + poly[n-1]
int horner(vector<int> v, int n, int x)
{
	int result = v[n - 1]; // Initialize result
 
	// Evaluate value of polynomial using Horner's method
	for (int i = n - 2; i >= 0; i--)
		result = result * x + v[i];
 
	return result;
}
 
// Driver program to test above function.
int main()
{
	// Let us evaluate value of 2x3 - 6x2 + 2x - 1 for x = 3
	vector<int> v {1,2,3,4};
	int x = 2;
	cout << "Value of polynomial is " << horner(v, v.size(), x) << endl;
	return 0;
}