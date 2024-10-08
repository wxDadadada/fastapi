// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	sdkm "github.com/iimeta/fastapi-sdk/model"
	"github.com/iimeta/fastapi/internal/model"
	mcommon "github.com/iimeta/fastapi/internal/model/common"
)

type (
	IAudio interface {
		// Speech
		Speech(ctx context.Context, params sdkm.SpeechRequest, fallbackModel *model.Model, retry ...int) (response sdkm.SpeechResponse, err error)
		// Transcriptions
		Transcriptions(ctx context.Context, params sdkm.AudioRequest, fallbackModel *model.Model, retry ...int) (response sdkm.AudioResponse, err error)
		// 保存日志
		SaveLog(ctx context.Context, reqModel, realModel, fallbackModel *model.Model, key *model.Key, completionsReq *sdkm.AudioRequest, completionsRes *model.CompletionsRes, retryInfo *mcommon.Retry, isSmartMatch ...bool)
	}
)

var (
	localAudio IAudio
)

func Audio() IAudio {
	if localAudio == nil {
		panic("implement not found for interface IAudio, forgot register?")
	}
	return localAudio
}

func RegisterAudio(i IAudio) {
	localAudio = i
}
