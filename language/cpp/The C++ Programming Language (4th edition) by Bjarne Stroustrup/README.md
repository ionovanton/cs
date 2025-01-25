## Progress
### 1 / 1

# The C++ Programming Language (4th edition) by Bjarne Stroustrup

## Classes 16

### 16.2.6 `explicit` Constructors

It is important to use `explicit` in single constructor argument.

```c++
#include "header.hpp"

class A {
private:
	int a;
	int b;
public:
	explicit A() : a(42), b(21) {};
	explicit A(int value) : a(value), b(value * value) {};
	A(int value, int another_value) : a(value), b(another_value) {};

	void inspect() { cout << a << endl; }
};

void foo(A a) {
	a.inspect();
}

int main()
{
	A x;
	A y {21};

	foo( {} );			// error: converting to ‘A’ from initializer list would use explicit constructor ‘A::A()’
	foo( {42} );		// error: converting to ‘A’ from initializer list would use explicit constructor ‘A::A(int)’
	foo( {36, 95} );	// ok

	x = 43;				// error: initialization does not do implicit conversion
	y = {43};			// error: initialization does not do implicit conversion
	y = {42, 52};		// ok
}
```

### 16.2.9.3 `mutable`

`mutable` attributes could be changed inside `const` objects.

```c++
#include "header.hpp"

struct X {
	X() : a(1), b(2), c(3) {};
	int a;
	const int b;
	mutable int c;
	void inspect() const { printf("%p\na: %d	b: %d	c: %d\n", this, a, b, c); }
};

int main() {
	X x;
	const X y;

	x.inspect();
	y.inspect();

	y.a++; // error
	y.b++; // error
	y.c++; // ok (because mutable)

	x.inspect();
	y.inspect();
}
```

Output:
```
0x7ffe554d3bc0
a: 1    b: 2    c: 3
0x7ffe554d3bcc
a: 1    b: 2    c: 3
0x7ffe554d3bc0
a: 1    b: 2    c: 3
0x7ffe554d3bcc
a: 1    b: 2    c: 4
```

### 16.4 Advice

5. By default declare single-argument constructors explicit; §16.2.6.

## Construction, Cleanup, Copy, and Move 17

### 17.2.5 `virtual` Destructors

A destructor can be declared to be virtual, and usually should be for a class with a virtual function.

```c++
struct A {
	virtual void f() = 0;
	~A() { cout << pfunc << '\n'; };
};

struct D : public A {
	D() { cout << pfunc << '\n'; };
	~D() { cout << pfunc << '\n'; };
	void f() override { cout << pfunc << '\n'; };
};

void foo(A *a) {
	a->f();
	delete a;
}

int main() {
	caret("stack D");
	{
		D x;
	}

	caret("heap D");
	{
		D *x = new D();
		delete x;
	}

	caret("heap D ... delete in function");
	{
		D *x = new D();
		foo(x);
	}
}
```

Output:
```
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
stack D
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
D::D()
D::~D()
A::~A()

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
heap D
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
D::D()
D::~D()
A::~A()

~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
heap D ... delete in function
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
D::D()
virtual void D::f()
A::~A() // <-- source for possible errors, since D::~D() should've been called first
```

### 17.3.2 Initialization Using Constructors

Note that the `{}`-initializer notation does not allow narrowing (§2.2.2). That is another reason to prefer the `{}` style over `()` or `=`.

### 17.3.4.1 `initializer_list` Constructor Disambiguation

The rules are:
• If either a default constructor or an initializer-list constructor could be invoked, prefer the default constructor.
• If both an initializer-list constructor and an ‘‘ordinary constructor’’ could be invoked, prefer the initializer-list constructor.

```c++
struct X {
	X(initializer_list<int> a) {};
	X() {};
	X(int) {};
};

int main() {
	X a {};		// default constructor
	X b {0};	// initializer_list constructor
}
```

Output:
```
X::X()
X::X(std::initializer_list<int>)
```

### 17.5.2 Move

Formal defenitions of *lvalue* and *rvalue*.

**lvalue and rvalue are properties of expressions, not objects!**

lvalue | rvalue
:- | :-
identifier | literal
binary `=, +=, -=, *=, %=, /=, &=, \|=, <<=, >>=` if not user-defined | binary `+, -, *, /, %, &, \|, <<, >>, <, >, <=, >=, ==` if not user-defined
prefix `++, --` | postfix `++, --`
unary `*` | unary `&, ~, !`
ternary `? :` if *both* operands are lvalue | ternary `? :` if *at least* one operand is rvalue
operator `,` if right operand is lvalue | operator `,` if right operand is rvalue
function call if return type is `T&` | function call if return type is `T` or `T&&`
cast-expression if type is `T&` (*lvalue reference*) | cast-expression if type is `T` (*no-reference*) or `T&&` (*rvalue reference*)

Reference collapsing rules
```c++
&   &   —>  &
&   &&  —>  &
&&  &   —>  &
&&  &&  —>  &&
```

Rvalue-references and their properties
```c++
int main() {
	int x = 5;
	int &a = x;
	int &&b = x; // error: an rvalue reference cannot be bound to an lvalue
	int &&c = std::move(x); // ok
	int &&d = 5; // ok

	/* rvalue-reference could be initialized only via rvalue */
	int &&e = c; // error: an rvalue reference cannot be bound to an lvalue
	int &f = c; // ok

	/* constness must be preserved  */
	const int &g = c; // ok
	const int &h = std::move(c); // ok
	int &&i = std::move(g); // error: qualifiers dropped in binding reference of type "int &&" to initializer of type "const int"
}
```

`std::move` possible implementation
```c++
template <typename T>
std::remove_reference_t<T> &&move(T &&x) {
	return static_cast<std::remove_reference_t<T>&&>(x);
}
```

universial reference can take **BOTH** lvalue and rvalue
```c++
template <typename T> // <— template is mandatory
... move(T &&x)... // <— type of parameter here is universial reference

...

template <typename T> // <— template is mandatory
void foo(T &&x) {
	cout << x << ": " << std::is_lvalue_reference<T>{} << endl;
}

int main() {
	int a = 42;
	int b = 21;
	foo(std::move(a));
	foo(b);
}
```

Output:
```
42: 0
21: 1
```

another instance of universial reference:
```c++
auto &&x = ...
```

`std::forward` possible implementation
```c++
T &&forward(std::remove_reference_t<T> &x) {
	return static_cast<T &&>(x);
}
```

`std::forward` preserves value category
```c++
void overloaded( int &arg ) { std::cout << "lvalue\n"; }
void overloaded( int &&arg ) { std::cout << "rvalue\n"; }

template <typename T>
T &&correct_forward(std::remove_reference_t<T> &x) {
	return static_cast<T &&>(x);
}

template <typename T>
T &&wrong_forward(T &&x) {
	return static_cast<T &&>(x);
}

template <typename T>
void f(T &&x) {
	cout << "correct_forward: ";
	overloaded(correct_forward<T>(x));
	cout << "wrong_forward: ";
	overloaded(wrong_forward(x));
	cout << "argument: ";
	overloaded(x);
}

int main() {
	int x = 42;

	cout << "pass by lvalue\n";
	f(x);
	cout << "pass by rvalue\n";
	f(21);
}
```

Output:
```
pass by lvalue
correct_forward: lvalue
wrong_forward: lvalue
argument: lvalue
pass by rvalue
correct_forward: rvalue
wrong_forward: lvalue
argument: lvalue
```

Hence the table for `correct_forward`

...|`std::forward`|`std::move`|`argument`
:-:|:-:|:-:|:-:
initial caller passes `rvalue`|`rvalue`|`rvalue`|`lvalue`
initial caller passes `lvalue`|`lvalue`|`rvalue`|`lvalue`

Why `correct_forward<T>(x)` works?

- After passing `<T>` from `main()` to `f()`, `correct_forward` instantiated like this:
```c++
/* source template */
template <typename T>
T &&correct_forward(std::remove_reference_t<T> &x)
{
	return static_cast<T &&>(x);
}

/* when passed by lvalue */
template<>
int & correct_forward<int &>(std::remove_reference_t<int &> & x)
{
	return static_cast<int &>(x);
}

/* when passed by rvalue */
template<>
int && correct_forward<int>(std::remove_reference_t<int> & x)
{
	return static_cast<int &&>(x);
}
```

- Then `x` is passed to `correct_forward` as lvalue (in both cases). Since we have `std::remove_reference_t` at our disposal, we will always be ending up with `T &x` in parameters. This will let us pass lvalue to `correct_forward`. It will look something like this:
```c++
	/* when passed by lvalue */
	template<>
	int & correct_forward<int &>(int &x)
	{
		return static_cast<int &>(x);
	}

	/* when passed by rvalue */
	template<>
	int && correct_forward<int>(int &x)
	{
		return static_cast<int &&>(x);
	}
```

- This way we may save value category.

Why `wrong_forward(x)` doesn't work?

- At the time of calling `wrong_forward()` it instantiated like this:
```c++
/* both when passed by lvalue or rvalue */
template<>
int & wrong_forward<int &>(int & x)
{
	return static_cast<int &>(x);
}
```

- It instantiating lvalue variant bacause **in both cases `x` is passed as lvalue**

- Since it is always instantiating lvalue variant we will always return an lvalue. Meaning it is wrong.

Why `wrong_forward<T>(x)` doesn't work?

- After passing `<T>` from `main()` to `f()`, `wrong_forward` instantiated like this:
```c++
/* when passed by lvalue */
template<>
int & wrong_forward<int &>(int & x)
{
	return static_cast<int &>(x);
}

/* when passed by rvalue */
template<>
int && wrong_forward<int &&>(int && x)
{
	return static_cast<int &&>(x);
}
```

- Since **lvalue always is passed from `f()` to `wrong_forward<T>()`** we will get compilation error on pass by rvalue: lvalue will be passed to synthesized function `int && wrong_forward<int &&>(int && x)` (because it previously was passed by rvalue).

Rule of thumb:
- `std::forward` — inside a templated function with an argument declared as a forwarding reference, where the argument is now `lvalue`, used to retrieve the original value category, that it was called with, and pass it on further down the call chain (perfect forwarding).
- `std::move` — ensure the argument passed to a function is an `rvalue` so that you can move from it (choose move semantics). By function could be an actual function or a constructor or an operator (e.g. assignment operator).


### 17.6.3.2 Maintaining Invariants

When designing a class we must provide it with invariant.

1) Initialize: establish an invariant
2) Maintain: on copy or move operations
3) Terminate: on destruction.

### 17.6.4 deleted Functions

`=delete` is used for disallowing function use.

### 17.7 Advice

4. If a class has a virtual function, it needs a virtual destructor; §17.2.5.
6. Prefer {} initialization over = and () initialization; §17.3.2.
10. If a class has a reference member, it probably needs copy operations (copy constructor and copy assignment); §17.4.1.1.

## Operator Overloading 18

### 18.2.3 Operators and User-Defined Types

```c++
struct X {
	int x;
	X() : x(42) {};
	X(int value) : x(value) {};
	X operator+(const X &other) {
		return X(x + other.x);
	};
	X operator+(X &&other) {
		return X(x + other.x);
	};
	X operator+=(const X &other) {
		x += other.x;
		return *this;
	}
};

ostream &operator<<(ostream &os, X &x) {
	cout << x.x;
	return os;
}

ostream &operator<<(ostream &os, X &&x) {
	cout << x.x;
	return os;
}

X foo(X a) {
	++a.x;
	return a;
}

int main() {
	X a(13);

	cout << a + 1 << endl; 	// ok: same as below
	cout << a.operator+(1) << endl; // ok: lvalue requires lvalue
	cout << 1 + a << endl; 	// error: there's no function as 2.operator+(a)
	X b{foo(a) += 2}; 		// ok: operator requires temporary value (probably xvalue)
}
```

### 18.2.4 Passing Objects

Typically, an operator returns a result. Returning a pointer or a reference to a newly created object is usually a very bad idea: using a pointer gives notational problems, and referring to an object on the free store (whether by a pointer or by a reference) results in memory management problems. Instead, return objects by value. For large objects, such as a Matrix, define move operations to make such transfers of values efficient.

Rely on move-constructor and copy elision when returning by value from function.

### 18.2.5 Operators in Namespaces

Note that in operator lookup no preference is given to members over nonmembers. This differs from lookup of named functions (§14.2.4). The lack of hiding of operators ensures that built-in operators are never inaccessible and that users can supply new meanings for an operator without modifying existing class declarations.

### 18.4.3 Ambiguities

In some cases, a value of the desired type can be constructed by repeated use of constructors or conversion operators. This must be handled by explicit conversions; only one level of user-defined implicit conversion is legal.

```c++
struct A {
	A() : x(1) {};
	A(int value) : x(value) {};
	int x;
};

struct B {
	B() : x(2) {};
	B(A value) : x(value.x) {};
	int x;
};

struct C {
	C() : x(3) {};
	C(B value) : x(value.x) {};
	int x;
};

void foo(C c) {}

int main() {
	foo(42); // error: no suitable constructor exists to convert from "int" to "C"
}

```

### 18.5 Advice

3. For large operands, use const reference argument types; §18.2.4.
4. For large results, use a move constructor; §18.2.4.
9. Use member functions to express operators that require an lvalue as their left-hand operand; §18.3.3.1.

## Special Operators 19

### 19.2.6 User-defined Literals

```c++
struct A {
	A(int value) : x(value) {};
	A() : A(42) {};
	int x;
};

A operator"" _x(unsigned long long x) {
	return A(x);
}

int main() {
	auto a = 123_x;
	cout << a.x << endl;
}
```

Output:
```
123
```

### 19.4.1 Finding Friends

A friend must be previously declared in an enclosing scope or defined in the non-class scope immediately enclosing the class that is declaring it to be a friend.

```c++
class C1 {};	// will become a friend
void f1() {}	// will become a friend

namespace N {

class C2 {};	// will become a friend
void f2() {}	// will become a friend

class C {
	int x;
public:
	friend class C1;	// ok: defined previously
	friend void f1();

	friend class C3;	// ok: defined in inclosing namespace
	friend void f3();

	friend class C4;	// First declared in N and assumed to be in N
	friend void f4();

};

class C3 {};	// will become a friend
void f3() {}	// will become a friend

} // namespace N

class C4 {};	// NOT a friend
void f4() {}	// NOT a friend
```

A friend function can be found through its arguments (§14.2.4) even if it was not declared in the immediately enclosing scope.

```c++
class X {
	int value = 42;
	friend void g(X&); // can be found through its arguments
public:
	int getValue() { return value; };
};

void g(X &x) {
	x.value = 21;
}

void h(X &x) {
	g(x);
}

int main() {
	X a;
	h(a);
	cout << a.getValue() << endl;
}
```

Output:
```
21
```

### 19.5 Advice

10. Use a friend function if you need a nonmember function to have access to the representation of a class (e.g., to improve notation or to access the representation of two classes); §19.4.

11. Prefer member functions to friend functions for granting access to the implementation of a class; §19.4.2.


## Derived Classes 20

### 20.3.2 Virtual Functions

Virtual calls.
```c++
struct A {
	A() { flog; }
	int x = 1;
	virtual void info(); 
};

struct B : public A {
	B() { flog; }
	int x = 2;
	void info() override;
};

struct C : public B {
	C() { flog; }
	int x = 3;
	void info() override;
};

void A::info() {
	flog;
	cout << x << endl;
}

void B::info() {
	flog;
	cout << x << endl;
}

void C::info() {
	flog;
	cout << x << endl;
}

void foo(A &e) {
	e.info();
}

int main() {
	A a;
	B b;
	C c;

	foo(a);
	foo(b);
	foo(c);
}
```

Output:
```
[ 0x7fff3034e9c0 ] A::A()
[ 0x7fff3034e9d0 ] A::A()
[ 0x7fff3034e9d0 ] B::B()
[ 0x7fff3034e9e0 ] A::A()
[ 0x7fff3034e9e0 ] B::B()
[ 0x7fff3034e9e0 ] C::C()
[ 0x7fff3034e9c0 ] virtual void A::info()
1
[ 0x7fff3034e9d0 ] virtual void B::info()
2
[ 0x7fff3034e9e0 ] virtual void C::info()
3
```

Non-virtual calls.
```c++
struct A {
	A() { flog; }
	int x = 1;
	void info(); 
};

struct B : public A {
	B() { flog; }
	int x = 2;
	void info();
};

struct C : public B {
	C() { flog; }
	int x = 3;
	void info();
};

void A::info() {
	flog;
	cout << x << endl;
}

void B::info() {
	flog;
	cout << x << endl;
}

void C::info() {
	flog;
	cout << x << endl;
}

void foo(A &e) {
	e.info();
}

int main() {
	A a;
	B b;
	C c;

	foo(a);
	foo(b);
	foo(c);
}
```

Output:
```
[ 0x7fffa267f7a0 ] A::A()
[ 0x7fffa267f7a4 ] A::A()
[ 0x7fffa267f7a4 ] B::B()
[ 0x7fffa267f7ac ] A::A()
[ 0x7fffa267f7ac ] B::B()
[ 0x7fffa267f7ac ] C::C()
[ 0x7fffa267f7a0 ] void A::info()
1
[ 0x7fffa267f7a4 ] void A::info()
1
[ 0x7fffa267f7ac ] void A::info()
1
```

Virtual calls, commenting out B function
```c++
struct A {
	A() { flog; }
	int x = 1;
	virtual void info(); 
};

struct B : public A {
	B() { flog; }
	int x = 2;
	// void info() override;
};

struct C : public B {
	C() { flog; }
	int x = 3;
	void info() override;
};

void A::info() {
	flog;
	cout << x << endl;
}

// void B::info() {
// 	flog;
// 	cout << x << endl;
// }

void C::info() {
	flog;
	cout << x << endl;
}

void foo(A &e) {
	e.info();
}

int main() {
	A a;
	B b;
	C c;

	foo(a);
	foo(b);
	foo(c);
}
```

Output:
```
[ 0x7ffd4bc3ced0 ] A::A()
[ 0x7ffd4bc3cee0 ] A::A()
[ 0x7ffd4bc3cee0 ] B::B()
[ 0x7ffd4bc3cef0 ] A::A()
[ 0x7ffd4bc3cef0 ] B::B()
[ 0x7ffd4bc3cef0 ] C::C()
[ 0x7ffd4bc3ced0 ] virtual void A::info()
1
[ 0x7ffd4bc3cee0 ] virtual void A::info()
1
[ 0x7ffd4bc3cef0 ] virtual void C::info()
3
```

Virtual calls, commenting out B function
```c++
struct A {
	A() { flog; }
	int x = 1;
	virtual void info(); 
};

struct B : public A {
	B() { flog; }
	int x = 2;
	void info() override;
};

struct C : public B {
	C() { flog; }
	int x = 3;
	// void info() override;
};

void A::info() {
	flog;
	cout << x << endl;
}

void B::info() {
	flog;
	cout << x << endl;
}

// void C::info() {
// 	flog;
// 	cout << x << endl;
// }

void foo(A &e) {
	e.info();
}

int main() {
	A a;
	B b;
	C c;

	foo(a);
	foo(b);
	foo(c);
}
```

Output:
```
[ 0x7ffcda558600 ] A::A()
[ 0x7ffcda558610 ] A::A()
[ 0x7ffcda558610 ] B::B()
[ 0x7ffcda558620 ] A::A()
[ 0x7ffcda558620 ] B::B()
[ 0x7ffcda558620 ] C::C()
[ 0x7ffcda558600 ] virtual void A::info()
1
[ 0x7ffcda558610 ] virtual void B::info()
2
[ 0x7ffcda558620 ] virtual void B::info()
2
```

`override` is *contextual keyword* so it can be used as variable name. Other contextual keyword is `final`. 
```c++
int override = 42;

struct A {
	A() { flog; }
	int override = 1;
	virtual void info() { flog; cout << override + ::override << endl; };
};

struct B : public A {
	B() { flog; }
	int override = 2;
	void info() override { flog; cout << override + ::override << endl; };
};

void foo(A &e) {
	e.info();
}

int main() {
	A a;
	B b;

	foo(a);
	foo(b);
}
```

Output:
```
[ 0x7ffc4238e540 ] A::A()
[ 0x7ffc4238e550 ] A::A()
[ 0x7ffc4238e550 ] B::B()
[ 0x7ffc4238e540 ] virtual void A::info()
43
[ 0x7ffc4238e550 ] virtual void B::info()
44
```

### 20.3.5 using Base Members

Functions do not overload across scopes (§12.3.3). For example:
```c++
struct Base {
	void bar(int) { flog; };
};

struct Derived : public Base {
	void bar(double) { flog; };
};

void foo(Derived d) {
	d.bar(1);
	Base &b = d;
	b.bar(1);
}

int main() {
	Derived d;

	foo(d);
}
```

Output:
```
[ 0x7fff7a3b69bf ] void Derived::bar(double)
[ 0x7fff7a3b69bf ] void Base::bar(int)
```

As for namespaces, using-declarations can be used to add a function to a scope.
```c++
struct Base {
	void bar(int) { flog; };
};

struct Derived : public Base {
	using Base::bar;
	void bar(double) { flog; };
};

void foo(Derived d) {
	d.bar(1);
	Base &b = d;
	b.bar(1);
}

int main() {
	Derived d;

	foo(d);
}
```

Output:
```
[ 0x7ffc4106be4f ] void Base::bar(int)
[ 0x7ffc4106be4f ] void Base::bar(int)
```

### 20.3.5.1 Inheriting Constructors

You can inherit constructors with `using`.
```c++
struct A {
	int x;
	A(int value) : x{value} { flog; };
};

struct B : public A {
	using A::A;
};

int main() {
	A a{21};
	B b{42};

	printf("%d\n%d\n", a.x, b.x);
}
```

Output:
```
[ 0x7ffe1b391740 ] A::A(int)
[ 0x7ffe1b391744 ] A::A(int)
21
42
```

### 20.5 Access Control

Don't use multiple access specifiers without good reason because compiler can reorder data in your class:

```c++
class A {
public:
	A(int i) : x{++i}, y{++i}, z{++i} {};
	void inspect() { printf("%d %d %d\n", x,y,z); };
public:
	int x;
public:
	int y;
public:
	int z;
};
```

A derived class can access a base class’s protected members only for objects of its own type. This prevents subtle errors that would otherwise occur when one derived class corrupts data belonging to other derived classes.

```c++
class A {
protected:
	int x {42};
};

class B : public A {
public:
	void f(B *b) {
		b->x = 21; // ok
	};
};

class C : public A {
public:
	void g(B *b) {
		b->x = 21; // error: protected member "A::x" (declared at line 3) is not accessible through a "B" pointer or object
	};
};
```

### 20.5.1.1 Use of protected Members

Members declared protected are far more open to abuse than members declared private. In particular, declaring data members protected is usually a design error. Placing significant amounts of data in a common class for all derived classes to use leaves that data open to corruption. Worse, protected data, like public data, cannot easily be restructured because there is no good way of finding every use. Thus, protected data becomes a software maintenance problem.

However, none of these objections are significant for protected member functions; protected is a fine way of specifying operations for use in derived classes.

### 20.5.3 `using`-Declarations and Access Control

A using-declaration (§14.2.2, §20.3.5) cannot be used to gain access to additional information. It is simply a mechanism for making accessible information more convenient to use. On the other hand, once access is available, it can be granted to other users.

```c++
class A {
private:
	int x;	
protected:
	int y;
public:
	int z;
};

class B : public A {
public:
	// using A::x; // error: A::x is private
	using A::y; // ok: makes A::y publicly avaliable through B
};

class C : private B {
	using A::x; // error: A::x is private
	using A::y; // ok
	using A::z; // ok
	void f() { cout << y << endl; }; // ok
};

int main() {
	A a;
	B b;
	C c;

	a.y = 2; // error: A::y is protected
	b.y = 2; // ok
}
```

### 20.6.1 Pointers to Function Members

```c++
class StdInterface {
public:
	virtual void suspend() = 0;
	virtual ~StdInterface() = default;
};

using IMemPtr = void (StdInterface::*)();

class A : public StdInterface {
	void suspend() override {
		flog;
	};
};


void f(StdInterface *p) {
	IMemPtr s = &StdInterface::suspend;
	p->suspend();
	(p->*s)();
}

int main() {
	A a;
	f(&a);
}
```

Output:
```
[ 0x7ffe400b9010 ] virtual void A::suspend()
[ 0x7ffe400b9010 ] virtual void A::suspend()
```

### 20.6.2 Pointers to Data Members

```c++
struct X {
	const char *s;
	int i;

	void print(int x) { cout << s << ' ' << x << '\n'; };
	int f1(int);
	void f2();

	X(const char *value) : s{value} {};
};

using PtrMemberFuncInt = void (X::*)(int);
using PtrMember = const char* X::*;

void f(X &a, X &b) {
	X *ptr = &b;
	PtrMemberFuncInt ptrfi = &X::print;
	PtrMember ptrm = &X::s;

	a.print(1);
	(a.*ptrfi)(2);

	(a.*ptrm) = "111";
	ptr->*ptrm = "222";

	b.print(3);
	(ptr->*ptrfi)(4);

	PtrMemberFuncInt ptrfi = &X::f1; // error: return type mismatch
	PtrMemberFuncInt ptrfi = &X::f2; // error: argument type mismatch
	PtrMember ptrm = &X::i; // error: type mismatch
}

int main() {
	X a("abc"), b("xyz");
	f(a, b);
}
```

Output:
```
abc 1
abc 2
222 3
222 4
```

### 20.6.3 Base and Derived Members

A derived class has at least the members that it inherits from its base classes. Often it has more. This implies that we can safely assign a pointer to a member of a base class to a pointer to a member of a derived class, but not the other way around. This property is often called *contravariance*.

```c++
class StdInterface {
public:
	virtual void suspend() = 0;
	virtual ~StdInterface() = default;
};

using IMemPtr = void (StdInterface::*)();

class A : public StdInterface {
public:
	void suspend() override {
		flog;
	};
};

int main() {
	void (StdInterface::*ptr_a)() = &A::suspend; // error: a value of type "void (A::*)()" cannot be used to initialize an entity of type "void (StdInterface::*)()"
	void (A::*ptr_b)() = &StdInterface::suspend; // ok
}
```

### More on *covariance*, *contravariance*, *bivariance* and *invariance*

```c++
class Base {};
class Derived : public Base {};

int main() {
	using BaseFunc = std::function<void(Base*)>;
	using DerivedFunc = std::function<void(Derived*)>;

	/* covariance — more derived type component can be assigned to a less derived type component */
	Base *base_ptr = new Derived();		// ok
	Derived *derived_ptr = new Base();	// error
	// covariant return rule
	auto BaseProducer = []() -> Base* { return new Base; };
	auto DerivedProducer = []() -> Derived* { return new Derived; };
	
	Base *base_ptr = DerivedProducer();		// ok
	Derived *derived_ptr = BaseProducer();	// error: a value of type "Base *" cannot be used to initialize an entity of type "Derived *"

	/* contravariance — less derived type component can be assigned to a more derived type component */
	auto LambdaBase = [](Base *){};
	auto LambdaDerived = [](Derived *){};

	BaseFunc a = LambdaDerived;		// error: no suitable user-defined conversion from "lambda [](Derived *)->void" to "BaseFunc" exists
	DerivedFunc b = LambdaBase;		// ok: contravariant

	/* bivariance — both covariance and contravariance can be maintained */

	/* invariance — no substitutability can occur */
	vector<Base> base_vector;
	vector<Derived> derived_vector;

	base_vector = derived_vector; // error
	derived_vector = base_vector; // error
}
```

### 20.7 Advice

3. Use abstract classes to focus design on the provision of clean interfaces; §20.4.
7. Use abstract classes to keep implementation details out of interfaces; §20.4.
8. A class with a virtual function should have a virtual destructor; §20.4.
13. Don’t declare data members protected; §20.5.1.1.

## Class Hierarchies 21

The way most C++ implementations work implies that a change in the size of a base class requires a recompilation of all derived classes.

### 21.2.2 Interface Inheritance

The user-interface system should be an implementation detail that is hidden from users who don’t want to know about it.

No recompilation of code using the Ival_box family of classes should be required after a change of the user-interface system.

In a derived class, if a virtual member function of a base class subobject has more than one final overrider the program is ill-formed.

```c++
struct X {
	virtual void foo() = 0;
};

struct A : virtual public X {
	void foo() override { flog; };
};

struct B : virtual public X {
	void foo() override { flog; };
};

struct C : public A, public B {
	// void foo() override { flog; };
};

int main() {
	A a;
	B b;
	C c;

	a.foo();
	b.foo();
}
```

Output:
```
main.cpp:13:8: error: no unique final overrider for ‘virtual void X::foo()’ in ‘C’
   13 | struct C : public A, public B {
      |        ^
```

### CRTP

Example, static polymorphism:
```c++
template<typename CRTP>
class Amount {
public:
	int getValue() const {
		return static_cast<CRTP const&>(*this).getValue();
	}
};

class Constant : public Amount<Constant> {
private:
	int value;
public:
	explicit Constant(const int x) : value{x} {};
	int getValue() const { return value; };
};

class Variable : public Amount<Variable> {
private:
	int value;
public:
	explicit Variable(const int x) : value{x} {};
	void setValue(const int x) { value = x; }
	int getValue() const { return value; };
};

template<typename T>
void print(Amount<T> const &amount) {
	cout << amount.getValue() << endl;
}

int main() {
	Constant c{42};
	Variable v{21};

	print(c);
	print(v);
}
```

Output:
```
42
21
```

Generelizing CRTP:
```c++
template<typename T>
struct CRTP {
	T &underlying() { return static_cast<T&>(*this); }
	T const &underlying() const { return static_cast<T const&>(*this); }
};

template<typename T>
class Amount : public CRTP<T> {
public:
	int getValue() const {
		return this->underlying().getValue();
	}
};

class Constant : public Amount<Constant> {
private:
	int value;
public:
	explicit Constant(const int x) : value{x} {};
	int getValue() const { return value; };
};

class Variable : public Amount<Variable> {
private:
	int value;
public:
	explicit Variable(const int x) : value{x} {};
	void setValue(const int x) { value = x; }
	int getValue() const { return value; };
};

template<typename T>
void print(Amount<T> const &amount) {
	cout << amount.getValue() << endl;
}

int main() {
	Constant c{42};
	Variable v{21};

	print(c);
	print(v);
}
```

Output:
```
42
21
```

### Mixins in C++

Mixin classes are template classes that define a generic behaviour, and are designed to inherit from the type you wish to plug their functionality onto.

```c++
class Name {
public:
	Name(std::string firstName, std::string lastName)
		: firstName(std::move(firstName))
		, lastName(std::move(lastName)) {}

	void print() const {
		cout << lastName << ", " << firstName << '\n';
	}
private:
	std::string firstName;
	std::string lastName;
};

struct A {
	void print() const {
		cout << "A\n";
	}
};

template<typename Printable>
struct RepeatPrint : Printable {
	explicit RepeatPrint(Printable const& printable) 
		: Printable(printable) {}
	
	void repeat(unsigned int n) const {
		while (n-- > 0)
			this->print();
	}
};

template<typename Printable>
RepeatPrint<Printable> repeatPrint(Printable const& printable) {
	return RepeatPrint<Printable>(printable);
}

int main() {
	Name arya("Arya", "Stark");

	arya.print();
	repeatPrint(arya).print();
	repeatPrint(arya).repeat(4);
	repeatPrint(A()).repeat(3);
}
```

Output:
```
Stark, Arya
Stark, Arya
Stark, Arya
Stark, Arya
Stark, Arya
Stark, Arya
A
A
A
```

The type of `RepeatPrint` when executing on `arya` is:

```dbg
ptype *this
```

Output:
```
type = const struct RepeatPrint<Name> [with Printable = Name] : public Printable {
  public:
    RepeatPrint(const Printable &);
    void repeat(unsigned int) const;
}
```

The type of `RepeatPrint` when executing on `A()` is:

```dbg
ptype *this
```

Output:
```
type = const struct RepeatPrint<A> [with Printable = A] : public Printable {
  public:
    RepeatPrint(const Printable &);
    void repeat(unsigned int) const;
}
```

### Mixins vs CRTP

**The CRTP:**

- Impacts the definition of the existing class, because it has to inherit from the CRTP.
- Client code uses the original class directly and benefits from its augmented functionalities.


**The mixin class:**

- Leaves the original class unchanged.
- Client code doesn’t use the original class directly, it needs to wrap it into the mixin to use the augmented functionality.
- Inherits from a the original class even if it doesn’t have a virtual destructor. This is ok unless the mixin class is deleted polymorphically through a pointer to the original class.

### 21.4 Advice

2. Avoid data members in base classes intended as interfaces; §21.2.1.1.
9. Use multiple inheritance to separate implementation from interface; §21.3.
10. Use a virtual base to represent something common to some, but not all, classes in a hierarchy; §21.3.5.

## Run-Time Type Information 22



### 22.7 Advice

1. Use virtual functions to ensure that the same operation is performed independently of which interface is used for an object; §22.1.
7. Don’t call virtual functions during construction or destruction; §22.4.





```c++

```

Output:
```

```

