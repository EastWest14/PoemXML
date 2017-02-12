package Cases.SmokeGroup;

import static org.junit.Assert.*;

import org.junit.Test;

import Configs.Config;
import Cases.*;

import utils.SimpleHTTPServer;

public class InfoCaseTest {

	@Test
	public void test() {
		SimpleHTTPServer mockServer = null;
		try {
			String responses[] = {"Hello, world!", "Wrong_response"};
			mockServer = new SimpleHTTPServer("/poems", responses);
		} catch(Exception e) {
			fail("Error creating server");
		}
		Config regressionConfig = new Config("http://localhost:" + SimpleHTTPServer.portNumber);
		InfoCase aCase = new InfoCase(regressionConfig);
		
		//Run pass case
		CaseRunResult result = aCase.run();
		assertEquals(result.passes(), true);
		
		//Run fail case
		result = aCase.run();
		assertEquals(result.passes(), false);
		
		//No response case
		mockServer.stop();
		result = aCase.run();
		assertEquals(result.passes(), false);
	}

}
