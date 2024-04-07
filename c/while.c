#include <stdio.h>

/* print Fahrenheit-Celsius table
 for fahr = 0, 20, ..., 300 */

int main() {
  float fahr, celsius;
  float lower, upper, step;

  // in C, all variables must be declared before they are used, usually at the
  // beginning of the function before any executable statement.
  lower = 0;   /* lower limit of temperature table */
  upper = 300; /* upper limit */
  step = 20;   /* step size */

  fahr = lower;
  while (fahr <= upper) {
    celsius = (5.0 / 9.0) * (fahr - 32.0);
    printf("%10.0f\t%7.1f\n", fahr, celsius); /* \t tab between them*/
    fahr = fahr + step;
  }
}
