#include <iostream>
#include <string>
#include <vector>
#include <iterator>
#include <cstdio>
#include <cstring>
#include <cassert>
#include <fstream>
#include <list>
#include <algorithm>
#include <numeric>
#include <functional>
#include <map>
#include <sstream>
#include <ranges>
#include <memory>
#include <set>
#include <stdlib.h>
#include <immintrin.h>
#include <type_traits>

#if __cplusplus > 199711L
#include <unordered_map>
#include <random>
#endif

#define pfunc __PRETTY_FUNCTION__
#define flog std::clog << "[ " << this << " ] " << __PRETTY_FUNCTION__ << std::endl

using namespace std;

void showFunc(const char *str) { cout << str << endl; }

void caret(std::string msg)
{
	const char c = '~';
	std::cout << "\n" << std::string(60, c) << "\n"
	<< msg << "\n" << std::string(60, c) << "\n";
}

