package common

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HandleError(l *log.Logger, w http.ResponseWriter, err error, errCode int) {
	l.Println(err.Error())
	http.Error(w, err.Error(), errCode)
}

func FieldToInt(raw interface{}) (int, error) {
	if mmrStr, ok := raw.(string); ok {
		if mmr, err := strconv.Atoi(mmrStr); err == nil {
			return mmr, nil
		} else {
			return 0, fmt.Errorf("field is not a valid integer")
		}
	} else {
		return 0, fmt.Errorf("field is not a number or string")
	}
}
