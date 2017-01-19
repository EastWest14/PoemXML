package Cases.SmokeGroup;

import Cases.*;
import java.net.*;
import java.util.Objects;
import java.io.*;


public class InfoCase implements RegressionTestCase {
	public CaseRunResult run() {
		URL url = null;
		try {
			url = new URL("http://localhost:8080/poems");
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

//You can make use of java.net.URL and/or java.net.URLConnection.



