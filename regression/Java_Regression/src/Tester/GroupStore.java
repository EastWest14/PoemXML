package Tester;

import Cases.*;
import Cases.AlwaysPassGroup.*;
import Cases.AlwaysFailGroup.*;

public class GroupStore {
	private RegressionGroup groupsAvailable[]; 
	GroupStore() {
		this.groupsAvailable = new RegressionGroup[2];
		this.groupsAvailable[0] = new AlwaysPassGroup();
		this.groupsAvailable[1] = new AlwaysFailGroup();
	}
	
	public RegressionGroup[] allGroups() {
		return this.groupsAvailable;
	}
}
