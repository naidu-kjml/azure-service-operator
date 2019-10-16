/*
Copyright 2019 microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package eventhubs

import (
	"context"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"

	"github.com/Azure/azure-sdk-for-go/services/eventhub/mgmt/2017-04-01/eventhub"
)

type azureConsumerGroupManager struct{}

func getConsumerGroupsClient() eventhub.ConsumerGroupsClient {
	consumerGroupClient := eventhub.NewConsumerGroupsClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	consumerGroupClient.Authorizer = auth
	consumerGroupClient.AddToUserAgent(config.UserAgent())
	return consumerGroupClient
}

// CreateConsumerGroup creates an Event Hub Consumer Group
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
// consumerGroupName - the consumer group name
// parameters - parameters supplied to create or update a consumer group resource.
func (_ *azureConsumerGroupManager) CreateConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error) {
	consumerGroupClient := getConsumerGroupsClient()

	parameters := eventhub.ConsumerGroup{}
	return consumerGroupClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		namespaceName,
		eventHubName,
		consumerGroupName,
		parameters,
	)

}

// DeleteConsumerGroup deletes an Event Hub Consumer Group
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// namespaceName - the Namespace name
// eventHubName - the Event Hub name
// consumerGroupName - the consumer group name
func (_ *azureConsumerGroupManager) DeleteConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (result autorest.Response, err error) {
	consumerGroupClient := getConsumerGroupsClient()
	return consumerGroupClient.Delete(
		ctx,
		resourceGroupName,
		namespaceName,
		eventHubName,
		consumerGroupName,
	)

}

//GetConsumerGroup gets consumer group description for the specified Consumer Group.
func (_ *azureConsumerGroupManager) GetConsumerGroup(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string) (eventhub.ConsumerGroup, error) {
	consumerGroupClient := getConsumerGroupsClient()
	return consumerGroupClient.Get(ctx, resourceGroupName, namespaceName, eventHubName, consumerGroupName)
}