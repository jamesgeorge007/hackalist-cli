package cmd

import (
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"

	"hackalist-cli/utils"
)

const API_URL = "http://www.hackalist.org/api/1.0/2019/01.json"

type Hackathon struct {
	Title         string `json:"title"`
	URL           string `json:"url"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	Year          string `json:"year"`
	City          string `json:"city"`
	Host          string `json:"host"`
	Length        string `json:"length"`
	Size          string `json:"size"`
	Travel        string `json:"travel"`
	Prize         string `json:"prize"`
	HighSchoolers string `json:"highSchoolers"`
	Cost          string `json:"cost"`
	FacebookURL   string `json:"facebookURL"`
	TwitterURL    string `json:"twitterURL"`
	GooglePlusURL string `json:"googlePlusURL"`
	Notes         string `json:"notes"`
}

func listHackathons() {

	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond) 
	utils.ClearScreen()
	utils.ShowBanner()
	color.Cyan(" \nFetching Data")
	spin.Start()
	time.Sleep(2 * time.Second)
	resp, err := http.Get(API_URL)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	respBody, err:= ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	spin.Stop()

	var responseJSON []Hackathon
	json.Unmarshal(respBody, &responseJSON)

	fmt.Println(len(responseJSON))
}

var hackathonCmd = &cobra.Command{
	Use: "hackathons",
	Short: "Lists all the hackathons available",
	Run: func(cmd *cobra.Command, args []string) {
		listHackathons()
	},
}

func init() {
	rootCmd.AddCommand(hackathonCmd)
}