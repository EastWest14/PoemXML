package Cases;

public final class CaseRunResult {
	private boolean passes;

	CaseRunResult() {
		this.passes = false;
	}

	CaseRunResult(boolean passes) {
		this.passes = passes;
	}

	public boolean passes() {
		return this.passes;
	}
	public void setPasses(boolean passes) {
		this.passes = passes;
	}
}
