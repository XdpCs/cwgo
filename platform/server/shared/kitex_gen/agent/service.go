// Code generated by thriftgo (0.3.4). DO NOT EDIT.

package agent

import (
	"context"
)

type AgentService interface {
	AddRepository(ctx context.Context, req *AddRepositoryReq) (r *AddRepositoryRes, err error)

	DeleteRepositories(ctx context.Context, req *DeleteRepositoriesReq) (r *DeleteRepositoriesRes, err error)

	UpdateRepository(ctx context.Context, req *UpdateRepositoryReq) (r *UpdateRepositoryRes, err error)

	GetRepositories(ctx context.Context, req *GetRepositoriesReq) (r *GetRepositoriesRes, err error)

	AddIDL(ctx context.Context, req *AddIDLReq) (r *AddIDLRes, err error)

	DeleteIDL(ctx context.Context, req *DeleteIDLsReq) (r *DeleteIDLsRes, err error)

	UpdateIDL(ctx context.Context, req *UpdateIDLReq) (r *UpdateIDLRes, err error)

	GetIDLs(ctx context.Context, req *GetIDLsReq) (r *GetIDLsRes, err error)

	SyncIDLsById(ctx context.Context, req *SyncIDLsByIdReq) (r *SyncIDLsByIdRes, err error)

	AddTemplate(ctx context.Context, req *AddTemplateReq) (r *AddTemplateRes, err error)

	DeleteTemplate(ctx context.Context, req *DeleteTemplateReq) (r *DeleteTemplateRes, err error)

	UpdateTemplate(ctx context.Context, req *UpdateTemplateReq) (r *UpdateTemplateRes, err error)

	GetTemplates(ctx context.Context, req *GetTemplatesReq) (r *GetTemplatesRes, err error)

	AddTemplateItem(ctx context.Context, req *AddTemplateItemReq) (r *AddTemplateItemRes, err error)

	DeleteTemplateItem(ctx context.Context, req *DeleteTemplateItemReq) (r *DeleteTemplateItemRes, err error)

	UpdateTemplateItem(ctx context.Context, req *UpdateTemplateItemReq) (r *UpdateTemplateItemRes, err error)

	GetTemplateItems(ctx context.Context, req *GetTemplateItemsReq) (r *GetTemplateItemsRes, err error)

	UpdateTasks(ctx context.Context, req *UpdateTasksReq) (r *UpdateTasksRes, err error)

	AddToken(ctx context.Context, req *AddTokenReq) (r *AddTokenRes, err error)

	DeleteToken(ctx context.Context, req *DeleteTokenReq) (r *DeleteTokenRes, err error)

	GetToken(ctx context.Context, req *GetTokenReq) (r *GetTokenRes, err error)
}