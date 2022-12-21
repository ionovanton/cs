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
	const int nl = m - l + 1;
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
			a[k] = y[j++];
			q += (nl - i);
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
	vector<int> a { 2, 3, 1 };
	int i = inversion_count(a, 0, a.size() - 1);

	for (auto e : a)
		cout << e << ' ';
	cout << endl << i << endl;
}



/*

// Merge sort in C++

// Merge two subarrays L and M into arr
void merge(vector<int> &arr, int p, int q, int r) {
  
  // Create L ← A[p..q] and M ← A[q+1..r]
  int n1 = q - p + 1;
  int n2 = r - q;

  int L[n1], M[n2];

  for (int i = 0; i < n1; i++)
	L[i] = arr[p + i];
  for (int j = 0; j < n2; j++)
	M[j] = arr[q + 1 + j];

  // Maintain current index of sub-arrays and main array
  int i, j, k;
  i = 0;
  j = 0;
  k = p;

  // Until we reach either end of either L or M, pick larger among
  // elements L and M and place them in the correct position at A[p..r]
  while (i < n1 && j < n2) {
	if (L[i] <= M[j]) {
	  arr[k] = L[i];
	  i++;
	} else {
	  arr[k] = M[j];
	  j++;
	}
	k++;
  }

  // When we run out of elements in either L or M,
  // pick up the remaining elements and put in A[p..r]
  while (i < n1) {
	arr[k] = L[i];
	i++;
	k++;
  }

  while (j < n2) {
	arr[k] = M[j];
	j++;
	k++;
  }
}

// Divide the array into two subarrays, sort them and merge them
void mergeSort(vector<int> &arr, int l, int r) {
  if (l < r) {
	// m is the point where the array is divided into two subarrays
	int m = l + (r - l) / 2;

	mergeSort(arr, l, m);
	mergeSort(arr, m + 1, r);

	// Merge the sorted subarrays
	merge(arr, l, m, r);
  }
}

// Print the array
void printArray(vector<int> arr, int size) {
  for (int i = 0; i < size; i++)
	cout << arr[i] << " ";
  cout << endl;
}

// Driver program
int main() {
  vector<int> arr { 1, 16, 4, 25, 9, 55, 2, 0 };
  int size = arr.size();

  mergeSort(arr, 0, size - 1);

  cout << "Sorted array: \n";
  printArray(arr, size);
  return 0;
}

*/