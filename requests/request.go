package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/Avengerist/ProxmoxGo/client"
)

func SendPost(c *client.Client, path string, data map[string]interface{}) {
	jsondata, _ := json.Marshal(data)
	request, _ := http.NewRequest("POST", c.Url+path, bytes.NewReader(jsondata))
	request.Header.Set("Content-type", "application/json")
	request.Header.Add("CSRFPreventionToken", c.Data.CSRFPreventionToken)
	request.AddCookie(&http.Cookie{Name: "PVEAuthCookie", Value: c.Data.Ticket})
	http_client := &http.Client{}
	output, _ := http_client.Do(request)
	fmt.Println(output)
}
