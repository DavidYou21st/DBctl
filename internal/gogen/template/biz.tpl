package biz

import (
	"context"

	"github.com/bullteam/zeus-core/logger"
    "fengyin/pkg/global"
)

{{.typeCode}}

type {{.bizName}}Repo interface {
	Create{{.bizName}}(ctx context.Context, u *{{.bizName}}) (*{{.bizName}}, error)
	Get{{.bizName}}(ctx context.Context, id int64) (*{{.bizName}}, error)
	Update{{.bizName}}(ctx context.Context, c *{{.bizName}}) (*{{.bizName}}, error)
	Delete{{.bizName}}(ctx context.Context, uid int64) error
	List{{.bizName}}(ctx context.Context, uid int64) ([]*{{.bizName}}, error)
}

type {{.bizName}}UseCase struct {
	repo {{.bizName}}Repo
	log  *logger.Helper
}

func New{{.bizName}}UseCase(repo {{.bizName}}Repo) *{{.bizName}}UseCase {
	return &{{.bizName}}UseCase{
    	repo: repo,
    	log:  global.LOG.WithFields(map[string]interface{}{"module": "usecase/{{.bizName}}"}),
    }
}

func (uc *{{.bizName}}UseCase) Create(ctx context.Context, uid int64, a *{{.bizName}}) (*{{.bizName}}, error) {
	return uc.repo.Create{{.bizName}}(ctx, a)
}

func (uc *{{.bizName}}UseCase) Get(ctx context.Context, id int64) (*{{.bizName}}, error) {
	return uc.repo.Get{{.bizName}}(ctx, id)
}

func (uc *{{.bizName}}UseCase) Update(ctx context.Context, u *{{.bizName}}) (*{{.bizName}}, error) {
	return uc.repo.Update{{.bizName}}(ctx, u)
}

func (uc *{{.bizName}}UseCase) Delete{{.bizName}}(ctx context.Context, uid int64) error {
	return uc.repo.Delete{{.bizName}}(ctx, uid)
}

func (uc *{{.bizName}}UseCase) List(ctx context.Context, uid int64) ([]*{{.bizName}}, error) {
	return uc.repo.List{{.bizName}}(ctx, uid)
}
