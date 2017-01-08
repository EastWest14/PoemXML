package Cases.AlwaysPassGroup;

import static org.junit.Assert.*;

import org.junit.Test;

import Cases.CaseRunResult;

public class AlwaysPassCaseTest {

	@Test
	public void test() {
		AlwaysPassCase aCase = new AlwaysPassCase();
		CaseRunResult result = aCase.run();
		assertEquals(result.passes(), true);
	}

}
