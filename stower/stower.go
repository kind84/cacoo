package stower

import (
	"encoding/json"
	"fmt"

	"github.com/kind84/cacoo/interfaces"
	"github.com/kind84/cacoo/models"
)

// Stower is responsible for storing data.
type Stower struct {
	repo interfaces.Repo
}

// NewStower ctor
func NewStower(repo interfaces.Repo) *Stower {
	return &Stower{repo}
}

// StowDgrams store diagrams grouped by folders.
func (s Stower) StowDgrams(dGrams []models.Diagram) {
	folders := make(map[string][]models.Diagram)
	for _, dGram := range dGrams {
		f := fmt.Sprintf("folder:%d", dGram.FolderID)
		folders[f] = append(folders[f], dGram)
	}

	for k, dGrams := range folders {
		for _, dg := range dGrams {
			v, _ := json.Marshal(dg)
			s.repo.SaveSet(k, v)
		}
	}
}
