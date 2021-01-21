package phone_data

import (
	"bytes"
	"io/ioutil"
)

func GetPhoneCallsDataGsql() (gsql string, err error) {
	var templateRendered bytes.Buffer

	err = GetCompaniesGsql(&templateRendered)
	if err != nil {
		return "", err
	}

	err = GetPeopleGsql(&templateRendered)
	if err != nil {
		return "", err
	}

	err = GetContractsGsql(&templateRendered)
	if err != nil {
		return "", err
	}

	err = GetCallsGsql(&templateRendered)
	if err != nil {
		return "", err
	}

	gsql = templateRendered.String()

	gsql = `insert $company isa company, has name "Telecom";`
	return gsql, nil
}

func GetPhoneCallsSchemaV1Gsql() (gsql string, err error) {
	file, err := ioutil.ReadFile("phone_data/phone-calls-schema-v1.gql")
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func GetPhoneCallsSchemaV2Gsql() (gsql string, err error) {
	file, err := ioutil.ReadFile("phone_data/phone-calls-schema-v2.gql")
	if err != nil {
		return "", err
	}

	return string(file), nil
}