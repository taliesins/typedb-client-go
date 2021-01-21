package phone_data

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"text/template"
)

type Call struct {
	CallerId string `json:"caller_id"`
	CalleeId string `json:"callee_id"`
	StartedAt string `json:"started_at"`
	Duration int `json:"duration"`
}

var insertCallTemplate = template.Must(template.New("InsertCall").Parse(`
match 
	$caller isa person, has phone-number "{{.CallerId}}";
	$callee isa person, has phone-number "{{.CalleeId}}";
insert 
	$call(caller: $caller, callee: $callee) isa call; 
	$call has started-at {{.StartedAt}}; 
	$call has duration {{.Duration}};
`))

func GetCallsGsql(templateRendered *bytes.Buffer) (err error) {
	file, err := ioutil.ReadFile("phone_data/calls.json")
	if err != nil {
		return err
	}

	var calls []Call
	err = json.Unmarshal([]byte(file), &calls)
	if err != nil {
		return err
	}

	 for _, call := range calls {
		 err = insertCallTemplate.Execute(templateRendered, call)

		 if err != nil {
			 return err
		 }
	}

	return nil
}