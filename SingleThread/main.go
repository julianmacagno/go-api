package main //Lineal

import (
	"github.com/go-resty/resty"
	"fmt"
	"encoding/json"
)

const (
	URL_SITES_ALL string = "https://api.mercadolibre.com/sites"
	URL_SITES_SINGLE string = "https://api.mercadolibre.com/sites/%s"
	URL_CURRENCYS_CONVERSIONS string = "https://api.mercadolibre.com/currency_conversions/search?from=%s&to=%s"
	CURRENCY_USD string = "USD"
)

type SitesSearch [] struct{
	Id					string 	`json="id"`
	Name 				string	`json="name"`
}

type Site struct {
	Id                 	string  `json:"id"`
	Name               	string  `json:"name"`
	DefaultCurrencyID  	string  `json:"default_currency_id"`
}

type ConversionRate struct {
	Ratio float64 `json:"ratio"`
}

func GetAllSites() (*SitesSearch, error){
	
	resp, err := resty.R().Get(URL_SITES_ALL)
	
	if err != nil{
		return nil, err
	}

	var sites SitesSearch
	err = json.Unmarshal(resp.Body(), &sites)
	
	if err != nil{
		return nil, err
	}

	return &sites, nil
}

func GetSite(siteId string) (*Site, error) {

	resp, err := resty.R().Get(fmt.Sprintf(URL_SITES_SINGLE, siteId))

	if err != nil{
		return nil, err
	}

	var site Site
	err = json.Unmarshal(resp.Body(), &site)

	if err != nil {
		return nil, err
	}

	return &site, nil
}

func GetRatio(from, to string) (*ConversionRate, error) {
	
	url := fmt.Sprintf(URL_CURRENCYS_CONVERSIONS, from, to)

	resp, err := resty.R().Get(url)

	if err != nil {
		return nil, err
	}

	var conversion ConversionRate
	err = json.Unmarshal(resp.Body(), &conversion)

	if err != nil {
		return nil, err
	}

	return &conversion, nil
}

func GetAllCurrencies() (map[string]float64, error) {
	
	sites, err := GetAllSites()
	
	if err != nil {
		return nil, err
	}

	result := make(map[string]float64)

	for _, site := range *sites {
		
		s, err := GetSite(site.Id)

		if err == nil {
			conversion, err := GetRatio(s.DefaultCurrencyID, CURRENCY_USD)
			f := "%s: %v"
			fmt.Println(fmt.Sprintf(f, s.Name, conversion.Ratio))

			if err == nil {
				result[s.Id] = conversion.Ratio
			} 
		}
	}	

	return result, nil
} 

func main() {

	c, err := GetAllCurrencies()

	if err != nil {
		return
	}
		
	fmt.Println(c)
}
