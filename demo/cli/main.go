// Copyright (c) 2023 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclient/users"
	"github.com/AccelByte/accelbyte-go-sdk/iam-sdk/pkg/iamclientmodels"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/factory"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/repository"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/service/iam"
	"github.com/AccelByte/accelbyte-go-sdk/services-api/pkg/utils/auth"

	lootboxrolldemo "cli/pkg"
)

func main() {
	config, err := lootboxrolldemo.GetConfig()
	if err != nil {
		log.Fatalf("Can't retrieve config: %s\n", err)
	}

	configRepo := auth.DefaultConfigRepositoryImpl()
	tokenRepo := auth.DefaultTokenRepositoryImpl()

	oauthService := &iam.OAuth20Service{
		Client:           factory.NewIamClient(configRepo),
		ConfigRepository: configRepo,
		TokenRepository:  tokenRepo,
	}

	fmt.Print("Login to AccelByte... ")
	err = oauthService.Login(config.ABUsername, config.ABPassword)
	if err != nil {
		log.Fatalf("Accelbyte account login failed: %s\n", err)
	}
	fmt.Println("[OK]")

	usersService := &iam.UsersService{
		Client:           factory.NewIamClient(configRepo),
		ConfigRepository: configRepo,
		TokenRepository:  tokenRepo,
	}
	userInfo, err := usersService.PublicGetMyUserV3Short(&users.PublicGetMyUserV3Params{})
	if err != nil {
		log.Fatalf("Get user info failed: %s\n", err)
	}
	fmt.Printf("\tUser: %s\n", userInfo.UserName)

	rand.Seed(time.Now().Unix())
	// Start testing
	err = startTesting(userInfo, config, configRepo, tokenRepo)
	if err != nil {
		fmt.Println("\n[FAILED]")
		log.Fatal(err)
	}
	fmt.Println("[SUCCESS]")
}

func startTesting(
	userInfo *iamclientmodels.ModelUserResponseV3,
	config *lootboxrolldemo.Config,
	configRepo repository.ConfigRepository,
	tokenRepo repository.TokenRepository) error {
	categoryPath := "/goLootboxRollPluginDemo"
	pdu := lootboxrolldemo.PlatformDataUnit{
		CLIConfig:    config,
		ConfigRepo:   configRepo,
		TokenRepo:    tokenRepo,
		CurrencyCode: "USD",
	}

	// clean up
	defer func() {
		fmt.Println("\nCleaning up...")
		fmt.Print("Deleting store... ")
		err := pdu.DeleteStore()
		if err != nil {
			return
		}
		fmt.Println("[OK]")

		err = pdu.UnsetPlatformServiceGrpcTarget()
		if err != nil {
			fmt.Printf("failed to unset platform service grpc plugin url")

			return
		}
	}()

	// 1.
	fmt.Printf("Configuring platform service grpc target... (%s) ", config.GRPCServerURL)
	err := pdu.SetPlatformServiceGrpcTarget()
	if err != nil {
		fmt.Println("[ERR]")

		return err
	}
	fmt.Println("[OK]")

	// 2.
	fmt.Print("Creating store... ")
	err = pdu.CreateStore(true)
	if err != nil {
		fmt.Println("[ERR]")

		return err
	}
	fmt.Println("[OK]")

	// 3.
	fmt.Print("Creating category... ")
	err = pdu.CreateCategory(categoryPath, true)
	if err != nil {
		fmt.Println("[ERR]")

		return err
	}
	fmt.Println("[OK]")

	// 4.
	fmt.Print("Creating lootbox item(s)... ")
	items, err := pdu.CreateLootboxItems(1, 5, categoryPath, true)
	if err != nil {
		fmt.Println("[ERR]")

		return err
	}
	fmt.Println("[OK]")
	items[0].WriteToConsole("  ")

	// 5.
	fmt.Printf("Granting item entitlement to user %s... ", userInfo.UserName)
	entitlementID, err := pdu.GrantEntitlement(lootboxrolldemo.Val(userInfo.UserID), items[0].ID, 1)
	if err != nil {
		fmt.Println("[ERR]")

		return err
	}
	fmt.Println("[OK]")
	fmt.Printf("\tEntitlement ID: %s\n", entitlementID)

	// 6.
	fmt.Print("Consuming entitlement... ")
	lootboxItemResult, err := pdu.ConsumeItemEntitlement(lootboxrolldemo.Val(userInfo.UserID), entitlementID, 2)
	if err != nil {
		fmt.Println("[ERR]")

		return err
	}
	fmt.Println("[OK]")
	lootboxItemResult.WriteToConsole("  ")

	return nil
}
