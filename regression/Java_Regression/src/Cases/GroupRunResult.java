package Cases;

public final class GroupRunResult {
	private boolean passes;

	public GroupRunResult() {
		this.passes = false;
	}

	public GroupRunResult(boolean passes) {
		this.passes = passes;
	}

	public boolean passes() {
		return this.passes;
	}
	public void setPasses(boolean passes) {
		this.passes = passes;
	}
}
