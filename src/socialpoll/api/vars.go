package main

import (
	"sync"
	"net/http"
)

var (
	varsLock sync.RWMutex
	vars map[*http.Request]map[string]interface{}
)

func openVars(r *http.Request)  {
	varsLock.Lock()
	if vars == nil {
		vars = map[*http.Request]map[string]interface{}{}
	}
	vars[r] = map[string]interface{}{}
	varsLock.Unlock()
}






