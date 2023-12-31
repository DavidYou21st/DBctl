// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"fctl/dal/model"
)

func newImFeedback(db *gorm.DB, opts ...gen.DOOption) imFeedback {
	_imFeedback := imFeedback{}

	_imFeedback.imFeedbackDo.UseDB(db, opts...)
	_imFeedback.imFeedbackDo.UseModel(&model.ImFeedback{})

	tableName := _imFeedback.imFeedbackDo.TableName()
	_imFeedback.ALL = field.NewAsterisk(tableName)
	_imFeedback.ID = field.NewInt64(tableName, "id")
	_imFeedback.UserID = field.NewString(tableName, "user_id")
	_imFeedback.NickeName = field.NewString(tableName, "nicke_name")
	_imFeedback.Contact = field.NewString(tableName, "contact")
	_imFeedback.FeedbackContent = field.NewString(tableName, "feedback_content")
	_imFeedback.Imgs = field.NewString(tableName, "imgs")
	_imFeedback.Status = field.NewString(tableName, "status")
	_imFeedback.HandleContent = field.NewString(tableName, "handle_content")
	_imFeedback.CreatedAt = field.NewTime(tableName, "created_at")
	_imFeedback.UpdatedAt = field.NewTime(tableName, "updated_at")
	_imFeedback.DeletedAt = field.NewField(tableName, "deleted_at")
	_imFeedback.HandlePerson = field.NewString(tableName, "handle_person")

	_imFeedback.fillFieldMap()

	return _imFeedback
}

type imFeedback struct {
	imFeedbackDo imFeedbackDo

	ALL             field.Asterisk
	ID              field.Int64  // 主键id
	UserID          field.String // 用户userid
	NickeName       field.String // 称呼
	Contact         field.String // 联系方式
	FeedbackContent field.String // 反馈内容
	Imgs            field.String // 图片url多张用,隔开
	Status          field.String // 状态 已处理和待处理
	HandleContent   field.String // 处理结果内容
	CreatedAt       field.Time   // 创建时间
	UpdatedAt       field.Time   // 更新时间
	DeletedAt       field.Field  // 删除时间
	HandlePerson    field.String // 操作人

	fieldMap map[string]field.Expr
}

func (i imFeedback) Table(newTableName string) *imFeedback {
	i.imFeedbackDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i imFeedback) As(alias string) *imFeedback {
	i.imFeedbackDo.DO = *(i.imFeedbackDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *imFeedback) updateTableName(table string) *imFeedback {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt64(table, "id")
	i.UserID = field.NewString(table, "user_id")
	i.NickeName = field.NewString(table, "nicke_name")
	i.Contact = field.NewString(table, "contact")
	i.FeedbackContent = field.NewString(table, "feedback_content")
	i.Imgs = field.NewString(table, "imgs")
	i.Status = field.NewString(table, "status")
	i.HandleContent = field.NewString(table, "handle_content")
	i.CreatedAt = field.NewTime(table, "created_at")
	i.UpdatedAt = field.NewTime(table, "updated_at")
	i.DeletedAt = field.NewField(table, "deleted_at")
	i.HandlePerson = field.NewString(table, "handle_person")

	i.fillFieldMap()

	return i
}

func (i *imFeedback) WithContext(ctx context.Context) *imFeedbackDo {
	return i.imFeedbackDo.WithContext(ctx)
}

func (i imFeedback) TableName() string { return i.imFeedbackDo.TableName() }

func (i imFeedback) Alias() string { return i.imFeedbackDo.Alias() }

func (i *imFeedback) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *imFeedback) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 12)
	i.fieldMap["id"] = i.ID
	i.fieldMap["user_id"] = i.UserID
	i.fieldMap["nicke_name"] = i.NickeName
	i.fieldMap["contact"] = i.Contact
	i.fieldMap["feedback_content"] = i.FeedbackContent
	i.fieldMap["imgs"] = i.Imgs
	i.fieldMap["status"] = i.Status
	i.fieldMap["handle_content"] = i.HandleContent
	i.fieldMap["created_at"] = i.CreatedAt
	i.fieldMap["updated_at"] = i.UpdatedAt
	i.fieldMap["deleted_at"] = i.DeletedAt
	i.fieldMap["handle_person"] = i.HandlePerson
}

func (i imFeedback) clone(db *gorm.DB) imFeedback {
	i.imFeedbackDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i imFeedback) replaceDB(db *gorm.DB) imFeedback {
	i.imFeedbackDo.ReplaceDB(db)
	return i
}

type imFeedbackDo struct{ gen.DO }

func (i imFeedbackDo) Debug() *imFeedbackDo {
	return i.withDO(i.DO.Debug())
}

func (i imFeedbackDo) WithContext(ctx context.Context) *imFeedbackDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i imFeedbackDo) ReadDB() *imFeedbackDo {
	return i.Clauses(dbresolver.Read)
}

func (i imFeedbackDo) WriteDB() *imFeedbackDo {
	return i.Clauses(dbresolver.Write)
}

func (i imFeedbackDo) Session(config *gorm.Session) *imFeedbackDo {
	return i.withDO(i.DO.Session(config))
}

func (i imFeedbackDo) Clauses(conds ...clause.Expression) *imFeedbackDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i imFeedbackDo) Returning(value interface{}, columns ...string) *imFeedbackDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i imFeedbackDo) Not(conds ...gen.Condition) *imFeedbackDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i imFeedbackDo) Or(conds ...gen.Condition) *imFeedbackDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i imFeedbackDo) Select(conds ...field.Expr) *imFeedbackDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i imFeedbackDo) Where(conds ...gen.Condition) *imFeedbackDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i imFeedbackDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *imFeedbackDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i imFeedbackDo) Order(conds ...field.Expr) *imFeedbackDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i imFeedbackDo) Distinct(cols ...field.Expr) *imFeedbackDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i imFeedbackDo) Omit(cols ...field.Expr) *imFeedbackDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i imFeedbackDo) Join(table schema.Tabler, on ...field.Expr) *imFeedbackDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i imFeedbackDo) LeftJoin(table schema.Tabler, on ...field.Expr) *imFeedbackDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i imFeedbackDo) RightJoin(table schema.Tabler, on ...field.Expr) *imFeedbackDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i imFeedbackDo) Group(cols ...field.Expr) *imFeedbackDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i imFeedbackDo) Having(conds ...gen.Condition) *imFeedbackDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i imFeedbackDo) Limit(limit int) *imFeedbackDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i imFeedbackDo) Offset(offset int) *imFeedbackDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i imFeedbackDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *imFeedbackDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i imFeedbackDo) Unscoped() *imFeedbackDo {
	return i.withDO(i.DO.Unscoped())
}

func (i imFeedbackDo) Create(values ...*model.ImFeedback) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i imFeedbackDo) CreateInBatches(values []*model.ImFeedback, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i imFeedbackDo) Save(values ...*model.ImFeedback) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i imFeedbackDo) First() (*model.ImFeedback, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFeedback), nil
	}
}

func (i imFeedbackDo) Take() (*model.ImFeedback, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFeedback), nil
	}
}

func (i imFeedbackDo) Last() (*model.ImFeedback, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFeedback), nil
	}
}

func (i imFeedbackDo) Find() ([]*model.ImFeedback, error) {
	result, err := i.DO.Find()
	return result.([]*model.ImFeedback), err
}

func (i imFeedbackDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ImFeedback, err error) {
	buf := make([]*model.ImFeedback, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i imFeedbackDo) FindInBatches(result *[]*model.ImFeedback, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i imFeedbackDo) Attrs(attrs ...field.AssignExpr) *imFeedbackDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i imFeedbackDo) Assign(attrs ...field.AssignExpr) *imFeedbackDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i imFeedbackDo) Joins(fields ...field.RelationField) *imFeedbackDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i imFeedbackDo) Preload(fields ...field.RelationField) *imFeedbackDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i imFeedbackDo) FirstOrInit() (*model.ImFeedback, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFeedback), nil
	}
}

func (i imFeedbackDo) FirstOrCreate() (*model.ImFeedback, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFeedback), nil
	}
}

func (i imFeedbackDo) FindByPage(offset int, limit int) (result []*model.ImFeedback, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i imFeedbackDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i imFeedbackDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i imFeedbackDo) Delete(models ...*model.ImFeedback) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *imFeedbackDo) withDO(do gen.Dao) *imFeedbackDo {
	i.DO = *do.(*gen.DO)
	return i
}
