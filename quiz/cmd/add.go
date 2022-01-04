package cmd

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "To add a new question.",
	Long:  `To add new question on db.json file.`,
	Run: func(cmd *cobra.Command, args []string) {
		question, _ := cmd.Flags().GetString("question")
		answer, _ := cmd.Flags().GetString("answer")

		saveVideos(&question, &answer)
	},
}

func init() {
	questionsCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().String("question", "", "The question")
	addCmd.PersistentFlags().String("answer", "", "The answer")
}

func saveVideos(question *string, answer *string) {

	// construct the body of the request as Json format
	postBody, _ := json.Marshal(map[string]string{
		"question": *question,
		"answer":   *answer,
	})

	// convert the encoded JSON data to a type implemented by the io.Reader interface
	requestBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:3000/questions", "application/json", requestBody) // send the request
	if err != nil {
		log.Fatalf("An error Occured %v", err)
	}

	defer resp.Body.Close() // close the response body
}
