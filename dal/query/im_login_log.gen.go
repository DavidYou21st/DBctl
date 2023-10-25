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

func newImLoginLog(db *gorm.DB, opts ...gen.DOOption) imLoginLog {
	_imLoginLog := imLoginLog{}

	_imLoginLog.imLoginLogDo.UseDB(db, opts...)
	_imLoginLog.imLoginLogDo.UseModel(&model.ImLoginLog{})

	tableName := _imLoginLog.imLoginLogDo.TableName()
	_imLoginLog.ALL = field.NewAsterisk(tableName)
	_imLoginLog.ID = field.NewInt32(tableName, "id")
	_imLoginLog.UserID = field.NewString(tableName, "user_id")
	_imLoginLog.Platform = field.NewInt32(tableName, "platform")
	_imLoginLog.LoginIP = field.NewString(tableName, "login_ip")
	_imLoginLog.Version = field.NewString(tableName, "version")
	_imLoginLog.CreateTime = field.NewTime(tableName, "create_time")
	_imLoginLog.UpdateTime = field.NewTime(tableName, "update_time")

	_imLoginLog.fillFieldMap()

	return _imLoginLog
}

type imLoginLog struct {
	imLoginLogDo imLoginLogDo

	ALL        field.Asterisk
	ID         field.Int32
	UserID     field.String
	Platform   field.Int32 // 平台，IOS:1、安卓:2、window:3、OSX:4、Web:5、小程序:6、Linux:7
	LoginIP    field.String
	Version    field.String
	CreateTime field.Time
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (i imLoginLog) Table(newTableName string) *imLoginLog {
	i.imLoginLogDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i imLoginLog) As(alias string) *imLoginLog {
	i.imLoginLogDo.DO = *(i.imLoginLogDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *imLoginLog) updateTableName(table string) *imLoginLog {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt32(table, "id")
	i.UserID = field.NewString(table, "user_id")
	i.Platform = field.NewInt32(table, "platform")
	i.LoginIP = field.NewString(table, "login_ip")
	i.Version = field.NewString(table, "version")
	i.CreateTime = field.NewTime(table, "create_time")
	i.UpdateTime = field.NewTime(table, "update_time")

	i.fillFieldMap()

	return i
}

func (i *imLoginLog) WithContext(ctx context.Context) *imLoginLogDo {
	return i.imLoginLogDo.WithContext(ctx)
}

func (i imLoginLog) TableName() string { return i.imLoginLogDo.TableName() }

func (i imLoginLog) Alias() string { return i.imLoginLogDo.Alias() }

func (i *imLoginLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *imLoginLog) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 7)
	i.fieldMap["id"] = i.ID
	i.fieldMap["user_id"] = i.UserID
	i.fieldMap["platform"] = i.Platform
	i.fieldMap["login_ip"] = i.LoginIP
	i.fieldMap["version"] = i.Version
	i.fieldMap["create_time"] = i.CreateTime
	i.fieldMap["update_time"] = i.UpdateTime
}

func (i imLoginLog) clone(db *gorm.DB) imLoginLog {
	i.imLoginLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i imLoginLog) replaceDB(db *gorm.DB) imLoginLog {
	i.imLoginLogDo.ReplaceDB(db)
	return i
}

type imLoginLogDo struct{ gen.DO }

func (i imLoginLogDo) Debug() *imLoginLogDo {
	return i.withDO(i.DO.Debug())
}

func (i imLoginLogDo) WithContext(ctx context.Context) *imLoginLogDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i imLoginLogDo) ReadDB() *imLoginLogDo {
	return i.Clauses(dbresolver.Read)
}

func (i imLoginLogDo) WriteDB() *imLoginLogDo {
	return i.Clauses(dbresolver.Write)
}

func (i imLoginLogDo) Session(config *gorm.Session) *imLoginLogDo {
	return i.withDO(i.DO.Session(config))
}

func (i imLoginLogDo) Clauses(conds ...clause.Expression) *imLoginLogDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i imLoginLogDo) Returning(value interface{}, columns ...string) *imLoginLogDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i imLoginLogDo) Not(conds ...gen.Condition) *imLoginLogDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i imLoginLogDo) Or(conds ...gen.Condition) *imLoginLogDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i imLoginLogDo) Select(conds ...field.Expr) *imLoginLogDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i imLoginLogDo) Where(conds ...gen.Condition) *imLoginLogDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i imLoginLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *imLoginLogDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i imLoginLogDo) Order(conds ...field.Expr) *imLoginLogDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i imLoginLogDo) Distinct(cols ...field.Expr) *imLoginLogDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i imLoginLogDo) Omit(cols ...field.Expr) *imLoginLogDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i imLoginLogDo) Join(table schema.Tabler, on ...field.Expr) *imLoginLogDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i imLoginLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *imLoginLogDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i imLoginLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *imLoginLogDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i imLoginLogDo) Group(cols ...field.Expr) *imLoginLogDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i imLoginLogDo) Having(conds ...gen.Condition) *imLoginLogDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i imLoginLogDo) Limit(limit int) *imLoginLogDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i imLoginLogDo) Offset(offset int) *imLoginLogDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i imLoginLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *imLoginLogDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i imLoginLogDo) Unscoped() *imLoginLogDo {
	return i.withDO(i.DO.Unscoped())
}

func (i imLoginLogDo) Create(values ...*model.ImLoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i imLoginLogDo) CreateInBatches(values []*model.ImLoginLog, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i imLoginLogDo) Save(values ...*model.ImLoginLog) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i imLoginLogDo) First() (*model.ImLoginLog, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImLoginLog), nil
	}
}

func (i imLoginLogDo) Take() (*model.ImLoginLog, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImLoginLog), nil
	}
}

func (i imLoginLogDo) Last() (*model.ImLoginLog, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImLoginLog), nil
	}
}

func (i imLoginLogDo) Find() ([]*model.ImLoginLog, error) {
	result, err := i.DO.Find()
	return result.([]*model.ImLoginLog), err
}

func (i imLoginLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ImLoginLog, err error) {
	buf := make([]*model.ImLoginLog, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i imLoginLogDo) FindInBatches(result *[]*model.ImLoginLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i imLoginLogDo) Attrs(attrs ...field.AssignExpr) *imLoginLogDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i imLoginLogDo) Assign(attrs ...field.AssignExpr) *imLoginLogDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i imLoginLogDo) Joins(fields ...field.RelationField) *imLoginLogDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i imLoginLogDo) Preload(fields ...field.RelationField) *imLoginLogDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i imLoginLogDo) FirstOrInit() (*model.ImLoginLog, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImLoginLog), nil
	}
}

func (i imLoginLogDo) FirstOrCreate() (*model.ImLoginLog, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImLoginLog), nil
	}
}

func (i imLoginLogDo) FindByPage(offset int, limit int) (result []*model.ImLoginLog, count int64, err error) {
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

func (i imLoginLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i imLoginLogDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i imLoginLogDo) Delete(models ...*model.ImLoginLog) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *imLoginLogDo) withDO(do gen.Dao) *imLoginLogDo {
	i.DO = *do.(*gen.DO)
	return i
}