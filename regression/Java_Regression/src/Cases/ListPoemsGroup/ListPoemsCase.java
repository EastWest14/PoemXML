package Cases.ListPoemsGroup;

import Cases.*;
import Configs.Config;

import java.net.*;
import java.util.Objects;
import java.io.*;


public class ListPoemsCase implements RegressionTestCase {
	private Config regressionConfig;
	
	private final String expectedResponse = "List of poems: 3 poems.\nID_1\nID_2\nID_3\n";
	
	public ListPoemsCase(Config regressionConfig) {
		this.regressionConfig = regressionConfig;
	}
	
	public CaseRunResult run() throws NullPointerException {
		URL url = null;
		try {
			url = new URL(this.regressionConfig.ServerUrl() + "/poems/list");
		} catch(NullPointerException e) {
			throw e;
		} catch(Exception e) {
			System.out.println("Unknown exception when building the URL: " + e);
		}
		
		try (BufferedReader reader = new BufferedReader(new InputStreamReader(url.openStream(), "UTF-8"))) {
		    for (String line; (line = reader.readLine()) != null;) {
		        if (!Objects.equals(line, expectedResponse)) {
		        	return new CaseRunResult(false);
		        }
		    }
		} catch (Exception e) {
			System.out.println("Unknown exception: " + e);
		}
		
		return new CaseRunResult(true);
	}
}
