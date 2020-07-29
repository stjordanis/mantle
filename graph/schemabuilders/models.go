package schemabuilders

import (
	"github.com/graphql-go/graphql"
	"github.com/terra-project/mantle/db/kvindex"
	"github.com/terra-project/mantle/graph"
	"github.com/terra-project/mantle/graph/generate"
	"github.com/terra-project/mantle/types"
)

func CreateModelSchemaBuilder(models ...types.ModelType) graph.SchemaBuilder {
	return func(fields *graphql.Fields) error {
		// handle module registration
		for _, model := range models {
			entityName, entityRoot := generate.GenerateGraphResolver(model)
			entityRoot.Args = generate.GenerateArgument(kvindex.NewKVIndex(model))
			(*fields)[entityName] = entityRoot
		}

		return nil
	}
}