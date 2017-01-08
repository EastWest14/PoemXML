package Cases;

public final class TestGroupRunResult {
	public static void main(String args[]) {
		boolean testPass = true;

		GroupRunResult grRunResultDefault = new GroupRunResult();
		GroupRunResult grRunResultPasses = new GroupRunResult(true);
		GroupRunResult grRunResultFails = new GroupRunResult(false);

		boolean passes = grRunResultDefault.passes();
		if (passes) {
			System.out.println("Test fail: Expected default group run result to be false, got true.");
			testPass = false;
		}
	
		passes = grRunResultPasses.passes();
		if (!passes) {
			System.out.println("Test fail: Expected group run result to be true, got false.");
			testPass = false;
		}

		passes = grRunResultFails.passes();
		if (passes) {
			System.out.println("Test fail: Expected group run result to be false, got true.");
			testPass = false;
		}

		grRunResultFails.setPasses(true);
		passes = grRunResultPasses.passes();
		if (!passes) {
			System.out.println("Test fail: Expected group run result to be reset to true, got false.");
			testPass = false;
		}

		grRunResultFails.setPasses(false);
		passes = grRunResultPasses.passes();
		if (!passes) {
			System.out.println("Test fail: Expected group run result to be reset to false, got true.");
			testPass = false;
		}

		if (testPass) {
			System.out.println("Test Pass");
		}
	}
}
