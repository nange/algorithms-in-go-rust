#ifndef __STACK_H__
#define __STACK_H__

typedef struct Stack {
  int top;
  int capacity;
  void **dataStore;

  int (*size)();
  bool (*isEmpty)();
  void (*push)();
  void *(*pop)();
  void *(*peek)();
  void (*clear)();
  void (*destory)();
} Stack;

/* Prototypes */
Stack *stackCreate();
int stackSize(Stack *s);
bool stackIsEmpty(Stack *s);
void stackPush(Stack *s, void *);
void *stackPop(Stack *s);
void *stackPeek(Stack *s);
void stackClear(Stack *s);
void stackDestory(Stack *s);

#endif /* __STACK_H__ */
