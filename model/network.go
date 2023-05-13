package model

type (
	Personal struct {
		Name string
		DOB  string
	}

	Account struct {
		AccountID   string
		AccountNo   string
		AccountName string
		Data        []Personal
	}

	Param struct {
		UserId    string
		CardNo    string
		AccountNo string
	}
)
