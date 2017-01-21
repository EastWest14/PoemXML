package Cases.SmokeGroup;

import Cases.*;
import Configs.Config;

public final class SmokeGroup implements RegressionGroup {
	private RegressionTestCase listOfCases[];
	
	public SmokeGroup(Config regressionConfig) {
		RegressionTestCase listOfCases[] = new RegressionTestCase[1];
		listOfCases[0] = new InfoCase(regressionConfig);
		this.listOfCases = listOfCases;
	}
	
	public GroupRunResult execute() {
		return CaseRunner.execute(this.listOfCases);
	}
}
