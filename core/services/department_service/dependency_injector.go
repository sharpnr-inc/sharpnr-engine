/*
Package departmentservice provides the business logic implementation
for handling department-related operations.

This package follows the dependency injection pattern by depending
on the DepartmentRepository interface instead of concrete implementations.
This design improves modularity, testability, and maintainability.
*/
package departmentservice

// // departmentService is the concrete implementation of the DepartmentService
// // interface. It delegates persistence-related operations to the
// // DepartmentRepository, ensuring a clean separation between business
// // logic and data access layers.
// type departmentService struct {
// 	repo interfaces.DepartmentRepository
// }

// // DepartmentService creates and returns a new instance of DepartmentService.
// //
// // This function acts as a dependency injector, wiring the DepartmentRepository
// // into the departmentService implementation. By injecting dependencies
// // from outside, this design makes it easy to swap repository
// // implementations (e.g., in-memory, SQL, mock) during testing or runtime.
// //
// // Example:
// //
// //	repo := repository.NewDepartmentRepository(db)
// //	service := departmentservice.DepartmentService(repo)
// //
// // Now `service` can be used wherever interfaces.DepartmentService is required.
// func DepartmentService(repo interfaces.DepartmentRepository) interfaces.DepartmentService {
// 	return &departmentService{repo: repo}
// }
