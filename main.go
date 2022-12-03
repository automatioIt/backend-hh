package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://api.hh.ru/"

type Vacancies struct {
	AlternateURL string      `json:"alternate_url"`
	Arguments    interface{} `json:"arguments"`
	Clusters     interface{} `json:"clusters"`
	Found        int64       `json:"found"`
	Items        []struct {
		AcceptTemporary bool `json:"accept_temporary"`
		Address         struct {
			Building    string      `json:"building"`
			City        string      `json:"city"`
			Description interface{} `json:"description"`
			ID          string      `json:"id"`
			Lat         float64     `json:"lat"`
			Lng         float64     `json:"lng"`
			Metro       struct {
				Lat         float64 `json:"lat"`
				LineID      string  `json:"line_id"`
				LineName    string  `json:"line_name"`
				Lng         float64 `json:"lng"`
				StationID   string  `json:"station_id"`
				StationName string  `json:"station_name"`
			} `json:"metro"`
			MetroStations []struct {
				Lat         float64 `json:"lat"`
				LineID      string  `json:"line_id"`
				LineName    string  `json:"line_name"`
				Lng         float64 `json:"lng"`
				StationID   string  `json:"station_id"`
				StationName string  `json:"station_name"`
			} `json:"metro_stations"`
			Raw    string `json:"raw"`
			Street string `json:"street"`
		} `json:"address"`
		AdvResponseURL    string `json:"adv_response_url"`
		AlternateURL      string `json:"alternate_url"`
		ApplyAlternateURL string `json:"apply_alternate_url"`
		Archived          bool   `json:"archived"`
		Area              struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"area"`
		Contacts   interface{} `json:"contacts"`
		CreatedAt  string      `json:"created_at"`
		Department interface{} `json:"department"`
		Employer   struct {
			AlternateURL string `json:"alternate_url"`
			ID           string `json:"id"`
			LogoUrls     struct {
				Two40    string `json:"240"`
				Nine0    string `json:"90"`
				Original string `json:"original"`
			} `json:"logo_urls"`
			Name         string `json:"name"`
			Trusted      bool   `json:"trusted"`
			URL          string `json:"url"`
			VacanciesURL string `json:"vacancies_url"`
		} `json:"employer"`
		HasTest                bool          `json:"has_test"`
		ID                     string        `json:"id"`
		InsiderInterview       interface{}   `json:"insider_interview"`
		Name                   string        `json:"name"`
		Premium                bool          `json:"premium"`
		PublishedAt            string        `json:"published_at"`
		Relations              []interface{} `json:"relations"`
		ResponseLetterRequired bool          `json:"response_letter_required"`
		ResponseURL            interface{}   `json:"response_url"`
		Salary                 struct {
			Currency string `json:"currency"`
			From     int64  `json:"from"`
			Gross    bool   `json:"gross"`
			To       int64  `json:"to"`
		} `json:"salary"`
		Schedule struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"schedule"`
		Snippet struct {
			Requirement    string `json:"requirement"`
			Responsibility string `json:"responsibility"`
		} `json:"snippet"`
		SortPointDistance interface{} `json:"sort_point_distance"`
		Type              struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"type"`
		URL                  string        `json:"url"`
		WorkingDays          []interface{} `json:"working_days"`
		WorkingTimeIntervals []interface{} `json:"working_time_intervals"`
		WorkingTimeModes     []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"working_time_modes"`
	} `json:"items"`
	Page    int64 `json:"page"`
	Pages   int64 `json:"pages"`
	PerPage int64 `json:"per_page"`
}

func main() {
	fmt.Printf("fdsfsdf")
	MakeRequest("vacancies", "?text=GoLang%20junior")
}

func MakeRequest(path string, params string) {
	requestURL := fmt.Sprintf(url+"%s"+"%s", path, params)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	var vacancies Vacancies

	err = json.Unmarshal(resBody, &vacancies)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%v", vacancies.Found)

	//// Now let's unmarshall the data into `payload`
	//var result map[string]interface{}
	//err = json.Unmarshal(resBody, &result)
	//if err != nil {
	//	log.Fatal("Error during Unmarshal(): ", err)
	//}
	//
	////fmt.Printf("client: response body: %s\n", resBody)
	//fmt.Printf("Total: %v\n", result["found"])
	//
	//fmt.Printf("Total: %v\n", result["items"])
	//
	//var name = result["items"]
	//
	//fmt.Printf("%v", name)

}
