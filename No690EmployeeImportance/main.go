package No690EmployeeImportance

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {

	length := len(employees)
	if length == 0 {
		return 0
	}

	cur := getEmployee(employees, id)
	if cur == nil {
		return 0
	}
	if len(cur.Subordinates) == 0 {
		return cur.Importance
	}
	result := cur.Importance
	for _, v := range cur.Subordinates {
		result += getImportance(employees, v)
	}
	return result

}

func getEmployee(employees []*Employee, id int) *Employee {
	for _, e := range employees {
		if id == e.Id {
			return e
		}
	}
	return nil
}
