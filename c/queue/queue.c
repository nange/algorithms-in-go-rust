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
	q->capacity = DEFAULT_QUEUE_SIZE;
	q->dataStore = malloc(sizeof(void *) * DEFAULT_QUEUE_SIZE);
	
	q->enqueue = queueEnqueue;
	q->dequeue = queueDequeue;
	q->size = queueSize;
	q->isEmpty = queueIsEmpty;
	q->font = queueFont;
	q->end = queueEnd;
	q->destory = queueDestory;
}

void queueShouldRealloc(Queue *q) {
	if (q->theSize == q->capacity) {
		q->capacity = q->capacity * 2;
		q->dataStore = realloc(q->dataStore, sizeof(void *) * q->capacity);
	}
}

void queueEnqueue(Queue *q, void *element) {
	queueShouldRealloc(q);
	q->dataStore[q->theSize++] = element;
}

void* queueDequeue(Queue *q) {
	if (q->theSize == 0) return NULL;
	void *d = q->dataStore[0];
	if (q->theSize > 1) {
		int i;
		for (i = 0; i < q->theSize-1; i++) {
			q->dataStore[i] = q->dataStore[i+1];
		}
	}
	q->theSize--;
	return d;
}

int queueSize(Queue *q) {
	return q->theSize;
}

bool queueIsEmpty(Queue *q) {
	return q->theSize == 0;
}

void *queueFont(Queue *q) {
	if (q->theSize == 0) return NULL;
	return q->dataStore[0];
}

void *queueEnd(Queue *q) {
	if (q->theSize == 0) return NULL;
	return q->dataStore[q->theSize-1];
}

void queueClear(Queue *q) {
	q->theSize = 0;
}

void queueDestory(Queue *q) {
	free(q->dataStore);
	free(q);
}
