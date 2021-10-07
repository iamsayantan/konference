package util

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var generator *snowflake.Node

func init() {
	var err error
	generator, err = snowflake.NewNode(1)

	if err != nil {
		panic(fmt.Sprintf("error creating id generator: %s", err.Error()))
	}
}

// GenerateNewId generates an unique 64 bit id.
func GenerateNewId() int64 {
	return generator.Generate().Int64()
}