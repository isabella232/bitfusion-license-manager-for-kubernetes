// vim: ts=4 sw=4 noexpandtab
//
// Copyright (c) 2015-2019 VMware, Inc. All Rights Reserved.
// Author: Tom Hite (thite@vmware.com)
//
// SPDX-License-Identifier: https://spdx.org/licenses/MIT.html
package template

import (
	"encoding/json"
	"log"
)

// Print out JSON errors on the log.
func jsonError(err error, jsonData string) {
	if err != nil {
		log.Printf("%T\n%s\n%#v\n", err, err, err)
		switch v := err.(type) {
		case *json.SyntaxError:
			log.Println(string(jsonData[v.Offset-40 : v.Offset]))
		}
	}
}

// Print out generic errors on the log.
func perror(err error) {
	if err != nil {
		panic(err)
	}
}
