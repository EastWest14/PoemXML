package Cases.SmokeGroup;

import Cases.*;
import Configs.Config;

import java.net.*;
import java.util.Objects;
import java.io.*;


public class InfoCase implements RegressionTestCase {
	private Config regressionConfig;
	private static String caseName = "Info Case";
	private static String expectedLine = "Hello, world!";
	
	public InfoCase(Config regressionConfig) {
		this.regressionConfig = regressionConfig;
	}
	
	public CaseRunResult run() throws NullPointerException {
		URL url = null;
		try {
			url = new URL(this.regressionConfig.ServerUrl() + "/");
		} catch(NullPointerException e) {
			throw e;
		} catch(Exception e) {
			System.out.println("Unknown exception when building the URL: " + e);
		}
		
		try (BufferedReader reader = new BufferedReader(new InputStreamReader(url.openStream(), "UTF-8"))) {
		    for (String line; (line = reader.readLine()) != null;) {
		        if (!Objects.equals(line, expectedLine)) {
		        	return new CaseRunResult(false, caseName, "Expected line: " + expectedLine + " got: " + line);
		        }
		    }
		} catch (Exception e) {
			System.out.println("Unknown exception: " + e);
			return new CaseRunResult(false, caseName, "Unknown exception: " + e);
		}
		
		return new CaseRunResult(true, caseName, "");
	}
}
