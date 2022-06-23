package utilities

import (
	"fmt"
	"reflect"
	"strings"
)

type EQUALS struct {
	Field string
	Value interface{}
}

func (equals EQUALS) BuildConditions() string {
	return fmt.Sprintf("`%s` = '%s'", equals.Field, equals.Value.(string))
}

type IS struct {
	Field string
	Value interface{}
}

func (is IS) BuildConditions() string {
	return fmt.Sprintf("`%s` IS %s", is.Field, is.Value.(string))
}

type LIKE struct {
	Field string
	Value interface{}
}

func (like LIKE) BuildConditions() string {
	return fmt.Sprintf("`%s` LIKE '%%%s%%'", like.Field, like.Value.(string))
}

type NOTEQUALS struct {
	Field string
	Value interface{}
}

func (notequals NOTEQUALS) BuildConditions() string {
	return fmt.Sprintf("`%s` <> '%s'", notequals.Field, notequals.Value.(string))
}

type IN struct {
	Field string
	Value string
}

func (in IN) BuildConditions() string {
	return fmt.Sprintf("`%s` IN (%s)", in.Field, in.Value)
}

type GREATER struct {
	Field string
	Value interface{}
}

func (greater GREATER) BuildConditions() string {
	return fmt.Sprintf("`%s` > '%s'", greater.Field, greater.Value.(string))
}

type LESS struct {
	Field string
	Value interface{}
}

func (less LESS) BuildConditions() string {
	return fmt.Sprintf("`%s` < '%s'", less.Field, less.Value.(string))
}

type GREATEREQUAL struct {
	Field string
	Value interface{}
}

func (greater GREATEREQUAL) BuildConditions() string {
	return fmt.Sprintf("`%s` >= '%s'", greater.Field, greater.Value.(string))
}

type LESSEQUAL struct {
	Field string
	Value interface{}
}

func (less LESSEQUAL) BuildConditions() string {
	return fmt.Sprintf("`%s` <= '%s'", less.Field, less.Value.(string))
}

type BETWEEN struct {
	Field       string
	LowerValue  interface{}
	HigherValue interface{}
}

func (between BETWEEN) BuildConditions() string {
	return fmt.Sprintf("`%s` BETWEEN '%s' AND '%s'", between.Field, between.LowerValue.(string), between.HigherValue.(string))
}

type AND struct {
	Conditions []interface{}
}

func (and AND) BuildConditions() string {
	conditions := ""
	conditions = getCondition(reflect.ValueOf(and.Conditions[0]).Type().Name(), and.Conditions[0])
	for index, condition := range and.Conditions {
		if index != 0 {
			conditions = fmt.Sprintf("%s AND %s", conditions, getCondition(reflect.ValueOf(condition).Type().Name(), condition))
		}
	}
	return conditions
}

type OR struct {
	Conditions []interface{}
}

func (or OR) BuildConditions() string {
	conditions := ""
	conditions = getCondition(reflect.ValueOf(or.Conditions[0]).Type().Name(), or.Conditions[0])
	for index, condition := range or.Conditions {
		if index != 0 {
			conditions = fmt.Sprintf("%s OR %s", conditions, getCondition(reflect.ValueOf(condition).Type().Name(), condition))
		}
	}
	return conditions
}

func BuildQuery(cond interface{}) interface{} {
	for key, value := range cond.(map[string]interface{}) {
		switch key {
		case "LIKE":
			like := LIKE{}
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "Field":
					like.Field = v.(string)
				case "Value":
					like.Value = v.(string)
				}
			}
			return like
		case "EQUALS":
			equals := EQUALS{}
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "Field":
					equals.Field = v.(string)
				case "Value":
					equals.Value = v.(string)
				}
			}
			return equals
		case "IS":
			equals := IS{}
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "Field":
					equals.Field = v.(string)
				case "Value":
					equals.Value = v.(string)
				}
			}
			return equals
		case "NOTEQUALS":
			notEquals := NOTEQUALS{}
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "Field":
					notEquals.Field = v.(string)
				case "Value":
					notEquals.Value = v.(string)
				}
			}
			return notEquals
		case "IN":
			in := IN{}
			values := ""
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "Field":
					in.Field = v.(string)
				case "Value":
					for index, value := range v.([]interface{}) {
						values += "'" + value.(string) + "'"
						if index != len(v.([]interface{}))-1 {
							values += ","
						}
					}
					in.Value = values
				}
			}
			return in
		case "GREATER":
			greater := GREATER{}
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "Field":
					greater.Field = v.(string)
				case "Value":
					greater.Value = v.(string)
				}
			}
			return greater
		case "LESS":
			less := LESS{}
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "Field":
					less.Field = v.(string)
				case "Value":
					less.Value = v.(string)
				}
			}
			return less
		case "GREATEREQUAL":
			greater := GREATEREQUAL{}
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "Field":
					greater.Field = v.(string)
				case "Value":
					greater.Value = v.(string)
				}
			}
			return greater
		case "LESSEQUAL":
			less := LESSEQUAL{}
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "Field":
					less.Field = v.(string)
				case "Value":
					less.Value = v.(string)
				}
			}
			return less
		case "BETWEEN":
			between := BETWEEN{}
			for k, v := range value.(map[string]interface{}) {
				switch k {
				case "Field":
					between.Field = v.(string)
				case "LowerValue":
					between.LowerValue = v.(string)
				case "HigherValue":
					between.HigherValue = v.(string)
				}
			}
			return between
		case "AND":
			and := AND{}
			for _, condition := range value.([]interface{}) {
				and.Conditions = append(and.Conditions, BuildQuery(condition))
			}
			return and
		case "OR":
			or := OR{}
			for _, condition := range value.([]interface{}) {
				or.Conditions = append(or.Conditions, BuildQuery(condition))
			}
			return or
		default:
			return value
		}
	}
	return nil
}

// Gets the request Body as map[string]interface{} and converts to SQL Query string which can be used as WHERE clause.
func ConvertJSONToSQL(body map[string]interface{}) string {
	sqlQuery := ""
	if !reflect.DeepEqual(map[string]interface{}{}, body) {
		switch reflect.ValueOf(BuildQuery(body)).Type().Name() {
		case "LIKE":
			sqlQuery += BuildQuery(body).(LIKE).BuildConditions()
		case "EQUALS":
			sqlQuery += BuildQuery(body).(EQUALS).BuildConditions()
		case "IS":
			sqlQuery += BuildQuery(body).(IS).BuildConditions()
		case "NOTEQUALS":
			sqlQuery += BuildQuery(body).(NOTEQUALS).BuildConditions()
		case "IN":
			sqlQuery += BuildQuery(body).(IN).BuildConditions()
		case "GREATER":
			sqlQuery += BuildQuery(body).(GREATER).BuildConditions()
		case "LESS":
			sqlQuery += BuildQuery(body).(LESS).BuildConditions()
		case "GREATEREQUAL":
			sqlQuery += BuildQuery(body).(GREATEREQUAL).BuildConditions()
		case "LESSEQUAL":
			sqlQuery += BuildQuery(body).(LESSEQUAL).BuildConditions()
		case "BETWEEN":
			sqlQuery += BuildQuery(body).(BETWEEN).BuildConditions()
		case "AND":
			sqlQuery += BuildQuery(body).(AND).BuildConditions()
		case "OR":
			sqlQuery += BuildQuery(body).(OR).BuildConditions()
		default:
			sqlQuery += ""
		}
	}
	return strings.ReplaceAll(sqlQuery, ".", "`.`")
}

func getCondition(conditionType string, condition interface{}) string {
	var queryString string
	switch conditionType {
	case "LIKE":
		queryString = condition.(LIKE).BuildConditions()
	case "EQUALS":
		queryString = condition.(EQUALS).BuildConditions()
	case "IS":
		queryString = condition.(IS).BuildConditions()
	case "NOTEQUALS":
		queryString = condition.(NOTEQUALS).BuildConditions()
	case "IN":
		queryString = condition.(IN).BuildConditions()
	case "GREATER":
		queryString = condition.(GREATER).BuildConditions()
	case "LESS":
		queryString = condition.(LESS).BuildConditions()
	case "GREATEREQUAL":
		queryString = condition.(GREATEREQUAL).BuildConditions()
	case "LESSEQUAL":
		queryString = condition.(LESSEQUAL).BuildConditions()
	case "BETWEEN":
		queryString = condition.(BETWEEN).BuildConditions()
	case "AND":
		queryString = fmt.Sprintf("(%s)", condition.(AND).BuildConditions())
	case "OR":
		queryString = fmt.Sprintf("(%s)", condition.(OR).BuildConditions())
	default:
		queryString = ""
	}
	return queryString
}
