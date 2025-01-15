package models

type Points struct {
	Points int `json:"points"`
}

type ProcessedReceipt struct {
	Id string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func CreatePointsResponse(points int) Points {
	return Points{
		Points: points,
	}
}

func CreateProcessedResponse(id string) ProcessedReceipt {
	return ProcessedReceipt{
		Id: id,
	}
}

func CreateErrorResponse(error string) ErrorResponse {
	return ErrorResponse{
		Error: error,
	}
}
