#include <assert.h>
#include <stdio.h>
#include "arraylist.c"

int main() {
	printf("\nRunning arraylist.c tests.\n\n");
	
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
	
	list->destory(list);
	list = NULL;
	assert(!list);
	
	int a1 = 1;
	int a2 = 2;
	int a3 = 3;
	int a4 = 4;
	ArrayList *list2 = listCreate();
	list2->append(list2, &a1);
	list2->append(list2, &a2);
	list2->append(list2, &a3);
	list2->insert(list2, 2, &a4);
	assert(list2->size(list2) == 4);
	
	int r1 = *(int *)list2->get(list2, 0);
	int r2 = *(int *)list2->get(list2, 1);
	int r3 = *(int *)list2->get(list2, 2);
	int r4 = *(int *)list2->get(list2, 3);
	assert(r1 == 1);
	assert(r2 == 2);
	assert(r3 == 4);
	assert(r4 == 3);
	
	list2->remove(list2, 0);
	r1 = *(int *)list2->get(list2, 0);
	assert(r1 == 2);
	assert(list2->size(list2) == 3);
	
	list2->clear(list2);
	assert(list2->size(list2) == 0);
	
	list2->destory(list2);
	list2 = NULL;
	assert(!list2);
	
	printf("All arraylist.c tests completed.\n\n");
	return 0;
}
