package Cases;

public final class CaseRunResult {
	private boolean passes;
	private String caseName;
	private String failMessage;

	private CaseRunResult() {
		this.passes = false;
	}

	private CaseRunResult(boolean passes) {
		this.passes = passes;
	}
	
	public CaseRunResult(boolean passes, String caseName, String failMessage) {
		this.passes = passes;
		this.caseName = caseName;
		this.failMessage = failMessage;
	}

	public boolean passes() {
		return this.passes;
	}
	public void setPasses(boolean passes) {
		this.passes = passes;
	}
	
	public String caseName() {
		return this.caseName;
	}
	
	public String failMessage() {
		return this.failMessage;
	}
}
