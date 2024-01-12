
#include "header.hpp"

template<typename T>
struct A {
	using value_type = T;
};

template<typename T>
struct B : public A<T> {
	using test = typename A<T>::value_type;

	test x{42.2};
};


int main() {
	B<float> x;

}
