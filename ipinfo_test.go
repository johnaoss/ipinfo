package ipinfo_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/johnaoss/ipinfo"
)

var token = os.Getenv("IPINFO_TOKEN")

// TestGetRequest asserts we can get the IP info for Google's DNS servers.
func TestGetRequest(t *testing.T) {
	client := ipinfo.NewClient(token)

	resp, err := client.Info("8.8.8.8")
	if err != nil {
		t.Fatalf("Failed to get client info: %v", err)
	}

	if resp.IP != "8.8.8.8" {
		t.Errorf("Expected IP to be 8.8.8.8, instead given: %s", resp.IP)
	}

	t.Logf("%+v\n", resp)
}

// TestUnauthorized asserts we can get the info from the unauthorized API.
func TestUnauthorized(t *testing.T) {
	resp, err := ipinfo.Info("8.8.8.8")
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	t.Logf("%+v\n", resp)
}

func ExampleInfo() {
	resp, _ := ipinfo.Info("8.8.8.8")
	fmt.Println(resp.City)
	// Output: Mountain View
}
