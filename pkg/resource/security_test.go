package resource

import (
    "net/http"
    "testing"
)

func TestAuthenticateRequest_EmptyAuthorizationHeader(t *testing.T) {
    requestHeader := http.Header{}
    requestHeader.Set("Authorization", "")

    code, err := authenticateRequest(requestHeader)

    if err != nil {
        expectedMsg := "Please provide authorize key."
        if err.Error() != expectedMsg {
            t.Errorf("Expected error message '%s', but was '%s'", expectedMsg, err.Error())
        }
    } else {
        t.Errorf("Expected error not null")
    }

    expectedCode := http.StatusBadRequest
    if code != expectedCode {
        t.Errorf("Expected %d, but got %d", expectedCode, code)
    }
}


func TestAuthenticateRequest_InvalidAuthorizationHeader(t *testing.T) {
    requestHeader := http.Header{}
    requestHeader.Set("Authorization", "12345")

    code, err := authenticateRequest(requestHeader)

    if err != nil {
        expectedMsg := "Unauthorized access."
        if err.Error() != expectedMsg {
            t.Errorf("Expected error message '%s', but was '%s'", expectedMsg, err.Error())
        }
    } else {
        t.Errorf("Expected error not nil")
    }


    expectedCode := http.StatusUnauthorized
    if code != expectedCode {
        t.Errorf("Expected %d, but got %d", expectedCode, code)
    }
}

func TestAuthenticateRequest_ValidAuthorizationHeader(t *testing.T) {
    requestHeader := http.Header{}
    // TODO: Very insecure test LOL
    requestHeader.Set("Authorization", "449a36b6689d841d7d27f31b4b7cc73a")

    code, err := authenticateRequest(requestHeader)

    if err != nil {
        t.Errorf("Error should be nil.")
    }

    expectedCode := 0
    if code != expectedCode {
        t.Errorf("Expected %d, but got %d", expectedCode, code)
    }
}