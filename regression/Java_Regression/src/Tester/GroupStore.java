package Tester;

import Cases.*;
import Cases.SmokeGroup.*;

public class GroupStore {
	private RegressionGroup groupsAvailable[]; 
	GroupStore() {
		this.groupsAvailable = new RegressionGroup[1];
		this.groupsAvailable[0] = new SmokeGroup();
	}
	
	public RegressionGroup[] allGroups() {
		return this.groupsAvailable;
	}
}
