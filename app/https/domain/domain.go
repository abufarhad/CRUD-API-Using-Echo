package domain

type Employee struct {
	Id     string `json:	"id"`
	Name   string `json:	"employee_name"`
	Age    string `json: 	"employee_age"`
	Salary string `json: 	"employee_salary"`
}

type Employees struct {
	Employees []Employee `json:"employee"`
}