#include <assert.h>
#include <stdio.h>
#include "queue.c"

int main() {
	printf("\nRunning queue.c tests.\n\n");

	Queue *q = queueCreate();
	assert(q->isEmpty(q));
  	assert(q->size(q) == 0);

	int n1 = 1, n2 = 2, n3 = 3;
	int *p1 = &n1, *p2 = &n2, *p3 = &n3;
	q->enqueue(q, p1);
	q->enqueue(q, p2);
	q->enqueue(q, p3);
	assert(q->size(q) == 3);
	
	int d = *(int*)(q->dequeue(q));
	assert(q->size(q) == 2);
	assert(d == 1);
	
	int d2 = *(int*)q->font(q);
	int d3 = *(int*)q->end(q);
	assert(d2 == 2);
	assert(d3 == 3);
	
	q->destory(q);

	printf("All queue.c tests completed.\n\n");
  	return 0;
}
