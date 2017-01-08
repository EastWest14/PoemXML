package Cases.AlwaysPassGroup;

import Cases.*;

public class AlwaysPassCase implements RegressionTestCase {
	public CaseRunResult run() {
		return new CaseRunResult(true);
	}
}