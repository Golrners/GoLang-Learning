package Models

import (
	_ "encoding/json"
)

type Project struct {
	Github      string   `json:"github"`
	TagLine     string   `json:"tagLine"`
	ProjectName string   `json:"projectName"`
	ImagesURL   []string `json:"images"`
	ID          string   `json:"ID"`
}

type Projects struct {
	Projects []Project `json:"Items"`
}
