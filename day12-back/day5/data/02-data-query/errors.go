package data_project_seq

import "github.com/labstack/gommon/log"

func HandleError(err error, reason string)  {
	log.Errorf("%v, %s\n", err, reason)
}