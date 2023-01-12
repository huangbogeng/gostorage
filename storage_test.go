package gostorage

import (
	"github.com/link-yundi/ytools/ylog"
	"testing"
)

/**
@Since 2022-12-02 13:34
@Author: Huang
@Description test storage
**/

func TestMapL1(t *testing.T) {

	m := NewMapL1()
	m.Set("1", 100)
	ylog.Info(m.Map)
}

func TestMapL2_Size(t *testing.T) {
	m := NewMapL2()
	m.Set("1", "2", 1000)

	ylog.Info(m.GetL2("1", "2"))

	m.Delete("1", "2")

	ylog.Info(m.GetL2("1", "2"))
}
