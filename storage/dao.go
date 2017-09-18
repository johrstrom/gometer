package core

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/johrstrom/gometer/core"
)

// TestPlanDAO is a DAO for testplans
type TestPlanDAO struct {
}

// GetByName Gets a test plan from file by the name of the file
func (dao *TestPlanDAO) GetByName(name string) (*core.TestPlan, error) {
	data, err := ioutil.ReadFile(name + ".json")
	if err != nil {
		return nil, err
	}

	tp := &core.TestPlan{}

	err = json.Unmarshal(data, tp)

	return tp, err
}

// Save Save's a testplan to a file
func (dao *TestPlanDAO) Save(plan *core.TestPlan) error {

	if planData, err := json.MarshalIndent(plan, "", " "); err == nil {
		ioutil.WriteFile(plan.Name+".json", planData, 0644)
	} else {
		return err
	}

	return nil
}

// Delete delets a test plan
func (dao *TestPlanDAO) Delete(name string) error {
	return os.Remove(name + ".json")
}
