package Cases.AlwaysPassGroup;

import Cases.*;

public final class TestAlwaysPassCase {
	public static void main(String args[]) {
		boolean testPass = true;

		AlwaysPassCase aCase = new AlwaysPassCase();
		CaseRunResult result = aCase.run();
		boolean passes = result.passes();
		if (!passes) {
			System.out.println("Test fail: Always pass case fails.");
			testPass = false;
		}

		if (testPass) {
			System.out.println("Test Pass");
		}
	}
} 