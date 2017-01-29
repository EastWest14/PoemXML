package Cases.SmokeGroup;

import static org.junit.Assert.*;

import org.junit.Before;
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
	  SimpleHttpServer() throws Exception {
	    HttpServer server = HttpServer.create(new InetSocketAddress(8000), 0);
	    server.createContext("/poem", new MyHandler());
	    server.setExecutor(null); // creates a default executor
	    server.start();
	  }

	  static class MyHandler implements HttpHandler {
	    public void handle(HttpExchange t) throws IOException {
	    	String response = "Hello, world!";
	    	t.sendResponseHeaders(200, response.length());
	    	OutputStream os = t.getResponseBody();
	    	os.write(response.getBytes());
	    	os.close();
	      }
	  }
}

public class InfoCaseTest {
	
	@Before
	public void setUp() throws Exception {
		System.out.println("Starting");

		System.out.println("Setup done");
	}

	@Test
	public void test() {
		try {
			SimpleHttpServer mockServer = new SimpleHttpServer();
		} catch(Exception e) {
			fail("Error creating server");
		}
		Config regressionConfig = new Config("http://localhost:8000");
		InfoCase aCase = new InfoCase(regressionConfig);
		
		System.out.println("About to run case");
		CaseRunResult result = aCase.run();
		System.out.println("Case ran: " + result.passes());
		
		//fail("Not yet implemented");
	}

}
