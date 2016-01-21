package component

import (
	"testing"

	"github.com/gh/stretchr/testify/assert"
	"k8s.io/kubernetes/pkg/api"
)

func TestContainerSimpleDeployment(t *testing.T) {
	kubeContainer := generateContainer(t)

	_, err := NewContainer(kubeContainer, "simpleTest")
	assert.NoError(t, err, "should be able to create container")
}

func generateContainer(t *testing.T) (container api.Container) {
	fuzzer(t).Fuzz(&container)
	return
}
