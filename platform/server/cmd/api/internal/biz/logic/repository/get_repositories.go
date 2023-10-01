/*
 *
 *  * Copyright 2022 CloudWeGo Authors
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *     http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package repository

import (
	"context"
	"github.com/cloudwego/cwgo/platform/server/cmd/api/internal/svc"
	"github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/repository"
)

const (
	successMsgGetRepositories = "" // TODO: to be filled...
)

type GetRepositoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRepositoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRepositoriesLogic {
	return &GetRepositoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRepositoriesLogic) GetRepositories(req *repository.GetRepositoriesReq) (res *repository.GetRepositoriesRes) {
	repos, err := l.svcCtx.DaoManager.Repository.GetRepositories(req.Page, req.Limit, req.Order, req.OrderBy)
	if err != nil {
		return &repository.GetRepositoriesRes{
			Code: 400,
			Msg:  err.Error(),
			Data: nil,
		}
	}

	return &repository.GetRepositoriesRes{
		Code: 0,
		Msg:  successMsgGetRepositories,
		Data: &repository.GetRepositoriesResData{Repositories: repos},
	}
}
