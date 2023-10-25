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

func newImUserSecurityQuestion(db *gorm.DB, opts ...gen.DOOption) imUserSecurityQuestion {
	_imUserSecurityQuestion := imUserSecurityQuestion{}

	_imUserSecurityQuestion.imUserSecurityQuestionDo.UseDB(db, opts...)
	_imUserSecurityQuestion.imUserSecurityQuestionDo.UseModel(&model.ImUserSecurityQuestion{})

	tableName := _imUserSecurityQuestion.imUserSecurityQuestionDo.TableName()
	_imUserSecurityQuestion.ALL = field.NewAsterisk(tableName)
	_imUserSecurityQuestion.UserID = field.NewString(tableName, "user_id")
	_imUserSecurityQuestion.QuestionID = field.NewInt32(tableName, "question_id")
	_imUserSecurityQuestion.Answer = field.NewString(tableName, "answer")
	_imUserSecurityQuestion.CreateTime = field.NewTime(tableName, "create_time")
	_imUserSecurityQuestion.UpdateTime = field.NewTime(tableName, "update_time")

	_imUserSecurityQuestion.fillFieldMap()

	return _imUserSecurityQuestion
}

type imUserSecurityQuestion struct {
	imUserSecurityQuestionDo imUserSecurityQuestionDo

	ALL        field.Asterisk
	UserID     field.String
	QuestionID field.Int32
	Answer     field.String
	CreateTime field.Time
	UpdateTime field.Time

	fieldMap map[string]field.Expr
}

func (i imUserSecurityQuestion) Table(newTableName string) *imUserSecurityQuestion {
	i.imUserSecurityQuestionDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i imUserSecurityQuestion) As(alias string) *imUserSecurityQuestion {
	i.imUserSecurityQuestionDo.DO = *(i.imUserSecurityQuestionDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *imUserSecurityQuestion) updateTableName(table string) *imUserSecurityQuestion {
	i.ALL = field.NewAsterisk(table)
	i.UserID = field.NewString(table, "user_id")
	i.QuestionID = field.NewInt32(table, "question_id")
	i.Answer = field.NewString(table, "answer")
	i.CreateTime = field.NewTime(table, "create_time")
	i.UpdateTime = field.NewTime(table, "update_time")

	i.fillFieldMap()

	return i
}

func (i *imUserSecurityQuestion) WithContext(ctx context.Context) *imUserSecurityQuestionDo {
	return i.imUserSecurityQuestionDo.WithContext(ctx)
}

func (i imUserSecurityQuestion) TableName() string { return i.imUserSecurityQuestionDo.TableName() }

func (i imUserSecurityQuestion) Alias() string { return i.imUserSecurityQuestionDo.Alias() }

func (i *imUserSecurityQuestion) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *imUserSecurityQuestion) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 5)
	i.fieldMap["user_id"] = i.UserID
	i.fieldMap["question_id"] = i.QuestionID
	i.fieldMap["answer"] = i.Answer
	i.fieldMap["create_time"] = i.CreateTime
	i.fieldMap["update_time"] = i.UpdateTime
}

func (i imUserSecurityQuestion) clone(db *gorm.DB) imUserSecurityQuestion {
	i.imUserSecurityQuestionDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i imUserSecurityQuestion) replaceDB(db *gorm.DB) imUserSecurityQuestion {
	i.imUserSecurityQuestionDo.ReplaceDB(db)
	return i
}

type imUserSecurityQuestionDo struct{ gen.DO }

func (i imUserSecurityQuestionDo) Debug() *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Debug())
}

func (i imUserSecurityQuestionDo) WithContext(ctx context.Context) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i imUserSecurityQuestionDo) ReadDB() *imUserSecurityQuestionDo {
	return i.Clauses(dbresolver.Read)
}

func (i imUserSecurityQuestionDo) WriteDB() *imUserSecurityQuestionDo {
	return i.Clauses(dbresolver.Write)
}

func (i imUserSecurityQuestionDo) Session(config *gorm.Session) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Session(config))
}

func (i imUserSecurityQuestionDo) Clauses(conds ...clause.Expression) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i imUserSecurityQuestionDo) Returning(value interface{}, columns ...string) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i imUserSecurityQuestionDo) Not(conds ...gen.Condition) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i imUserSecurityQuestionDo) Or(conds ...gen.Condition) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i imUserSecurityQuestionDo) Select(conds ...field.Expr) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i imUserSecurityQuestionDo) Where(conds ...gen.Condition) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i imUserSecurityQuestionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *imUserSecurityQuestionDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i imUserSecurityQuestionDo) Order(conds ...field.Expr) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i imUserSecurityQuestionDo) Distinct(cols ...field.Expr) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i imUserSecurityQuestionDo) Omit(cols ...field.Expr) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i imUserSecurityQuestionDo) Join(table schema.Tabler, on ...field.Expr) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i imUserSecurityQuestionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i imUserSecurityQuestionDo) RightJoin(table schema.Tabler, on ...field.Expr) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i imUserSecurityQuestionDo) Group(cols ...field.Expr) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i imUserSecurityQuestionDo) Having(conds ...gen.Condition) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i imUserSecurityQuestionDo) Limit(limit int) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i imUserSecurityQuestionDo) Offset(offset int) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i imUserSecurityQuestionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i imUserSecurityQuestionDo) Unscoped() *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Unscoped())
}

func (i imUserSecurityQuestionDo) Create(values ...*model.ImUserSecurityQuestion) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i imUserSecurityQuestionDo) CreateInBatches(values []*model.ImUserSecurityQuestion, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i imUserSecurityQuestionDo) Save(values ...*model.ImUserSecurityQuestion) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i imUserSecurityQuestionDo) First() (*model.ImUserSecurityQuestion, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImUserSecurityQuestion), nil
	}
}

func (i imUserSecurityQuestionDo) Take() (*model.ImUserSecurityQuestion, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImUserSecurityQuestion), nil
	}
}

func (i imUserSecurityQuestionDo) Last() (*model.ImUserSecurityQuestion, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImUserSecurityQuestion), nil
	}
}

func (i imUserSecurityQuestionDo) Find() ([]*model.ImUserSecurityQuestion, error) {
	result, err := i.DO.Find()
	return result.([]*model.ImUserSecurityQuestion), err
}

func (i imUserSecurityQuestionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ImUserSecurityQuestion, err error) {
	buf := make([]*model.ImUserSecurityQuestion, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i imUserSecurityQuestionDo) FindInBatches(result *[]*model.ImUserSecurityQuestion, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i imUserSecurityQuestionDo) Attrs(attrs ...field.AssignExpr) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i imUserSecurityQuestionDo) Assign(attrs ...field.AssignExpr) *imUserSecurityQuestionDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i imUserSecurityQuestionDo) Joins(fields ...field.RelationField) *imUserSecurityQuestionDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i imUserSecurityQuestionDo) Preload(fields ...field.RelationField) *imUserSecurityQuestionDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i imUserSecurityQuestionDo) FirstOrInit() (*model.ImUserSecurityQuestion, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImUserSecurityQuestion), nil
	}
}

func (i imUserSecurityQuestionDo) FirstOrCreate() (*model.ImUserSecurityQuestion, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImUserSecurityQuestion), nil
	}
}

func (i imUserSecurityQuestionDo) FindByPage(offset int, limit int) (result []*model.ImUserSecurityQuestion, count int64, err error) {
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

func (i imUserSecurityQuestionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i imUserSecurityQuestionDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i imUserSecurityQuestionDo) Delete(models ...*model.ImUserSecurityQuestion) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *imUserSecurityQuestionDo) withDO(do gen.Dao) *imUserSecurityQuestionDo {
	i.DO = *do.(*gen.DO)
	return i
}
