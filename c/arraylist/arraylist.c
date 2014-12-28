#include <stdlib.h>
#include <string.h>
#include <assert.h>

#include "arraylist.h"

const int DEFAULT_CAPACITY = 10;

/**
 * create a new empty arraylist
 */
ArrayList *listCreate() {
	ArrayList *list = malloc(sizeof(ArrayList));
	list->theSize = 0;
	list->dataStore = malloc(sizeof(void *) * DEFAULT_CAPACITY);
	
	list->size = listSize;
	
	return list;
}

int listSize(ArrayList *list) {
	return list->theSize;
}



