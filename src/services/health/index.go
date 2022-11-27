package health_service

type Status struct {
	Status string `json:"status"`
}

func GetHealth() Status {
	return Status{Status: "healthy"}
}
