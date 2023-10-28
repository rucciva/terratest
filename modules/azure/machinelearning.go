package azure

import (
	"context"

	machinelearning "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/machinelearning/armmachinelearning/v3"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// ResourceGroupExists indicates whether a resource group exists within a subscription; otherwise false
// This function would fail the test if there is an error.
func MachinelearningWorkspaceExists(t testing.TestingT, resourceGroupName string, workspaceName string, subscriptionID string) bool {
	result, err := MachinelearningWorkspaceExistsE(t, resourceGroupName, workspaceName, subscriptionID)
	require.NoError(t, err)
	return result
}

// ResourceGroupExistsE indicates whether a resource group exists within a subscription
func MachinelearningWorkspaceExistsE(t testing.TestingT, resourceGroupName, workspaceName, subscriptionID string) (bool, error) {
	workspace, err := GetMachinelearningWorkspaceE(t, resourceGroupName, workspaceName, subscriptionID)
	if err != nil {
		if ResourceNotFoundErrorExists(err) {
			return false, nil
		}
		return false, err
	}
	return (workspaceName == *workspace.Name), nil

}

// GetMachinelearningWorkspace is a helper function that gets the machinelearning workspace.
// This function would fail the test if there is an error.
func GetMachinelearningWorkspace(t testing.TestingT, resGroupName string, workspaceName string, subscriptionID string) *machinelearning.Workspace {
	workspace, err := GetMachinelearningWorkspaceE(t, resGroupName, workspaceName, subscriptionID)
	require.NoError(t, err)

	return workspace
}

// GetMachinelearningWorkspaceE is a helper function that gets the machinelearning workspace.
func GetMachinelearningWorkspaceE(t testing.TestingT, resGroupName string, workspaceName string, subscriptionID string) (*machinelearning.Workspace, error) {
	// Create a ml workspace client
	workspaceClient, err := GetMachinelearningWorkspaceClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the corresponding workspace client
	workspace, err := workspaceClient.Get(context.Background(), resGroupName, workspaceName, nil)
	if err != nil {
		return nil, err
	}

	return &workspace.Workspace, nil
}

// GetMachinelearningWorkspaceClientE is a helper function that will setup a machine learning workspace client.
func GetMachinelearningWorkspaceClientE(subscriptionID string) (*machinelearning.WorkspacesClient, error) {
	// Create a new Subnet client from client factory
	client, err := CreateMachinelearningWorkspaceClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	return client, nil
}
