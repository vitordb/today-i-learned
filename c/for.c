#include <stdio.h>

#define LOWER 0
#define UPPER 300
#define STEP 20

int main() {
  int fahr;

  // for is usually appropriate for loops in which the initialization and the
  // increment are single statements and logically related, sinc eit is more
  // compact than while and it keep the loop control statements together in one
  // place.
  for (fahr = LOWER; fahr <= UPPER; fahr = fahr + STEP)
    printf("%3d %6.1f\n", fahr, (5.0 / 9.0 * (fahr - 32)));
}
