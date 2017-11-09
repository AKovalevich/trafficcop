package endpoint

type Endpoint struct {
	Pattern			string `json:"pattern"`
	SchemaVersion	string `json:"schema_version"`
	Async			string `json:"async"`
	Command			string `json:"command"`
	Method			string `json:"method"`
}

type Endpoints []Endpoint
