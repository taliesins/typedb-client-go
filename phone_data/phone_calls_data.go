package phone_data

import (
	"bytes"
	"io/ioutil"
)

func GetCallsGql() (gql []string, err error){
	gql = []string{}
	calls, err := GetCalls()
	if err != nil {
		return nil, err
	}

	for _, call := range calls {
		templateRendered := bytes.Buffer{}
		err = GetCallGql(&call, &templateRendered)
		if err != nil {
			return nil, err
		}
		gql = append(gql, templateRendered.String())
	}

	return gql, err
}

func GetCompaniesGql() (gql []string, err error){
	gql = []string{}
	companies, err := GetCompanies()
	if err != nil {
		return nil, err
	}

	for _, company := range companies {
		templateRendered := bytes.Buffer{}
		err = GetCompanyGql(&company, &templateRendered)
		if err != nil {
			return nil, err
		}
		gql = append(gql, templateRendered.String())
	}

	return gql, err
}

func GetContractsGql() (gql []string, err error){
	gql = []string{}
	contracts, err := GetContracts()
	if err != nil {
		return nil, err
	}

	for _, contract := range contracts {
		templateRendered := bytes.Buffer{}
		err = GetContractGql(&contract, &templateRendered)
		if err != nil {
			return nil, err
		}
		gql = append(gql, templateRendered.String())
	}

	return gql, err
}

func GetPeopleGql() (gql []string, err error){
	gql = []string{}
	people, err := GetPeople()
	if err != nil {
		return nil, err
	}

	for _, person := range people {
		templateRendered := bytes.Buffer{}
		err = GetPersonGql(&person, &templateRendered)
		if err != nil {
			return nil, err
		}
		gql = append(gql, templateRendered.String())
	}

	return gql, err
}

func GetPhoneCallsDataGql() (gql []string, err error){
	gql = []string{}

	companiesGql, err := GetCompaniesGql()
	if err != nil {
		return nil, err
	}
	gql = append(gql, companiesGql...)

	peopleGql, err := GetPeopleGql()
	if err != nil {
		return nil, err
	}
	gql = append(gql, peopleGql...)

	contractsGql, err := GetContractsGql()
	if err != nil {
		return nil, err
	}
	gql = append(gql, contractsGql...)

	callsGql, err := GetCallsGql()
	if err != nil {
		return nil, err
	}
	gql = append(gql, callsGql...)

	return gql, err
}

func GetPhoneCallsSchemaV1Gql() (gsql string, err error) {
	file, err := ioutil.ReadFile("phone_data/phone-calls-schema-v1.gql")
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func GetPhoneCallsSchemaV2Gql() (gsql string, err error) {
	file, err := ioutil.ReadFile("phone_data/phone-calls-schema-v2.gql")
	if err != nil {
		return "", err
	}

	return string(file), nil
}