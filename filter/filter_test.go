package filter

import (
	"encoding/json"
	"testing"
)

func TestFilter(t *testing.T) {
	filter := Status("Status").Equal("waiting").
		And(MultiSelect("Tag").Contain("test").
			Or(Checkbox("Done").Equal(true))).
		And(RichText("Name").Contain("test"))

	data, _ := json.MarshalIndent(filter, "", "  ")
	t.Log(string(data))

	properties := Status("Status").Set("name", "publish")
	data2, _ := json.MarshalIndent(properties, "", "  ")
	t.Log(string(data2))
}
