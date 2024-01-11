#include <iostream>

using namespace std;

int main()
/* B1 */
{
	int a = 1;
	int b = 1;
	/* B2 */
	{
		int b = 2;
		/* B3 */
		{
			int a = 3;
			cout << a << b << endl;
		}
		/* B4 */
		{
			int b = 4;
			cout << a << b << endl;
		}
		cout << a << b << endl;
	}
	cout << a << b << endl;
}

/*

Expecting:
32
14
12
11

Answer:
32
14
12
11

*/

