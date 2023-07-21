package service

import (
	"github.com/sarahrajabazdeh/DreamPilot/dreamerr"
)

func handleError(err error) error {
	if err != nil {
		return dreamerr.PropagateError(err, 2)
	}
	return nil
}
