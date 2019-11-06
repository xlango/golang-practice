package config

const (
	QPS = 20

	EsUrl = "http://192.168.10.222:9200"

	EsIndex = "crawler_dist"

	ItemSaverRpc = "ItemSaverService.Save"
	WorkerRpc    = "CrawlService.Process"

	ParserKeshi         = "ParserKeshiList"
	ParserDoctor        = "ParserDoctor"
	ParserDoctorProfile = "DoctorParser"
	NilParser           = "NilParser"
)
