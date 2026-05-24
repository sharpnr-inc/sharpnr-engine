package orchestrator_routes

import (
	orchestrator_controllers "sharpnr-engine/orchestrator/controllers"

	"github.com/gin-gonic/gin"
)

func registerLegalEntityRoutes(r *gin.RouterGroup, co *orchestrator_controllers.OrchestratorController) {

	ler := r.Group("/legal-entity")
	{
		// --- Root Legal Entity CRUD ---
		ler.GET("", co.ListAllLegalEntity)
		ler.POST("", co.AddLegalEntity)
		ler.GET("/:legalEntityID", co.GetLegalEntityDetailsHandler)
		ler.PUT("/:legalEntityID")
		ler.DELETE("/:legalEntityID")

		/* ------------------------------------------------------------------ */
		/*                         Address Routes                              */
		/* ------------------------------------------------------------------ */
		addr := ler.Group("/:legalEntityID/addresses")
		{
			addr.GET("", co.ListAllAddressesOfLegalEntity)
			addr.POST("", co.AddAddressesInLegalEntityHandler)
			addr.PUT("/:addressID")
			addr.DELETE("/:addressID")
			addr.GET("/:addressID")
		}

		/* ------------------------------------------------------------------ */
		/*                       Bank Account Routes                           */
		/* ------------------------------------------------------------------ */
		bank := ler.Group("/:legalEntityID/bank-accounts")
		{
			bank.GET("", co.ListAllBankAccountsOfLegalEntity)
			bank.POST("", co.AddBankAccountDetailsInLegalEntityHandler)
			bank.PUT("/:bankAccountID")
			bank.DELETE("/:bankAccountID")
		}

		/* ------------------------------------------------------------------ */
		/*                        Identity Routes                              */
		/* ------------------------------------------------------------------ */
		identity := ler.Group("/:legalEntityID/identities")
		{
			identity.GET("", co.ListAllIdentifiersOfLegalEntity)
			identity.POST("", co.AddIdentifiersInLegalEntityHandler)
			identity.PUT("/:identityID")
			identity.DELETE("/:identityID")
		}

		/* ------------------------------------------------------------------ */
		/*                         Contact Routes                              */
		/* ------------------------------------------------------------------ */
		contact := ler.Group("/:legalEntityID/contacts")
		{
			contact.GET("", co.ListAllContactsOfLegalEntity)
			contact.POST("", co.AddContactsInLegalEntityHandler)
			contact.PUT("/:contactID")
			contact.DELETE("/:contactID")
		}

		/* ------------------------------------------------------------------ */
		/*                        License & Compliance                         */
		/* ------------------------------------------------------------------ */
		license := ler.Group("/:legalEntityID/licenses")
		{
			license.GET("", co.ListAllLicenseAndComplianceHandler)
			license.POST("", co.AddLicenseAndComplianceHandler)
			license.PUT("/:licenseID")
			license.DELETE("/:licenseID")
		}

		/* ------------------------------------------------------------------ */
		/*                        Tax Registration                             */
		/* ------------------------------------------------------------------ */
		tax := ler.Group("/:legalEntityID/tax-registrations")
		{
			tax.GET("", co.ListAllTaxRegistrationOfLegalEntityHandler)
			tax.POST("", co.AddTaxRegistrationInLegalEntityHandler)
			tax.PUT("/:taxRegID")
			tax.DELETE("/:taxRegID")
		}

		tax_categories := ler.Group("/:legalEntityID/tax-categories")
		{
			tax_categories.GET("")
			tax_categories.POST("")
			tax_categories.PUT("/:taxCatID")
			tax_categories.DELETE("/:taxCatID")
		}

		/* ------------------------------------------------------------------ */
		/*                          Accounting                                 */
		/* ------------------------------------------------------------------ */
		accounting := ler.Group("/:legalEntityID/accounting")
		{
			accounting.GET("")
		}

		/* ------------------------------------------------------------------ */
		/*                        Utility/Lookup Routes                        */
		/* ------------------------------------------------------------------ */
		utils := ler.Group("/utils")
		{
			utils.GET("/entity-types", co.GetEntityTypes)
			utils.GET("/gst-types", co.GetGSTRegistrationTypes)
			utils.GET("/address-types", co.GetLegalEntityAddressTypes)
			utils.GET("/license-types", co.GetLegalEntityLicenseTypes)
			utils.GET("/contact-types", co.GetLegalEntityContactTypes)
			utils.GET("/identifier-types", co.GetLegalEntityIdentifierTypes)
			utils.GET("/tax-component-types", co.GetLegalEntityTaxComponentTypes)
			utils.GET("/bank-account-types", co.GetLegalEntityBankAccountTypes)
			utils.GET("/tax-types", co.GetLegalEntityTaxTypes)
		}
	}
}
