package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	Url  string
	Data struct {
		Username            string `json:"username"`
		Ticket              string `json:"ticket"`
		CSRFPreventionToken string `json:"CSRFPreventionToken"`
	} `json:"data"`
}

func GetTicket(url string, username string, password string) *Client {
	var credentials = map[string]string{
		"username": username,
		"password": password,
	}
	var client = &Client{Url: url}
	data, _ := json.Marshal(credentials)
	ticket, _ := http.Post(url+"/api2/json/access/ticket", "application/json", bytes.NewReader(data))
	r, _ := io.ReadAll(ticket.Body)
	json.Unmarshal(r, &client)
	fmt.Println(client.Url)
	return client
}

