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

type AgentNestedData struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	MainImgLink   string `json:"mainImgLink"`
	CompanyName   string `json:"companyName"`
	ContactNumber string `json:"contactNumber"`
	PageNumber    int    `json:"pageNumber"`
}
type Agents map[string]AgentNestedData

type Settings struct {
	LastSuccessPage int `json:"lastSuccessPage"`
	CurrentPage     int `json:"currentPage"`
}
