package Cases.SmokeGroup;

import static org.junit.Assert.*;

import org.junit.Test;

import Configs.Config;
import Cases.*;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;

class SimpleHttpServer {
	public final static int portNumber = 8000;
	
	private HttpServer server;
	
	  SimpleHttpServer(String path, String responses[]) throws Exception {
	    server = HttpServer.create(new InetSocketAddress(portNumber), 0);
	    server.createContext(path, new MyHandler(responses));
	    server.setExecutor(null); // creates a default executor
	    server.start();
	  }
	  
	  public void stop() {
		  server.stop(0);
	  }

	  class MyHandler implements HttpHandler {
		private String responses[];
		private int responseNum = 0;
		
		MyHandler(String responses[]) {
			this.responses = responses;
		}
		
	    public void handle(HttpExchange t) throws IOException, IndexOutOfBoundsException {
	    	if (this.responses.length <= this.responseNum) {
	    		throw new IndexOutOfBoundsException("Not enough responses specified for the server");
	    	}
	    	String response = this.responses[this.responseNum];
	    	this.responseNum++;
	    	t.sendResponseHeaders(200, response.length());
	    	OutputStream os = t.getResponseBody();
	    	os.write(response.getBytes());
	    	os.close();
	      }
	  }
}

public class InfoCaseTest {

	@Test
	public void test() {
		SimpleHttpServer mockServer = null;
		try {
			String responses[] = {"Hello, world!", "Wrong_response"};
			mockServer = new SimpleHttpServer("/poems", responses);
		} catch(Exception e) {
			fail("Error creating server");
		}
		Config regressionConfig = new Config("http://localhost:" + SimpleHttpServer.portNumber);
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
