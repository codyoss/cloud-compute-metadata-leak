package ml

import (
	"testing"
	"time"

	"ml.test.com/leakcheck"

	_ "cloud.google.com/go/compute/metadata"
)

func TestML(t *testing.T) {

	defer leakcheck.Check(t)

	time.Sleep(1 * time.Second)
}
