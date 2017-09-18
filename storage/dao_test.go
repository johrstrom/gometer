package core

import (
	"testing"

	gmetertest "github.com/johrstrom/gometer/test"
	"github.com/stretchr/testify/assert"
)

func TestSave(test *testing.T) {
	dao := &TestPlanDAO{}
	tp := gmetertest.NewSimpleTestPlan()
	err := dao.Save(tp)
	assert.Nil(test, err, "Should not have gotten error but got ", err)

	tpFromFile, err := dao.GetByName(tp.Name)
	assert.Nil(test, err, "Should not have gotten error but got ", err)

	assert.Equal(test, tp, tpFromFile, "Test plan from file doesn't match what was saved")

	//err = dao.Delete(tp.Name)
	assert.Nil(test, err, "Should not have gotten error but got ", err)
}
