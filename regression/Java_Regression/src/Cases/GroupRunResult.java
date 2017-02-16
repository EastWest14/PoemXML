package Cases;

public final class GroupRunResult {
	private boolean passes;
	private String failMessage;
	private String groupName;

	private GroupRunResult() {
		this.passes = false;
	}

	public GroupRunResult(boolean passes) {
		this.passes = passes;
		this.failMessage = "";
		this.groupName = "Default group name";
	}

	public GroupRunResult(String groupName, boolean passes, String failMessage) {
		this.groupName = groupName;
		this.passes = passes;
		this.failMessage = failMessage;
	}
	
	public boolean passes() {
		return this.passes;
	}
	public void setPasses(boolean passes) {
		this.passes = passes;
	}
	
	public String failMessage() {
		return this.failMessage;
	}
	
	public String groupName() {
		return this.groupName;
	}
}
