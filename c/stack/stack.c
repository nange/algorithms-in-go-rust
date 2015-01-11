#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include <stdbool.h>
#include "stack.h"

#define DEFAULT_STACK_SIZE 4

/**
 * create a new empty stack
 */
Stack *stackCreate() {
  Stack *stack = malloc(sizeof(Stack));
  assert(stack);
  stack->top = 0;
  stack->capacity = DEFAULT_STACK_SIZE;
  stack->dataStore = malloc(sizeof(void *)*DEFAULT_STACK_SIZE);

  stack->size = stackSize;
  stack->isEmpty = stackIsEmpty;
  stack->push = stackPush;
  stack->pop = stackPop;
  stack->peek = stackPeek;
  stack->clear = stackClear;
  stack->destory = stackDestory;
}

void stackShouldRealloc(Stack *s) {
  if (s->top == s->capacity) {
    s->capacity = s->capacity*2;
    s->dataStore = realloc(s->dataStore, sizeof(void *)*s->capacity);
  }
}

int stackSize(Stack *s) {
  return s->top;
}

bool stackIsEmpty(Stack *s) {
  return s->top == 0;
}

void stackPush(Stack *s, void *element) {
  stackShouldRealloc(s);
  s->dataStore[s->top++] = element;
}

void *stackPop(Stack *s) {
  if (s->isEmpty(s)) {
    return NULL;
  }

  return s->dataStore[--s->top];
}

void *stackPeek(Stack *s) {
  if (s->isEmpty(s)) {
    return NULL;
  }

  return s->dataStore[s->top-1];
}

void stackClear(Stack *s) {
    s->top = 0;
}

void stackDestory(Stack *s) {
  free(s->dataStore);
  free(s);
}
