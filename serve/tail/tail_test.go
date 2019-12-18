package tail

import (
	"github.com/woshihot/go-lib/serve/mux"
	"testing"
)

func TestRest(t *testing.T) {

	mux.Init("8081", CreateTailHandler("c:\\box\\data\\logs\\agentServer.log", "/tail", "/follow", "g"))
	select {}
}
