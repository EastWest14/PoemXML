package Cases;

import static org.junit.Assert.*;

import org.junit.Test;

public class GroupRunResultTest {

	@Test
	public void test() {
		GroupRunResult grRunResultPasses = new GroupRunResult(true);
		GroupRunResult grRunResultFails = new GroupRunResult(false);
	
		assertEquals(grRunResultPasses.passes(), true);

		assertEquals(grRunResultFails.passes(), false);

		grRunResultFails.setPasses(true);
		assertEquals(grRunResultPasses.passes(), true);

		grRunResultFails.setPasses(false);
		assertEquals(grRunResultPasses.passes(), true);
	}
}
