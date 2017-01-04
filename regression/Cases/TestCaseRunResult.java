package Cases;

public final class TestCaseRunResult {
	public static void main(String args[]) {
		boolean testPass = true;

		CaseRunResult caseRunResultDefault = new CaseRunResult();
		CaseRunResult caseRunResultPasses = new CaseRunResult(true);
		CaseRunResult caseRunResultFails = new CaseRunResult(false);

		boolean passes = caseRunResultDefault.passes();
		if (passes) {
			System.out.println("Test fail: Expected default case run result to be false, got true.");
			testPass = false;
		}
	
		passes = caseRunResultPasses.passes();
		if (!passes) {
			System.out.println("Test fail: Expected case run result to be true, got false.");
			testPass = false;
		}

		passes = caseRunResultFails.passes();
		if (passes) {
			System.out.println("Test fail: Expected case run result to be false, got true.");
			testPass = false;
		}

		caseRunResultFails.setPasses(true);
		passes = caseRunResultPasses.passes();
		if (!passes) {
			System.out.println("Test fail: Expected case run result to be reset to true, got false.");
			testPass = false;
		}

		caseRunResultFails.setPasses(false);
		passes = caseRunResultPasses.passes();
		if (!passes) {
			System.out.println("Test fail: Expected case run result to be reset to false, got true.");
			testPass = false;
		}

		if (testPass) {
			System.out.println("Test Pass");
		}
	}
}