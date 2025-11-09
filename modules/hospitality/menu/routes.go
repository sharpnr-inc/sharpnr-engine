package hospitalitymenu

import (
	menucontroller "sharpnr-engine/modules/hospitality/menu/menu_controller"
	menurepo "sharpnr-engine/modules/hospitality/menu/menu_repo"
	menuservice "sharpnr-engine/modules/hospitality/menu/menu_service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MenuRoutes(v1 *gin.RouterGroup, db *gorm.DB) {

	newMeneRepo := menurepo.NewHospitalityMenuRepository(db)
	newMenuService := menuservice.NewHospitalityMenuService(newMeneRepo)
	newMenuController := menucontroller.NewMenuController(newMenuService)

	menuRoutes := v1.Group("/hospitality-menu")

	// -------------------------------
	// 🌐 MENU GROUPS (Bar, Hotel, Restaurant)
	// -------------------------------
	menuRoutes.POST("/groups", newMenuController.CreateMenuGroup)
	menuRoutes.GET("/groups", newMenuController.GetAllMenuGroup) // Get All Menu Groups
	menuRoutes.GET("/groups/:id", func(c *gin.Context) {})       // Get One Group
	menuRoutes.PUT("/groups/:id", func(c *gin.Context) {})       // Update Group
	menuRoutes.DELETE("/groups/:id", func(c *gin.Context) {})    // Delete Group

	// -------------------------------
	// 🗂️ MENU CATEGORIES
	// -------------------------------
	menuRoutes.POST("/categories", newMenuController.CreateCategoryInMenu)           // Create Category
	menuRoutes.GET("/categories", newMenuController.GetAllCategoryInMenu)            // Get All Categories
	menuRoutes.GET("/categories/:id", newMenuController.GetCategoryInMenuByID)       // Get Category by ID
	menuRoutes.GET("/groups/:id/categories", newMenuController.GetCategoriesByGroup) // Get Categories by Group
	menuRoutes.PUT("/categories/:id", newMenuController.UpdateCategoryInMenu)        // Update Category
	menuRoutes.DELETE("/categories/:id", newMenuController.DeleteCategoryInMenu)     // Delete Category

	// -------------------------------
	// 📂 MENU SUBCATEGORIES
	// -------------------------------
	menuRoutes.POST("/subcategories", newMenuController.CreateSubCategoryInCategory) // Create SubCategory
	// menuRoutes.GET("/subcategories", func(c *gin.Context) {})                        // Get All SubCategories
	menuRoutes.GET("/subcategories/:id", newMenuController.GetSubCategoryById) // Get SubCategory by ID
	menuRoutes.GET("/categories/:id/subcategories", func(c *gin.Context) {})   // Get SubCategories by Category
	menuRoutes.PUT("/subcategories/:id", func(c *gin.Context) {})              // Update SubCategory
	menuRoutes.DELETE("/subcategories/:id", func(c *gin.Context) {})           // Delete SubCategory

	// -------------------------------
	// 🍽️ MENU ITEMS
	// -------------------------------
	menuRoutes.POST("/items", func(c *gin.Context) {})                  // Add Menu Item
	menuRoutes.GET("/items", func(c *gin.Context) {})                   // Get All Items
	menuRoutes.GET("/items/:id", func(c *gin.Context) {})               // Get Item by ID
	menuRoutes.GET("/subcategories/:id/items", func(c *gin.Context) {}) // Get Items by SubCategory
	menuRoutes.PUT("/items/:id", func(c *gin.Context) {})               // Update Item
	menuRoutes.DELETE("/items/:id", func(c *gin.Context) {})            // Delete Item
	menuRoutes.POST("/items/:id/images", func(c *gin.Context) {})       // Add Image Links to Item

	// -------------------------------
	// 🧩 VARIANTS (Half, Full, etc.)
	// -------------------------------
	menuRoutes.POST("/items/:id/variants", func(c *gin.Context) {}) // Add Variant
	menuRoutes.GET("/items/:id/variants", func(c *gin.Context) {})  // Get Variants for Item
	menuRoutes.PUT("/variants/:id", func(c *gin.Context) {})        // Update Variant
	menuRoutes.DELETE("/variants/:id", func(c *gin.Context) {})     // Delete Variant

	// -------------------------------
	// ➕ ADDON GROUPS (e.g., Toppings, Drinks)
	// -------------------------------
	menuRoutes.POST("/items/:id/addon-groups", func(c *gin.Context) {}) // Create Addon Group
	menuRoutes.GET("/items/:id/addon-groups", func(c *gin.Context) {})  // Get Addon Groups for Item
	menuRoutes.PUT("/addon-groups/:id", func(c *gin.Context) {})        // Update Addon Group
	menuRoutes.DELETE("/addon-groups/:id", func(c *gin.Context) {})     // Delete Addon Group

	// -------------------------------
	// 🧀 ADDONS (e.g., "Extra Cheese", "Coke")
	// -------------------------------
	menuRoutes.POST("/addon-groups/:id/addons", func(c *gin.Context) {}) // Create Addon
	menuRoutes.GET("/addon-groups/:id/addons", func(c *gin.Context) {})  // Get Addons by Group
	menuRoutes.PUT("/addons/:id", func(c *gin.Context) {})               // Update Addon
	menuRoutes.DELETE("/addons/:id", func(c *gin.Context) {})            // Delete Addon
}
