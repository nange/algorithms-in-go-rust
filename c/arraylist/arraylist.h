#ifndef __ARRAYLIST_H__
#define __ARRAYLIST_H__

typedef struct ArrayList {
	int theSize;
	int capacity;
	void **dataStore;
	
	int (*size)(struct ArrayList *);
	void *(*get)(struct ArrayList *, unsigned int);
	void (*set)(struct ArrayList *, unsigned int, void *);
	void (*append)(struct ArrayList *, void *);
	void (*insert)(struct ArrayList *, unsigned int, void *);
	void (*remove)(struct ArrayList *, unsigned int);
	void (*clear)(struct ArrayList *);
	void (*destory)(struct ArrayList *);
} ArrayList;

/* Prototypes */
ArrayList *listCreate();
int listSize(ArrayList *list);
void *listGet(ArrayList *list, unsigned int index);
void listSet(ArrayList *list, unsigned int index, void *value);
void listAppend(ArrayList *list, void *value);
void listInsert(ArrayList *list, unsigned int index, void *value);
void listRemove(ArrayList *list, unsigned int index);
void listClear(ArrayList *list);
void listDestory(ArrayList *list);

#endif	/* __ARRAYLIST_H__ */
