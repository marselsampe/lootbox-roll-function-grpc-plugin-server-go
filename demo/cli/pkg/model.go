// Copyright (c) 2023 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package lootboxrolldemo

import "fmt"

type SimpleItemInfo struct {
	ID    string
	SKU   string
	Title string
}

type SimpleLootboxItem struct {
	ID          string
	SKU         string
	Title       string
	Diff        string
	RewardItems []SimpleItemInfo
}

func (i *SimpleLootboxItem) WriteToConsole(indent string) {
	fmt.Printf("%sLootbox Item ID: %s\n", indent, i.ID)
	if i.RewardItems != nil {
		fmt.Printf("%sReward Items:\n", indent)
		for _, item := range i.RewardItems {
			fmt.Printf("\t%s%s : %s : %s\n", indent, item.ID, item.SKU, item.Title)
		}
	}
}
