package entities
type RecommendAMS struct {
	Max        int
	AMSVersion string
	AMSRevison string
}

type AMS struct {
	AMSServer   string `json:"server"`
	AMSVersion  string `json:"version"`
	AMSRevision string `json:"revision"`
	AMSUsername string `json:"username"`
	AMSPassword string `json:"password"`
}

