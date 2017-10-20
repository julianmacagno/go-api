package main //Paralelo

import (
	"github.com/go-resty/resty"
	"fmt"
	"encoding/json"
	"sync"
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
	From string
	To string
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

	conversion.From = from
	conversion.To = to

	return &conversion, nil
}

func GetConversionToUSD(c chan *ConversionRate, siteId string) {

	s, err := GetSite(siteId)

	if err != nil { 
		c <- nil
		return
	}
	conversion, err := GetRatio(s.DefaultCurrencyID, CURRENCY_USD)

	if err != nil {
		c <- nil
		return
	}
	
	c <- conversion
}

func HandleResults(c chan *ConversionRate, wg *sync.WaitGroup, m map[string]float64) {
	for i:=0; i<20; i++ {
		ConversionRate := <-c

		if ConversionRate != nil  {
			m[ConversionRate.From] = ConversionRate.Ratio
		}

		wg.Done()			
	}
}

func GetAllCurrencies() (map[string]float64, error) {
	
	sites, err := GetAllSites()
	
	if err != nil {
		return nil, err
	}

	result := make(map[string]float64)

	c := make(chan *ConversionRate)
	
	var wg sync.WaitGroup	
	var i = 0

	go HandleResults(c, &wg, result)

	for _, site := range *sites {
		
		//aca hacemos paralelismo; para ello creamos canales para comunicar las distinas go
		//rutines con el hilo principal.
		wg.Add(1)
		i++
		
		go	GetConversionToUSD(c, site.Id)
	}	

	wg.Wait()

	fmt.Println(result)

	return result, nil
} 

func main() {

	c, err := GetAllCurrencies()

	if err != nil {
		return
	}
		
	fmt.Println(c)
}
