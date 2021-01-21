package phone_data

import (
	"encoding/json"
	"io"
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

func GetCompanies() (companies []Company, err error) {
	file, err := ioutil.ReadFile("phone_data/companies.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(file), &companies)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func GetCompanyGql(company *Company, wr io.Writer) (err error) {
	err = insertCompanyTemplate.Execute(wr, company)
	if err != nil {
		return err
	}
	return err
}
