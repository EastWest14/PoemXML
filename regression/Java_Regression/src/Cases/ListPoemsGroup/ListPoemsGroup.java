package Cases.ListPoemsGroup;

import Cases.*;
import Configs.Config;

public class ListPoemsGroup implements RegressionGroup {
	private RegressionTestCase listOfCases[];
	private static String groupName = "List Poems Group";
	
	public ListPoemsGroup(Config regressionConfig) {
		RegressionTestCase listOfCases[] = new RegressionTestCase[1];
		listOfCases[0] = new ListPoemsCase(regressionConfig);
		this.listOfCases = listOfCases;
	}
	
	public GroupRunResult execute() {
		return CaseRunner.execute(groupName, this.listOfCases);
	}
}
