package gopenapi

type ApiAdapter struct {
	openid  string
	openkey string
	pf      string
	api     *OpenApi
}

func (this *ApiAdapter) SetApi(api *OpenApi) {
	this.api = api
}

func (this *ApiAdapter) makeParams(p map[string]string) map[string]string {
	params := make(map[string]string)
	params["openid"] = this.openid
	params["openkey"] = this.openkey
	params["pf"] = this.pf
	for k, v := range p {
		params[k] = v
	}
	return params
}
