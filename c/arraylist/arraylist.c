#include <stdlib.h>
#include <string.h>
#include <assert.h>
#include "arraylist.h"

#define ARRAYLIST_DEAUFALT_SIZE 4

/**
 * create a new empty arraylist
 */
ArrayList *listCreate() {
	ArrayList *list = malloc(sizeof(ArrayList));
	
	list->dataStore = malloc(sizeof(void *) * ARRAYLIST_DEAUFALT_SIZE);
	assert(list->dataStore);
	list->theSize = 0;
	list->capacity = ARRAYLIST_DEAUFALT_SIZE;
	
	list->size = listSize;
	list->get = listGet;
	list->set = listSet;
	list->append = listAppend;
	
	return list;
}

void listIsRealloc(ArrayList *list) {
	if (list->theSize == list->capacity) {
		unsigned int newCapa = list->capacity * 2;
		list->capacity = newCapa;
		list->dataStore = realloc(list->dataStore, sizeof(void *) * newCapa);
		assert(list->dataStore);
	}
}

int listSize(ArrayList *list) {
	return list->theSize;
}

void *listGet(ArrayList *list, unsigned int index) {
	assert(index < list->theSize);
	return list->dataStore[index];
}

void listSet(ArrayList *list, unsigned int index, void *value) {
	assert(index < list->theSize);
	list->dataStore[index] = value;
}

void listAppend(ArrayList *list, void *value) {
	listIsRealloc(list);
	list->dataStore[list->theSize++] = value;
}






