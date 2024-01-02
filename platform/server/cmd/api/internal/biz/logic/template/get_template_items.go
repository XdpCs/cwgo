/*
 *
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package template

import (
	"context"

	"github.com/cloudwego/cwgo/platform/server/cmd/api/internal/svc"
	"github.com/cloudwego/cwgo/platform/server/shared/consts"
	"github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/agent"
	"github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/template"
	"github.com/cloudwego/cwgo/platform/server/shared/log"
	"go.uber.org/zap"
)

const (
	successMsgGetTemplateItems = "get template items successfully"
)

type GetTemplateItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTemplateItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTemplateItemsLogic {
	return &GetTemplateItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTemplateItemsLogic) GetTemplateItems(req *template.GetTemplateItemsReq) (res *template.GetTemplateItemsRes) {
	client, err := l.svcCtx.Manager.GetAgentClient()
	if err != nil {
		log.Error(consts.ErrMsgRpcGetClient, zap.Error(err))
		return &template.GetTemplateItemsRes{
			Code: consts.ErrNumRpcGetClient,
			Msg:  consts.ErrMsgRpcGetClient,
		}
	}

	rpcRes, err := client.GetTemplateItems(l.ctx, &agent.GetTemplateItemsReq{
		TemplateId: req.Id,
	})
	if err != nil {
		log.Error(consts.ErrMsgRpcConnectClient, zap.Error(err))
		return &template.GetTemplateItemsRes{
			Code: consts.ErrNumRpcConnectClient,
			Msg:  consts.ErrMsgRpcConnectClient,
		}
	}
	if rpcRes.Code != 0 {
		return &template.GetTemplateItemsRes{
			Code: rpcRes.Code,
			Msg:  rpcRes.Msg,
		}
	}

	return &template.GetTemplateItemsRes{
		Code: 0,
		Msg:  successMsgGetTemplateItems,
		Data: &template.GetTemplateItemsResData{TemplateItems: rpcRes.Data.TemplateItems},
	}
}