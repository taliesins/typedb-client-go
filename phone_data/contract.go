package phone_data

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"text/template"
)

type Contract struct {
	CompanyName string `json:"company_name"`
	PersonId string `json:"person_id"`
}

var insertContractTemplate = template.Must(template.New("InsertContract").Parse(`
match 
	$company isa company, 
		has name "{{.CompanyName}}"; 
	$customer isa person, 
		has phone-number "{{.PersonId}}";
insert 
	(provider: $company, customer: $customer) isa contract;
`))


func GetContractsGsql(templateRendered *bytes.Buffer) (err error) {
	file, err := ioutil.ReadFile("phone_data/contracts.json")
	if err != nil {
		return err
	}

	var contracts []Contract
	err = json.Unmarshal([]byte(file), &contracts)
	if err != nil {
		return err
	}

	for _, contract := range contracts {
		err = insertContractTemplate.Execute(templateRendered, contract)

		if err != nil {
			return err
		}
	}

	return nil
}