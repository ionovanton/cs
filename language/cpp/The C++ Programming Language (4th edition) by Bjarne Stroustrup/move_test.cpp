#include "header.hpp"

struct X {
	X(int value) : x(value) { cout << pfunc << endl; }
	X() : X(42) { cout << pfunc << endl; }

	X(const X &other) : X(other.x) { cout << pfunc << endl; }
	X &operator=(const X &other) {
		cout << pfunc << endl;
		x = other.x;
		return *this;
	}

	X(X &&other) : X(other.x) { other.x = 0; cout << pfunc << endl; }
	X &operator=(X &&other) {
		cout << pfunc << endl;
		x = other.x;
		other.x = 0;
		return *this;
	}

	X operator+(const X &a) {
		cout << pfunc << endl;
		return X(x + a.x);
	};

	int x;
};

ostream &operator<<(ostream &os, X &rhs) {
	os << rhs.x;
	return os;
}

ostream &operator<<(ostream &os, X &&rhs) {
	os << rhs.x;
	return os;
}

X foo() {
	return X(21);
}

int main() {
	X a = foo();
	X b = foo();

	// b = move(bar());
	// b = bar();
}


