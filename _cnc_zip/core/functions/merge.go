package functions

import (
	"Nosviak2/core/configs/models"
	"Nosviak2/core/database"
	"Nosviak2/core/sources/ranks"
	"time"
)

// MergeFieldsWithUser will merge the product fields with the db user
func MergeFieldsWithUser(product *models.Product, user *database.User) error {
	if product.Fields.Ranks != nil {
		r := ranks.MakeRank(user.Username)
		err := r.SyncWithString(user.Ranks)
		if err != nil {
			return err
		}
	
		for _, i := range *product.Fields.Ranks {
			r.GiveRank(i)
		}

		user.Ranks, err = r.MakeString()
		if err != nil {
			return err
		}
	}

	// Modifys the theme field if not equl to nil
	if product.Fields.Theme != nil {
		user.Theme = *product.Fields.Theme
	}

	// Modifys the maxtime field if not equal to nil
	if product.Fields.Maxtime != nil {
		user.MaxTime = *product.Fields.Maxtime
	}

	// Modifys the concurrents field if not equal to nil
	if product.Fields.Concurrents != nil {
		user.Concurrents = *product.Fields.Concurrents
	}

	// Modifys the cooldown field if not equal to nil
	if product.Fields.Cooldown != nil {
		user.Cooldown = *product.Fields.Cooldown
	}

	// Modifys the maxSessions field if not equal to nil
	if product.Fields.MaxSessions != nil {
		user.MaxSessions = *product.Fields.MaxSessions
	}

	// Modifys/Expands the expiry field if not equal to nil
	if product.Fields.Expiry != nil {
		switch product.Fields.Expiry.Type {

		case "add": // Adds ontop of the users current expiry
			user.Expiry = time.Unix(user.Expiry, 0).Add((time.Hour * 24) * time.Duration(product.Fields.Expiry.Value)).Unix()
		
		case "set": // Sets the users expiry to x amount of days
			user.Expiry = time.Now().Add((time.Hour * 24) * time.Duration(product.Fields.Expiry.Value)).Unix()
		}
	}

	return nil
}