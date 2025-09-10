package models

type OsagoResponse struct {
	Rez []OSAGO `json:"rez"`
}

type OSAGO struct {
	Seria      string `json:"seria"`
	Nomer      string `json:"nomer"`
	OrgOsago   string `json:"orgosago"`
	Status     string `json:"status"`
	Term       string `json:"term"`
	TermStart  string `json:"termStart"`
	TermStop   string `json:"termStop"`
	StartPolis string `json:"startPolis"`
	StopPolis  string `json:"stopPolis"`
	BrandModel string `json:"brandmodel"`
	RegNum     string `json:"regnum"`
	DopBelarus string `json:"dopbelarus"`
}
