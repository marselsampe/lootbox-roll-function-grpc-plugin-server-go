// Copyright (c) 2023 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package server

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"

	pb "lootbox-roll-function-grpc-plugin-server-go/pkg/pb"
)

type LootboxServiceServer struct {
	pb.UnimplementedLootBoxServer
}

func NewLootboxServiceServer() *LootboxServiceServer {
	rand.Seed(time.Now().Unix())

	return &LootboxServiceServer{}
}

func (s *LootboxServiceServer) RollLootBoxRewards(_ context.Context, req *pb.RollLootBoxRewardsRequest) (*pb.RollLootBoxRewardsResponse, error) {
	logJSON("RollLootBoxRewards Request: ", req)
	rewards := req.GetItemInfo().GetLootBoxRewards()
	rewardWeightSum := 0
	for _, r := range rewards {
		rewardWeightSum += int(r.Weight)
	}

	var finalItems []*pb.RewardObject
	for i := int32(0); i < req.GetQuantity(); i++ {
		selIdx := 0
		for r := int(random(rewardWeightSum)); selIdx < len(rewards); selIdx++ {
			r -= int(rewards[selIdx].GetWeight())
			if r <= 0.0 {
				break
			}
		}

		selReward := rewards[selIdx]
		itemCount := len(selReward.GetItems())

		selItemIdx := int(math.Round(random(itemCount - 1)))
		selItem := selReward.GetItems()[selItemIdx]

		finalItems = append(finalItems, &pb.RewardObject{
			ItemId:  selItem.ItemId,
			ItemSku: selItem.ItemSku,
			Count:   selItem.Count,
		})
	}

	response := &pb.RollLootBoxRewardsResponse{Rewards: finalItems}
	logJSON("RollLootBoxRewards Response: ", response)

	return response, nil
}

func random(max int) float64 {
	return rand.Float64() * float64(max)
}

func logJSON(prefix string, jsonData interface{}) {
	r, _ := json.MarshalIndent(jsonData, "", "  ")
	if r != nil {
		fmt.Printf("%s%s\n", prefix, r)
	}
}
