package Cases;

import static org.junit.Assert.*;

import org.junit.Test;

public class AlwaysFailCaseTest {

	@Test
	public void test() {
		AlwaysFailCase aCase = new AlwaysFailCase();
		CaseRunResult result = aCase.run();
		assertEquals(result.passes(), false);
	}

}
