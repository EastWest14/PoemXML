package Cases.SmokeGroup;

import Cases.*;
import Configs.Config;

import java.net.*;
import java.util.Objects;
import java.io.*;


public class InfoCase implements RegressionTestCase {
	private Config regressionConfig;
	
	public InfoCase(Config regressionConfig) {
		this.regressionConfig = regressionConfig;
	}
	
	public CaseRunResult run() throws NullPointerException {
		URL url = null;
		try {
			url = new URL(this.regressionConfig.ServerUrl() + "/poems");
		} catch(NullPointerException e) {
			System.out.println("here");
			throw e;
		} catch(Exception e) {
			System.out.println("Unknown exception when building the URL: " + e);
		}
		
		try (BufferedReader reader = new BufferedReader(new InputStreamReader(url.openStream(), "UTF-8"))) {
		    for (String line; (line = reader.readLine()) != null;) {
		        if (!Objects.equals(line, "Hello, world!")) {
		        	System.out.println("Fail: " + line);
		        	return new CaseRunResult(false);
		        }
		    }
		} catch (Exception e) {
			System.out.println("Unknown exception: " + e);
		}
		
		System.out.println("Pass");
		return new CaseRunResult(true);
	}
}
