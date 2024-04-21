package main

import (
	"L2/develop/dev11/internal/models"
	"L2/develop/dev11/internal/server"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestService(t *testing.T){
	var events []models.EventCreate
	filename, err := os.Open("test_data.json")
	if err != nil{
		t.Errorf("cannot read testing data: %v", err)
	}
	defer filename.Close()
	buf, err := io.ReadAll(filename)
	if err != nil{
		t.Errorf("cannot read testing data: %v", err)
	}
	if err := json.Unmarshal(buf, &events); err != nil{
		t.Errorf("cannot read testing data: %v", err)
	}
	for _, event :=  range events{
		c := http.Client{}
		q := url.Values{}
		u := url.URL{
			Scheme: "http",
			Host: "localhost:8080",
			Path: "/create_event",
			RawQuery: q.Encode(),
		}
		body, err := json.Marshal(event)
		if err != nil{
			t.Errorf("%v", err)
		}

		req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
		if err != nil{
			t.Errorf("%v", err)
		}
		resp, err := c.Do(req)
		if err != nil{
			t.Errorf("%v", err)
		}
		if resp.StatusCode != http.StatusCreated{
			t.Errorf("Incorrect Status code")
		}
		for{
		    bs := make([]byte, 1014)
		    n, err := resp.Body.Read(bs)
		    fmt.Println(string(bs[:n]))
		    if n == 0 || err != nil{
		        break
		    }
		}
	}
	t.Run("update test", func(t *testing.T) {
		for _, event :=  range events{
			c := http.Client{}
			q := url.Values{}
			u := url.URL{
				Scheme: "http",
				Host: "localhost:8080",
				Path: "/update_event",
				RawQuery: q.Encode(),
			}
			event.DataEvent.Title += " updated" 
			body, err := json.Marshal(event)
			if err != nil{
				t.Errorf("%v", err)
			}
	
			req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
			if err != nil{
				t.Errorf("%v", err)
			}
			resp, err := c.Do(req)
			if err != nil{
				t.Errorf("%v", err)
			}
			if resp.StatusCode != http.StatusOK{
				t.Errorf("Incorrect Status code")
			}
			for{
				bs := make([]byte, 1014)
				n, err := resp.Body.Read(bs)
				fmt.Println(string(bs[:n]))
				if n == 0 || err != nil{
					break
				}
			}
		}

	})

	t.Run("event for day", func(t *testing.T) {
		for _, event :=  range events{
			
			c := http.Client{}
			q := url.Values{}
			q.Add("user_id", event.UserID)
			q.Add("date", event.Date)
			u := url.URL{
				Scheme: "http",
				Host: "localhost:8080",
				Path: "/events_for_day",
				RawQuery: q.Encode(),
			}
			//event.DataEvent.Title += "updated" 
			if err != nil{
				t.Errorf("%v", err)
			}
	
			req, err := http.NewRequest(http.MethodGet, u.String(), nil)
			if err != nil{
				t.Errorf("%v", err)
			}
			resp, err := c.Do(req)
			if err != nil{
				t.Errorf("%v", err)
			}
			if resp.StatusCode != http.StatusOK{
				t.Errorf("Incorrect Status code")
			}
			
			bs := make([]byte, 1014)
			n, err := resp.Body.Read(bs)
			var data server.ResultResponse
			err = json.Unmarshal(bs[0:n], &data)
			if err != nil{
				t.Error(err)
			}
			count := 0
			for _, e := range(data.Result){
				if compare(e, event.DataEvent){
					count += 1
				}
			}
			if count == 0{
				t.Errorf("responce data not equal request data")
			}
			
			if n == 0 || err != nil{
				break
			}
			
		}

	})


	t.Run("event for month", func(t *testing.T) {
			event := events[0]
			c := http.Client{}
			q := url.Values{}
			q.Add("user_id", event.UserID)
			q.Add("date", event.Date)
			u := url.URL{
				Scheme: "http",
				Host: "localhost:8080",
				Path: "/events_for_month",
				RawQuery: q.Encode(),
			}
			if err != nil{
				t.Errorf("%v", err)
			}
	
			req, err := http.NewRequest(http.MethodGet, u.String(), nil)
			if err != nil{
				t.Errorf("%v", err)
			}
			resp, err := c.Do(req)
			if err != nil{
				t.Errorf("%v", err)
			}
			if resp.StatusCode != http.StatusOK{
				t.Errorf("Incorrect Status code")
			}
			
			bs := make([]byte, 1014)
			n, err := resp.Body.Read(bs)
			var data server.ResultResponse
			err = json.Unmarshal(bs[0:n], &data)
			if err != nil{
				t.Error(err)
			}
			count := 0
			for _, e := range(data.Result){
				if compare(e, event.DataEvent){
					count += 1
				}
			}
			if count == 0{
				t.Errorf("responce data not equal request data")
			}
			

	})

	t.Run("event for week", func(t *testing.T) {
		event := events[0]
		c := http.Client{}
		q := url.Values{}
		q.Add("user_id", event.UserID)
		q.Add("date", event.Date)
		u := url.URL{
			Scheme: "http",
			Host: "localhost:8080",
			Path: "/events_for_week",
			RawQuery: q.Encode(),
		}
		if err != nil{
			t.Errorf("%v", err)
		}

		req, err := http.NewRequest(http.MethodGet, u.String(), nil)
		if err != nil{
			t.Errorf("%v", err)
		}
		resp, err := c.Do(req)
		if err != nil{
			t.Errorf("%v", err)
		}
		if resp.StatusCode != http.StatusOK{
			t.Errorf("Incorrect Status code")
		}
		
		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		var data server.ResultResponse
		err = json.Unmarshal(bs[0:n], &data)
		if err != nil{
			t.Error(err)
		}
		count := 0
		for _, e := range(data.Result){
			if compare(e, event.DataEvent){
				count += 1
			}
		}
		if count == 0{
			t.Errorf("responce data not equal request data")
		}
		

})

	t.Run("delete test", func(t *testing.T) {
		for _, event :=  range events{
			c := http.Client{}
			q := url.Values{}
			u := url.URL{
				Scheme: "http",
				Host: "localhost:8080",
				Path: "/delete_event",
				RawQuery: q.Encode(),
			}
			//event.DataEvent.Title += "updated" 
			body, err := json.Marshal(event)
			if err != nil{
				t.Errorf("%v", err)
			}
	
			req, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(body))
			if err != nil{
				t.Errorf("%v", err)
			}
			resp, err := c.Do(req)
			if err != nil{
				t.Errorf("%v", err)
			}
			if resp.StatusCode != http.StatusNoContent{
				t.Errorf("Incorrect Status code")
			}
			for{
				bs := make([]byte, 1014)
				n, err := resp.Body.Read(bs)
				fmt.Println(string(bs[:n]))
				if n == 0 || err != nil{
					break
				}
			}
		}

	})


}


func compare(event1, event2 models.EventData) bool {
	if event1.Title != event2.Title + " updated" {
		return false
	}

	if event1.Description != event2.Description {
		return false
	}

	return true
}
