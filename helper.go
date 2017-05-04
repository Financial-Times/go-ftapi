package ftapi

import "regexp"

type ontology struct {
	Article        string
	Brand          string
	Classification	string
	Company        string
	Concept		string
	Genre          string
	Location	string
	Organisation   string
	Person         string
	PublicCompany  string
	Section		string
	SpecialReport	string
	Subject		string
	Thing          string
	Topic		string

	About		string
	MajorMentions   string
	Mentions       string
	HasAuthor	string
	IsClassifiedBy string
	IsPrimarilyClassifiedBy string
	CreatedBy string
}

var Ontology = ontology{
	Article:        "http://www.ft.com/ontology/content/Article",
	Brand:          "http://www.ft.com/ontology/product/Brand",
	Classification: "http://www.ft.com/ontology/classification/Classification",
	Company:        "http://www.ft.com/ontology/company/Company",
	Concept:        "http://www.ft.com/ontology/concept/Concept",
	Genre:          "http://www.ft.com/ontology/content/Genre",
	Location:	"http://www.ft.com/ontology/location/Location",
	Organisation:   "http://www.ft.com/ontology/organisation/Organisation",
	Person:         "http://www.ft.com/ontology/person/Person",
	PublicCompany:  "http://www.ft.com/ontology/company/PublicCompany",
	Section:        "http://www.ft.com/ontology/classification/Section",
	SpecialReport:	"http://www.ft.com/ontology/classification/SpecialReport",
	Subject:        "http://www.ft.com/ontology/classification/Subject",
	Thing:          "http://www.ft.com/ontology/core/Thing",
	Topic:          "http://www.ft.com/ontology/topic/Topic",

	About:          "http://www.ft.com/ontology/annotation/about",
	MajorMentions:       "http://www.ft.com/ontology/annotation/majorMentions",
	Mentions:       "http://www.ft.com/ontology/annotation/mentions",
	HasAuthor:       "http://www.ft.com/ontology/annotation/hasAuthor",
	IsClassifiedBy: "http://www.ft.com/ontology/annotation/isClassifiedBy",
	IsPrimarilyClassifiedBy: "http://www.ft.com/ontology/annotation/isPrimarilyClassifiedBy",
	CreatedBy:      "http://www.ft.com/ontology/annotation/createdBy",
}

func (o *ontology) AllTypes(t string) []string {
    switch t {
    case o.PublicCompany:
        return []string{o.Thing, o.Concept, o.Organisation, o.Company, t}
    case o.Company:
        return []string{o.Thing, o.Concept, o.Organisation, t}
    case o.Mentions:
        return []string{o.Thing, t}
    case o.IsClassifiedBy:
        return []string{o.Thing, t}
    case o.Section:
        return []string{o.Thing, o.Concept, o.Classification, t}
    case o.SpecialReport:
        return []string{o.Thing, o.Concept, o.Classification, t}
    case o.Subject:
        return []string{o.Thing, o.Concept, o.Classification, t}
    default:
        return []string{o.Thing, o.Concept, t}
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
