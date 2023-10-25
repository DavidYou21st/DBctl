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

func newImPrivilegeUserLog(db *gorm.DB, opts ...gen.DOOption) imPrivilegeUserLog {
	_imPrivilegeUserLog := imPrivilegeUserLog{}

	_imPrivilegeUserLog.imPrivilegeUserLogDo.UseDB(db, opts...)
	_imPrivilegeUserLog.imPrivilegeUserLogDo.UseModel(&model.ImPrivilegeUserLog{})

	tableName := _imPrivilegeUserLog.imPrivilegeUserLogDo.TableName()
	_imPrivilegeUserLog.ALL = field.NewAsterisk(tableName)
	_imPrivilegeUserLog.FromUserid = field.NewString(tableName, "from_userid")
	_imPrivilegeUserLog.AddType = field.NewString(tableName, "add_type")
	_imPrivilegeUserLog.Reson = field.NewString(tableName, "reson")
	_imPrivilegeUserLog.TargetObj = field.NewString(tableName, "target_obj")
	_imPrivilegeUserLog.Times = field.NewTime(tableName, "times")

	_imPrivilegeUserLog.fillFieldMap()

	return _imPrivilegeUserLog
}

type imPrivilegeUserLog struct {
	imPrivilegeUserLogDo imPrivilegeUserLogDo

	ALL        field.Asterisk
	FromUserid field.String // 来源id
	AddType    field.String // 类型添加群或者添加好友
	Reson      field.String // 原因
	TargetObj  field.String // 目标对象好友或者groupid
	Times      field.Time

	fieldMap map[string]field.Expr
}

func (i imPrivilegeUserLog) Table(newTableName string) *imPrivilegeUserLog {
	i.imPrivilegeUserLogDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i imPrivilegeUserLog) As(alias string) *imPrivilegeUserLog {
	i.imPrivilegeUserLogDo.DO = *(i.imPrivilegeUserLogDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *imPrivilegeUserLog) updateTableName(table string) *imPrivilegeUserLog {
	i.ALL = field.NewAsterisk(table)
	i.FromUserid = field.NewString(table, "from_userid")
	i.AddType = field.NewString(table, "add_type")
	i.Reson = field.NewString(table, "reson")
	i.TargetObj = field.NewString(table, "target_obj")
	i.Times = field.NewTime(table, "times")

	i.fillFieldMap()

	return i
}

func (i *imPrivilegeUserLog) WithContext(ctx context.Context) *imPrivilegeUserLogDo {
	return i.imPrivilegeUserLogDo.WithContext(ctx)
}

func (i imPrivilegeUserLog) TableName() string { return i.imPrivilegeUserLogDo.TableName() }

func (i imPrivilegeUserLog) Alias() string { return i.imPrivilegeUserLogDo.Alias() }

func (i *imPrivilegeUserLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *imPrivilegeUserLog) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 5)
	i.fieldMap["from_userid"] = i.FromUserid
	i.fieldMap["add_type"] = i.AddType
	i.fieldMap["reson"] = i.Reson
	i.fieldMap["target_obj"] = i.TargetObj
	i.fieldMap["times"] = i.Times
}

func (i imPrivilegeUserLog) clone(db *gorm.DB) imPrivilegeUserLog {
	i.imPrivilegeUserLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i imPrivilegeUserLog) replaceDB(db *gorm.DB) imPrivilegeUserLog {
	i.imPrivilegeUserLogDo.ReplaceDB(db)
	return i
}

type imPrivilegeUserLogDo struct{ gen.DO }

func (i imPrivilegeUserLogDo) Debug() *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Debug())
}

func (i imPrivilegeUserLogDo) WithContext(ctx context.Context) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i imPrivilegeUserLogDo) ReadDB() *imPrivilegeUserLogDo {
	return i.Clauses(dbresolver.Read)
}

func (i imPrivilegeUserLogDo) WriteDB() *imPrivilegeUserLogDo {
	return i.Clauses(dbresolver.Write)
}

func (i imPrivilegeUserLogDo) Session(config *gorm.Session) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Session(config))
}

func (i imPrivilegeUserLogDo) Clauses(conds ...clause.Expression) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i imPrivilegeUserLogDo) Returning(value interface{}, columns ...string) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i imPrivilegeUserLogDo) Not(conds ...gen.Condition) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i imPrivilegeUserLogDo) Or(conds ...gen.Condition) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i imPrivilegeUserLogDo) Select(conds ...field.Expr) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i imPrivilegeUserLogDo) Where(conds ...gen.Condition) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i imPrivilegeUserLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *imPrivilegeUserLogDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i imPrivilegeUserLogDo) Order(conds ...field.Expr) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i imPrivilegeUserLogDo) Distinct(cols ...field.Expr) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i imPrivilegeUserLogDo) Omit(cols ...field.Expr) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i imPrivilegeUserLogDo) Join(table schema.Tabler, on ...field.Expr) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i imPrivilegeUserLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i imPrivilegeUserLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i imPrivilegeUserLogDo) Group(cols ...field.Expr) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i imPrivilegeUserLogDo) Having(conds ...gen.Condition) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i imPrivilegeUserLogDo) Limit(limit int) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i imPrivilegeUserLogDo) Offset(offset int) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i imPrivilegeUserLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i imPrivilegeUserLogDo) Unscoped() *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Unscoped())
}

func (i imPrivilegeUserLogDo) Create(values ...*model.ImPrivilegeUserLog) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i imPrivilegeUserLogDo) CreateInBatches(values []*model.ImPrivilegeUserLog, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i imPrivilegeUserLogDo) Save(values ...*model.ImPrivilegeUserLog) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i imPrivilegeUserLogDo) First() (*model.ImPrivilegeUserLog, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImPrivilegeUserLog), nil
	}
}

func (i imPrivilegeUserLogDo) Take() (*model.ImPrivilegeUserLog, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImPrivilegeUserLog), nil
	}
}

func (i imPrivilegeUserLogDo) Last() (*model.ImPrivilegeUserLog, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImPrivilegeUserLog), nil
	}
}

func (i imPrivilegeUserLogDo) Find() ([]*model.ImPrivilegeUserLog, error) {
	result, err := i.DO.Find()
	return result.([]*model.ImPrivilegeUserLog), err
}

func (i imPrivilegeUserLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ImPrivilegeUserLog, err error) {
	buf := make([]*model.ImPrivilegeUserLog, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i imPrivilegeUserLogDo) FindInBatches(result *[]*model.ImPrivilegeUserLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i imPrivilegeUserLogDo) Attrs(attrs ...field.AssignExpr) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i imPrivilegeUserLogDo) Assign(attrs ...field.AssignExpr) *imPrivilegeUserLogDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i imPrivilegeUserLogDo) Joins(fields ...field.RelationField) *imPrivilegeUserLogDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i imPrivilegeUserLogDo) Preload(fields ...field.RelationField) *imPrivilegeUserLogDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i imPrivilegeUserLogDo) FirstOrInit() (*model.ImPrivilegeUserLog, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImPrivilegeUserLog), nil
	}
}

func (i imPrivilegeUserLogDo) FirstOrCreate() (*model.ImPrivilegeUserLog, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImPrivilegeUserLog), nil
	}
}

func (i imPrivilegeUserLogDo) FindByPage(offset int, limit int) (result []*model.ImPrivilegeUserLog, count int64, err error) {
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

func (i imPrivilegeUserLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i imPrivilegeUserLogDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i imPrivilegeUserLogDo) Delete(models ...*model.ImPrivilegeUserLog) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *imPrivilegeUserLogDo) withDO(do gen.Dao) *imPrivilegeUserLogDo {
	i.DO = *do.(*gen.DO)
	return i
}
