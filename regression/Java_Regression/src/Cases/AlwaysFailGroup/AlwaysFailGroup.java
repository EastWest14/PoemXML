package Cases.AlwaysFailGroup;

import Cases.*;

public final class AlwaysFailGroup implements RegressionGroup {
	private static String groupName = "Always Fail Group";
	private RegressionTestCase listOfCases[];
	
	public AlwaysFailGroup() {
		RegressionTestCase listOfCases[] = new RegressionTestCase[1];
		listOfCases[0] = new AlwaysFailCase();
		this.listOfCases = listOfCases;
	}
	
	public GroupRunResult execute() {
		return CaseRunner.execute(groupName, this.listOfCases);
	}
}
