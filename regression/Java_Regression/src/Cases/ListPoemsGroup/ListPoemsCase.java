package Cases.ListPoemsGroup;

import Cases.*;
import Configs.Config;

import java.net.*;
import java.util.Objects;
import java.io.*;


public class ListPoemsCase implements RegressionTestCase {
	private Config regressionConfig;
	private static String caseName = "List Poems Case";
	
	private final String[] expectedResponses = {"List of poems: 3 poems.", "ID_1", "ID_2", "ID_3"};
	
	public ListPoemsCase(Config regressionConfig) {
		this.regressionConfig = regressionConfig;
	}
	
	public CaseRunResult run() throws NullPointerException {
		URL url = null;
		try {
			url = new URL(this.regressionConfig.ServerUrl() + "/poems");
		} catch(NullPointerException e) {
			throw e;
		} catch(Exception e) {
			System.out.println("Unknown exception when building the URL: " + e);
		}
		
		try (BufferedReader reader = new BufferedReader(new InputStreamReader(url.openStream(), "UTF-8"))) {
			int expectedNumberOfLines = expectedResponses.length;
			int currentExpectedLine = 0;

		    for (String line; (line = reader.readLine()) != null;) {
		    	if (currentExpectedLine >= expectedNumberOfLines) {
		    		return new CaseRunResult(false, caseName, "Expected " + expectedNumberOfLines + " lines, got more than that.");
		    	}

		    	String expectedLine = expectedResponses[currentExpectedLine];

		        if (!Objects.equals(line, expectedLine)) {
		        	return new CaseRunResult(false, caseName, "Expected: " + expectedLine + " got: " + line);
		        }
		        currentExpectedLine++;
		    }
		} catch (Exception e) {
			System.out.println("Unknown exception: " + e);
		}
		
		return new CaseRunResult(true, caseName, "");
	}
}
