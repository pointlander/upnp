package upnp

import (
	"testing"
)

func TestFuctions(t *testing.T) {
	u, err := NewUPNP()
	if err != nil {
		t.Error("Unexpected set up error: ", err)
		return
	}

	err = u.AddPortMapping(5050, 5050, "UDP")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}

	err = u.AddPortMapping(1337, 1337, "UDP")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}

	e, err := u.GetPortMappings()
	if err != nil {
		t.Error("Unexpected error: ", err)
	}
	if len(e) != 2 {
		t.Error("there should be 2 entries")
	}
	t.Log(e)

	err = u.DelPortMapping(5050, "UDP")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}

	err = u.DelPortMapping(1337, "UDP")
	if err != nil {
		t.Error("Unexpected error: ", err)
	}

	ip, _ := u.ExternalIPAddress()
	if ip == nil {
		t.Error("Missing external IP")
	}

}
