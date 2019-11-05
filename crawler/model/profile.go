package model

type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string
	Education  string
	Occupation string
	Hokou      string
	Xinzuo     string
	House      string
	Car        string
}

type Score struct {
	School     string
	ScoreLines []ScoreLine
}

type ScoreLine struct {
	Year     string
	Top      string
	Lower    string
	Averrage string
	Archive  string
	Person   string
	Pici     string
}

type Property struct {
	Name           string
	Type           string
	Address        string
	HouseType      string
	OpenTime       string
	UnitPrice      string
	TotalPriceUp   string
	TotalPriceDown string
}

type Doctor struct {
	Name       string
	Zhicheng   string
	Hospital   string
	Department string
	Disease    string
	WebSite    string
	Tel        string
	Post       string
	Email      string
	Fax        string
	Address    string
}
