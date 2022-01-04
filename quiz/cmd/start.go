package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

type question struct {
	Id                       int
	Question                 string
	Description              string
	Answers                  map[string]string
	Multiple_correct_answers string
	Correct_answers          map[string]string
	Explanation              string
	Tip                      string
	Tags                     []map[string]string
	Category                 string
	Difficulty               string
}

type answer struct {
	Answer  string
	Correct string
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "To start new Quiz",
	Long:  `To start new Quiz, this command will show the list of categories of Quiz`,
	Run: func(cmd *cobra.Command, args []string) {

		// print header
		pterm.Println() // spacer
		pterm.DefaultHeader.WithFullWidth().Println("Welcom to my demo")
		pterm.Println() // spacer

		s, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString("QUIZ")).Srender()
		pterm.DefaultCenter.Println(s)
		pterm.DefaultCenter.WithCenterEachLineSeparately().Println("Developped by KARIM ABDI\nAtos 2021")

		// print header
		pterm.Println() // spacer
		pterm.DefaultHeader.WithFullWidth().Println("Let's go :-)")
		pterm.Println() // spacer

		category := getCategory()
		time.Sleep(1 * time.Second)
		pterm.DefaultSection.Println("Selected Category: ", category)
		difficulty := getDifficulty()
		time.Sleep(1 * time.Second)
		pterm.DefaultSection.Println("Selected Difficulty: ", difficulty)
		questions := getQuizByCategory(category, difficulty)
		runQuiz(questions)

	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

// user need select the Quiz category
func getCategory() (category string) {

	prompt := promptui.Select{
		Label: "Select Category: ",
		Items: []string{"Linux", "DevOps", "Docker"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	return result
}

// user need select the Quiz Difficulty
func getDifficulty() (difficulty string) {

	prompt := promptui.Select{
		Label: "Select Difficulty: ",
		Items: []string{"Easy", "Medium", "Hard"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	return result
}

// uget the quiz questions
func getQuizByCategory(category string, difficulty string) (questions []question) {
	apiKey := "Sv0vh47jftb4Mst3PznidthNclgDnIAnSgEDKwhl"
	Url := fmt.Sprintf("https://quizapi.io/api/v1/questions?apiKey=%s&category=%s&difficulty=%s", apiKey, category, difficulty) // construct the url with the parameters
	resp, err := http.Get(Url)                                                                                                  // send the request
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body) // reading the response
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &questions)
	if err != nil {
		panic(err)
	}

	return questions

}

// start displaying questions
func runQuiz(questions []question) {
	for _, question := range questions {

		var items []string
		correctAnswer := ""

		for k, value := range question.Answers {
			if value != "" && k != "" {
				if question.Correct_answers[k+"_correct"] == "true" {
					correctAnswer = value
				}
				items = append(items, value)
			}
		}

		templates := &promptui.SelectTemplates{
			Label:    "{{ . | cyan }}",
			Active:   "\U0001F336 {{ . | green  }}",
			Inactive: "  {{ .Answer }}",
			Selected: " ",
		}

		prompt := promptui.Select{
			Label:     question.Question,
			Items:     items,
			Templates: templates,
		}

		_, selected, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		pterm.DefaultSection.Println(question.Question)
		spinnerSuccess, _ := pterm.DefaultSpinner.Start("Checking your answer")

		pterm.NewRGB(15, 199, 209).Println("Your answer:  ", selected)
		time.Sleep(time.Second * 2) // Simulate 2 seconds of processing.
		if correctAnswer == selected {
			spinnerSuccess.Success("Correct answer")
		} else {
			spinnerSuccess.Fail("Wrong answer")
			fmt.Println("The correct answer is: ", correctAnswer)
		}

		fmt.Println("-------------------------------------- \n\n")
	}

}
