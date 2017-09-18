package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/johrstrom/gometer/core"
	"github.com/johrstrom/gometer/samplers"
)

func main() {
	plan := core.NewTestPlan("temp")
	group := plan.AddThreadGroup()

	samplerNode := core.NewTestElementNode(samplers.DefaultHTTPSampler())

	group.AddNodeUnder(samplerNode)

	if planData, err := json.MarshalIndent(plan, "", " "); err == nil {
		ioutil.WriteFile("temp.json", planData, 0644)
	} else {
		fmt.Println("Error occured: ", err)
	}

}
