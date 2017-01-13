package Cases;

public class AlwaysPassCase implements RegressionTestCase {
	public CaseRunResult run() {
		return new CaseRunResult(true);
	}
}