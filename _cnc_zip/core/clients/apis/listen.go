package apis

import (
	"Nosviak2/core/sources/layouts/toml"
	//"Nosviak2/core/sources/tools"
	"net/http"
)

//ListenAndServe will run the api
func ListenAndServe() error {

	http.HandleFunc("/attack", Attack)

	http.HandleFunc("/ongoing", Ongoing)
	
	http.HandleFunc("/edition", Edition)

	http.HandleFunc("/method", Methods)

	http.HandleFunc("/autobuy", Autobuy)


	//TLS mode enabled on api
	if toml.ApiToml.TLS.TLS { //starts listening on the tls side
		return http.ListenAndServeTLS(toml.ApiToml.API.Host, toml.ApiToml.TLS.Certification, toml.ApiToml.TLS.Key, nil)
	} else { //non tls driver will start here
		return http.ListenAndServe(toml.ApiToml.API.Host, nil)
	}
}