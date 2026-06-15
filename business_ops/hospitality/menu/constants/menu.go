package menu_constants

type MenuItemDietaryType string

const (
	Vegetarian    MenuItemDietaryType = "VEGETARIAN"
	NonVegetarian MenuItemDietaryType = "NON_VEGETARIAN"
	Vegan         MenuItemDietaryType = "VEGAN"
	Egg           MenuItemDietaryType = "EGG"
	Jain          MenuItemDietaryType = "JAIN"
	GlutenFree    MenuItemDietaryType = "GLUTEN_FREE"
)
