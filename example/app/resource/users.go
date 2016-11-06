package resource

import (
	"log"

	"github.com/raframework/rago/rahttp"
)

type Users struct {
}

func (u *Users) Create(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Users.Create...")
}

func (u *Users) Update(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Users.Update...")
}

func (u *Users) Get(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Users.Get...")
}

func (u *Users) Delete(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Users.Delete...")
}

func (u *Users) List(request *rahttp.Request, response *rahttp.Response) {
	log.Println("example: Users.List...")
}
