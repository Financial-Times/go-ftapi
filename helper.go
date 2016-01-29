package ftapi

import "regexp"

var UuidRegexp = regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
var FinalUuidRegexp = regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")

func GetUuid(url string) string {
    return UuidRegexp.FindString(url)
}

func GetFinalUuid(url string) string {
    return FinalUuidRegexp.FindString(url)
}

