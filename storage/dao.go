package core

import (
	"bytes"
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

	//if planData, err := json.Marshal(plan); err == nil {
	if planData, err := dao.MarshalPlan(plan); err == nil {
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

func (dao *TestPlanDAO) MarshalPlan(plan *core.TestPlan) ([]byte, error) {
	buffer := bytes.NewBufferString("{\"name\":\"")
	buffer.WriteString(plan.Name + "\",\"root\":")

	root := plan.GetRootNode()
	buffer.Write(marshalNode(root))

	//for idx, item := range root.SubTree {

	//}

	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func marshalProperties(props core.Properties) []byte {
	buffer := bytes.NewBufferString("\"properties\":")
	propsData, _ := json.Marshal(props)

	buffer.Write(propsData)
	//buffer.WriteString("}")
	return buffer.Bytes()
}

func marshalNode(node *core.TestElementNode) []byte {
	buffer := bytes.NewBufferString("{\"node\": {")
	buffer.Write(marshalProperties(node.Properties()))
	buffer.WriteString(", \"subTree\":[")

	for _, subNode := range node.SubTree {
		buffer.Write(marshalNode(subNode))
	}

	buffer.WriteString("]}}")
	return buffer.Bytes()
}
