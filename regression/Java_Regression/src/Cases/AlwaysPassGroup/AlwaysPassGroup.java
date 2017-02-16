package Cases.AlwaysPassGroup;

import Cases.*;

public final class AlwaysPassGroup implements RegressionGroup {
	private RegressionTestCase listOfCases[];
	private static String groupName = "Always Pass Group";
	
	public AlwaysPassGroup() {
		RegressionTestCase listOfCases[] = new RegressionTestCase[1];
		listOfCases[0] = new AlwaysPassCase();
		this.listOfCases = listOfCases;
	}
	
	public GroupRunResult execute() {
		return CaseRunner.execute(groupName, this.listOfCases);
	}
}
