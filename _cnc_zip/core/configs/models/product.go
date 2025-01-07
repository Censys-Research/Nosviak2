package models

type Product struct {
	Action string `json:"action"`
	Fields struct {
		Ranks       *[]string	`json:"ranks"`
		Theme       *string        	`json:"theme"`
		Maxtime     *int           	`json:"maxtime"`
		Cooldown    *int           	`json:"cooldown"`
		Concurrents *int           	`json:"concurrents"`
		MaxSessions *int           	`json:"max_sessions"`
		Expiry      *struct {
			Type  string 	`json:"type"`
			Value int    	`json:"value"`
		} `json:"expiry"`
	} `json:"fields"`
}
