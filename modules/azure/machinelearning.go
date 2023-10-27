package azure

import (
	"context"

	"github.com/Azure/azure-SDK-for-go/services/machinelearningservices/mgmt/2019-05-01/machinelearningservices"
	"github.com/gruntwork-io/terratest/modules/testing"
	"github.com/stretchr/testify/require"
)

// GetMachinelearningWorkspace is a helper function that gets the machinelearning workspace.
// This function would fail the test if there is an error.
func GetMachinelearningWorkspace(t testing.TestingT, resGroupName string, workspaceName string, subscriptionID string) machinelearning.Workspace {
	wo, err := GetMachinelearningWorkspaceE(t, subscriptionID, resGroupName, serverName, dbName)
	require.NoError(t, err)

	return database
}

// GetMachinelearningWorkspaceE is a helper function that gets the machinelearning workspace.
func GetMachinelearningWorkspaceE(t testing.TestingT, subscriptionID string, resGroupName string, workspaceName string) (*machinelearning.Workspace, error) {
	// Create a ml workspace client
	mlWorkspaceClient, err := GetMachinelearningWorkspaceClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Get the corresponding workspace client
	mlWorkspace, err := mlWorkspaceClient.Get(context.Background(), resGroupName, workspaceName)
	if err != nil {
		return nil, err
	}

	//Return workspace
	return &mlWorkspace, nil
}

// GetMachinelearningWorkspaceClientE is a helper function that will setup a machine learning workspace client.
func GetMachinelearningWorkspaceClientE(subscriptionID string) (*machinelearning.WorkspacesClient, error) {
	// Create a new Subnet client from client factory
	client, err := CreateMachinelearningWorkspaceClientE(subscriptionID)
	if err != nil {
		return nil, err
	}

	// Create an authorizer
	authorizer, err := NewAuthorizer()
	if err != nil {
		return nil, err
	}
	client.Authorizer = *authorizer

	return client, nil
}
