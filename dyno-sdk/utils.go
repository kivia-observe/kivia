package dynosdk

import (
	"io"
	"net/http"
)

func getPublicIP(httpClient *http.Client) string {
    res, err := httpClient.Get("https://api.ipify.org")
    if err != nil {
        return ""
    }
    defer res.Body.Close()
    ip, _ := io.ReadAll(res.Body)
    return string(ip)
}