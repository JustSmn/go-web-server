package routes

import (
	"testing"
)

func TestRouterCreation(t *testing.T) {
	router := Setup(nil)

	if router == nil {
		t.Error("Router should not be nil")
	}
}
