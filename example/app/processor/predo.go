package processor

import (
	"log"

	"github.com/raframework/rago/rahttp"
)

type Predo struct {
}

func (p *Predo) Process(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example.app.processor: predo processing")

	response.WithHeader("Content-Type", "application/json;charset=utf-8")
}
