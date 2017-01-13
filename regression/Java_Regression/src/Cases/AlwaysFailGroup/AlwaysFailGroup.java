package Cases.AlwaysFailGroup;

import Cases.*;

public final class AlwaysFailGroup implements RegressionGroup {
	private RegressionTestCase listOfCases[];
	
	public AlwaysFailGroup() {
		RegressionTestCase listOfCases[] = new RegressionTestCase[1];
		listOfCases[0] = new AlwaysFailCase();
		this.listOfCases = listOfCases;
	}
	
	public GroupRunResult execute() {
		return CaseRunner.execute(this.listOfCases);
	}
}
