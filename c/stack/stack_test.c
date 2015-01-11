#include <assert.h>
#include <stdio.h>
#include "stack.c"

int main() {
  printf("\nRunning stack.c tests.\n\n");

  Stack *s = stackCreate();
  assert(s->isEmpty(s));
  assert(s->size(s) == 0);

  char *c1 = "12";
  char *c2 = "23";
  char *c3 = "34";
  s->push(s, c1);
  s->push(s, c2);
  s->push(s, c3);
  assert(s->size(s) == 3);

  char *c4 = "34";
  assert(strcmp(c4, (char *)s->pop(s)) == 0);
  char *c5 = "23";
  assert(strcmp(c5, (char *)s->peek(s)) == 0);

  s->clear(s);
  assert(s->isEmpty(s));

  s->destory(s);

  printf("All stack.c tests completed.\n\n");
  return 0;
}
