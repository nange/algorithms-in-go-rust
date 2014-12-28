#include <stdio.h>

#include "arraylist.c"

int main() {
	ArrayList *list = listCreate();
	printf("list size: %d\n", list->size(list));
}
