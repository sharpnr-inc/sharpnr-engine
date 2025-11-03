package authcontrollers

import (
	"net/http"
	dto "sharpnr-engine/core/dtos"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *AuthController) GetOrgUIRenderer(ctx *gin.Context) {

	uiRender := map[string]any{
		"organisation_id":           "",
		"ui_renderer_tree_id":       "",
		"orgnisation_business_type": "hospitality",
		"main_menu_render_tree": map[string]any{

			"sharp_hospitality": map[string]any{
				"has_access": true,
				"sub_module_section": map[string]any{
					"hotel": map[string]any{
						"has_access": true,
						"sub_section_access": map[string]bool{
							"overview":        true,
							"booking":         true,
							"request":         false,
							"room_management": true,
						},
					},
					"bar": map[string]any{
						"has_access": true,
						"sub_section_access": map[string]bool{
							"menu":      true,
							"orders":    true,
							"inventory": true,
						},
					},
					"restaurent": map[string]any{
						"has_access": true,
						"sub_section_access": map[string]bool{
							"table_management":        true,
							"orders":                  true,
							"pickup":                  true,
							"reservation_and_booking": true,
						},
					},
					"services":     map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"menu":         map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"kitchen":      map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"inventory":    map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"housekeeping": map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"settings":     map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"orgConfig":    map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
				},
			},
			"sharp_accounting": map[string]any{
				"has_access": true,
				"sub_module_section": map[string]any{
					"transactions":        map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"expenses":            map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"invoices":            map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"ledger":              map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"customers":           map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"suppliersandvendors": map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"reports":             map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
				},
			},
			"sharp_hr": map[string]any{
				"has_access": true,
				"sub_module_section": map[string]any{
					"employee":          map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"payroll":           map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"leaves_management": map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"reports":           map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
				},
			},
			"sharp_marketing": map[string]any{
				"has_access": true,
				"sub_module_section": map[string]any{
					"campaigns":               map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"leads":                   map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"reports":                 map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"social_media_management": map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
				},
			},
			"employee": map[string]any{
				"has_access": true,
				"sub_module_section": map[string]any{
					"profile":                 map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"leads":                   map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"reports":                 map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
					"social_media_management": map[string]any{"has_access": true, "sub_section_access": map[string]bool{}},
				},
			},
		},
	}

	var data any = uiRender

	ctx.JSON(http.StatusOK, dto.Response[any]{
		StatusCode: 200,
		Message:    "Successfully fetched UI Renderer Tree",
		Error: dto.ErrorResponse{
			Message: "",
			IsError: false,
		},
		Data:      &data,
		Timestamp: time.Now(),
	})
}
