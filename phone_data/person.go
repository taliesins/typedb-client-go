package phone_data

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"text/template"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	City string `json:"city"`
	Age int `json:"age"`
}

var insertPersonTemplate = template.Must(template.New("InsertPerson").Parse(`
insert 
	$person isa person, 
{{if .FirstName }}		has first-name "{{.FirstName}}",{{end}}
{{if .LastName }}		has last-name "{{.LastName}}",{{end}}
{{if .City }}		has city "{{.City}}",{{end}}
{{if .Age }}		has age {{.Age}},{{end}}
		has phone-number "{{.PhoneNumber}}";
`))

func GetPeopleGsql(templateRendered *bytes.Buffer) (err error) {
	file, err := ioutil.ReadFile("phone_data/people.json")
	if err != nil {
		return err
	}

	var people []Person
	err = json.Unmarshal([]byte(file), &people)
	if err != nil {
		return err
	}

	for _, person := range people {
		err = insertPersonTemplate.Execute(templateRendered, person)

		if err != nil {
			return err
		}
	}

	return nil
}
