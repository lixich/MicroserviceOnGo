package main

type Transaction struct {
	Sum                 float64 `json:"Sum"`
	YearEnd             int     `json:"YearEnd"`
	Hour                float64 `json:"Hour"`
	WeekDay             int     `json:"WeekDay"`
	HaveDeviceId        int     `json:"HaveDeviceId"`
	EuropeAsiaCountryIp int     `json:"EuropeAsiaCountryIp"`
	Refund              int     `json:"Refund"`
}

// TODO repeated code.
var TransactionKeys []string = []string{"Sum", "YearEnd", "Hour", "WeekDay", "HaveDeviceId", "EuropeAsiaCountryIp", "Refund"}
