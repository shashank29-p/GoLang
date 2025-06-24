//interfaces declaration
type Company interface {
	department() string
	employeeDetails(empId int) string
}

type Employee struct{
	EmpId   int
	EmpName string
	Department string
}

Employee := Employee{
	EmpId: 1,
	EmpName: "John Doe",
	Department: "Engineering",
}

func (e Employee) department() string {
	return e.Department
}

func (e Employee) employeeDetails(empId int) string {
	if e.EmpId == empId {
		return "Employee ID: " + string(e.EmpId) + ", Name: " + e.EmpName + ", Department: " + e.Department
	}
	return "Employee not found"
}