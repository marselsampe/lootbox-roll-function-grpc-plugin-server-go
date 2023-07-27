// Copyright (c) 2023 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package service

import (
	"context"
	"math"
	"math/rand"
	"time"

	pb "lootbox-roll-function-grpc-plugin-server-go/pkg/pb"
)

type LootBoxServiceServer struct {
	pb.UnimplementedLootBoxServer
}

func NewLootBoxServiceServer() *LootBoxServiceServer {
	rand.Seed(time.Now().Unix())

	return &LootBoxServiceServer{}
}

func (s *LootBoxServiceServer) RollLootBoxRewards(_ context.Context, req *pb.RollLootBoxRewardsRequest) (*pb.RollLootBoxRewardsResponse, error) {
	rewards := req.GetItemInfo().GetLootBoxRewards()
	rewardWeightSum := 0
	for _, r := range rewards {
		rewardWeightSum += int(r.Weight)
	}

	var resultItems []*pb.RewardObject
	for i := int32(0); i < req.GetQuantity(); i++ {
		selectedIdx := 0
		for r := int(random(rewardWeightSum)); selectedIdx < len(rewards); selectedIdx++ {
			r -= int(rewards[selectedIdx].GetWeight())
			if r <= 0.0 {
				break
			}
		}

		selectedReward := rewards[selectedIdx]
		selectedRewardItemCount := len(selectedReward.GetItems())

		selectedItemIdx := int(math.Round(random(selectedRewardItemCount - 1)))
		selectedItem := selectedReward.GetItems()[selectedItemIdx]

		resultItems = append(resultItems, &pb.RewardObject{
			ItemId:  selectedItem.ItemId,
			ItemSku: selectedItem.ItemSku,
			Count:   selectedItem.Count,
		})
	}

	return &pb.RollLootBoxRewardsResponse{Rewards: resultItems}, nil
}

func random(max int) float64 {
	return rand.Float64() * float64(max)
}
