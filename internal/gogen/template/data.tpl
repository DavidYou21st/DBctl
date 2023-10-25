package data

import (
	"context"
	"fengyin/internal/{{lowCamelCase .appName}}/biz"
    "fengyin/pkg/global"
	"github.com/bullteam/zeus-core/logger"
)

var _ biz.{{.dataName}}Repo = (*{{lowCamelCase .dataName}}Repo)(nil)

type {{lowCamelCase .dataName}}Repo struct {
	data *Data
	log  *logger.Helper
}

func New{{.dataName}}Repo(data *Data) biz.{{.dataName}}Repo {
	return &{{lowCamelCase .dataName}}Repo{
		data: data,
		log:  global.LOG.WithFields(map[string]interface{}{"module": "data/{{.dataName}}"}),
	}
}

func (r *{{lowCamelCase .dataName}}Repo) Create{{.dataName}}(ctx context.Context, a *biz.{{.dataName}}) (*biz.{{.dataName}}, error) {
	po, err := r.data.db.{{.dataName}}.
		Create().
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.{{.dataName}}{ {{range $i, $v := .fieldNames}}
        {{ . }} : po.{{ . }},{{end}}
	}, nil
}

func (r *{{lowCamelCase .dataName}}Repo) Get{{.dataName}}(ctx context.Context, id int64) (*biz.{{.dataName}}, error) {
	po, err := r.data.db.{{.dataName}}.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.{{.dataName}}{ {{range $i, $v := .fieldNames}}
        {{ . }} : po.{{ . }},{{end}}
	}, nil
}

func (r *{{lowCamelCase .dataName}}Repo) List{{.dataName}}(ctx context.Context, uid int64) ([]*biz.{{.dataName}}, error) {
	pos, err := r.data.db.{{.dataName}}.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.{{.dataName}}, 0)
	for _, po := range pos {
		rv = append(rv, &biz.{{.dataName}}{ {{range $i, $v := .fieldNames}}
             {{ . }} : po.{{ . }},{{end}}
		})
	}
	return rv, nil
}


func (r *{{lowCamelCase .dataName}}Repo) Update{{.dataName}}(ctx context.Context, b *biz.{{.dataName}}) (*biz.{{.dataName}}, error) {
	o := {{.dataName}}{}
	result := r.data.db.WithContext(ctx).First(&o, b.Id)
	if result.Error != nil {
		return nil, result.Error
	}
	o.UserID = b.UserID
	result = r.data.db.WithContext(ctx).Save(&o)
	if result.Error != nil {
		return nil, result.Error
	}
	return &biz.{{.dataName}}{
		ID: o.ID,
	}, nil
}

func (r *{{lowCamelCase .dataName}}Repo) Delete{{.dataName}}(ctx context.Context, uid int64) error {
    // todo delete data
	return nil
}