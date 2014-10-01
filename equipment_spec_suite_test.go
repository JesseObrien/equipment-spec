package equipment_spec_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEquipmentSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EquipmentSpec Suite")
}
