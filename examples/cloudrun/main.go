package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/meant4/reamaze-go/reamaze"
	"github.com/rs/cors"
	"github.com/spf13/viper"
)

var (
	c *viper.Viper
)

type SendMessageBody struct {
	FirstName   string `json:"firstName" validate:"true"`
	LastName    string `json:"lastName" validate:"true"`
	Email       string `json:"email" validate:"true"`
	PhoneNumber string `json:"phoneNumber"`
	CompanyName string `json:"companyName"`
	HearAboutUs string `json:"hearAboutUs" validate:"true"`
	TellUs      string `json:"tellUs" validate:"true"`
	Nda         bool   `json:"nda"`
	Marketing   bool   `json:"marketing"`
}

type ResponseMsg struct {
	ErrCode int      `json:"errCode,omitempty"`
	ErrMsg  string   `json:"errMsg,omitempty"`
	Fields  []string `json:"fields,omitempty"`
	Msg     string   `json:"msg,omitempty"`
	Id      string   `json:"id,omitempty"`
}

const (
	MissingRequired = iota + 100
	InvalidJSON
	WebserviceError
	ReamazeError
)

func init() {
	fmt.Printf("Initializing api %s\n", time.Now().String())
	c = viper.New()
	c.AddConfigPath(".")
	c.SetConfigName("dev")
	c.SetConfigType("env")
	if err := c.ReadInConfig(); err != nil {
		log.Println("Error reading env file", err)
	}
	c.AutomaticEnv()
	// Setting up reamaze variables
	c.SetDefault("REAMAZE_EMAIL", "test@example.com")
	c.SetDefault("REAMAZE_API_TOKEN", "dummy")
	c.SetDefault("REAMAZE_BRAND", "example")
	c.SetDefault("REAMAZE_CATEGORY", "relations")
	c.WatchConfig()
	fmt.Printf("Initialized %d variables: %s\n", len(c.AllKeys()), strings.ToUpper(strings.Join(c.AllKeys(), ", ")))
}

func main() {
	// Initializing router
	router := httprouter.New()
	router.POST("/v1/sendmessage", SendReamazeMessage)
	// Setting default CORS headers
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func SendReamazeMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var incomingMsg SendMessageBody
	w.Header().Add("Content-Type", "application/json")

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}
	defer r.Body.Close()
	if len(body) > 0 {
		err := json.Unmarshal(body, &incomingMsg)
		if err != nil {
			respMsg := ResponseMsg{
				ErrCode: InvalidJSON,
				ErrMsg:  fmt.Sprintf("Invalid JSON request. %s", err),
				Msg:     fmt.Sprintf("Invalid JSON request. %s", err),
			}
			resp, _ := json.Marshal(respMsg)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, string(resp))
			return
		}

		// Initializing new reamaze client
		reamazeClient, err := reamaze.NewClient(c.GetString("REAMAZE_EMAIL"), c.GetString("REAMAZE_API_TOKEN"), c.GetString("REAMAZE_BRAND"))
		if err != nil {
			log.Println(err)
		}

		// Preparing reamaze payload to create new conversation
		payload := &reamaze.CreateConversationRequest{}
		payload.Conversation.Category = c.GetString("REAMAZE_CATEGORY")
		payload.Conversation.Message.Body = incomingMsg.TellUs
		payload.Conversation.User.Email = incomingMsg.Email
		payload.Conversation.User.Name = incomingMsg.FirstName + " " + incomingMsg.LastName
		payload.Conversation.Data = struct {
			Phone       string `json:"phone,omitempty"`
			CompanyName string `json:"company_name,omitempty"`
			HearAboutUs string `json:"hear_about_us,omitempty"`
			Nda         bool   `json:"nda"`
			Marketing   bool   `json:"marketing_approval"`
			FirstName   string `json:"first_name,omitempty"`
			LastName    string `json:"last_name,omitempty"`
		}{
			Phone:       incomingMsg.PhoneNumber,
			CompanyName: incomingMsg.CompanyName,
			HearAboutUs: incomingMsg.HearAboutUs,
			Nda:         incomingMsg.Nda,
			Marketing:   incomingMsg.Marketing,
			FirstName:   incomingMsg.FirstName,
			LastName:    incomingMsg.LastName,
		}
		payload.Conversation.User.Data = struct {
			Phone       string `json:"phone,omitempty"`
			CompanyName string `json:"company_name,omitempty"`
			Nda         bool   `json:"nda"`
			Marketing   bool   `json:"marketing_approval"`
			FirstName   string `json:"first_name,omitempty"`
			LastName    string `json:"last_name,omitempty"`
		}{
			Phone:       incomingMsg.PhoneNumber,
			CompanyName: incomingMsg.CompanyName,
			Nda:         incomingMsg.Nda,
			Marketing:   incomingMsg.Marketing,
			FirstName:   incomingMsg.FirstName,
			LastName:    incomingMsg.LastName,
		}

		_, err = reamazeClient.CreateConversation(payload)
		if err != nil {
			log.Println("error from CreateConversation:", err)
			respMsg := ResponseMsg{
				ErrCode: ReamazeError,
				ErrMsg:  fmt.Sprintf("Internal communication error: %s", err),
				Msg:     fmt.Sprintf("Internal communication error: %s", err),
			}

			resp, _ := json.Marshal(respMsg)
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, string(resp))
			return
		}

		respMsg := ResponseMsg{
			Msg: "OK",
		}
		resp, _ := json.Marshal(respMsg)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(resp))
		return
	}
}
