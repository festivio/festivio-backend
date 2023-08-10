package domain

type Team struct {
	Count struct {
		Animators  int `json:"animators"`
		Managers   int `json:"managers"`
		Founders   int `json:"founders"`
		TotalCount int `json:"total_count"`
	} `json:"count"`
	Employees []*ShortUserInfo `json:"employees"`
}
