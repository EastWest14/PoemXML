package Tester;

import Cases.*;

public class Tester {
	public static void main(String[] args) {
		GroupStore store = new GroupStore();
		RegressionGroup allGroups[] = store.allGroups();
		GroupRunResult runResults[] = runGroups(allGroups);
		outputResultsToConsole(runResults);
	}
	
	private static GroupRunResult[] runGroups(RegressionGroup groups[]) {
		GroupRunResult results[] = new GroupRunResult[groups.length];
		int i = 0;
		for (RegressionGroup group: groups) {
			results[i] = group.execute();
			i++;
		}
		return results;
	}
	
	private static void outputResultsToConsole(GroupRunResult results[]) {
		boolean regressionSuitPasses = true;
		int numGrRunned = 0;
		int numFailedGroups = 0;
		for (GroupRunResult result: results) {
			numGrRunned++;
			if (!result.passes()) {
				numFailedGroups++;
				regressionSuitPasses = false;
			}
		}
		
		if (regressionSuitPasses) {
			System.out.println("===");
			System.out.println("Regression suit passed!");
			System.out.println("Number of cases run: " + numGrRunned);
			System.out.println("===");
			System.exit(0);
		}
		
		System.out.println("xxx");
		System.out.println("Regression suit failed!");
		System.out.println("Number of cases run: " + numGrRunned);
		System.out.println("Number of groups failed: " + numFailedGroups);
		System.out.println("xxx");
		System.exit(1);
	}
}
