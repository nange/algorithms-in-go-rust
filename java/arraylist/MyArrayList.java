package arraylist;

import java.lang.ArrayIndexOutOfBoundsException;

public class MyArrayList<AnyType> {
	private static final int DEFAULT_CAPACITY = 10;

	private int size;
	private AnyType[] items;

	public MyArrayList() {
		clear();
	}

	public void clear() {
		this.size = 0;
		this.ensureCapacity(DEFAULT_CAPACITY);
	}

	public int size() {
		return this.size;
	}

	public AnyType get(int index) {
		if (index < 0 || index >= this.size()) {
			throw new ArrayIndexOutOfBoundsException();
		}
		return this.items[index];
	}

	public void set(int index, AnyType newVal) {
		if (index < 0 || index >= this.size()) {
			throw new ArrayIndexOutOfBoundsException();
		}
		this.items[index] = newVal;
	}

	public void add(int index, AnyType item) {
		if (this.size() == this.items.length) {
			ensureCapacity(2 * this.size());
		}

		for (int i = this.size; i > index; i--) {
			this.items[i] = this.items[i - 1];
		}
		this.items[index] = item;

		this.size++;
	}

	public void add(AnyType item) {
		this.add(this.size(), item);
	}

	public AnyType remove(int index) {
		if (index < 0 || index >= this.size()) {
			throw new ArrayIndexOutOfBoundsException();
		}
		AnyType removedItem = this.items[index];
		for (int i = index; i < this.size(); i++) {
			this.items[i] = this.items[i + 1];
		}

		this.size--;
		return removedItem;
	}

	@SuppressWarnings("unchecked")
	private void ensureCapacity(int newCapacity) {
		if (newCapacity < this.size)
			return;

		AnyType[] old = this.items;
		items = (AnyType[]) new Object[newCapacity];
		for (int i = 0; i < this.size(); i++) {
			items[i] = old[i];
		}
	}
}