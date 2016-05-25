package ftapi

import "regexp"

type ontology struct {
	Thing          string
	Organisation   string
	Company        string
	PublicCompany  string
	Person         string
	Brand          string
	Article        string
	Mentions       string
	IsClassifiedBy string
}

var Ontology = ontology{
	Thing:          "http://www.ft.com/ontology/thing/Thing",
	Organisation:   "http://www.ft.com/ontology/organisation/Organisation",
	Company:        "http://www.ft.com/ontology/company/Company",
	PublicCompany:  "http://www.ft.com/ontology/company/PublicCompany",
	Person:         "http://www.ft.com/ontology/person/Person",
	Brand:          "http://www.ft.com/ontology/product/Brand",
	Article:        "http://www.ft.com/ontology/content/Article",
	Mentions:       "http://www.ft.com/ontology/annotation/mentions",
	IsClassifiedBy: "http://www.ft.com/ontology/annotation/isClassifiedBy",
}

func (o *ontology) GetTypes(t string) []string {
    switch t {
    case o.PublicCompany:
        return []string{o.Thing, o.Organisation, o.Company, o.PublicCompany}
    case o.Company:
        return []string{o.Thing, o.Organisation, o.Company}
    default:
        return []string{o.Thing, t}
    }
}

var FinalComponentRegexp = regexp.MustCompile("[A-Za-z]+$")
var UUIDRegexp = regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
var FinalUUIDRegexp = regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")

func UUID(url string) string {
	return UUIDRegexp.FindString(url)
}

func FinalUUID(url string) string {
	return FinalUUIDRegexp.FindString(url)
}

func FinalComponent(url string) string {
    return FinalComponentRegexp.FindString(url)
}
