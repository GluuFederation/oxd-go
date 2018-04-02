//
//  Copyright Sagiton
//  Author: Michał Kępkowski
//  Date: 02/01/17
//
package utils

import (
	"github.com/juju/loggo"
)



func CheckError(prefix string, message string, err error) {
	if err != nil {
		loggo.GetLogger("prefix").Errorf(message,err.Error())
		//os.Exit(1)
	}
}
