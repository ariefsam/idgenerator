package idgenerator

import "github.com/teris-io/shortid"

var InstanceNumber uint8 = 0
var sid *shortid.Shortid

func init() {
	var err error
	sid, err = shortid.New(InstanceNumber, shortid.DefaultABC, 23)
	if err != nil {
		panic(err)
	}
}

func Generate() string {
	// then either:
	uid, _ := sid.Generate()
	return uid
}
