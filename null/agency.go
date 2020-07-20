package main

func main() {
	colleage1 := createColleage1()
	colleage2 := createColleage3()
	totalProfessors := 0
	departmentArray := []string{"computerscience", "mechanical", "civil", "electronics"}

	for _, departmentName := range departmentArray {
		d := colleage1.getDepartment(departmentName)
	}

}
