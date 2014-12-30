#include <assert.h>
#include <stdio.h>
#include "arraylist.c"

int main() {
	ArrayList *list = listCreate();
	assert(list->size(list) == 0);
	
	int a = 1;
	list->append(list, &a);
	assert(list->size(list) == 1);
	
	int item1 = *(int *)list->get(list, 0);
	assert(item1 == 1);
	
	char c = 'c';
	list->append(list, &c);
	assert(list->size(list) == 2);
	
	char item2 = *(char *)list->get(list, 1);
	assert(item2 == 'c');
	
	int b = 2;
	list->set(list, 1, &b);
	assert(list->size(list) == 2);
	
	int item2_new = *(int *)list->get(list, 1);
	assert(item2_new == 2);
	
	printf("\n ArrayList test pass...\n\n");
	return 0;
}

