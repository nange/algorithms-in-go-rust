package queue;

public class Queue<T> {
	private static final int DEFAULT_CAPACITY = 10;
	
	private T[] dataStore;
	private int theSize;
	
	public Queue() {
		this.theSize = 0;
		this.ensureCapacity(DEFAULT_CAPACITY);
	}
	
	public void enqueue(T element) {
		if (this.theSize == this.dataStore.length) {
			this.ensureCapacity(this.theSize *2);
		}
		this.dataStore[this.theSize++] = element;
	}
	
	public T dequeue() {
		if (this.size() == 0) return null;
		
		T d = this.dataStore[0];
		if (this.size() > 1) {
			for (int i = 0; i < this.theSize-1; i++) {
				this.dataStore[i] = this.dataStore[i+1];
			}
		}
		this.theSize--;
		
		return d;
	}
	
	public int size() {
		return this.theSize;
	}
	
	public boolean isEmpty() {
		return this.theSize == 0;
	}
	
	public T font() {
		if (this.size() == 0) return null;
		return this.dataStore[0];
	}
	
	public T end() {
		if (this.size() == 0) return null;
		return this.dataStore[this.theSize-1];
	}
	
	public void clear() {
		this.theSize = 0;
	}
	
	@SuppressWarnings("unchecked")
	private void ensureCapacity(int newCapa) {
		if (this.theSize > newCapa) return;
		
		T[] old = this.dataStore;
		this.dataStore = (T[])new Object[newCapa];
		for (int i = 0; i < this.theSize; i++) {
			this.dataStore[i] = old[i];
		}
	}

}
