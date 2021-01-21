package phone_data

import (
	"encoding/json"
	"io"
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

func GetCalls() (companies []Call, err error) {
	file, err := ioutil.ReadFile("phone_data/calls.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(file), &companies)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func GetCallGql(call *Call, wr io.Writer) (err error) {
	err = insertCallTemplate.Execute(wr, call)
	if err != nil {
		return err
	}
	return err
}
