package login

import (
    "bytes"
    "encoding/json"
	"fmt"
	"net/http"
	"io"

	"amplify-ob/types"
)

func LoginHandler(user string, pass string) error{
	url := "https://services.sandbox.ampint.axwaytest.net/api/auth/login"
	payload := types.LoginPostRequest{
		Domain: "open-banking",
		DuplicateSession: false,
		Email: user,
		Password: pass,
		Remember: true,
	}
	bodyBytes, err := json.Marshal(payload)
    if err != nil {
        return fmt.Errorf("could not marshal payload: %w", err)
    }
	resp, err := http.Post(url, "application/json", bytes.NewReader(bodyBytes));
	if err != nil {
		fmt.Println("error in Post");
	}
    defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	fmt.Println("Response JSON:", string(data))	
	// fmt.Println(string(data))

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("login failed: %s — %s", resp.Status, string(data))
    }
	return nil
}