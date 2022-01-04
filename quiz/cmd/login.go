package cmd

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/fatih/color"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with Okta",
	Long:  `When you run this command you will be redirected to the browser to login on Okta.`,
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("config") // name of config file (without extension)
		viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath("./config/")
		//renewToken()
		ctx = context.Background()
		conf = &oauth2.Config{
			ClientID:     "0oa7ym8feNw9SpRse696",
			ClientSecret: "CXgpok4K298k8I9Tkj4U9lQDhcWaEer64vlW0FJR",
			Scopes:       []string{"openid profile email groups", "offline_access"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "atos-poc.okta.com/oauth2/v1/authorize",
				TokenURL: "https://atos-poc.okta.com/oauth2/v1/token",
			},
			// my own callback URL
			RedirectURL: "http://localhost:8080/authorization-code/callback",
		}

		// add transport for self-signed certificate to context
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		sslcli := &http.Client{Transport: tr}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, sslcli)

		// Redirect user to consent page to ask for permission
		// for the scopes specified above.
		url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)

		log.Println(color.CyanString("You will now be taken to your browser for authentication"))
		time.Sleep(1 * time.Second)
		open.Run(url)
		time.Sleep(1 * time.Second)
		log.Printf("Authentication URL: %s\n", url)

		http.HandleFunc("/authorization-code/callback", callbackHandler)
		log.Fatal(http.ListenAndServe(":8080", nil))

	},
}

var (
	conf *oauth2.Config
	ctx  context.Context
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

// `/authorization-code/callback` handler

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	queryParts, _ := url.ParseQuery(r.URL.RawQuery)

	// Use the authorization code that is pushed to the redirect
	// URL.
	code := queryParts["code"][0]
	log.Printf("code: %s\n\n", code)

	// Exchange will do the handshake to retrieve the initial access token.
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}

	viper.Set("access_token", tok.AccessToken)
	viper.Set("refresh_token", tok.RefreshToken)
	viper.WriteConfig() // writes current config to predefined path set by 'viper.AddConfigPath()' and 'viper.SetConfigName'
	log.Printf("Access Token: %s \n\n", viper.Get("access_token"))
	log.Printf("Refresh Token: %s \n\n", viper.Get("refresh_token"))

	// The HTTP Client returned by conf.Client will refresh the token as necessary.
	client := conf.Client(ctx, tok)

	resp, err := client.Get("https://atos-poc.okta.com/")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(color.CyanString("Authentication successful"))
	}
	defer resp.Body.Close()

	// show succes page
	msg := "<p><strong>Success!</strong></p>"
	msg = msg + "<p>You are authenticated and can now return to the CLI.</p>"
	fmt.Fprintf(w, msg)
}
