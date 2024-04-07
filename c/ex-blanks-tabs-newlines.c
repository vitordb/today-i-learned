#include <stdio.h>

// write a program to count blank, tabs, and newlines.

int main(void) {

  int blanks = 0;
  int tabs = 0;
  int newlines = 0;

  char c;

  //  for (fahr = LOWER; fahr <= UPPER; fahr = fahr + STEP)
  //    printf("%3d %6.1f\n", fahr, (5.0 / 9.0 * (fahr - 32)));

  while ((c = getchar()) != EOF) {
    if (c == ' ') {
      ++blanks;
    } else if (c == '\t') {
      ++tabs;
    } else if (c == '\n') {
      ++newlines;
    }
  }

  printf("blanks: %d\ntabs: %d\nnewlines: %d\n", blanks, tabs, newlines);

  return 0;
}
