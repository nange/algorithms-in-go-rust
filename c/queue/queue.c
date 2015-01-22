#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include <stdbool.h>
#include "queue.h"

#define DEFAULT_QUEUE_SIZE 4

Queue *queueCreate() {
	Queue *q = malloc(sizeof(Queue));
	assert(q);
	q->theSize = 0;
	q->capacity = 0;
	q->dataStore
	
}


