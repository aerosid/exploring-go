package codepen

import (
	"fmt"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	status, _, _, _ := NewClient().SetUrl("http://www.ifonepay.com").Go()
	if status != 200 {
		t.Log("status: expected=200 actual=" + fmt.Sprint(status))
	} else {
		t.Log("success")
	}
}

func TestServer(t *testing.T) {
	srv := NewServer().Start(false)
	<-time.After(5 * time.Second)
	status, _, body, err := NewClient().SetUrl("http://localhost:8000").Go()
	if status == 200 {
		t.Log("Response: " + body)
	} else {
		t.Log("Error: " + err.Error())
	}
	srv.Stop()

}
