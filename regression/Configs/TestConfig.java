package Configs;

public final class TestConfig {
	public static void main(String args[]) throws Exception {
		boolean testPass = true;
		
		try {
			final String initialServerUrl = "Initial server URL";
			Config aConfig = new Config(initialServerUrl);

			try {
				String initialSetUrl = aConfig.ServerUrl();
				if (initialSetUrl != initialServerUrl) {
					System.out.println("Test fail: Initial server URL set/read incorrectly.");
					testPass = false;
				}

				final String serverUrl = "Server URL";

				aConfig.SetServerUrl(serverUrl);
				String returnedServerUrl = aConfig.ServerUrl();
				if (returnedServerUrl != serverUrl) {
					System.out.println("Test fail: Server URL set/read incorrectly.");
					testPass = false;
				}

				if (testPass) {
					System.out.println("Test Pass");
				}
			} catch (Exception e) {
				System.out.println("Unexpected exception");
				System.out.println("Test Fail");
			}
		} catch (Exception e) {
			System.out.println("Exception when initializing config");
			throw e;
		}
	}
}
