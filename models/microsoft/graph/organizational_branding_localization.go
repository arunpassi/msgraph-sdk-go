package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// OrganizationalBrandingLocalization 
type OrganizationalBrandingLocalization struct {
    OrganizationalBrandingProperties
}
// NewOrganizationalBrandingLocalization instantiates a new organizationalBrandingLocalization and sets the default values.
func NewOrganizationalBrandingLocalization()(*OrganizationalBrandingLocalization) {
    m := &OrganizationalBrandingLocalization{
        OrganizationalBrandingProperties: *NewOrganizationalBrandingProperties(),
    }
    return m
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OrganizationalBrandingLocalization) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.OrganizationalBrandingProperties.GetFieldDeserializers()
    return res
}
func (m *OrganizationalBrandingLocalization) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *OrganizationalBrandingLocalization) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.OrganizationalBrandingProperties.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
