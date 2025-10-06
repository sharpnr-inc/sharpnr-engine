package mappers

// ToDepartmentModel maps an AddDepartmentRequestDTO to a Department model.
//
// This function takes the data from the request DTO and returns a pointer
// to a new Department model ready to be saved in the database.
//
// Parameters:
//   - req: AddDepartmentRequestDTO containing the department details.
//
// Returns:
//   - *models.Department: pointer to the Department model populated with
//     the data from the request DTO.
// func ToDepartmentModel(req dto.AddDepartmentRequestDTO) *models.Department {
// 	return &models.Department{
// 		Name:        req.Name,
// 		Description: req.Description,
// 	}
// }

// func ToAdd_Dep_Res(org *models.Department) error {
// 	return nil
// }
