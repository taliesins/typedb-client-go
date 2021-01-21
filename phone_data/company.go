package phone_data

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"text/template"
)

type Company struct {
	Name string `json:"name"`
}

var insertCompanyTemplate = template.Must(template.New("InsertCompany").Parse(`
insert 
	$company isa company, 
		has name "{{.Name}}";
`))

func GetCompaniesGsql(templateRendered *bytes.Buffer) (err error) {
	file, err := ioutil.ReadFile("phone_data/companies.json")
	if err != nil {
		return err
	}

	var companies []Company
	err = json.Unmarshal([]byte(file), &companies)
	if err != nil {
		return err
	}

	for _, company := range companies {
		err = insertCompanyTemplate.Execute(templateRendered, company)

		if err != nil {
			return err
		}
	}

	return nil
}