package db

import (
	"fmt"
	"os"

	"github.com/deta/deta-go/deta"
	"github.com/deta/deta-go/service/base"
)

var Deta *deta.Deta
var BaseUser *base.Base

func init() {
	d, err := deta.New(deta.WithProjectKey(os.Getenv("DETA_COLLECTION_KEY")))
	if err != nil {
		fmt.Println("failed to init new Deta instance:", err)
		return
	}

	Deta = d

	user, err := base.New(d, "User")
	if err != nil {
		fmt.Println("failed to init new Base instance:", err)
		return
	}

	BaseUser = user
}
