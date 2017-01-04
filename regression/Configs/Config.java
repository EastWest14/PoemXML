package Configs;

public final class Config {
	private String serverUrl;

	Config() throws Exception {
		throw new Exception("Initializer disabled");
	}

	Config(String serverUrl) {
		this.serverUrl = serverUrl;
	}

	public void SetServerUrl(String serverUrl) {
		this.serverUrl = serverUrl;
	}

	public String ServerUrl() {
		return this.serverUrl;
	}
}
