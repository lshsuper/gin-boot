package core

import "net/http"

type SuperBind struct {

}

func (b *SuperBind)Name()  string{

	return ""
}

func (b *SuperBind)Bind(req *http.Request,p interface{}) error {
   return nil
}
