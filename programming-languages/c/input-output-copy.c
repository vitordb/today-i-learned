#include <stdio.h>

int main() {
  double nc;

  for (nc = 0; getchar() != EOF; ++nc)
    ; // null statement because for in C needs a body
  printf("%.0f\n", nc);
}
