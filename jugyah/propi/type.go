package propi

type BuildingNestedData struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	ShortBhk     string `json:"shortBhk"`
	DetailLink   string `json:"detailLink"`
	MainImgLink  string `json:"mainImgLink"`
	ShortAddress string `json:"shortAddress"`
}
type Buildings map[string]BuildingNestedData

type OldAgentNestedData struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	MainImgLink   string `json:"mainImgLink"`
	CompanyName   string `json:"companyName"`
	ContactNumber string `json:"contactNumber"`
	PageNumber    int    `json:"pageNumber"`
	Link          string `json:"link"`
}
type OldAgentsMap map[string]OldAgentNestedData

type Settings struct {
	LastSuccessPage int `json:"lastSuccessPage"`
	CurrentPage     int `json:"currentPage"`
}

type AgentUrlDetail struct {
	Url        string `json:"url"`
	PageNumber int    `json:"pageNumber"`
	Fetched    bool   `json:"fetched"`
}
type AgentsUrlMap map[string]AgentUrlDetail

type AgentDetail struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	MainImgLink   string `json:"mainImgLink"`
	CompanyName   string `json:"companyName"`
	ContactNumber string `json:"contactNumber"`
	PageNumber    int    `json:"pageNumber"`
	Link          string `json:"link"`

	CertificationNumber string   `json:"certificationNumber"`
	Specialization      []string `json:"specialization"`
	AreaOfOperation     []string `json:"areaOfOperation"`
	ActiveInBuildings   []string `json:"activeInBuildings"`
	ActiveInLocations   []string `json:"activeInLocations"`

	Json map[string]interface{} `json:"extraData"`
}
type AgentsDetailMap map[string]AgentDetail
