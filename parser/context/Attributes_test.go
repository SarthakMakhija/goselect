//go:build unit
// +build unit

package context

import (
	"reflect"
	"testing"
)

func TestAttributeSizeIsASupportedAttribute(t *testing.T) {
	attributes := NewAttributes()
	isSupportedAttribute := attributes.IsASupportedAttribute(AttributeSize)

	if isSupportedAttribute != true {
		t.Fatalf("Expected AttributeSize to be a supported attribute but was not")
	}
}

func TestDescriptionOfASupportedAttribute(t *testing.T) {
	attributes := NewAttributes()
	description := attributes.DescriptionOf(AttributeSize)

	if description == "" {
		t.Fatalf("Expected description of AttributeSize to not be blank but was blank")
	}
}

func TestDescriptionOfAnUnSupportedAttribute(t *testing.T) {
	attributes := NewAttributes()
	description := attributes.DescriptionOf("unknown")

	if description != "" {
		t.Fatalf("Expected description of unknown to be blank but was not blank")
	}
}

func TestAliasesOfASupportedAttribute(t *testing.T) {
	attributes := NewAttributes()
	aliases := attributes.aliasesFor(AttributeSize)

	if len(aliases) == 0 {
		t.Fatalf("Expected aliases of AttributeSize to not be empty but was empty")
	}
}

func TestAliasesOfAnUnSupportedAttribute(t *testing.T) {
	attributes := NewAttributes()
	aliases := attributes.aliasesFor("unknown")

	if len(aliases) != 0 {
		t.Fatalf("Expected aliases of unknown to be blank but was not blank")
	}
}

func TestAttributeDefinitionOfASupportedAttribute(t *testing.T) {
	attributes := NewAttributes()
	definition := attributes.attributeDefinitionFor(AttributeSize)

	if definition == nil {
		t.Fatalf("Expected definition of AttributeSize to not be nil but was nil")
	}
}

func TestAttributeDefinitionOfAnUnSupportedAttribute(t *testing.T) {
	attributes := NewAttributes()
	definition := attributes.attributeDefinitionFor("unknown")

	if definition != nil {
		t.Fatalf("Expected definition of unknown to be nil but was not nil")
	}
}

func TestAllAttributesWithAliases(t *testing.T) {
	attributes := NewAttributes()
	allAttributesWithAliases := attributes.AllAttributeWithAliases()

	if len(allAttributesWithAliases) == 0 {
		t.Fatalf("Expected allAttributesWithAliases be non-empty but was empty")
	}
}

func TestIsTheAttributeAWildCard(t *testing.T) {
	isAWildcard := IsAWildcardAttribute("*")

	if isAWildcard != true {
		t.Fatalf("Expected isAWildcard to be true but was not")
	}
}

func TestIsTheAttributeNotAWildCard(t *testing.T) {
	isAWildcard := IsAWildcardAttribute("name")

	if isAWildcard != false {
		t.Fatalf("Expected isAWildcard to be false but was true")
	}
}

func TestAttributesOnWildcard(t *testing.T) {
	attributes := AttributesOnWildcard()
	expected := []string{AttributeName, AttributeExtension, AttributeSize, AttributeAbsolutePath}

	if !reflect.DeepEqual(expected, attributes) {
		t.Fatalf("Expected attributes on wildcard to be %v, received %v", expected, attributes)
	}
}
