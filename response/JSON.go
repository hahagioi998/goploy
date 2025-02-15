// Copyright 2022 The Goploy Authors. All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package response

import (
	"encoding/json"
	"net/http"
)

type JSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// response code
const (
	Pass             = 0
	Deny             = 1
	Error            = 2
	AccountDisabled  = 10000
	IllegalRequest   = 10001
	NamespaceInvalid = 10002
	IllegalParam     = 10003
	LoginExpired     = 10086
)

//JSON response
func (j JSON) Write(w http.ResponseWriter) error {
	return json.NewEncoder(w).Encode(j)
}
