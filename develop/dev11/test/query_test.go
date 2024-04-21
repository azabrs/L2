package main

import (
	"L2/develop/dev11/internal/models"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"
	"bytes"
	"fmt"
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
		q.Add("user_id", event.UserID)
		q.Add("date", event.Date)
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

		req, err := http.NewRequest(http.MethodGet, u.String(), bytes.NewBuffer(body))
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
			q.Add("user_id", event.UserID)
			q.Add("date", event.Date)
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
	
			req, err := http.NewRequest(http.MethodGet, u.String(), bytes.NewBuffer(body))
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

	})
}