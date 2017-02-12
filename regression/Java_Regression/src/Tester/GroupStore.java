package Tester;

import Cases.*;
import Cases.SmokeGroup.*;
import Cases.ListPoemsGroup.*;
import Configs.*;

public class GroupStore {
	private RegressionGroup groupsAvailable[]; 
	GroupStore(Config regressionConfig) {
		this.groupsAvailable = new RegressionGroup[2];
		this.groupsAvailable[0] = new SmokeGroup(regressionConfig);
		this.groupsAvailable[1] = new ListPoemsGroup(regressionConfig);
	}
	
	public RegressionGroup[] allGroups() {
		return this.groupsAvailable;
	}
}
