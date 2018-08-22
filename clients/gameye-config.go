package clients

/*
GameyeClientConfig configures the GameyeClient
*/
type GameyeClientConfig struct {
	Endpoint string
	Token    string
}

func (config GameyeClientConfig) validate() {
	// TODO
}
