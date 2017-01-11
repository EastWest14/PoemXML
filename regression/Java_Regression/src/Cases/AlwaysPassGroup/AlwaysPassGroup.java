package Cases.AlwaysPassGroup;

import Cases.*;

public final class AlwaysPassGroup implements RegressionGroup {
	private RegressionTestCase listOfCases[];
	
	public AlwaysPassGroup() {
		RegressionTestCase listOfCases[] = new RegressionTestCase[1];
		listOfCases[0] = new AlwaysPassCase();
		this.listOfCases = listOfCases;
	}
	
	public GroupRunResult execute() {
		for (RegressionTestCase rCase: listOfCases) {
			CaseRunResult result = rCase.run();
			if (!result.passes()) {
				return new GroupRunResult(false);
			}
		}
		return new GroupRunResult(true);
	}
}