package model

// ThePackage is root of application definition
type ThePackage struct {
	ApplicationName string      ``                   // name of application
	PackagePath     string      `yaml:"packagePath"` // golang path of application
	Entities        []TheEntity `yaml:"entities"`    // list of entity used in this apps
	Enums           []TheEnum   `yaml:"enums"`       // list of enum used in this apps
}

type (

	// TheEntity is component that define the blue print of object
	TheEntity struct {
		Name            string     `yaml:"name"`      // MANDATORY. name of the entity
		Fields          []TheField `yaml:"fields"`    // MANDATORY. all field under the entity
		TableName       string     `yaml:"tableName"` // OPTIONAL. if empty then the name will be as same as entity name
		HasAutocomplete bool       ``                 // internal used purposed
		HasEnum         bool       ``                 // internal used purposed
	}

	// TheField is
	TheField struct {
		Name            string         `yaml:"name"`            // MANDATORY. name of the field
		DataType        string         `yaml:"dataType"`        // MANDATORY. golang type like int, string, float64, bool, time.Time
		EnumReference   string         `yaml:"enumReference"`   // MANDATORY WITH CONDITION dataType is enum. Refered to enums field
		EntityReference string         `yaml:"entityReference"` // MANDATORY WITH CONDITION dataType is entity. Refered to other entity name
		EntityField     string         `yaml:"entityField"`     // MANDATORY WITH CONDITION dataType is entity. Refered to other entity field
		DefaultValue    string         `yaml:"defaultValue"`    // value of the field
		Sortable        string         `yaml:"sortable"`        // not implemented yet
		Filterable      string         `yaml:"filterable"`      // not implemented yet
		Regex           string         `yaml:"regex"`           // not implemented yet
		Required        bool           `yaml:"required"`        // not implemented yet
		EnumValues      []TextAndValue ``                       // internal used purposed
	}
)

type (

	// TheEnum is static choice of value
	TheEnum struct {
		Name   string         `yaml:"name"`   // MANDATORY the name of enum
		Values []TextAndValue `yaml:"values"` // the value contained in this enum
	}

	// TextAndValue is used in enum. If value is not defined then the value value's will be as same as text
	TextAndValue struct {
		Text  string `yaml:"text"`  // MANDATORY
		Value string `yaml:"value"` // OPTIONAL. If empty then the value is same as Text
	}
)
