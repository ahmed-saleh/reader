package main

import (
    "testing"

)

func TestParsingRequest(t *testing.T) {

	var a Body

	a.Name = "modules-final/config-final-message"
	a.Description = "config-final-message ran successfully"
	a.Event_type = "finish"
	a.Origin = "cloudinit"
	a.Result = "SUCCESS"
	//a.Timestamp = time.Now().Unix()
	var b = Parse([]byte("{\"name\": \"modules-final/config-final-message\", \"description\": \"config-final-message ran successfully\", \"event_type\": \"finish\", \"origin\": \"cloudinit\", \"result\": \"SUCCESS\"}"))
	if a != b {
		t.Fatalf("expected: %+v\n received %+v\n", a, b)
	}
}
