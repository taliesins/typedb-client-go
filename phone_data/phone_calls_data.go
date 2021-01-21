package phone_data

import (
	"io/ioutil"
)

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