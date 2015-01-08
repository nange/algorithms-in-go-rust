package stack;

public class Stack<T> {
	private static final int DEFAULT_CAPACITY = 10;
	
	private T[] dataStore;
	private int top;
	
	public Stack() {
		this.top = 0;
		this.ensureCapacity(DEFAULT_CAPACITY);
	}
	
	public int size() {
		return this.top;
	}
	
	public boolean isEmpty() {
		return this.top == 0;
	}
	
	public void push(T element) {
		if (this.top == this.dataStore.length) {
			this.ensureCapacity(2*this.size());
		}
		this.dataStore[this.top++] = element;
	}
	
	public T pop() {
		if (this.isEmpty())  return null;
		
		return this.dataStore[--this.top];
	}
	
	public T peek() {
		if (this.isEmpty())  return null;
		
		return this.dataStore[this.top-1];
	}
	
	public void clear() {
		this.top = 0;
	}

	@SuppressWarnings("unchecked")
	private void ensureCapacity(int newCapacity) {
		if (newCapacity < this.size()) return;
		
		T[] old = this.dataStore;
		this.dataStore = (T[])new Object[newCapacity];
		for (int i = 0; i < this.size(); i++) {
			this.dataStore[i] = old[i];
		}
	}
}
