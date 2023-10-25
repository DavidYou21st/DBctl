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

func newImFriend(db *gorm.DB, opts ...gen.DOOption) imFriend {
	_imFriend := imFriend{}

	_imFriend.imFriendDo.UseDB(db, opts...)
	_imFriend.imFriendDo.UseModel(&model.ImFriend{})

	tableName := _imFriend.imFriendDo.TableName()
	_imFriend.ALL = field.NewAsterisk(tableName)
	_imFriend.OwnerUserID = field.NewString(tableName, "owner_user_id")
	_imFriend.FriendUserID = field.NewString(tableName, "friend_user_id")
	_imFriend.Remark = field.NewString(tableName, "remark")
	_imFriend.CreateTime = field.NewTime(tableName, "create_time")
	_imFriend.AddSource = field.NewInt32(tableName, "add_source")
	_imFriend.OperatorUserID = field.NewString(tableName, "operator_user_id")
	_imFriend.Ex = field.NewString(tableName, "ex")
	_imFriend.Status = field.NewInt32(tableName, "status")
	_imFriend.IsDeleted = field.NewInt32(tableName, "is_deleted")

	_imFriend.fillFieldMap()

	return _imFriend
}

type imFriend struct {
	imFriendDo imFriendDo

	ALL            field.Asterisk
	OwnerUserID    field.String // 用户的UserID
	FriendUserID   field.String // 好友的UserID
	Remark         field.String // 好友备注
	CreateTime     field.Time   // 添加时间
	AddSource      field.Int32  // 添加来源，1-ID,2-扫码,3-名片，4-群聊，5-其他
	OperatorUserID field.String // 处理者的UserID
	Ex             field.String // 拓展字段
	Status         field.Int32  // 好友状态 0=在黑名单中，1=正常
	IsDeleted      field.Int32  // 是否是已删除的好友 1=否，2=是

	fieldMap map[string]field.Expr
}

func (i imFriend) Table(newTableName string) *imFriend {
	i.imFriendDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i imFriend) As(alias string) *imFriend {
	i.imFriendDo.DO = *(i.imFriendDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *imFriend) updateTableName(table string) *imFriend {
	i.ALL = field.NewAsterisk(table)
	i.OwnerUserID = field.NewString(table, "owner_user_id")
	i.FriendUserID = field.NewString(table, "friend_user_id")
	i.Remark = field.NewString(table, "remark")
	i.CreateTime = field.NewTime(table, "create_time")
	i.AddSource = field.NewInt32(table, "add_source")
	i.OperatorUserID = field.NewString(table, "operator_user_id")
	i.Ex = field.NewString(table, "ex")
	i.Status = field.NewInt32(table, "status")
	i.IsDeleted = field.NewInt32(table, "is_deleted")

	i.fillFieldMap()

	return i
}

func (i *imFriend) WithContext(ctx context.Context) *imFriendDo { return i.imFriendDo.WithContext(ctx) }

func (i imFriend) TableName() string { return i.imFriendDo.TableName() }

func (i imFriend) Alias() string { return i.imFriendDo.Alias() }

func (i *imFriend) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *imFriend) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 9)
	i.fieldMap["owner_user_id"] = i.OwnerUserID
	i.fieldMap["friend_user_id"] = i.FriendUserID
	i.fieldMap["remark"] = i.Remark
	i.fieldMap["create_time"] = i.CreateTime
	i.fieldMap["add_source"] = i.AddSource
	i.fieldMap["operator_user_id"] = i.OperatorUserID
	i.fieldMap["ex"] = i.Ex
	i.fieldMap["status"] = i.Status
	i.fieldMap["is_deleted"] = i.IsDeleted
}

func (i imFriend) clone(db *gorm.DB) imFriend {
	i.imFriendDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i imFriend) replaceDB(db *gorm.DB) imFriend {
	i.imFriendDo.ReplaceDB(db)
	return i
}

type imFriendDo struct{ gen.DO }

func (i imFriendDo) Debug() *imFriendDo {
	return i.withDO(i.DO.Debug())
}

func (i imFriendDo) WithContext(ctx context.Context) *imFriendDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i imFriendDo) ReadDB() *imFriendDo {
	return i.Clauses(dbresolver.Read)
}

func (i imFriendDo) WriteDB() *imFriendDo {
	return i.Clauses(dbresolver.Write)
}

func (i imFriendDo) Session(config *gorm.Session) *imFriendDo {
	return i.withDO(i.DO.Session(config))
}

func (i imFriendDo) Clauses(conds ...clause.Expression) *imFriendDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i imFriendDo) Returning(value interface{}, columns ...string) *imFriendDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i imFriendDo) Not(conds ...gen.Condition) *imFriendDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i imFriendDo) Or(conds ...gen.Condition) *imFriendDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i imFriendDo) Select(conds ...field.Expr) *imFriendDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i imFriendDo) Where(conds ...gen.Condition) *imFriendDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i imFriendDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *imFriendDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i imFriendDo) Order(conds ...field.Expr) *imFriendDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i imFriendDo) Distinct(cols ...field.Expr) *imFriendDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i imFriendDo) Omit(cols ...field.Expr) *imFriendDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i imFriendDo) Join(table schema.Tabler, on ...field.Expr) *imFriendDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i imFriendDo) LeftJoin(table schema.Tabler, on ...field.Expr) *imFriendDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i imFriendDo) RightJoin(table schema.Tabler, on ...field.Expr) *imFriendDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i imFriendDo) Group(cols ...field.Expr) *imFriendDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i imFriendDo) Having(conds ...gen.Condition) *imFriendDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i imFriendDo) Limit(limit int) *imFriendDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i imFriendDo) Offset(offset int) *imFriendDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i imFriendDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *imFriendDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i imFriendDo) Unscoped() *imFriendDo {
	return i.withDO(i.DO.Unscoped())
}

func (i imFriendDo) Create(values ...*model.ImFriend) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i imFriendDo) CreateInBatches(values []*model.ImFriend, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i imFriendDo) Save(values ...*model.ImFriend) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i imFriendDo) First() (*model.ImFriend, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFriend), nil
	}
}

func (i imFriendDo) Take() (*model.ImFriend, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFriend), nil
	}
}

func (i imFriendDo) Last() (*model.ImFriend, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFriend), nil
	}
}

func (i imFriendDo) Find() ([]*model.ImFriend, error) {
	result, err := i.DO.Find()
	return result.([]*model.ImFriend), err
}

func (i imFriendDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ImFriend, err error) {
	buf := make([]*model.ImFriend, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i imFriendDo) FindInBatches(result *[]*model.ImFriend, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i imFriendDo) Attrs(attrs ...field.AssignExpr) *imFriendDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i imFriendDo) Assign(attrs ...field.AssignExpr) *imFriendDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i imFriendDo) Joins(fields ...field.RelationField) *imFriendDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i imFriendDo) Preload(fields ...field.RelationField) *imFriendDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i imFriendDo) FirstOrInit() (*model.ImFriend, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFriend), nil
	}
}

func (i imFriendDo) FirstOrCreate() (*model.ImFriend, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFriend), nil
	}
}

func (i imFriendDo) FindByPage(offset int, limit int) (result []*model.ImFriend, count int64, err error) {
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

func (i imFriendDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i imFriendDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i imFriendDo) Delete(models ...*model.ImFriend) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *imFriendDo) withDO(do gen.Dao) *imFriendDo {
	i.DO = *do.(*gen.DO)
	return i
}