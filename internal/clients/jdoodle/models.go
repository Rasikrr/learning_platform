package jdoodle

type ExecuteRequest struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Script       string `json:"script"`
	Stdin        string `json:"stdin,omitempty"`
	Language     string `json:"language"`
	VersionIndex string `json:"versionIndex"`
	CompileOnly  bool   `json:"compileOnly"`
}

type ExecuteResponse struct {
	Output            string `json:"output"`
	StatusCode        int    `json:"statusCode"`
	Memory            string `json:"memory"`
	CPUTime           string `json:"cpuTime"`
	CompilationStatus int    `json:"compilationStatus"`
}
