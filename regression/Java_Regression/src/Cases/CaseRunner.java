package Cases;

public class CaseRunner {
	public static GroupRunResult execute(String groupName, RegressionTestCase[] cases) {
		boolean passes = true;
		String failMessage = "";
		
		for (RegressionTestCase aCase: cases) {
			CaseRunResult result = aCase.run();
			if (!result.passes()) {
				passes = false;
				failMessage += "\t{" + result.caseName() + ":  " + result.failMessage() + "}\n";
			}
		}
		
		return new GroupRunResult(groupName, passes, failMessage);
	}
}
