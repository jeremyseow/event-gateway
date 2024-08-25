package schema

import (
	"fmt"
	"strings"

	"github.com/jeremyseow/event-gateway/dto"
	"github.com/santhosh-tekuri/jsonschema/v6"
	"go.uber.org/zap"
)

const (
	logTag = "processor.schema"
)

type Validator struct {
	logger         *zap.Logger
	schemaCache    map[string]*jsonschema.Schema
	schemaCompiler *jsonschema.Compiler
}

func NewValidator(logger *zap.Logger) *Validator {
	schemaDefs := map[string]string{
		"ui_event": `{
			"type": "object",
			"properties": {
				"event_time": { "type": "integer" },
				"etype": { 
					"enum": ["click", "scroll", "log"]
				}
			},
			"required": ["event_time", "etype"]
		}`,
		"booking": `{
			"type": "object",
			"properties": {
				"booking_code": { 
					"type": "string",
					"pattern": "(?:^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[a-f0-9]{4}-[a-f0-9]{12}$)|(?:^0{8}-0{4}-0{4}-0{4}-0{12}$)" 
				},
				"price": { "type": "number" }
			},
			"required": ["booking_code", "price"]
		}`,
		"search": `{
			"type": "object",
			"properties": {
				"search_term": { 
					"type": "string",
					"minLength": 3,
					"maxLength": 20 
				}
			},
			"required": ["search_term"]
		}`,
	}

	schemaCache := make(map[string]*jsonschema.Schema)
	schemaCompiler := jsonschema.NewCompiler()

	for schemaName, schemaDef := range schemaDefs {
		unmarshalledSchema, err := jsonschema.UnmarshalJSON(strings.NewReader(schemaDef))
		if err != nil {
			logger.Error("error unmarshalling schema", zap.String("logTag", logTag), zap.String("function", "NewValidator"), zap.String("error", err.Error()))
			continue
		}

		if err := schemaCompiler.AddResource(schemaName, unmarshalledSchema); err != nil {
			logger.Error("error storing schema", zap.String("logTag", logTag), zap.String("function", "NewValidator"), zap.String("error", err.Error()))
			continue
		}

		compiledSchema, err := schemaCompiler.Compile(schemaName)
		if err != nil {
			logger.Error("error compiling schema", zap.String("logTag", logTag), zap.String("function", "NewValidator"), zap.String("error", err.Error()))
			continue
		}

		schemaCache[schemaName] = compiledSchema
	}

	return &Validator{
		logger:         logger,
		schemaCache:    schemaCache,
		schemaCompiler: schemaCompiler,
	}
}

func (v *Validator) ValidateEntity(entity dto.Entity) error {
	schemaName := entity.Schema
	schema, ok := v.schemaCache[schemaName]
	if !ok {
		return fmt.Errorf("schema not found: %s", schemaName)
	}

	return schema.Validate(entity.Parameters)
}
