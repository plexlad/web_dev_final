package lib

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

// TODO: Design better tests that check both valid and marks any other value
// as invalid

func TestSchemaJSONSerialization(t *testing.T) {
	schema := Schema{
		ID:          "schema-123",
		Version:     1,
		UserVersion: 2,
		Name:        "Test Schema",
		Description: "A test schema",
		CreatedAt:   time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
		UpdatedAt:   time.Date(2024, 1, 2, 12, 0, 0, 0, time.UTC),
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal schema: %v", err)
	}

	jsonStr := string(jsonData)
	t.Logf("Schema JSON:\n%s", jsonStr)

	// Check for correct JSON keys
	requiredKeys := []string{
		`"_id"`,
		`"version"`,
		`"user_version"`,
		`"name"`,
		`"description"`,
		`"created_at"`,
		`"updated_at"`,
	}

	for _, key := range requiredKeys {
		if !strings.Contains(jsonStr, key) {
			t.Errorf("JSON missing required key: %s", key)
		}
	}

	// Check for specific values
	if !strings.Contains(jsonStr, `"_id": "schema-123"`) {
		t.Error("JSON missing or incorrect _id field")
	}
	if !strings.Contains(jsonStr, `"version": 1`) {
		t.Error("JSON missing or incorrect version field")
	}
	if !strings.Contains(jsonStr, `"user_version": 2`) {
		t.Error("JSON missing or incorrect user_version field")
	}
	if !strings.Contains(jsonStr, `"name": "Test Schema"`) {
		t.Error("JSON missing or incorrect name field")
	}

	// Unmarshal back to verify round-trip
	var decoded Schema
	err = json.Unmarshal(jsonData, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal schema: %v", err)
	}

	// Verify key fields
	if decoded.ID != schema.ID {
		t.Errorf("ID mismatch: got %v, want %v", decoded.ID, schema.ID)
	}
	if decoded.Name != schema.Name {
		t.Errorf("Name mismatch: got %v, want %v", decoded.Name, schema.Name)
	}
	if decoded.Version != schema.Version {
		t.Errorf("Version mismatch: got %v, want %v", decoded.Version, schema.Version)
	}
	if decoded.UserVersion != schema.UserVersion {
		t.Errorf("UserVersion mismatch: got %v, want %v", decoded.UserVersion, schema.UserVersion)
	}
}

func TestSchemaUnmarshalFromJSON(t *testing.T) {
	// Test unmarshaling from a known-good JSON string
	apiJSON := `{
		"_id": "schema-456",
		"version": 3,
		"user_version": 1,
		"name": "API Schema",
		"description": "From API",
		"created_at": "2024-01-01T00:00:00Z",
		"updated_at": "2024-01-02T00:00:00Z"
	}`

	var schema Schema
	err := json.Unmarshal([]byte(apiJSON), &schema)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	// Verify all fields were correctly populated
	if schema.ID != "schema-456" {
		t.Errorf("Expected ID 'schema-456', got '%s'", schema.ID)
	}
	if schema.Version != 3 {
		t.Errorf("Expected Version 3, got %d", schema.Version)
	}
	if schema.UserVersion != 1 {
		t.Errorf("Expected UserVersion 1, got %d", schema.UserVersion)
	}
	if schema.Name != "API Schema" {
		t.Errorf("Expected Name 'API Schema', got '%s'", schema.Name)
	}
	if schema.Description != "From API" {
		t.Errorf("Expected Description 'From API', got '%s'", schema.Description)
	}
}

func TestInstanceJSONSerialization(t *testing.T) {
	instance := Instance{
		ID:     "instance-456",
		SchemaID: "schema-12",
//		Visualization: Visualization{
//			Name: "Main",
//			Type: VisDefault,
//		},
		UserID: "user-789",
		Name:   "My Character",
//		VariableValues: map[string]any{
//			"health": 100.0,
//			"name":   "Hero",
//		},
//		ActiveFeatures: []string{"combat"},
//		ActiveModules:  []string{"weapons"},
		CreatedAt:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:      time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(instance, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal instance: %v", err)
	}

	jsonStr := string(jsonData)
	t.Logf("Instance JSON:\n%s", jsonStr)

	// Check for correct JSON keys
	requiredKeys := []string{
		`"_id"`,
		`"schema_id"`,
		`"visualization"`,
		`"user_id"`,
		`"name"`,
		`"variable_values"`,
		`"active_features"`,
		`"active_modules"`,
		`"created_at"`,
		`"updated_at"`,
	}

	for _, key := range requiredKeys {
		if !strings.Contains(jsonStr, key) {
			t.Errorf("JSON missing required key: %s", key)
		}
	}

	// Check specific values
	if !strings.Contains(jsonStr, `"_id": "instance-456"`) {
		t.Error("JSON missing or incorrect _id field")
	}
	if !strings.Contains(jsonStr, `"user_id": "user-789"`) {
		t.Error("JSON missing or incorrect user_id field")
	}
	if !strings.Contains(jsonStr, `"name": "My Character"`) {
		t.Error("JSON missing or incorrect name field")
	}

	// Unmarshal back
	var decoded Instance
	err = json.Unmarshal(jsonData, &decoded)
	if err != nil {
		t.Fatalf("Failed to unmarshal instance: %v", err)
	}

	// Verify fields
	if decoded.ID != instance.ID {
		t.Errorf("ID mismatch: got %v, want %v", decoded.ID, instance.ID)
	}
	if decoded.UserID != instance.UserID {
		t.Errorf("UserID mismatch: got %v, want %v", decoded.UserID, instance.UserID)
	}
}

func TestInstanceUnmarshalFromJSON(t *testing.T) {
	apiJSON := `{
		"_id": "instance-789",
		"schema": {
			"_id": "schema-123",
			"version": 1,
			"user_version": 0,
			"name": "Test",
			"description": "",
			"created_at": "2024-01-01T00:00:00Z",
			"updated_at": "2024-01-01T00:00:00Z"
		},
		"visualization": {
			"name": "Main",
			"type": "default",
			"child_visualizations": null
		},
		"user_id": "user-123",
		"name": "Test Instance",
		"variable_values": {
			"health": 50
		},
		"active_features": ["feature1"],
		"active_modules": ["module1"],
		"created_at": "2024-01-01T00:00:00Z",
		"updated_at": "2024-01-01T00:00:00Z"
	}`

	var instance Instance
	err := json.Unmarshal([]byte(apiJSON), &instance)
	if err != nil {
		t.Fatalf("Failed to unmarshal: %v", err)
	}

	if instance.ID != "instance-789" {
		t.Errorf("Expected ID 'instance-789', got '%s'", instance.ID)
	}
	if instance.UserID != "user-123" {
		t.Errorf("Expected UserID 'user-123', got '%s'", instance.UserID)
	}
	if instance.Name != "Test Instance" {
		t.Errorf("Expected Name 'Test Instance', got '%s'", instance.Name)
	}
//	if len(instance.ActiveFeatures) != 1 || instance.ActiveFeatures[0] != "feature1" {
//		t.Errorf("ActiveFeatures not correctly unmarshaled: %v", instance.ActiveFeatures)
//	}
}

//func TestVariableTypes(t *testing.T) {
//	tests := []struct {
//		name         string
//		variable     Variable
//		expectedKeys []string
//	}{
//		{
//			name: "Number variable",
//			variable: Variable{
//				Type:    TypeNumber,
//				Default: 42.0,
//			},
//			expectedKeys: []string{`"type"`, `"number"`, `"default"`},
//		},
//		{
//			name: "String variable",
//			variable: Variable{
//				Type:    TypeString,
//				Default: "test",
//			},
//			expectedKeys: []string{`"type"`, `"string"`, `"default"`},
//		},
//		{
//			name: "Boolean variable",
//			variable: Variable{
//				Type:    TypeBoolean,
//				Default: true,
//			},
//			expectedKeys: []string{`"type"`, `"boolean"`, `"default"`},
//		},
//		{
//			name: "Enum variable",
//			variable: Variable{
//				Type:    TypeEnum,
//				Default: "option1",
//				Options: []string{"option1", "option2", "option3"},
//			},
//			expectedKeys: []string{`"type"`, `"enum"`, `"options"`},
//		},
//		{
//			name: "Array variable",
//			variable: Variable{
//				Type: TypeArray,
//				Items: &Variable{
//					Type: TypeNumber,
//				},
//			},
//			expectedKeys: []string{`"type"`, `"array"`, `"items"`},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			jsonData, err := json.Marshal(tt.variable)
//			if err != nil {
//				t.Fatalf("Failed to marshal variable: %v", err)
//			}
//
//			jsonStr := string(jsonData)
//			t.Logf("%s JSON: %s", tt.name, jsonStr)
//
//			// Check for expected keys
//			for _, key := range tt.expectedKeys {
//				if !strings.Contains(jsonStr, key) {
//					t.Errorf("JSON missing expected key: %s", key)
//				}
//			}
//
//			// Round-trip test
//			var decoded Variable
//			err = json.Unmarshal(jsonData, &decoded)
//			if err != nil {
//				t.Fatalf("Failed to unmarshal variable: %v", err)
//			}
//
//			if decoded.Type != tt.variable.Type {
//				t.Errorf("Type mismatch: got %v, want %v", decoded.Type, tt.variable.Type)
//			}
//		})
//	}
//}
//
//func TestVariableUnmarshalFromJSON(t *testing.T) {
//	tests := []struct {
//		name     string
//		json     string
//		validate func(*testing.T, Variable)
//	}{
//		{
//			name: "Number with min/max",
//			json: `{"type":"number","default":50,"min":0,"max":100}`,
//			validate: func(t *testing.T, v Variable) {
//				if v.Type != TypeNumber {
//					t.Errorf("Expected type number, got %v", v.Type)
//				}
//				if v.Min == nil || *v.Min != 0 {
//					t.Error("Min not correctly unmarshaled")
//				}
//				if v.Max == nil || *v.Max != 100 {
//					t.Error("Max not correctly unmarshaled")
//				}
//			},
//		},
//		{
//			name: "Enum with options",
//			json: `{"type":"enum","options":["red","green","blue"]}`,
//			validate: func(t *testing.T, v Variable) {
//				if v.Type != TypeEnum {
//					t.Errorf("Expected type enum, got %v", v.Type)
//				}
//				if len(v.Options) != 3 {
//					t.Errorf("Expected 3 options, got %d", len(v.Options))
//				}
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			var v Variable
//			err := json.Unmarshal([]byte(tt.json), &v)
//			if err != nil {
//				t.Fatalf("Failed to unmarshal: %v", err)
//			}
//			tt.validate(t, v)
//		})
//	}
//}
//
//func TestPropertyJSONSerialization(t *testing.T) {
//	prop := Property{
//		Formula: "health * 1.5",
//		Format:  FormatRound,
//	}
//
//	jsonData, err := json.Marshal(prop)
//	if err != nil {
//		t.Fatalf("Failed to marshal property: %v", err)
//	}
//
//	jsonStr := string(jsonData)
//	t.Logf("Property JSON: %s", jsonStr)
//
//	// Check for correct keys
//	if !strings.Contains(jsonStr, `"formula"`) {
//		t.Error("JSON missing 'formula' key")
//	}
//	if !strings.Contains(jsonStr, `"format"`) {
//		t.Error("JSON missing 'format' key")
//	}
//	if !strings.Contains(jsonStr, `"round"`) {
//		t.Error("JSON missing 'round' format value")
//	}
//
//	// Round-trip test
//	var decoded Property
//	err = json.Unmarshal(jsonData, &decoded)
//	if err != nil {
//		t.Fatalf("Failed to unmarshal: %v", err)
//	}
//
//	if decoded.Formula != prop.Formula {
//		t.Errorf("Formula mismatch: got %v, want %v", decoded.Formula, prop.Formula)
//	}
//	if decoded.Format != prop.Format {
//		t.Errorf("Format mismatch: got %v, want %v", decoded.Format, prop.Format)
//	}
//}
//
//func TestFeatureJSONSerialization(t *testing.T) {
//	feature := Feature{
//		Name:        "Combat System",
//		Description: "Adds combat capabilities",
//		AddsModules: []string{"weapons", "armor"},
//	}
//
//	jsonData, err := json.Marshal(feature)
//	if err != nil {
//		t.Fatalf("Failed to marshal feature: %v", err)
//	}
//
//	jsonStr := string(jsonData)
//	t.Logf("Feature JSON: %s", jsonStr)
//
//	// Check for correct keys
//	requiredKeys := []string{`"name"`, `"description"`, `"adds_modules"`}
//	for _, key := range requiredKeys {
//		if !strings.Contains(jsonStr, key) {
//			t.Errorf("JSON missing key: %s", key)
//		}
//	}
//
//	// Round-trip test
//	var decoded Feature
//	err = json.Unmarshal(jsonData, &decoded)
//	if err != nil {
//		t.Fatalf("Failed to unmarshal: %v", err)
//	}
//
//	if decoded.Name != feature.Name {
//		t.Errorf("Name mismatch: got %v, want %v", decoded.Name, feature.Name)
//	}
//	if len(decoded.AddsModules) != len(feature.AddsModules) {
//		t.Errorf("AddsModules length mismatch: got %v, want %v", len(decoded.AddsModules), len(feature.AddsModules))
//	}
//}
//
//func TestModuleJSONSerialization(t *testing.T) {
//	module := Module{
//		Name:        "Weapons",
//		Description: "Weapon system",
//		AddsVariables: map[string]Variable{
//			"attack": {
//				Type:    TypeNumber,
//				Default: 10.0,
//			},
//		},
//		AddsProperties: map[string]Property{
//			"total_damage": {
//				Formula: "attack * 2",
//				Format:  FormatRound,
//			},
//		},
//	}
//
//	jsonData, err := json.Marshal(module)
//	if err != nil {
//		t.Fatalf("Failed to marshal module: %v", err)
//	}
//
//	jsonStr := string(jsonData)
//	t.Logf("Module JSON: %s", jsonStr)
//
//	// Check for correct keys
//	requiredKeys := []string{`"name"`, `"adds_variables"`, `"adds_display_values"`}
//	for _, key := range requiredKeys {
//		if !strings.Contains(jsonStr, key) {
//			t.Errorf("JSON missing key: %s", key)
//		}
//	}
//
//	// Round-trip test
//	var decoded Module
//	err = json.Unmarshal(jsonData, &decoded)
//	if err != nil {
//		t.Fatalf("Failed to unmarshal: %v", err)
//	}
//
//	if decoded.Name != module.Name {
//		t.Errorf("Name mismatch: got %v, want %v", decoded.Name, module.Name)
//	}
//}
//
//func TestVisualizationJSONSerialization(t *testing.T) {
//	vis := Visualization{
//		Name: "Main Layout",
//		Type: VisGrid,
//		ChildVisualizations: []Visualization{
//			{
//				Name: "Stats Section",
//				Type: VisCard,
//			},
//			{
//				Name: "Inventory",
//				Type: VisAccordion,
//			},
//		},
//		Config: json.RawMessage(`{"columns": 2, "gap": 16}`),
//	}
//
//	jsonData, err := json.MarshalIndent(vis, "", "  ")
//	if err != nil {
//		t.Fatalf("Failed to marshal visualization: %v", err)
//	}
//
//	jsonStr := string(jsonData)
//	t.Logf("Visualization JSON:\n%s", jsonStr)
//
//	// Check for correct keys
//	requiredKeys := []string{`"name"`, `"type"`, `"child_visualizations"`}
//	for _, key := range requiredKeys {
//		if !strings.Contains(jsonStr, key) {
//			t.Errorf("JSON missing key: %s", key)
//		}
//	}
//
//	// Check for specific values
//	if !strings.Contains(jsonStr, `"grid"`) {
//		t.Error("JSON missing 'grid' type value")
//	}
//
//	// Round-trip test
//	var decoded Visualization
//	err = json.Unmarshal(jsonData, &decoded)
//	if err != nil {
//		t.Fatalf("Failed to unmarshal visualization: %v", err)
//	}
//
//	if decoded.Name != vis.Name {
//		t.Errorf("Name mismatch: got %v, want %v", decoded.Name, vis.Name)
//	}
//	if decoded.Type != vis.Type {
//		t.Errorf("Type mismatch: got %v, want %v", decoded.Type, vis.Type)
//	}
//	if len(decoded.ChildVisualizations) != len(vis.ChildVisualizations) {
//		t.Errorf("ChildVisualizations length mismatch: got %v, want %v", 
//			len(decoded.ChildVisualizations), len(vis.ChildVisualizations))
//	}
//}
//
//func TestInitializationJSONSerialization(t *testing.T) {
//	init := Initialization{
//		Steps: []InitializationStep{
//			{
//				Title: "Character Setup",
//				Fields: []Field{
//					{
//						Prompt:       "Enter your name",
//						VariableName: "character_name",
//						Formula:      "input",
//					},
//					{
//						Prompt:       "Choose your class",
//						VariableName: "class",
//						Formula:      "input",
//					},
//				},
//			},
//		},
//	}
//
//	jsonData, err := json.MarshalIndent(init, "", "  ")
//	if err != nil {
//		t.Fatalf("Failed to marshal initialization: %v", err)
//	}
//
//	jsonStr := string(jsonData)
//	t.Logf("Initialization JSON:\n%s", jsonStr)
//
//	// Check for correct keys
//	requiredKeys := []string{`"steps"`, `"title"`, `"fields"`, `"prompt"`, `"variable_name"`, `"formula"`}
//	for _, key := range requiredKeys {
//		if !strings.Contains(jsonStr, key) {
//			t.Errorf("JSON missing key: %s", key)
//		}
//	}
//
//	// Round-trip test
//	var decoded Initialization
//	err = json.Unmarshal(jsonData, &decoded)
//	if err != nil {
//		t.Fatalf("Failed to unmarshal: %v", err)
//	}
//
//	if len(decoded.Steps) != len(init.Steps) {
//		t.Errorf("Steps length mismatch: got %v, want %v", len(decoded.Steps), len(init.Steps))
//	}
//}

// Commented out - helper methods not yet implemented
// func TestHelperMethods(t *testing.T) { ... }

// Commented out schema- instance methods not yet implemented
// func TestInstanceMethods(t *testing.T) { ... }
