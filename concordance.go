package ftapi

type Identifier struct {
    Authority string `json:"authority"`
    IdentifierValue string `json:"identifierValue"`
}

type Concordance struct {
    Concept *Thing `json:"concept"`
    Identifier *Identifier `json:"identifier"`
}

// TODO api calls

