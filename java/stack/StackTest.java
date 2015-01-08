package stack;

import static org.junit.Assert.*;

import org.junit.Test;

public class StackTest {

	@Test
	public void testStack() {
		Stack<String> stack = new Stack<String>();
		assertTrue(stack.isEmpty());
		assertEquals(0, stack.size());
		
		stack.push("1");
		stack.push("2");
		stack.push("3");
		assertEquals(3, stack.size());
		
		String p1 = stack.pop();
		assertEquals("3", p1);
		assertEquals(2, stack.size());
		
		String p2 = stack.peek();
		assertEquals("2", p2);
		assertEquals(2, stack.size());
		
		stack.clear();
		assertTrue(stack.isEmpty());
		
	}

}
