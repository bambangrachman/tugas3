package model

type Data struct {
	Status struct {
		Water string `json:"water"`
		Wind  string `json:"wind"`
	} `json:"status"`
}

type Result struct {
	Weather Data   `json:"weather"`
	Status  string `json:"status"`
}
