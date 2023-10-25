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

func newImGroupMember(db *gorm.DB, opts ...gen.DOOption) imGroupMember {
	_imGroupMember := imGroupMember{}

	_imGroupMember.imGroupMemberDo.UseDB(db, opts...)
	_imGroupMember.imGroupMemberDo.UseModel(&model.ImGroupMember{})

	tableName := _imGroupMember.imGroupMemberDo.TableName()
	_imGroupMember.ALL = field.NewAsterisk(tableName)
	_imGroupMember.GroupID = field.NewString(tableName, "group_id")
	_imGroupMember.UserID = field.NewString(tableName, "user_id")
	_imGroupMember.RoleLevel = field.NewInt32(tableName, "role_level")
	_imGroupMember.JoinTime = field.NewTime(tableName, "join_time")
	_imGroupMember.OperatorUserID = field.NewString(tableName, "operator_user_id")
	_imGroupMember.Status = field.NewInt32(tableName, "status")
	_imGroupMember.NotificationStatus = field.NewInt32(tableName, "notification_status")
	_imGroupMember.GroupName = field.NewString(tableName, "group_name")
	_imGroupMember.MemberType = field.NewInt32(tableName, "member_type")

	_imGroupMember.fillFieldMap()

	return _imGroupMember
}

type imGroupMember struct {
	imGroupMemberDo imGroupMemberDo

	ALL                field.Asterisk
	GroupID            field.String // 群组ID
	UserID             field.String // 用户ID
	RoleLevel          field.Int32  // 成员级别，1-群主，2-群管理员，3-群成员
	JoinTime           field.Time   // 入群时间
	OperatorUserID     field.String // 操作者的用户ID
	Status             field.Int32  // 1- 正常，2 - 禁言
	NotificationStatus field.Int32  // 群公告状态，1未读，2已读
	GroupName          field.String // 在群群昵称
	MemberType         field.Int32  // 用户类型，1-真实用户，2-虚拟用户

	fieldMap map[string]field.Expr
}

func (i imGroupMember) Table(newTableName string) *imGroupMember {
	i.imGroupMemberDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i imGroupMember) As(alias string) *imGroupMember {
	i.imGroupMemberDo.DO = *(i.imGroupMemberDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *imGroupMember) updateTableName(table string) *imGroupMember {
	i.ALL = field.NewAsterisk(table)
	i.GroupID = field.NewString(table, "group_id")
	i.UserID = field.NewString(table, "user_id")
	i.RoleLevel = field.NewInt32(table, "role_level")
	i.JoinTime = field.NewTime(table, "join_time")
	i.OperatorUserID = field.NewString(table, "operator_user_id")
	i.Status = field.NewInt32(table, "status")
	i.NotificationStatus = field.NewInt32(table, "notification_status")
	i.GroupName = field.NewString(table, "group_name")
	i.MemberType = field.NewInt32(table, "member_type")

	i.fillFieldMap()

	return i
}

func (i *imGroupMember) WithContext(ctx context.Context) *imGroupMemberDo {
	return i.imGroupMemberDo.WithContext(ctx)
}

func (i imGroupMember) TableName() string { return i.imGroupMemberDo.TableName() }

func (i imGroupMember) Alias() string { return i.imGroupMemberDo.Alias() }

func (i *imGroupMember) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *imGroupMember) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 9)
	i.fieldMap["group_id"] = i.GroupID
	i.fieldMap["user_id"] = i.UserID
	i.fieldMap["role_level"] = i.RoleLevel
	i.fieldMap["join_time"] = i.JoinTime
	i.fieldMap["operator_user_id"] = i.OperatorUserID
	i.fieldMap["status"] = i.Status
	i.fieldMap["notification_status"] = i.NotificationStatus
	i.fieldMap["group_name"] = i.GroupName
	i.fieldMap["member_type"] = i.MemberType
}

func (i imGroupMember) clone(db *gorm.DB) imGroupMember {
	i.imGroupMemberDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i imGroupMember) replaceDB(db *gorm.DB) imGroupMember {
	i.imGroupMemberDo.ReplaceDB(db)
	return i
}

type imGroupMemberDo struct{ gen.DO }

func (i imGroupMemberDo) Debug() *imGroupMemberDo {
	return i.withDO(i.DO.Debug())
}

func (i imGroupMemberDo) WithContext(ctx context.Context) *imGroupMemberDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i imGroupMemberDo) ReadDB() *imGroupMemberDo {
	return i.Clauses(dbresolver.Read)
}

func (i imGroupMemberDo) WriteDB() *imGroupMemberDo {
	return i.Clauses(dbresolver.Write)
}

func (i imGroupMemberDo) Session(config *gorm.Session) *imGroupMemberDo {
	return i.withDO(i.DO.Session(config))
}

func (i imGroupMemberDo) Clauses(conds ...clause.Expression) *imGroupMemberDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i imGroupMemberDo) Returning(value interface{}, columns ...string) *imGroupMemberDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i imGroupMemberDo) Not(conds ...gen.Condition) *imGroupMemberDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i imGroupMemberDo) Or(conds ...gen.Condition) *imGroupMemberDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i imGroupMemberDo) Select(conds ...field.Expr) *imGroupMemberDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i imGroupMemberDo) Where(conds ...gen.Condition) *imGroupMemberDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i imGroupMemberDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *imGroupMemberDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i imGroupMemberDo) Order(conds ...field.Expr) *imGroupMemberDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i imGroupMemberDo) Distinct(cols ...field.Expr) *imGroupMemberDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i imGroupMemberDo) Omit(cols ...field.Expr) *imGroupMemberDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i imGroupMemberDo) Join(table schema.Tabler, on ...field.Expr) *imGroupMemberDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i imGroupMemberDo) LeftJoin(table schema.Tabler, on ...field.Expr) *imGroupMemberDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i imGroupMemberDo) RightJoin(table schema.Tabler, on ...field.Expr) *imGroupMemberDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i imGroupMemberDo) Group(cols ...field.Expr) *imGroupMemberDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i imGroupMemberDo) Having(conds ...gen.Condition) *imGroupMemberDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i imGroupMemberDo) Limit(limit int) *imGroupMemberDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i imGroupMemberDo) Offset(offset int) *imGroupMemberDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i imGroupMemberDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *imGroupMemberDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i imGroupMemberDo) Unscoped() *imGroupMemberDo {
	return i.withDO(i.DO.Unscoped())
}

func (i imGroupMemberDo) Create(values ...*model.ImGroupMember) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i imGroupMemberDo) CreateInBatches(values []*model.ImGroupMember, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i imGroupMemberDo) Save(values ...*model.ImGroupMember) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i imGroupMemberDo) First() (*model.ImGroupMember, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImGroupMember), nil
	}
}

func (i imGroupMemberDo) Take() (*model.ImGroupMember, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImGroupMember), nil
	}
}

func (i imGroupMemberDo) Last() (*model.ImGroupMember, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImGroupMember), nil
	}
}

func (i imGroupMemberDo) Find() ([]*model.ImGroupMember, error) {
	result, err := i.DO.Find()
	return result.([]*model.ImGroupMember), err
}

func (i imGroupMemberDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ImGroupMember, err error) {
	buf := make([]*model.ImGroupMember, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i imGroupMemberDo) FindInBatches(result *[]*model.ImGroupMember, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i imGroupMemberDo) Attrs(attrs ...field.AssignExpr) *imGroupMemberDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i imGroupMemberDo) Assign(attrs ...field.AssignExpr) *imGroupMemberDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i imGroupMemberDo) Joins(fields ...field.RelationField) *imGroupMemberDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i imGroupMemberDo) Preload(fields ...field.RelationField) *imGroupMemberDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i imGroupMemberDo) FirstOrInit() (*model.ImGroupMember, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImGroupMember), nil
	}
}

func (i imGroupMemberDo) FirstOrCreate() (*model.ImGroupMember, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImGroupMember), nil
	}
}

func (i imGroupMemberDo) FindByPage(offset int, limit int) (result []*model.ImGroupMember, count int64, err error) {
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

func (i imGroupMemberDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i imGroupMemberDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i imGroupMemberDo) Delete(models ...*model.ImGroupMember) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *imGroupMemberDo) withDO(do gen.Dao) *imGroupMemberDo {
	i.DO = *do.(*gen.DO)
	return i
}