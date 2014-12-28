#ifndef __ARRAYLIST_H__
#define __ARRAYLIST_H__

typedef struct ArrayList {
	unsigned int theSize;
	void *dataStore;
	
	int (*size)();
	//void *(*get)(ArrayList *, unsigned int);
	//void (*set)(ArrayList *, unsigned int, void *);
	//void (*append)(ArrayList *, void *);
	//void (*insertAfter)(ArrayList *, unsigned int, void *);
	//void (*remove)(ArrayList *, unsigned int);
	//void (*clear)(ArrayList *);
} ArrayList;

/* Prototypes */
ArrayList *listCreate();
int listSize(ArrayList *list);
void *listGet(ArrayList *list, unsigned int index);
void listSet(ArrayList *list, unsigned int index, void *value);
void listAppend(ArrayList *list, void *value);
void listInsertAfter(ArrayList *list, unsigned int index, void *value);
void listRemove(ArrayList *list, unsigned int index);
void listClear(ArrayList *list);

#endif	/* __ARRAYLIST_H__ */
