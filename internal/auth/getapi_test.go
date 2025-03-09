package auth

import (
    "testing"
    "net/http"
)



func TestGetAPIKey(t *testing.T) {
    headers := make(http.Header)
    //headers.Add("Authorization", "123")

    _, err := GetAPIKey(headers)

    if err != ErrNoAuthHeaderIncluded {
        t.Fatalf("No error!")
    }
}

