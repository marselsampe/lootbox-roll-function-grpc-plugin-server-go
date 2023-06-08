// Copyright (c) 2023 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package platformservice

import (
	"context"
	"net/url"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"cli/pkg/client/platformservice/openapi2/client"
	"cli/pkg/client/platformservice/openapi2/client/service_plugin_config"
	"cli/pkg/client/platformservice/openapi2/models"
)

type Client struct {
	tokenRepo         repository.TokenRepository
	platformSvcClient *client.JusticePlatformService
}

func NewClient(serviceAddress string, tokenRepo repository.TokenRepository) (*Client, error) {
	u, err := url.Parse(serviceAddress)
	if err != nil {
		return nil, err
	}
	platformSvcClient := client.New(httptransport.New(u.Host, "platform", []string{u.Scheme}), strfmt.Default)
	return &Client{
		tokenRepo:         tokenRepo,
		platformSvcClient: platformSvcClient,
	}, nil
}

func (c *Client) UpdateLootBoxPluginConfig(namespace string, config *models.LootBoxPluginConfigUpdate) error {
	token, err := c.tokenRepo.GetToken()
	if err != nil {
		return err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	bearerToken := httptransport.BearerToken(*token.AccessToken)
	_, err = c.platformSvcClient.ServicePluginConfig.UpdateLootBoxPluginConfig(&service_plugin_config.UpdateLootBoxPluginConfigParams{
		Namespace: namespace,
		Body:      config,
		Context:   ctx,
	}, bearerToken)
	return err
}

func (c *Client) DeleteLootBoxPluginConfig(namespace string) error {
	token, err := c.tokenRepo.GetToken()
	if err != nil {
		return err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	bearerToken := httptransport.BearerToken(*token.AccessToken)
	_, err = c.platformSvcClient.ServicePluginConfig.DeleteLootBoxPluginConfig(&service_plugin_config.DeleteLootBoxPluginConfigParams{
		Namespace: namespace,
		Context:   ctx,
	}, bearerToken)
	return err
}
