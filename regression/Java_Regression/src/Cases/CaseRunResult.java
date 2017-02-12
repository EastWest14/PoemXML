package Cases;

public final class CaseRunResult {
	private boolean passes;
	private String failMessage;

	private CaseRunResult() {
		this.passes = false;
	}

	public CaseRunResult(boolean passes) {
		this.passes = passes;
	}
	
	public CaseRunResult(boolean passes, String failMessage) {
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
}
