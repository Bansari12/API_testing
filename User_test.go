package main

import (
	"testing"
	"net/http"	
	"https://github.com/gavv/httpexpect"

	) 
func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users?ID", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"ID":1,"Name":"Bansari","Address":"xyz Road","MobileNumber":999999999,"City":"Rajkot","Email":"bansari@improwised.com","MobileNumber":0987654321},{"ID":2,"Name":"Divya","Address":"ABC Road","MobileNumber":888888888,"City";"Ahmedabad","Email":"divya@improwised.com"}]`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}