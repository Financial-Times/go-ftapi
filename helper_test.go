package ftapi_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
    ftapi "github.com/Financial-Times/go-ftapi"
)

func TestFinalUuid(t *testing.T) {
    a := assert.New(t)

    result := ftapi.FinalUuid("deadbeef-a7da-11e5-9700-2b669a5aeb83/ccca5db0-a7da-11e5-9700-2b669a5aeb83")

    a.Equal("ccca5db0-a7da-11e5-9700-2b669a5aeb83", result)
}

