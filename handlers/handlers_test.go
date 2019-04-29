package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kind84/cacoo/models"

	"github.com/julienschmidt/httprouter"
)

var d1 = models.Diagram{
	URL:            "https://cacoo.com/diagrams/8PQ05Yp2unrpN996",
	ImageURL:       "https://cacoo.com/diagrams/8PQ05Yp2unrpN996.png",
	ImageURLForAPI: "https://cacoo.com/api/v1/diagrams/8PQ05Yp2unrpN996.png",
	DiagramID:      "8PQ05Yp2unrpN996",
	Title:          "Ready 2 Rumble",
	Description:    "",
	Security:       "private",
	Type:           "normal",
	OwnerName:      "6xRj7b6IfZGJR6R5",
	OwnerNickname:  "Hulk Hogan",
	Owner: models.User{
		Name:     "6xRj7b6IfZGJR6R5",
		Nickname: "Hulk Hogan",
		Type:     "other",
		ImageURL: "https://cacoo.com/account/6xRj7b6IfZGJR6R5/image/32x32",
	},
	Editing:          false,
	Own:              true,
	Shared:           false,
	FolderID:         0,
	FolderName:       "",
	ProjectID:        "",
	ProjectName:      "",
	OrganizationKey:  "",
	OrganizationName: "",
	SheetCount:       1,
	Created:          "Thu, 25 Apr 2019 15:56:57 +0200",
	Updated:          "Thu, 25 Apr 2019 15:57:33 +0200",
}

type repo struct{}

func (r repo) Save(k interface{}, v interface{})       {}
func (r repo) SaveSet(k interface{}, v ...interface{}) {}

func (r repo) GetASet(k interface{}) []string {
	js, err := json.Marshal(d1)
	if err != nil {
		panic(err)
	}
	return []string{string(js)}
}

func TestGetFolder(t *testing.T) {
	r := repo{}

	router := httprouter.New()
	router.GET("/api/folder/:id", GetFolder(r))

	req, _ := http.NewRequest("GET", "/api/folder/123", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status, got %v, want %v\n", status, http.StatusOK)
	}

	decoder := json.NewDecoder(rr.Body)
	var f struct {
		ID       string           `json:"id"`
		Diagrams []models.Diagram `json:"diagrams"`
	}
	decoder.Decode(&f)

	if f.ID != "123" {
		t.Errorf("Wrong FolderID: expected 123, got %v", f.ID)
	}
	if len(f.Diagrams) != 1 {
		t.Errorf("Wrong quantity of diagrams: expected 1, got %d", len(f.Diagrams))
	}
	if f.Diagrams[0].DiagramID != d1.DiagramID {
		t.Errorf("Wrong quantity of diagrams: expected %s, got %d", d1.DiagramID, len(f.Diagrams))
	}
}
