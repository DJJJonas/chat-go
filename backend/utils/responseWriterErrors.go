// This file has functions to Write Erros (WE) to Writer interfaces
package utils

import (
	"fmt"
	"net/http"
)

func WEUpgradeRequired(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusUpgradeRequired)
	fmt.Fprintf(w, "An error ocurred while trying to upgrade: %s", e.Error())
}
