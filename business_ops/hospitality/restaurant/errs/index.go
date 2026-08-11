// here we will list all the Erros regarding the erros for Restaurant Modules,
package hospitalityRestaurantErrs

import "errors"

var (
	NoLegalEntityAttachedToRestaurnt = errors.New("no legal entity attached to the restaurant")
	NoKitchenAttachedToRestaurant    = errors.New("no kitchen is attached to the restaurant")
)
