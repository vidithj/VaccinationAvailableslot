package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	cowinurlhost  = "https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByPin?pincode="
	smsgatewayurl = ""
	authtoken     = ""
	senderid      = "TXTIND"
	language      = "english"
)

var (
	minage    = 18
	maxage    = 100
	pincode   = "226010"
	timedelay = 1
	phone     string
)

type userInfo struct {
	Pincode   string `json:"PINCODE"`
	Email     string `json:"EMAIL"`
	MinAge    int    `json:"MINAGE"`
	MaxAge    int    `json:"MAXAGE"`
	TimeDelay int    `json:"TIMEDELAY"`
	Phone     string `json:"PHONE"`
}
type slotInfo struct {
	Sessions []Sessions `json:"sessions"`
}
type Sessions struct {
	Name              string   `json:"name"`
	State_Name        string   `json:"state_name"`
	District_Name     string   `json:"district_name"`
	Pincode           int      `json:"pincode"`
	From              string   `json:"from"`
	To                string   `json:"to"`
	Date              string   `json:"date"`
	AvailableCapacity int      `json:"available_capacity"`
	Fee               string   `json:"fee"`
	Age               int      `json:"min_age_limit"`
	Vaccine           string   `json:"vaccine"`
	Slots             []string `json:"slots"`
}

func main() {
	arg := os.Args
	var cowinurl string
	if len(arg) > 1 {
		cowinurl = cowinurlhost + arg[1] + "&date="
		minage, _ = strconv.Atoi(arg[2])
		maxage, _ = strconv.Atoi(arg[3])
		timedelay, _ = strconv.Atoi(arg[4])
		phone = arg[5]
	} else {
		user, err := readInfo()
		if err != nil {
			fmt.Println("Error while fetching userinfo : ", err)
		}
		cowinurl = cowinurlhost + user[0].Pincode + "&date="
		minage = user[0].MinAge
		maxage = user[0].MaxAge
		timedelay = user[0].TimeDelay
		phone = user[0].Phone
	}
	for {
		weekdates := getDates()
		for _, val := range weekdates {
			found := GetValidSlots(cowinurl + val)
			if !found {
				fmt.Println("No slot found for :", val)
			}
		}
		time.Sleep(time.Duration(timedelay) * time.Hour)
	}
}
func readInfo() ([]userInfo, error) {
	user := make([]userInfo, 0)
	userinfoFile, err := os.Open("assets/Userinfo.json")
	if err != nil {
		return user, err
	}
	defer userinfoFile.Close()
	userinfoValue, _ := ioutil.ReadAll(userinfoFile)
	json.Unmarshal(userinfoValue, &user)
	return user, nil
}
func getDates() []string {
	today := time.Now()
	weekdates := make([]string, 0)
	for i := 0; i < 7; i++ {
		dateString := today.Format("02-01-2006")
		weekdates = append(weekdates, dateString)
		today = today.AddDate(0, 0, 1)
	}
	return weekdates
}
func GetValidSlots(formattedurl string) bool {
	cowinurlwithdate, _ := url.Parse(formattedurl)
	req := &http.Request{
		Method: "GET",
		URL:    cowinurlwithdate,
		Header: map[string][]string{
			"accept":          {"application/json"},
			"Accept-Language": {"hi_IN"},
			"user-agent":      {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36"},
		},
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error while fetching slots : ", err)
	}
	data, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	resp := slotInfo{}
	validslots := make([]Sessions, 0)
	json.Unmarshal(data, &resp)
	for _, val := range resp.Sessions {
		if val.Age >= minage && val.Age < maxage && val.AvailableCapacity > 0 {
			validslots = append(validslots, val)
		}
	}
	if len(validslots) > 0 {
		return notify(validslots)
	}
	return false

}
func notify(validslots []Sessions) bool {
	var smsbody []string
	for ind, val := range validslots {
		fmt.Println("-------------------------------")
		fmt.Println("Vaccination Information #", ind+1)
		fmt.Println("--------------------------------")
		fmt.Println("Vaccination Center Name:", val.Name)
		fmt.Println("Location:", val.State_Name, " ,", val.District_Name, " ,", val.Pincode)
		fmt.Println("Date: ", val.Date)
		fmt.Println("Vaccination Center Timings:", val.To, " - ", val.From)
		fmt.Println("Vaccine available capacity:", val.AvailableCapacity)
		fmt.Println("Vaccination Fees:", val.Fee)
		fmt.Println("Vaccine:", val.Vaccine)
		fmt.Println("Availble Slots:", strings.Join(val.Slots, " ,"))
		fmt.Println("--------------------------------------------------------------")
		smsbody = append(smsbody, "center- "+val.Name+" total slots- "+strconv.Itoa(len(val.Slots))+" vaccine capacity- "+strconv.Itoa(val.AvailableCapacity))
	}

	return notifyviasms(smsbody)
}

func notifyviasms(smsbody []string) bool {
	smsurl, _ := url.Parse(smsgatewayurl)
	postBody, _ := json.Marshal(map[string]string{
		"route":     "v3",
		"sender_id": senderid,
		"message":   strings.Join(smsbody, "\n"),
		"language":  language,
		"flash":     "0",
		"numbers":   phone,
	})
	reqBody := ioutil.NopCloser(bytes.NewBuffer(postBody))
	req := &http.Request{
		Method: "POST",
		URL:    smsurl,
		Header: map[string][]string{
			"authorization": {authtoken},
			"Content-Type":  {"application/json"},
		},
		Body: reqBody,
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("SMS Gateway issues. Cant deliver msg : ", err)
	}
	if res.StatusCode >= 200 && res.StatusCode < 400 {
		return true
	}
	fmt.Println("SMS Gateway issues. Cant deliver msg")
	return true
}
