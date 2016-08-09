package ftapi

import "regexp"

type ontology struct {
	Article        string
	Brand          string
	Company        string
	Genre          string
	Location	string
	Organisation   string
	Person         string
	PublicCompany  string
	Section		string
	Subject		string
	Thing          string
	Topic		string

	Mentions       string
	IsClassifiedBy string
}

var Ontology = ontology{
	Article:        "http://www.ft.com/ontology/content/Article",
	Brand:          "http://www.ft.com/ontology/product/Brand",
	Company:        "http://www.ft.com/ontology/company/Company",
	Genre:          "http://www.ft.com/ontology/content/Genre",
	Location:	"http://www.ft.com/ontology/location/Location",
	Organisation:   "http://www.ft.com/ontology/organisation/Organisation",
	Person:         "http://www.ft.com/ontology/person/Person",
	PublicCompany:  "http://www.ft.com/ontology/company/PublicCompany",
	Section:        "http://www.ft.com/ontology/classification/Section",
	Subject:        "http://www.ft.com/ontology/classification/Subject",
	Thing:          "http://www.ft.com/ontology/thing/Thing",
	Topic:          "http://www.ft.com/ontology/classification/Topic",

	Mentions:       "http://www.ft.com/ontology/annotation/mentions",
	IsClassifiedBy: "http://www.ft.com/ontology/annotation/isClassifiedBy",
}

func (o *ontology) AllTypes(t string) []string {
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
