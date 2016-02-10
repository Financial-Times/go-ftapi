package ftapi

import "regexp"

type ontology struct {
	Organisation   string
	PublicCompany  string
	Person         string
	Brand          string
	Article        string
	Mentions       string
	IsClassifiedBy string
}

var Ontology = ontology{
	Organisation:   "http://www.ft.com/ontology/organisation/Organisation",
	PublicCompany:  "http://www.ft.com/ontology/company/PublicCompany",
	Person:         "http://www.ft.com/ontology/person/Person",
	Brand:          "http://www.ft.com/ontology/product/Brand",
	Article:        "http://www.ft.com/ontology/content/Article",
	Mentions:       "http://www.ft.com/ontology/annotation/mentions",
	IsClassifiedBy: "http://www.ft.com/ontology/annotation/isClassifiedBy",
}

var UUIDRegexp = regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
var FinalUUIDRegexp = regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")

func UUID(url string) string {
	return UUIDRegexp.FindString(url)
}

func FinalUUID(url string) string {
	return FinalUUIDRegexp.FindString(url)
}
