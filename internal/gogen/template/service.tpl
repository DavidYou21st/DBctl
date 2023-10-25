package service

import (
	"context"
    "zeuscloud/internal/{{lowCamelCase .appName}}/dto"
    "zeuscloud/internal/{{lowCamelCase .appName}}/biz"
)

func (s *{{.appName}}Service) Create{{.serviceName}}(ctx context.Context, req dto.Create{{.serviceName}}Req) (*dto.Create{{.serviceName}}Reply, error) {
	rv, err := s.{{.serviceStr}}.Create(ctx, req.UID, &biz.{{.serviceName}}{ {{range $i, $v := .fieldNames}}
        {{ . }} : req.{{ . }},{{end}}
	})
	if err != nil {
		return nil, err
	}

	return &dto.Create{{.serviceName}}Reply{
		ID: rv.ID,
	}, nil
}

func (s *{{.appName}}Service) Get{{.serviceName}}(ctx context.Context, req dto.Get{{.serviceName}}Req) (*dto.Get{{.serviceName}}Reply, error) {
	x, err := s.{{.serviceStr}}.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	return &dto.Get{{.serviceName}}Reply{ {{range $i, $v := .getReply}}
          {{ . }} : x.{{ . }},{{end}}
	}, nil
}

func (s *{{.appName}}Service) List{{.serviceName}}(ctx context.Context, req dto.List{{.serviceName}}Req) (*dto.List{{.serviceName}}Reply, error) {
	rv, err := s.{{.serviceStr}}.List(ctx, req.UID)
	if err != nil {
		return nil, err
	}

	rs := make([]*dto.List{{.serviceName}}Reply, 0)
	for _, x := range rv {
		rs = append(rs, &dto.List{{.serviceName}}Reply { {{range $i, $v := .listReply}}
            {{ . }} : x.{{ . }},{{end}}
		})
	}
	return &dto.List{{.serviceName}}Reply{
		Results: rs,
	}, nil
}

func (s *{{.appName}}Service) Update{{.serviceName}}(ctx context.Context, req *dto.Update{{.serviceName}}Req) (*dto.Update{{.serviceName}}Reply, error) {
	x, err := s.{{.serviceStr}}.Update(ctx, &biz.{{.serviceName}}{})
	if err != nil {
		return nil, err
	}

	return &dto.Update{{.serviceName}}Reply{
		ID: x.ID,
	}, nil
}