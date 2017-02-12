package Cases.ListPoemsGroup;

import Cases.*;
import Configs.Config;

public class ListPoemsGroup implements RegressionGroup {
	private RegressionTestCase listOfCases[];
	
	public ListPoemsGroup(Config regressionConfig) {
		RegressionTestCase listOfCases[] = new RegressionTestCase[1];
		listOfCases[0] = new ListPoemsCase(regressionConfig);
		this.listOfCases = listOfCases;
	}
	
	public GroupRunResult execute() {
		return CaseRunner.execute(this.listOfCases);
	}
}
