package Cases;

import static org.junit.Assert.*;
import org.junit.Test;

public class CaseRunResultTest {
	@Test
	public void test() {
		CaseRunResult caseRunResultDefault = new CaseRunResult();
		CaseRunResult caseRunResultPasses = new CaseRunResult(true);
		CaseRunResult caseRunResultFails = new CaseRunResult(false);

		boolean passes = caseRunResultDefault.passes();
		assertEquals(passes, false);

		passes = caseRunResultPasses.passes();
		assertEquals(passes, true);
		
		passes = caseRunResultFails.passes();
		assertEquals(passes, false);

		caseRunResultFails.setPasses(true);
		passes = caseRunResultPasses.passes();
		assertEquals(passes, true);

		caseRunResultFails.setPasses(false);
		passes = caseRunResultPasses.passes();
		assertEquals(passes, true);
	}

}
