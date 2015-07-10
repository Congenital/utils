package utils

import (
	"github.com/Congenital/log/v0.2/log"
	"testing"
)

func TestFuzzyFile(t *testing.T) {
	log.Info(FuzzyQuery("./", "utils"))
}
