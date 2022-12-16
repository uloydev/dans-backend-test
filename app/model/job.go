package model

type JobFilter struct {
	FullTime    string `query:"full_time"`
	Description string `query:"description"`
	Location    string `query:"location"`
	Page        string `query:"page"`
}

type JobResponse struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	Company     string `json:"company"`
	CompanyUrl  string `json:"company_url"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	CompanyLogo string `json:"company_logo"`
	Location    string `json:"location"`
}

type JobRequest struct {
}
