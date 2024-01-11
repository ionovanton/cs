int main() {
	int a = 1;
	int b = 16;
	{
		b++;
		int a = 89;
		b++;
	}
	{
		b++;
		a = 42;
		b++;
	}
}

