// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type InvalidParameterStandard struct {
	Field string `json:"field"`
	// invalid parameters rules
	Rule   *InvalidRules `json:"rule,omitempty"`
	Source *string       `json:"source,omitempty"`
	Reason string        `json:"reason"`
}

func (o *InvalidParameterStandard) GetField() string {
	if o == nil {
		return ""
	}
	return o.Field
}

func (o *InvalidParameterStandard) GetRule() *InvalidRules {
	if o == nil {
		return nil
	}
	return o.Rule
}

func (o *InvalidParameterStandard) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *InvalidParameterStandard) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}
