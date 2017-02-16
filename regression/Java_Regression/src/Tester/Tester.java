package Tester;

import Cases.*;
import Configs.*;

public class Tester {
	public static void main(String[] args) {
		Config regressionConfig = new Config("http://localhost:8080");
		
		GroupStore store = new GroupStore(regressionConfig);
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
		
		System.out.println("===");
		//Run, List groups one by one with pass messages and error messages
		for (GroupRunResult result: results) {
			numGrRunned++;
			boolean passes = result.passes();
			if (!passes) {
				numFailedGroups++;
				regressionSuitPasses = false;
			}
			System.out.println("\"" + result.groupName() + "\": " + passes);
		}
		System.out.println("");
		
		if (regressionSuitPasses) {
			System.out.println("Regression suit passed!");
		} else {
			System.out.println("Regression suit failed!");
		}
		System.out.println("Number of cases run: " + numGrRunned);
		System.out.println("Number of groups failed: " + numFailedGroups);
		System.out.println("===");
		
		if (regressionSuitPasses) {
			System.exit(0);
		} else {
			System.exit(1);
		}
		
		/*if (regressionSuitPasses) {
			
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
		System.exit(1);*/
	}
}
