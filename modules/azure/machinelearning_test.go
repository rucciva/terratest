//go:build azure
// +build azure

package azure

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMachinelearningWorkspaceExistsE(t *testing.T) {
	t.Parallel()

	subscriptionID := ""

	_, err := MachinelearningWorkspaceExistsE(subscriptionID)
	require.Error(t, err)
}

func TestGetMachinelearningWorkspaceE(t *testing.T) {
	t.Parallel()

	subscriptionID := ""

	_, err := GetMachinelearningWorkspaceE(subscriptionID)
	require.Error(t, err)
}
