#ifndef __QUEUE_H__
#define __QUEUE_H__

typedef struct Queue {
	int theSize;
	int capacity;
	void **dataStore;
	
	void (*enqueue)();
	void *(*dequeue)();
	int (*size)();
	bool (*isEmpty)();
	void *(font)();
	void *(end)();
	void (clear)();
} Queue;

/* Prototypes */
Queue *queueCreate();
void queueEnqueue(Queue *, void *);
void *queueDequeue(Queue *);
int queueSize(Queue *);
bool queueIsEmpty(Queue *);
void *queueFont(Queue *);
void *queueEnd(Queue *);
void queueClear(Queue *);

#endif __QUEUE_H__
