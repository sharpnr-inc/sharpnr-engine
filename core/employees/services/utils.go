package coreEmployeeServices

func buildEmployeeName(firstName, lastName, nickName string) string {
	if nickName != "" {
		return nickName
	}

	if lastName != "" {
		return firstName + " " + lastName
	}

	return firstName
}

/* ---------------------------------- <***> --------------------------------- */
