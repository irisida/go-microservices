package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	req := CreateRepoRequest{
		Name:        "createRepoRequestAsJson",
		Description: "A test repository description",
		Homepage:    "www.testreporeq.com",
		Private:     false,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     true,
	}

	// Marshal attempts to create a valid json string
	// from an input interface.
	bytes, err := json.Marshal(req)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	assert.EqualValues(t, false, req.Private)
	assert.EqualValues(t, false, req.HasIssues)
	assert.EqualValues(t, false, req.HasProjects)
	assert.EqualValues(t, true, req.HasWiki)
}

func TestCreateRepoRequestUnmarshalJson(t *testing.T) {
	req := CreateRepoRequest{
		Name:        "createRepoRequestUnmarshalJson",
		Description: "A test repository description",
		Homepage:    "www.testreporeq.com",
		Private:     false,
		HasIssues:   false,
		HasProjects: false,
		HasWiki:     true,
	}

	// Marshal the test request
	bytes, err := json.Marshal(req)
	assert.Nil(t, err)

	// Unmarshal test
	var target CreateRepoRequest

	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)
	assert.EqualValues(t, target.Name, req.Name)
	assert.EqualValues(t, target.Description, req.Description)
	assert.EqualValues(t, target.Homepage, req.Homepage)
	assert.EqualValues(t, target.Private, req.Private)
	assert.EqualValues(t, target.HasIssues, req.HasIssues)
	assert.EqualValues(t, target.HasProjects, req.HasProjects)
	assert.EqualValues(t, target.HasWiki, req.HasWiki)

}
