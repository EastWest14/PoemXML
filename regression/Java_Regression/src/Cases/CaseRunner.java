package Cases;

public class CaseRunner {
	public static GroupRunResult execute(RegressionTestCase[] cases) {
		boolean passes = true;
		
		for (RegressionTestCase aCase: cases) {
			CaseRunResult result = aCase.run();
			if (!result.passes()) {
				passes = false;
			}
		}
		
		return new GroupRunResult(passes);
	}
}
