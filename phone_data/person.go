package phone_data

import (
	"bytes"
	"encoding/json"
	"io"
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
	$_ isa person, 
{{if .FirstName }}		has first-name "{{.FirstName}}",{{end}}
{{if .LastName }}		has last-name "{{.LastName}}",{{end}}
{{if .City }}		has city "{{.City}}",{{end}}
{{if .Age }}		has age {{.Age}},{{end}}
		has phone-number "{{.PhoneNumber}}";
`))

var insertPeopleTemplate = template.Must(template.New("InsertPeople").Parse(`
insert
{{range .}}
	$_ isa person,{{if .FirstName }} has first-name "{{.FirstName}}",{{end}}{{if .LastName }} has last-name "{{.LastName}}",{{end}}{{if .City }} has city "{{.City}}",{{end}}{{if .Age }} has age {{.Age}},{{end}} has phone-number "{{.PhoneNumber}}";
{{end}}
`))

func GetPeople() (people []Person, err error) {
	file, err := ioutil.ReadFile("phone_data/people.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(file), &people)
	if err != nil {
		return nil, err
	}

	return people, nil
}

func GetPersonGql(person *Person, wr io.Writer) (err error) {
	err = insertPersonTemplate.Execute(wr, person)
	if err != nil {
		return err
	}
	return err
}

func GetPeopleGql() (gql []string, err error){
	gql = []string{}
	people, err := GetPeople()
	if err != nil {
		return nil, err
	}

	templateRendered := bytes.Buffer{}
	err = insertPeopleTemplate.Execute(&templateRendered, people)
	if err != nil {
		return nil, err
	}
	gql = append(gql, templateRendered.String())
	return gql, err
}