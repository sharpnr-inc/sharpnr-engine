package menu_helpers

import (
	menu_constants "sharpnr-engine/property_business_ops/hospitality/menu/constants"
	hospitality_menu_contracts "sharpnr-engine/property_business_ops/hospitality/menu/contracts"
)

var DefaultMenuCategories = []hospitality_menu_contracts.AddMenuCategoryInput{

	// -------------------------------------------------
	// STARTERS
	// -------------------------------------------------

	{
		CategoryName:        "Starters",
		CategoryDescription: "Light appetizers to begin your meal",
		SubCategories: &[]hospitality_menu_contracts.AddSubCategoryInput{

			{
				SubCatName:        "Veg Starters",
				SubCatDescription: "Vegetarian appetizers",
				MenuItem: &[]hospitality_menu_contracts.AddMenuItemInput{

					{
						Name:        "Paneer Tikka",
						Description: "Grilled cottage cheese marinated with spices",
						BasePrice:   250,
						Unit:        "plate",
						DietaryType: menu_constants.Vegetarian,

						AddOns: &[]hospitality_menu_contracts.AddAddOnInput{
							{Name: "Extra Mint Chutney", Price: 20},
							{Name: "Extra Paneer", Price: 60},
						},
					},

					{
						Name:        "Veg Spring Rolls",
						Description: "Crispy rolls stuffed with vegetables",
						BasePrice:   180,
						Unit:        "plate",
						DietaryType: menu_constants.Vegetarian,
					},

					{
						Name:        "Chilli Paneer",
						Description: "Indo-Chinese spicy paneer cubes",
						BasePrice:   260,
						Unit:        "plate",
						DietaryType: menu_constants.Vegetarian,
					},
				},
			},

			{
				SubCatName:        "Non Veg Starters",
				SubCatDescription: "Chicken and seafood starters",
				MenuItem: &[]hospitality_menu_contracts.AddMenuItemInput{

					{
						Name:        "Chicken Tikka",
						Description: "Charcoal grilled chicken",
						BasePrice:   320,
						Unit:        "plate",
						DietaryType: menu_constants.NonVegetarian,

						VariantGroup: &[]hospitality_menu_contracts.AddVariantGroupInput{
							{
								VariantGroupName: "Portion Size",
								Variant: &[]hospitality_menu_contracts.AddVariantInput{
									{Name: "Half", Price: 0},
									{Name: "Full", Price: 180},
								},
							},
						},
					},

					{
						Name:        "Chicken Wings",
						Description: "Spicy fried chicken wings",
						BasePrice:   300,
						Unit:        "plate",
						DietaryType: menu_constants.NonVegetarian,
					},

					{
						Name:        "Fish Fingers",
						Description: "Crispy fried fish served with dip",
						BasePrice:   340,
						Unit:        "plate",
						DietaryType: menu_constants.NonVegetarian,
					},
				},
			},
		},
	},

	// -------------------------------------------------
	// PIZZA
	// -------------------------------------------------

	{
		CategoryName:        "Pizza",
		CategoryDescription: "Stone baked pizzas",
		SubCategories: &[]hospitality_menu_contracts.AddSubCategoryInput{

			{
				SubCatName:        "Veg Pizza",
				SubCatDescription: "Vegetarian pizzas",
				MenuItem: &[]hospitality_menu_contracts.AddMenuItemInput{

					{
						Name:        "Margherita Pizza",
						Description: "Classic cheese pizza",
						BasePrice:   250,
						Unit:        "pizza",
						DietaryType: menu_constants.Vegetarian,

						VariantGroup: &[]hospitality_menu_contracts.AddVariantGroupInput{
							{
								VariantGroupName: "Size",
								Variant: &[]hospitality_menu_contracts.AddVariantInput{
									{Name: "Small", Price: 0},
									{Name: "Medium", Price: 120},
									{Name: "Large", Price: 220},
								},
							},
						},

						AddOns: &[]hospitality_menu_contracts.AddAddOnInput{
							{Name: "Extra Cheese", Price: 40},
							{Name: "Olives", Price: 30},
							{Name: "Jalapenos", Price: 30},
						},
					},

					{
						Name:        "Farmhouse Pizza",
						Description: "Loaded with vegetables",
						BasePrice:   320,
						Unit:        "pizza",
						DietaryType: menu_constants.Vegetarian,
					},
				},
			},

			{
				SubCatName:        "Non Veg Pizza",
				SubCatDescription: "Chicken pizzas",
				MenuItem: &[]hospitality_menu_contracts.AddMenuItemInput{

					{
						Name:        "Chicken Supreme Pizza",
						Description: "Pizza topped with grilled chicken",
						BasePrice:   380,
						Unit:        "pizza",
						DietaryType: menu_constants.NonVegetarian,
					},

					{
						Name:        "BBQ Chicken Pizza",
						Description: "Smoky BBQ chicken pizza",
						BasePrice:   400,
						Unit:        "pizza",
						DietaryType: menu_constants.NonVegetarian,
					},
				},
			},
		},
	},

	// -------------------------------------------------
	// BURGERS
	// -------------------------------------------------

	{
		CategoryName:        "Burgers",
		CategoryDescription: "Classic burgers and sandwiches",
		SubCategories: &[]hospitality_menu_contracts.AddSubCategoryInput{

			{
				SubCatName:        "Veg Burgers",
				SubCatDescription: "Vegetarian burgers",
				MenuItem: &[]hospitality_menu_contracts.AddMenuItemInput{

					{
						Name:        "Veg Burger",
						Description: "Grilled veggie patty burger",
						BasePrice:   160,
						Unit:        "piece",
						DietaryType: menu_constants.Vegetarian,

						AddOns: &[]hospitality_menu_contracts.AddAddOnInput{
							{Name: "Extra Cheese", Price: 30},
							{Name: "Extra Patty", Price: 50},
						},
					},
				},
			},

			{
				SubCatName:        "Chicken Burgers",
				SubCatDescription: "Chicken burgers",
				MenuItem: &[]hospitality_menu_contracts.AddMenuItemInput{

					{
						Name:        "Chicken Burger",
						Description: "Crispy chicken patty burger",
						BasePrice:   200,
						Unit:        "piece",
						DietaryType: menu_constants.NonVegetarian,
					},
				},
			},
		},
	},

	// -------------------------------------------------
	// MAIN COURSE
	// -------------------------------------------------

	{
		CategoryName:        "Main Course",
		CategoryDescription: "Hearty Indian dishes",
		SubCategories: &[]hospitality_menu_contracts.AddSubCategoryInput{

			{
				SubCatName:        "Veg Curries",
				SubCatDescription: "Vegetarian main dishes",
				MenuItem: &[]hospitality_menu_contracts.AddMenuItemInput{

					{
						Name:        "Paneer Butter Masala",
						Description: "Paneer in rich tomato gravy",
						BasePrice:   320,
						Unit:        "bowl",
						DietaryType: menu_constants.Vegetarian,
					},

					{
						Name:        "Dal Tadka",
						Description: "Yellow lentils tempered with spices",
						BasePrice:   220,
						Unit:        "bowl",
						DietaryType: menu_constants.Vegetarian,
					},
				},
			},

			{
				SubCatName:        "Non Veg Curries",
				SubCatDescription: "Chicken and mutton curries",
				MenuItem: &[]hospitality_menu_contracts.AddMenuItemInput{

					{
						Name:        "Butter Chicken",
						Description: "Creamy tomato chicken curry",
						BasePrice:   380,
						Unit:        "bowl",
						DietaryType: menu_constants.NonVegetarian,
					},

					{
						Name:        "Mutton Rogan Josh",
						Description: "Traditional Kashmiri mutton curry",
						BasePrice:   420,
						Unit:        "bowl",
						DietaryType: menu_constants.NonVegetarian,
					},
				},
			},
		},
	},

	// -------------------------------------------------
	// DESSERTS
	// -------------------------------------------------

	{
		CategoryName:        "Desserts",
		CategoryDescription: "Sweet dishes",
		SubCategories: &[]hospitality_menu_contracts.AddSubCategoryInput{

			{
				SubCatName:        "Indian Desserts",
				SubCatDescription: "Traditional sweets",
				MenuItem: &[]hospitality_menu_contracts.AddMenuItemInput{

					{
						Name:        "Gulab Jamun",
						Description: "Milk dumplings soaked in sugar syrup",
						BasePrice:   120,
						Unit:        "portion",
						DietaryType: menu_constants.Vegetarian,
					},

					{
						Name:        "Rasmalai",
						Description: "Soft cheese patties soaked in sweet milk",
						BasePrice:   140,
						Unit:        "portion",
						DietaryType: menu_constants.Vegetarian,
					},
				},
			},
		},
	},

	// -------------------------------------------------
	// BEVERAGES
	// -------------------------------------------------

	{
		CategoryName:        "Beverages",
		CategoryDescription: "Hot and cold drinks",
		SubCategories: &[]hospitality_menu_contracts.AddSubCategoryInput{

			{
				SubCatName:        "Coffee",
				SubCatDescription: "Fresh brewed coffee",
				MenuItem: &[]hospitality_menu_contracts.AddMenuItemInput{

					{
						Name:        "Cappuccino",
						Description: "Fresh brewed cappuccino",
						BasePrice:   120,
						Unit:        "cup",
						DietaryType: menu_constants.Vegetarian,

						VariantGroup: &[]hospitality_menu_contracts.AddVariantGroupInput{
							{
								VariantGroupName: "Size",
								Variant: &[]hospitality_menu_contracts.AddVariantInput{
									{Name: "Small", Price: 0},
									{Name: "Regular", Price: 20},
									{Name: "Large", Price: 40},
								},
							},
						},
					},

					{
						Name:        "Cold Coffee",
						Description: "Iced coffee with cream",
						BasePrice:   160,
						Unit:        "glass",
						DietaryType: menu_constants.Vegetarian,
					},
				},
			},

			{
				SubCatName:        "Tea",
				SubCatDescription: "Traditional teas",
				MenuItem: &[]hospitality_menu_contracts.AddMenuItemInput{

					{
						Name:        "Masala Chai",
						Description: "Indian spiced tea",
						BasePrice:   40,
						Unit:        "cup",
						DietaryType: menu_constants.Vegetarian,
					},

					{
						Name:        "Green Tea",
						Description: "Healthy herbal tea",
						BasePrice:   60,
						Unit:        "cup",
						DietaryType: menu_constants.Vegan,
					},
				},
			},
		},
	},
}
