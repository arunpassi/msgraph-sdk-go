package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PatternedRecurrence 
type PatternedRecurrence struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The frequency of an event. Do not specify for a one-time access review.
    pattern *RecurrencePattern;
    // The duration of an event.
    range_escaped *RecurrenceRange;
}
// NewPatternedRecurrence instantiates a new patternedRecurrence and sets the default values.
func NewPatternedRecurrence()(*PatternedRecurrence) {
    m := &PatternedRecurrence{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PatternedRecurrence) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetPattern gets the pattern property value. The frequency of an event. Do not specify for a one-time access review.
func (m *PatternedRecurrence) GetPattern()(*RecurrencePattern) {
    if m == nil {
        return nil
    } else {
        return m.pattern
    }
}
// GetRange_escaped gets the range_escaped property value. The duration of an event.
func (m *PatternedRecurrence) GetRange_escaped()(*RecurrenceRange) {
    if m == nil {
        return nil
    } else {
        return m.range_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PatternedRecurrence) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["pattern"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecurrencePattern() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPattern(val.(*RecurrencePattern))
        }
        return nil
    }
    res["range_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewRecurrenceRange() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRange_escaped(val.(*RecurrenceRange))
        }
        return nil
    }
    return res
}
func (m *PatternedRecurrence) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PatternedRecurrence) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("pattern", m.GetPattern())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("range_escaped", m.GetRange_escaped())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PatternedRecurrence) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetPattern sets the pattern property value. The frequency of an event. Do not specify for a one-time access review.
func (m *PatternedRecurrence) SetPattern(value *RecurrencePattern)() {
    m.pattern = value
}
// SetRange_escaped sets the range_escaped property value. The duration of an event.
func (m *PatternedRecurrence) SetRange_escaped(value *RecurrenceRange)() {
    m.range_escaped = value
}
