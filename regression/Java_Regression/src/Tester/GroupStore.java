package Tester;

import Cases.*;
import Cases.SmokeGroup.*;
import Configs.*;

public class GroupStore {
	private RegressionGroup groupsAvailable[]; 
	GroupStore(Config regressionConfig) {
		this.groupsAvailable = new RegressionGroup[1];
		this.groupsAvailable[0] = new SmokeGroup(regressionConfig);
	}
	
	public RegressionGroup[] allGroups() {
		return this.groupsAvailable;
	}
}
