package arraylist;

import static org.junit.Assert.*;

import org.junit.Before;
import org.junit.Test;

public class MyArrayListTest {
	
	private static MyArrayList<String> list = new MyArrayList<String>();
	
	@Before
	public void setUp() throws Exception {
		list.clear();
	}

	@Test
	public void testClear() {
		list.clear();
		assertEquals(0, list.size());
	}

	@Test
	public void testSize() {
		assertEquals(0, list.size());
		list.add("1");
		assertEquals(1, list.size());
	}

	@Test
	public void testGet() {
		list.add("1");
		list.add("2");
		list.add("3");
		String item = list.get(0);
		assertEquals("1", item);
	}

	@Test
	public void testSet() {
		list.add("1");
		list.add("2");
		list.add("3");
		list.set(2, "4");
		assertEquals("4", list.get(2));
	}

	@Test
	public void testAddIntAnyType() {
		list.add("1");
		list.add("2");
		list.add("3");
		list.add(0, "4");
		assertEquals("4", list.get(0));
	}

	@Test
	public void testAddAnyType() {
		list.add("1");
		list.add("2");
		list.add("3");
		assertEquals(3, list.size());
		assertEquals("3", list.get(2));
	}

	@Test
	public void testRemove() {
		list.add("1");
		list.add("2");
		list.add("3");
		list.remove(2);
		assertEquals(2, list.size());
		assertEquals("2", list.get(1));
	}
	
}
