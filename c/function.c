
#include <stdio.h>

int power(int m, int n);

int main() {
  int i;

  for (i = 0; i < 10; ++i)
    printf("%d %d %d\n", i, power(2, i), power(-3, i));
  return 0;
}

int power(int base, int n) {
  int i, p;

  p = 1;
  for (i = 1; i <= n; ++i)
    p = p * base;
  return p;
}

// using call by value, the parameter n is used as a temporary variable, and is
// counted down. A copy of the value from the arguments is passed to the
// function.
int power2(int base, int n) {
  int p;
  for (p = 1; n > 0; --n)
    p = p * base;
  return p;
}
