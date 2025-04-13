package main

import (
	"oprc_core/src"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cf := src.Config{}
	cf.LoadConfig()
	t.Logf("%v", cf)
}
