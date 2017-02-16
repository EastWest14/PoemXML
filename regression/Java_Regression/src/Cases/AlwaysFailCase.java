package Cases;

public class AlwaysFailCase implements RegressionTestCase {
	public CaseRunResult run() {
		return new CaseRunResult(false, "Always Fail Case", "Always fails");
	}
}
