package Cases;

public class AlwaysFailCase implements RegressionTestCase {
	public CaseRunResult run() {
		return new CaseRunResult(false);
	}
}
