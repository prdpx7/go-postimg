package imgur

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
)

const ImgurURL = "https://api.imgur.com/3/image"

type Headers struct {
	Authorization string `json:"Authorization"`
}

type Payload struct {
	Image string `json:"image"`
}

type ImageData struct {
	Account_id int    `json:"account_id"`
	Animated   bool   `json:"animated"`
	Bandwidth  int    `json:"bandwidth"`
	DateTime   int    `json:"datetime"`
	Deletehash string `json:"deletehash"`
	Favorite   bool   `json:"favorite"`
	Height     int    `json:"height"`
	Id         string `json:"id"`
	In_gallery bool   `json:"in_gallery"`
	Is_ad      bool   `json:"is_ad"`
	Link       string `json:"link"`
	Name       string `json:"name"`
	Size       int    `json:"size"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	Views      int    `json:"views"`
	Width      int    `json:"width"`
}

type ImgurResponse struct {
	Status  int       `json:"status"`
	Data    ImageData `json:"data"`
	Success bool      `json:"success"`
}

func UploadImage(imgPath string) string {
	normalizedImgPath := NormalizePath(imgPath)
	content, err := os.ReadFile(normalizedImgPath)
	if err != nil {
		panic(err)
	}
	encodeImage := base64.StdEncoding.EncodeToString(content)

	payload, _ := json.Marshal(Payload{Image: encodeImage})

	req, err := http.NewRequest("POST", ImgurURL, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Client-ID 2b8986ab0193370")
	req.Header.Set("Content-Type", "application/json")

	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	var response ImgurResponse
	json.NewDecoder(resp.Body).Decode(&response)
	if response.Success {
		return response.Data.Link
	}
	panic("Unable to upload image")
}
