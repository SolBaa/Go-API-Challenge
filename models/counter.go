package models

type CounterEndpoints struct {
	GetUsers         int
	GetUserByID      int
	AddCompanyToUser int
	DeleteUser       int
	EndCounter       int
}

func (m *CounterEndpoints) EndpointCounter(endpoint string) *CounterEndpoints {

	switch endpoint {
	case "getUsers":
		m.GetUsers += 1
	case "getUsersID":
		m.GetUserByID += 1
	case "addCompany":
		m.AddCompanyToUser += 1
	case "deleteUser":
		m.DeleteUser += 1
	case "endCounter":
		m.EndCounter += 1

	}

	return m
}
