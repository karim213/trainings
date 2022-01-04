package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

type local_question struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "To get questions.",
	Long:  `To get questions from the db.json file`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt("id")
		all, _ := cmd.Flags().GetBool("all")
		getVideos(&all, &id)
	},
}

func init() {
	questionsCmd.AddCommand(getCmd)

	getCmd.PersistentFlags().Int("id", 0, "Get question by ID")
	getCmd.PersistentFlags().Bool("all", false, "Get All questions")

}

func getVideos(all *bool, id *int) {
	if !*all && *id == 0 {
		pterm.Error.Println("Some parameters are missing!")
		return
	}

	Url := "http://localhost:3000/questions" // construct the url with the parameters

	if !*all && *id != 0 {
		Url = fmt.Sprintf("http://localhost:3000/questions/%d", *id)
	}

	resp, err := http.Get(Url) // send the request
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body) // reading the response
	if err != nil {
		log.Fatalln(err)
	}
	pterm.Println() // Blank line
	if *all {
		var data []local_question
		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			fmt.Println(err.Error())
		}
		d := pterm.TableData{{"ID", "Question", "Answer"}}
		for _, value := range data {
			d = append(d, []string{fmt.Sprint(value.Id), value.Question, value.Answer})
		}
		pterm.DefaultTable.WithHasHeader().WithData(d).Render()

	} else {
		var data local_question
		err = json.Unmarshal([]byte(body), &data)
		if err != nil {
			fmt.Println(err.Error())
		}
		d := pterm.TableData{{"ID", "Question", "Answer"}}
		d = append(d, []string{fmt.Sprint(data.Id), data.Question, data.Answer})

		pterm.DefaultTable.WithHasHeader().WithData(d).Render()
	}
	pterm.Println() // Blank line

}
