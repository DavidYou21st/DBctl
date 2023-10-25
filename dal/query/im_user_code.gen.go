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

func newImUserCode(db *gorm.DB, opts ...gen.DOOption) imUserCode {
	_imUserCode := imUserCode{}

	_imUserCode.imUserCodeDo.UseDB(db, opts...)
	_imUserCode.imUserCodeDo.UseModel(&model.ImUserCode{})

	tableName := _imUserCode.imUserCodeDo.TableName()
	_imUserCode.ALL = field.NewAsterisk(tableName)
	_imUserCode.ID = field.NewInt32(tableName, "id")
	_imUserCode.Account = field.NewString(tableName, "account")
	_imUserCode.UserID = field.NewString(tableName, "user_id")
	_imUserCode.Code = field.NewString(tableName, "code")
	_imUserCode.IsUsed = field.NewInt32(tableName, "is_used")
	_imUserCode.CreatedAt = field.NewTime(tableName, "created_at")
	_imUserCode.ExpiredAt = field.NewTime(tableName, "expired_at")

	_imUserCode.fillFieldMap()

	return _imUserCode
}

type imUserCode struct {
	imUserCodeDo imUserCodeDo

	ALL       field.Asterisk
	ID        field.Int32  // id
	Account   field.String // 账号名
	UserID    field.String // 用户id
	Code      field.String // 验证码
	IsUsed    field.Int32  // 是否已经使用，1-未使用，2-已使用
	CreatedAt field.Time   // 创建时间
	ExpiredAt field.Time   // 过期时间

	fieldMap map[string]field.Expr
}

func (i imUserCode) Table(newTableName string) *imUserCode {
	i.imUserCodeDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i imUserCode) As(alias string) *imUserCode {
	i.imUserCodeDo.DO = *(i.imUserCodeDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *imUserCode) updateTableName(table string) *imUserCode {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt32(table, "id")
	i.Account = field.NewString(table, "account")
	i.UserID = field.NewString(table, "user_id")
	i.Code = field.NewString(table, "code")
	i.IsUsed = field.NewInt32(table, "is_used")
	i.CreatedAt = field.NewTime(table, "created_at")
	i.ExpiredAt = field.NewTime(table, "expired_at")

	i.fillFieldMap()

	return i
}

func (i *imUserCode) WithContext(ctx context.Context) *imUserCodeDo {
	return i.imUserCodeDo.WithContext(ctx)
}

func (i imUserCode) TableName() string { return i.imUserCodeDo.TableName() }

func (i imUserCode) Alias() string { return i.imUserCodeDo.Alias() }

func (i *imUserCode) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *imUserCode) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 7)
	i.fieldMap["id"] = i.ID
	i.fieldMap["account"] = i.Account
	i.fieldMap["user_id"] = i.UserID
	i.fieldMap["code"] = i.Code
	i.fieldMap["is_used"] = i.IsUsed
	i.fieldMap["created_at"] = i.CreatedAt
	i.fieldMap["expired_at"] = i.ExpiredAt
}

func (i imUserCode) clone(db *gorm.DB) imUserCode {
	i.imUserCodeDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i imUserCode) replaceDB(db *gorm.DB) imUserCode {
	i.imUserCodeDo.ReplaceDB(db)
	return i
}

type imUserCodeDo struct{ gen.DO }

func (i imUserCodeDo) Debug() *imUserCodeDo {
	return i.withDO(i.DO.Debug())
}

func (i imUserCodeDo) WithContext(ctx context.Context) *imUserCodeDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i imUserCodeDo) ReadDB() *imUserCodeDo {
	return i.Clauses(dbresolver.Read)
}

func (i imUserCodeDo) WriteDB() *imUserCodeDo {
	return i.Clauses(dbresolver.Write)
}

func (i imUserCodeDo) Session(config *gorm.Session) *imUserCodeDo {
	return i.withDO(i.DO.Session(config))
}

func (i imUserCodeDo) Clauses(conds ...clause.Expression) *imUserCodeDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i imUserCodeDo) Returning(value interface{}, columns ...string) *imUserCodeDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i imUserCodeDo) Not(conds ...gen.Condition) *imUserCodeDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i imUserCodeDo) Or(conds ...gen.Condition) *imUserCodeDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i imUserCodeDo) Select(conds ...field.Expr) *imUserCodeDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i imUserCodeDo) Where(conds ...gen.Condition) *imUserCodeDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i imUserCodeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *imUserCodeDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i imUserCodeDo) Order(conds ...field.Expr) *imUserCodeDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i imUserCodeDo) Distinct(cols ...field.Expr) *imUserCodeDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i imUserCodeDo) Omit(cols ...field.Expr) *imUserCodeDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i imUserCodeDo) Join(table schema.Tabler, on ...field.Expr) *imUserCodeDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i imUserCodeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *imUserCodeDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i imUserCodeDo) RightJoin(table schema.Tabler, on ...field.Expr) *imUserCodeDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i imUserCodeDo) Group(cols ...field.Expr) *imUserCodeDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i imUserCodeDo) Having(conds ...gen.Condition) *imUserCodeDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i imUserCodeDo) Limit(limit int) *imUserCodeDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i imUserCodeDo) Offset(offset int) *imUserCodeDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i imUserCodeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *imUserCodeDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i imUserCodeDo) Unscoped() *imUserCodeDo {
	return i.withDO(i.DO.Unscoped())
}

func (i imUserCodeDo) Create(values ...*model.ImUserCode) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i imUserCodeDo) CreateInBatches(values []*model.ImUserCode, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i imUserCodeDo) Save(values ...*model.ImUserCode) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i imUserCodeDo) First() (*model.ImUserCode, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImUserCode), nil
	}
}

func (i imUserCodeDo) Take() (*model.ImUserCode, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImUserCode), nil
	}
}

func (i imUserCodeDo) Last() (*model.ImUserCode, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImUserCode), nil
	}
}

func (i imUserCodeDo) Find() ([]*model.ImUserCode, error) {
	result, err := i.DO.Find()
	return result.([]*model.ImUserCode), err
}

func (i imUserCodeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ImUserCode, err error) {
	buf := make([]*model.ImUserCode, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i imUserCodeDo) FindInBatches(result *[]*model.ImUserCode, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i imUserCodeDo) Attrs(attrs ...field.AssignExpr) *imUserCodeDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i imUserCodeDo) Assign(attrs ...field.AssignExpr) *imUserCodeDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i imUserCodeDo) Joins(fields ...field.RelationField) *imUserCodeDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i imUserCodeDo) Preload(fields ...field.RelationField) *imUserCodeDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i imUserCodeDo) FirstOrInit() (*model.ImUserCode, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImUserCode), nil
	}
}

func (i imUserCodeDo) FirstOrCreate() (*model.ImUserCode, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImUserCode), nil
	}
}

func (i imUserCodeDo) FindByPage(offset int, limit int) (result []*model.ImUserCode, count int64, err error) {
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

func (i imUserCodeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i imUserCodeDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i imUserCodeDo) Delete(models ...*model.ImUserCode) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *imUserCodeDo) withDO(do gen.Dao) *imUserCodeDo {
	i.DO = *do.(*gen.DO)
	return i
}