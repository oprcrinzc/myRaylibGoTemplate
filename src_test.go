package main

import (
	t "oprc_core/src/types"
	"testing"
)

func TestLoadConfig(T *testing.T) {
	cf := t.Config{}
	cf.LoadConfig()
	T.Logf("%v", cf)
}
