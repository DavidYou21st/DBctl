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

func newImSensitiveLib(db *gorm.DB, opts ...gen.DOOption) imSensitiveLib {
	_imSensitiveLib := imSensitiveLib{}

	_imSensitiveLib.imSensitiveLibDo.UseDB(db, opts...)
	_imSensitiveLib.imSensitiveLibDo.UseModel(&model.ImSensitiveLib{})

	tableName := _imSensitiveLib.imSensitiveLibDo.TableName()
	_imSensitiveLib.ALL = field.NewAsterisk(tableName)
	_imSensitiveLib.ID = field.NewString(tableName, "id")
	_imSensitiveLib.Name = field.NewString(tableName, "name")
	_imSensitiveLib.Describe = field.NewString(tableName, "Describe")
	_imSensitiveLib.CreateTime = field.NewTime(tableName, "create_time")
	_imSensitiveLib.Suggestion = field.NewString(tableName, "Suggestion")

	_imSensitiveLib.fillFieldMap()

	return _imSensitiveLib
}

type imSensitiveLib struct {
	imSensitiveLibDo imSensitiveLibDo

	ALL        field.Asterisk
	ID         field.String // 库ID
	Name       field.String // 库名称
	Describe   field.String // 库描述
	CreateTime field.Time   // 创建时间
	Suggestion field.String // 库建议

	fieldMap map[string]field.Expr
}

func (i imSensitiveLib) Table(newTableName string) *imSensitiveLib {
	i.imSensitiveLibDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i imSensitiveLib) As(alias string) *imSensitiveLib {
	i.imSensitiveLibDo.DO = *(i.imSensitiveLibDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *imSensitiveLib) updateTableName(table string) *imSensitiveLib {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewString(table, "id")
	i.Name = field.NewString(table, "name")
	i.Describe = field.NewString(table, "Describe")
	i.CreateTime = field.NewTime(table, "create_time")
	i.Suggestion = field.NewString(table, "Suggestion")

	i.fillFieldMap()

	return i
}

func (i *imSensitiveLib) WithContext(ctx context.Context) *imSensitiveLibDo {
	return i.imSensitiveLibDo.WithContext(ctx)
}

func (i imSensitiveLib) TableName() string { return i.imSensitiveLibDo.TableName() }

func (i imSensitiveLib) Alias() string { return i.imSensitiveLibDo.Alias() }

func (i *imSensitiveLib) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *imSensitiveLib) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 5)
	i.fieldMap["id"] = i.ID
	i.fieldMap["name"] = i.Name
	i.fieldMap["Describe"] = i.Describe
	i.fieldMap["create_time"] = i.CreateTime
	i.fieldMap["Suggestion"] = i.Suggestion
}

func (i imSensitiveLib) clone(db *gorm.DB) imSensitiveLib {
	i.imSensitiveLibDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i imSensitiveLib) replaceDB(db *gorm.DB) imSensitiveLib {
	i.imSensitiveLibDo.ReplaceDB(db)
	return i
}

type imSensitiveLibDo struct{ gen.DO }

func (i imSensitiveLibDo) Debug() *imSensitiveLibDo {
	return i.withDO(i.DO.Debug())
}

func (i imSensitiveLibDo) WithContext(ctx context.Context) *imSensitiveLibDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i imSensitiveLibDo) ReadDB() *imSensitiveLibDo {
	return i.Clauses(dbresolver.Read)
}

func (i imSensitiveLibDo) WriteDB() *imSensitiveLibDo {
	return i.Clauses(dbresolver.Write)
}

func (i imSensitiveLibDo) Session(config *gorm.Session) *imSensitiveLibDo {
	return i.withDO(i.DO.Session(config))
}

func (i imSensitiveLibDo) Clauses(conds ...clause.Expression) *imSensitiveLibDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i imSensitiveLibDo) Returning(value interface{}, columns ...string) *imSensitiveLibDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i imSensitiveLibDo) Not(conds ...gen.Condition) *imSensitiveLibDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i imSensitiveLibDo) Or(conds ...gen.Condition) *imSensitiveLibDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i imSensitiveLibDo) Select(conds ...field.Expr) *imSensitiveLibDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i imSensitiveLibDo) Where(conds ...gen.Condition) *imSensitiveLibDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i imSensitiveLibDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *imSensitiveLibDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i imSensitiveLibDo) Order(conds ...field.Expr) *imSensitiveLibDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i imSensitiveLibDo) Distinct(cols ...field.Expr) *imSensitiveLibDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i imSensitiveLibDo) Omit(cols ...field.Expr) *imSensitiveLibDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i imSensitiveLibDo) Join(table schema.Tabler, on ...field.Expr) *imSensitiveLibDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i imSensitiveLibDo) LeftJoin(table schema.Tabler, on ...field.Expr) *imSensitiveLibDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i imSensitiveLibDo) RightJoin(table schema.Tabler, on ...field.Expr) *imSensitiveLibDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i imSensitiveLibDo) Group(cols ...field.Expr) *imSensitiveLibDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i imSensitiveLibDo) Having(conds ...gen.Condition) *imSensitiveLibDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i imSensitiveLibDo) Limit(limit int) *imSensitiveLibDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i imSensitiveLibDo) Offset(offset int) *imSensitiveLibDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i imSensitiveLibDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *imSensitiveLibDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i imSensitiveLibDo) Unscoped() *imSensitiveLibDo {
	return i.withDO(i.DO.Unscoped())
}

func (i imSensitiveLibDo) Create(values ...*model.ImSensitiveLib) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i imSensitiveLibDo) CreateInBatches(values []*model.ImSensitiveLib, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i imSensitiveLibDo) Save(values ...*model.ImSensitiveLib) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i imSensitiveLibDo) First() (*model.ImSensitiveLib, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImSensitiveLib), nil
	}
}

func (i imSensitiveLibDo) Take() (*model.ImSensitiveLib, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImSensitiveLib), nil
	}
}

func (i imSensitiveLibDo) Last() (*model.ImSensitiveLib, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImSensitiveLib), nil
	}
}

func (i imSensitiveLibDo) Find() ([]*model.ImSensitiveLib, error) {
	result, err := i.DO.Find()
	return result.([]*model.ImSensitiveLib), err
}

func (i imSensitiveLibDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ImSensitiveLib, err error) {
	buf := make([]*model.ImSensitiveLib, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i imSensitiveLibDo) FindInBatches(result *[]*model.ImSensitiveLib, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i imSensitiveLibDo) Attrs(attrs ...field.AssignExpr) *imSensitiveLibDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i imSensitiveLibDo) Assign(attrs ...field.AssignExpr) *imSensitiveLibDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i imSensitiveLibDo) Joins(fields ...field.RelationField) *imSensitiveLibDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i imSensitiveLibDo) Preload(fields ...field.RelationField) *imSensitiveLibDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i imSensitiveLibDo) FirstOrInit() (*model.ImSensitiveLib, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImSensitiveLib), nil
	}
}

func (i imSensitiveLibDo) FirstOrCreate() (*model.ImSensitiveLib, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImSensitiveLib), nil
	}
}

func (i imSensitiveLibDo) FindByPage(offset int, limit int) (result []*model.ImSensitiveLib, count int64, err error) {
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

func (i imSensitiveLibDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i imSensitiveLibDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i imSensitiveLibDo) Delete(models ...*model.ImSensitiveLib) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *imSensitiveLibDo) withDO(do gen.Dao) *imSensitiveLibDo {
	i.DO = *do.(*gen.DO)
	return i
}