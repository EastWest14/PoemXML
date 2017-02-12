package utils;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;

public class SimpleHTTPServer {
	public final static int portNumber = 8000;
	
	private HttpServer server;
	
	  public SimpleHTTPServer(String path, String responses[]) throws Exception {
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