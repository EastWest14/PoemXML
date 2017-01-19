package Cases.SmokeGroup;

import Cases.*;

public final class SmokeGroup implements RegressionGroup {
	private RegressionTestCase listOfCases[];
	
	public SmokeGroup() {
		RegressionTestCase listOfCases[] = new RegressionTestCase[1];
		listOfCases[0] = new InfoCase();
		this.listOfCases = listOfCases;
	}
	
	public GroupRunResult execute() {
		return CaseRunner.execute(this.listOfCases);
	}
}
