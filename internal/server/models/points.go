package models

type Points struct {
	Points int `json:"points"`
}

func CreatePointsResponse(points int) Points {
	return Points{
		Points: points,
	}
}
