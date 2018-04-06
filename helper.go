package ftapi

import "regexp"

type ontology struct {
	Article        string
	Brand          string
	Classification	string
	Company        string
	Concept		string
	Content		string
	ContentPackage		string
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
	ImplicitlyAbout		string
	MajorMentions   string
	Mentions       string
	HasAuthor	string
	IsClassifiedBy string
	ImplicitlyClassifiedBy string
	IsPrimarilyClassifiedBy string
	CreatedBy string
}


var Ontology = ontology{
	Article:        "http://www.ft.com/ontology/content/Article",
	Brand:          "http://www.ft.com/ontology/product/Brand",
	Classification: "http://www.ft.com/ontology/classification/Classification",
	Company:        "http://www.ft.com/ontology/company/Company",
	Concept:        "http://www.ft.com/ontology/concept/Concept",
	Content:        "http://www.ft.com/ontology/content/Content",
	ContentPackage:        "http://www.ft.com/ontology/content/ContentPackage",
	Genre:          "http://www.ft.com/ontology/Genre",
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
	ImplicitlyAbout:          "http://www.ft.com/ontology/annotation/implicitlyAbout",
	MajorMentions:       "http://www.ft.com/ontology/annotation/majorMentions",
	Mentions:       "http://www.ft.com/ontology/annotation/mentions",
	HasAuthor:       "http://www.ft.com/ontology/annotation/hasAuthor",
	IsClassifiedBy: "http://www.ft.com/ontology/classification/isClassifiedBy",
	ImplicitlyClassifiedBy: "http://www.ft.com/ontology/implicitlyClassifiedBy",
	IsPrimarilyClassifiedBy: "http://www.ft.com/ontology/classification/isPrimarilyClassifiedBy",
	CreatedBy:      "http://www.ft.com/ontology/annotation/createdBy",
}

// At some point, we intend to drop the 'intermediate' ontologies in these URIs.

var NewOntology = ontology{
	Article:        "http://www.ft.com/ontology/Article",
	Brand:          "http://www.ft.com/ontology/Brand",
	Classification: "http://www.ft.com/ontology/Classification",
	Company:        "http://www.ft.com/ontology/Company",
	Concept:        "http://www.ft.com/ontology/Concept",
	Content:        "http://www.ft.com/ontology/Content",
	ContentPackage:        "http://www.ft.com/ontology/ContentPackage",
	Genre:          "http://www.ft.com/ontology/Genre",
	Location:	"http://www.ft.com/ontology/Location",
	Organisation:   "http://www.ft.com/ontology/Organisation",
	Person:         "http://www.ft.com/ontology/Person",
	PublicCompany:  "http://www.ft.com/ontology/PublicCompany",
	Section:        "http://www.ft.com/ontology/Section",
	SpecialReport:	"http://www.ft.com/ontology/SpecialReport",
	Subject:        "http://www.ft.com/ontology/Subject",
	Thing:          "http://www.ft.com/ontology/Thing",
	Topic:          "http://www.ft.com/ontology/Topic",

	About:          "http://www.ft.com/ontology/about",
	ImplicitlyAbout:          "http://www.ft.com/ontology/implicitlyAbout",
	MajorMentions:       "http://www.ft.com/ontology/majorMentions",
	Mentions:       "http://www.ft.com/ontology/mentions",
	HasAuthor:       "http://www.ft.com/ontology/hasAuthor",
	IsClassifiedBy: "http://www.ft.com/ontology/isClassifiedBy",
	ImplicitlyClassifiedBy: "http://www.ft.com/ontology/implicitlyClassifiedBy",
	IsPrimarilyClassifiedBy: "http://www.ft.com/ontology/isPrimarilyClassifiedBy",
	CreatedBy:      "http://www.ft.com/ontology/createdBy",
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
    case o.Brand:
        return []string{o.Thing, o.Concept, o.Classification, t}
    case o.Genre:
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
