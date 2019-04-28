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
	URL              string `json:"url" redis:"url"`
	ImageURL         string `json:"imageUrl" redis:"imageUrl"`
	ImageURLForAPI   string `json:"imageUrlForApi" redis:"imageUrlForApi"`
	DiagramID        string `json:"diagramId" redis:"diagramId"`
	Title            string `json:"title" redis:"title"`
	Description      string `json:"description" redis:"description"`
	Security         string `json:"security" redis:"security"`
	Type             string `json:"type" redis:"type"`
	OwnerName        string `json:"ownerName" redis:"ownerName"`
	OwnerNickname    string `json:"ownerNickname" redis:"ownerNickname"`
	Owner            User   `json:"owner" redis:"owner"`
	Editing          bool   `json:"editing" redis:"editing"`
	Own              bool   `json:"own" redis:"own"`
	Shared           bool   `json:"shared" redis:"shared"`
	FolderID         int    `json:"folderId" redis:"folderId"`
	FolderName       string `json:"folderName" redis:"folderName"`
	ProjectID        string `json:"projectId" redis:"projectId"`
	ProjectName      string `json:"projectName" redis:"projectName"`
	OrganizationKey  string `json:"organizationKey" redis:"organizationKey"`
	OrganizationName string `json:"organizationName" redis:"organizationName"`
	SheetCount       int    `json:"sheetCount" redis:"sheetCount"`
	Created          string `json:"created" redis:"created"`
	Updated          string `json:"updated" redis:"updated"`
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
