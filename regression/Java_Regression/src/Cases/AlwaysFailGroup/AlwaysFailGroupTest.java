package Cases.AlwaysFailGroup;

import static org.junit.Assert.*;
import Cases.*;

import org.junit.Test;

public class AlwaysFailGroupTest {

	@Test
	public void test() {
		AlwaysFailGroup afGroup = new AlwaysFailGroup();
		GroupRunResult result = afGroup.execute();
		
		assertEquals(result.passes(), false);
	}

}
