package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"hackalist-cli/utils"
)

var year, month, day = time.Now().Date()

const rootApiURL = "http://www.hackalist.org/api/1.0/2019"

type Hackathons map[string][]struct {
	// Hackathons is a hash which maps names of months to arrays of events
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

	var computedMonth string
	currentMonth := strconv.Itoa(int(month))
	if int(month) < 10 {
		computedMonth = "0" + currentMonth
	} else {
		computedMonth = currentMonth
	}
	API_URL := fmt.Sprintf("%s/%s.json", rootApiURL, computedMonth)

	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)

	utils.ClearScreen()
	utils.ShowBanner()

	color.Cyan(" \nFetching Data")
	spin.Start()
	time.Sleep(2 * time.Second)

	client := &http.Client{
		Timeout: time.Duration(1 * time.Minute),
	}

	resp, err := client.Get(API_URL)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	var hackathon Hackathons

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(body, &hackathon)
		if err != nil {
			log.Println(err)
		}

		spin.Stop()

		for month, events := range hackathon {
			fmt.Println(month)
			for index, event := range events {
				fmt.Printf("%v => %v \nwebsite: %s\n\n", index, event.Title, event.URL)
			}
		}
	}
}

var hackathonCmd = &cobra.Command{
	Use:   "hackathons",
	Short: "Lists all the hackathons available",
	Run: func(cmd *cobra.Command, args []string) {
		listHackathons()
	},
}

func init() {
	rootCmd.AddCommand(hackathonCmd)
}
