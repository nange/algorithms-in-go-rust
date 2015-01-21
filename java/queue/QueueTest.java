package queue;

import static org.junit.Assert.*;

import org.junit.Before;
import org.junit.Test;

public class QueueTest {
	private static Queue<String> q = new Queue<String>();

	@Before
	public void setUp() throws Exception {
		q = new Queue<String>();
	}

	@Test
	public void testEnqueue() {
		q.enqueue("1");
		q.enqueue("2");
		
		assertEquals(2, q.size());
		assertEquals(false, q.isEmpty());
	}

	@Test
	public void testDequeue() {
		q.enqueue("2");
		q.enqueue("3");
		
		assertEquals("2", q.dequeue());
		assertEquals(1, q.size());
	}

	@Test
	public void testSize() {
		q.enqueue("3");
		q.enqueue("4");
		assertEquals(2, q.size());
		
		q.dequeue();
		assertEquals(1, q.size());
		
		q.dequeue();
		assertEquals(0, q.size());
		assertEquals(true, q.isEmpty());
	}

	@Test
	public void testIsEmpty() {
		assertEquals(true, q.isEmpty());
		q.enqueue("test");
		assertEquals(false, q.isEmpty());
	}

	@Test
	public void testFont() {
		q.enqueue("5");
		q.enqueue("6");
		q.enqueue("7");
		
		assertEquals("5", q.font());
	}

	@Test
	public void testEnd() {
		q.enqueue("5");
		q.enqueue("6");
		q.enqueue("7");
		
		assertEquals("7", q.end());
	}

	@Test
	public void testClear() {
		q.enqueue("5");
		q.enqueue("7");
		
		q.clear();
		assertEquals(true, q.isEmpty());
	}
}
