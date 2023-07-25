// Copyright (c) 2023 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package server

import (
	"context"
	"math"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"

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
	logrus.Infof("RollLootBoxRewards Request: %s", logJSONFormatter(req))
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

	response := &pb.RollLootBoxRewardsResponse{Rewards: resultItems}
	logrus.Infof("RollLootBoxRewards Response: %s", logJSONFormatter(response))

	return response, nil
}

func random(max int) float64 {
	return rand.Float64() * float64(max)
}
