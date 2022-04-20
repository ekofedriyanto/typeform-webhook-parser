package parser

import (
	"encoding/json"
	"github.com/ahmetb/go-linq/v3"
	"reflect"
)

type TypeFormParser interface {
	ToMap() (returnData map[string]interface{}, err error)
}

type typeForm struct {
	TypeFormData *TypeFormData
}

var _ TypeFormParser = (*typeForm)(nil)

func NewParser(typeFormData *TypeFormData) (returnData TypeFormParser) {
	return &typeForm{
		TypeFormData: typeFormData,
	}
}

func (t *typeForm) ToMap() (returnData map[string]interface{}, err error) {
	returnData = make(map[string]interface{})

	// Populate data from typeform standard fields
	{
		if t.TypeFormData.EventID != "" {
			if tag, ok := t.getJSONTag(*t.TypeFormData, "EventID"); ok {
				returnData[tag] = t.TypeFormData.EventID
			}
		}

		if t.TypeFormData.EventType != "" {
			if tag, ok := t.getJSONTag(*t.TypeFormData, "EventType"); ok {
				returnData[tag] = t.TypeFormData.EventType
			}
		}

		if t.TypeFormData.FormResponse.FormID != "" {
			if tag, ok := t.getJSONTag(*t.TypeFormData.FormResponse, "FormID"); ok {
				returnData[tag] = t.TypeFormData.FormResponse.FormID
			}
		}

		if t.TypeFormData.FormResponse.Token != "" {
			if tag, ok := t.getJSONTag(*t.TypeFormData.FormResponse, "Token"); ok {
				returnData[tag] = t.TypeFormData.FormResponse.Token
			}
		}

		if t.TypeFormData.FormResponse.LandedAt != nil {
			if tag, ok := t.getJSONTag(*t.TypeFormData.FormResponse, "LandedAt"); ok {
				returnData[tag] = t.TypeFormData.FormResponse.LandedAt
			}
		}

		if t.TypeFormData.FormResponse.SubmittedAt != nil {
			if tag, ok := t.getJSONTag(*t.TypeFormData.FormResponse, "SubmittedAt"); ok {
				returnData[tag] = t.TypeFormData.FormResponse.SubmittedAt
			}
		}

		if t.TypeFormData.FormResponse.Definition.Title != "" {
			if tag, ok := t.getJSONTag(*t.TypeFormData.FormResponse.Definition, "Title"); ok {
				returnData[tag] = t.TypeFormData.FormResponse.Definition.Title
			}
		}

	}

	// Populate data from hidden fields
	linq.From(t.TypeFormData.FormResponse.Hidden).
		ToMap(&returnData)

	// Populate Questions/Definitions
	linq.From(t.TypeFormData.FormResponse.Definition.Fields).
		ToMapBy(&returnData,
			func(a interface{}) interface{} {
				return a.(Fields).Ref
			},
			func(b interface{}) interface{} {
				returnCheck := linq.From(t.TypeFormData.FormResponse.Answers).
					Where(func(c interface{}) bool {
						return c.(Answers).Field.ID == b.(Fields).ID && c.(Answers).Field.Ref == b.(Fields).Ref
					}).
					Select(func(d interface{}) (rt interface{}) {
						switch d.(Answers).Type {
						case "phone_number":
							rt = d.(Answers).PhoneNumber
						case "email":
							rt = d.(Answers).Email
						case "url":
							rt = d.(Answers).URL
						case "text":
							rt = d.(Answers).Text
						case "number":
							rt = d.(Answers).Number
						case "date":
							rt = d.(Answers).Date
						case "file_url":
							rt = d.(Answers).FileURL
						case "choice":
							rt = d.(Answers).Choice.Label
						case "choices":
							{
								rtJson, err := json.Marshal(&d.(Answers).Choices.Labels)

								if err != nil {
									rt = ""
								}

								rt = string(rtJson)

							}
						case "boolean":
							rt = d.(Answers).Boolean
						default:
							rt = ""
						}

						return
					}).
					First()
				return returnCheck
			})

	return
}

func (t *typeForm) getJSONTag(v interface{}, fieldName string) (returnData string, ok bool) {
	tof := reflect.TypeOf(v)
	sf, ok := tof.FieldByName(fieldName)
	if !ok {
		return "", false
	}
	return sf.Tag.Lookup("json")
}
