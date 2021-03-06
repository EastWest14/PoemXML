package Cases;

import static org.junit.Assert.*;

import org.junit.Test;

public class CaseRunnerTest {

	@Test
	public void test() {
		RegressionTestCase[] cases = new RegressionTestCase[1];
		cases[0] = new AlwaysPassCase();
		GroupRunResult resultShouldPass = CaseRunner.execute("", cases); 
		assertEquals(resultShouldPass.passes(), true);
		
		cases = new RegressionTestCase[1];
		cases[0] = new AlwaysFailCase();
		resultShouldPass = CaseRunner.execute("", cases); 
		assertEquals(resultShouldPass.passes(), false);
		
		cases = new RegressionTestCase[2];
		cases[0] = new AlwaysFailCase();
		cases[1] = new AlwaysPassCase();
		resultShouldPass = CaseRunner.execute("", cases); 
		assertEquals(resultShouldPass.passes(), false);
		
		cases = new RegressionTestCase[2];
		cases[0] = new AlwaysPassCase();
		cases[1] = new AlwaysPassCase();
		resultShouldPass = CaseRunner.execute("", cases); 
		assertEquals(resultShouldPass.passes(), true);
	}

}
