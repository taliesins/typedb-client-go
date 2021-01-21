package phone_data

import (
	"encoding/json"
	"io"
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

func GetContracts() (companies []Contract, err error) {
	file, err := ioutil.ReadFile("phone_data/contracts.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(file), &companies)
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func GetContractGql(contract *Contract, wr io.Writer) (err error) {
	err = insertContractTemplate.Execute(wr, contract)
	if err != nil {
		return err
	}
	return err
}
