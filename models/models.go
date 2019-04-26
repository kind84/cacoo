package models

// User representation.
type User struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Type     string `json:"type"`
	ImageURL string `json:"imageUrl"`
}

// Diagram representation.
type Diagram struct {
	URL              string `json:"url"`
	ImageURL         string `json:"imageUrl"`
	ImageURLForAPI   string `json:"imageUrlForApi"`
	DiagramID        string `json:"diagramId"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Security         string `json:"security"`
	Type             string `json:"type"`
	OwnerName        string `json:"ownerName"`
	OwnerNickname    string `json:"ownerNickname"`
	Owner            User   `json:"owner"`
	Editing          bool   `json:"editing"`
	Own              bool   `json:"own"`
	Shared           bool   `json:"shared"`
	FolderID         int    `json:"folderId"`
	FolderName       string `json:"folderName"`
	ProjectID        string `json:"projectId"`
	ProjectName      string `json:"projectName"`
	OrganizationKey  string `json:"organizationKey"`
	OrganizationName string `json:"organizationName"`
	SheetCount       int    `json:"sheetCount"`
	Created          string `json:"created"`
	Updated          string `json:"updated"`
}

// DiagramDetail representation.
type DiagramDetail struct {
	Diagram
	Sheets []struct {
		URL            string `json:"url"`
		ImageURL       string `json:"imageUrl"`
		ImageURLForAPI string `json:"imageUrlForApi"`
		UID            string `json:"uid"`
		Name           string `json:"name"`
		Width          int    `json:"width"`
		Height         int    `json:"height"`
	} `json:"sheets"`
	Comments []struct {
		User    User   `json:"user"`
		Content string `json:"content"`
		Created string `json:"created"`
		Updated string `json:"updated"`
	} `json:"comments"`
}

// Diagrams representation.
type Diagrams struct {
	Result []Diagram `json:"result"`
	Count  int       `json:"count"`
}

// Folder representation.
type Folder struct {
	FolderID   int    `json:"folderId"`
	FolderName string `json:"folderName"`
	Type       string `json:"type"`
	Created    string `json:"created"`
	Updated    string `json:"updated"`
}

// Folders representation.
type Folders struct {
	Result []Folder `json:"result"`
}
