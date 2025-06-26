package filter

import (
	"encoding/json"
	"testing"
	"time"
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

	updateData := Date("Publish Time").Set("start", time.Now().Format("2006-01-02")).And(Status("Status").Set("name", "publish"))
	data3, _ := json.MarshalIndent(updateData, "", "  ")
	t.Log(string(data3))
}
