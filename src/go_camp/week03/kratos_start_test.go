package go_camp_week03

import (
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func TestKratosStart(t *testing.T) {
	hs := http.NewServer()
	gs := grpc.NewServer()
	app := kratos.New(
		Name("kratos"),
		Version("v1.0.0"),
		Server(hs, gs),
	)
	time.AfterFunc(time.Second, func(){
		app.Stop()
	})
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
