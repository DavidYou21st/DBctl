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

func newImFunctionalConfig(db *gorm.DB, opts ...gen.DOOption) imFunctionalConfig {
	_imFunctionalConfig := imFunctionalConfig{}

	_imFunctionalConfig.imFunctionalConfigDo.UseDB(db, opts...)
	_imFunctionalConfig.imFunctionalConfigDo.UseModel(&model.ImFunctionalConfig{})

	tableName := _imFunctionalConfig.imFunctionalConfigDo.TableName()
	_imFunctionalConfig.ALL = field.NewAsterisk(tableName)
	_imFunctionalConfig.ID = field.NewInt32(tableName, "id")
	_imFunctionalConfig.ProhibitRegister = field.NewInt32(tableName, "prohibitRegister")
	_imFunctionalConfig.ProhibitAddFriend = field.NewInt32(tableName, "prohibitAddFriend")
	_imFunctionalConfig.ProhibitCreateGroup = field.NewInt32(tableName, "prohibitCreateGroup")
	_imFunctionalConfig.ProhibitGroupSendPicture = field.NewInt32(tableName, "prohibitGroupSendPicture")
	_imFunctionalConfig.ProhibitGroupSendVideo = field.NewInt32(tableName, "prohibitGroupSendVideo")
	_imFunctionalConfig.ProhibitGroupSendCard = field.NewInt32(tableName, "prohibitGroupSendCard")
	_imFunctionalConfig.ProhibitSearch = field.NewInt32(tableName, "prohibitSearch")
	_imFunctionalConfig.MultiplePlatform = field.NewInt32(tableName, "multiplePlatform")
	_imFunctionalConfig.Android = field.NewInt32(tableName, "android")
	_imFunctionalConfig.Iphone = field.NewInt32(tableName, "iphone")
	_imFunctionalConfig.Windows = field.NewInt32(tableName, "windows")
	_imFunctionalConfig.Mac = field.NewInt32(tableName, "mac")
	_imFunctionalConfig.ProhibitUserAddUser = field.NewInt32(tableName, "prohibitUserAddUser")
	_imFunctionalConfig.ProhibitUserAddVip = field.NewInt32(tableName, "prohibitUserAddVip")
	_imFunctionalConfig.ProhibitVipAddVip = field.NewInt32(tableName, "prohibitVipAddVip")

	_imFunctionalConfig.fillFieldMap()

	return _imFunctionalConfig
}

type imFunctionalConfig struct {
	imFunctionalConfigDo imFunctionalConfigDo

	ALL                      field.Asterisk
	ID                       field.Int32
	ProhibitRegister         field.Int32 // 禁止客户端注册，1-开启，2-没开启
	ProhibitAddFriend        field.Int32 // 禁止普通用户添加好友，1-开启，2-没开启
	ProhibitCreateGroup      field.Int32 // 禁止普通用户建群，1-开启，2-没开启
	ProhibitGroupSendPicture field.Int32 // 禁止群内发送照片，1-开启，2-没开启
	ProhibitGroupSendVideo   field.Int32 // 禁止群内发送视频，1-开启，2-没开启
	ProhibitGroupSendCard    field.Int32 // 禁止群内发送名片，1-开启，2-没开启
	ProhibitSearch           field.Int32 // 禁止搜索好友，1-开启，2-没开启
	MultiplePlatform         field.Int32 // 登录设置，1-单平台，2-多平台
	Android                  field.Int32 // 允许安卓多平台登录，1-开启，2-没开启
	Iphone                   field.Int32 // 允许iphone多平台登录，1-开启，2-没开启
	Windows                  field.Int32 // 允许window多平台登录，1-开启，2-没开启
	Mac                      field.Int32 // 允许mac多平台登录，1-开启，2-没开启
	ProhibitUserAddUser      field.Int32 // 禁止普通用户添加普通用户，1-开启，2-没开启
	ProhibitUserAddVip       field.Int32 // 禁止普通用户添加特权用户，1-开启，2-没开启
	ProhibitVipAddVip        field.Int32 // 禁止特权用户添加特权用户，1-开启，2-没开启

	fieldMap map[string]field.Expr
}

func (i imFunctionalConfig) Table(newTableName string) *imFunctionalConfig {
	i.imFunctionalConfigDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i imFunctionalConfig) As(alias string) *imFunctionalConfig {
	i.imFunctionalConfigDo.DO = *(i.imFunctionalConfigDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *imFunctionalConfig) updateTableName(table string) *imFunctionalConfig {
	i.ALL = field.NewAsterisk(table)
	i.ID = field.NewInt32(table, "id")
	i.ProhibitRegister = field.NewInt32(table, "prohibitRegister")
	i.ProhibitAddFriend = field.NewInt32(table, "prohibitAddFriend")
	i.ProhibitCreateGroup = field.NewInt32(table, "prohibitCreateGroup")
	i.ProhibitGroupSendPicture = field.NewInt32(table, "prohibitGroupSendPicture")
	i.ProhibitGroupSendVideo = field.NewInt32(table, "prohibitGroupSendVideo")
	i.ProhibitGroupSendCard = field.NewInt32(table, "prohibitGroupSendCard")
	i.ProhibitSearch = field.NewInt32(table, "prohibitSearch")
	i.MultiplePlatform = field.NewInt32(table, "multiplePlatform")
	i.Android = field.NewInt32(table, "android")
	i.Iphone = field.NewInt32(table, "iphone")
	i.Windows = field.NewInt32(table, "windows")
	i.Mac = field.NewInt32(table, "mac")
	i.ProhibitUserAddUser = field.NewInt32(table, "prohibitUserAddUser")
	i.ProhibitUserAddVip = field.NewInt32(table, "prohibitUserAddVip")
	i.ProhibitVipAddVip = field.NewInt32(table, "prohibitVipAddVip")

	i.fillFieldMap()

	return i
}

func (i *imFunctionalConfig) WithContext(ctx context.Context) *imFunctionalConfigDo {
	return i.imFunctionalConfigDo.WithContext(ctx)
}

func (i imFunctionalConfig) TableName() string { return i.imFunctionalConfigDo.TableName() }

func (i imFunctionalConfig) Alias() string { return i.imFunctionalConfigDo.Alias() }

func (i *imFunctionalConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *imFunctionalConfig) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 16)
	i.fieldMap["id"] = i.ID
	i.fieldMap["prohibitRegister"] = i.ProhibitRegister
	i.fieldMap["prohibitAddFriend"] = i.ProhibitAddFriend
	i.fieldMap["prohibitCreateGroup"] = i.ProhibitCreateGroup
	i.fieldMap["prohibitGroupSendPicture"] = i.ProhibitGroupSendPicture
	i.fieldMap["prohibitGroupSendVideo"] = i.ProhibitGroupSendVideo
	i.fieldMap["prohibitGroupSendCard"] = i.ProhibitGroupSendCard
	i.fieldMap["prohibitSearch"] = i.ProhibitSearch
	i.fieldMap["multiplePlatform"] = i.MultiplePlatform
	i.fieldMap["android"] = i.Android
	i.fieldMap["iphone"] = i.Iphone
	i.fieldMap["windows"] = i.Windows
	i.fieldMap["mac"] = i.Mac
	i.fieldMap["prohibitUserAddUser"] = i.ProhibitUserAddUser
	i.fieldMap["prohibitUserAddVip"] = i.ProhibitUserAddVip
	i.fieldMap["prohibitVipAddVip"] = i.ProhibitVipAddVip
}

func (i imFunctionalConfig) clone(db *gorm.DB) imFunctionalConfig {
	i.imFunctionalConfigDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i imFunctionalConfig) replaceDB(db *gorm.DB) imFunctionalConfig {
	i.imFunctionalConfigDo.ReplaceDB(db)
	return i
}

type imFunctionalConfigDo struct{ gen.DO }

func (i imFunctionalConfigDo) Debug() *imFunctionalConfigDo {
	return i.withDO(i.DO.Debug())
}

func (i imFunctionalConfigDo) WithContext(ctx context.Context) *imFunctionalConfigDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i imFunctionalConfigDo) ReadDB() *imFunctionalConfigDo {
	return i.Clauses(dbresolver.Read)
}

func (i imFunctionalConfigDo) WriteDB() *imFunctionalConfigDo {
	return i.Clauses(dbresolver.Write)
}

func (i imFunctionalConfigDo) Session(config *gorm.Session) *imFunctionalConfigDo {
	return i.withDO(i.DO.Session(config))
}

func (i imFunctionalConfigDo) Clauses(conds ...clause.Expression) *imFunctionalConfigDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i imFunctionalConfigDo) Returning(value interface{}, columns ...string) *imFunctionalConfigDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i imFunctionalConfigDo) Not(conds ...gen.Condition) *imFunctionalConfigDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i imFunctionalConfigDo) Or(conds ...gen.Condition) *imFunctionalConfigDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i imFunctionalConfigDo) Select(conds ...field.Expr) *imFunctionalConfigDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i imFunctionalConfigDo) Where(conds ...gen.Condition) *imFunctionalConfigDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i imFunctionalConfigDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *imFunctionalConfigDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i imFunctionalConfigDo) Order(conds ...field.Expr) *imFunctionalConfigDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i imFunctionalConfigDo) Distinct(cols ...field.Expr) *imFunctionalConfigDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i imFunctionalConfigDo) Omit(cols ...field.Expr) *imFunctionalConfigDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i imFunctionalConfigDo) Join(table schema.Tabler, on ...field.Expr) *imFunctionalConfigDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i imFunctionalConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) *imFunctionalConfigDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i imFunctionalConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) *imFunctionalConfigDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i imFunctionalConfigDo) Group(cols ...field.Expr) *imFunctionalConfigDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i imFunctionalConfigDo) Having(conds ...gen.Condition) *imFunctionalConfigDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i imFunctionalConfigDo) Limit(limit int) *imFunctionalConfigDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i imFunctionalConfigDo) Offset(offset int) *imFunctionalConfigDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i imFunctionalConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *imFunctionalConfigDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i imFunctionalConfigDo) Unscoped() *imFunctionalConfigDo {
	return i.withDO(i.DO.Unscoped())
}

func (i imFunctionalConfigDo) Create(values ...*model.ImFunctionalConfig) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i imFunctionalConfigDo) CreateInBatches(values []*model.ImFunctionalConfig, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i imFunctionalConfigDo) Save(values ...*model.ImFunctionalConfig) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i imFunctionalConfigDo) First() (*model.ImFunctionalConfig, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFunctionalConfig), nil
	}
}

func (i imFunctionalConfigDo) Take() (*model.ImFunctionalConfig, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFunctionalConfig), nil
	}
}

func (i imFunctionalConfigDo) Last() (*model.ImFunctionalConfig, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFunctionalConfig), nil
	}
}

func (i imFunctionalConfigDo) Find() ([]*model.ImFunctionalConfig, error) {
	result, err := i.DO.Find()
	return result.([]*model.ImFunctionalConfig), err
}

func (i imFunctionalConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ImFunctionalConfig, err error) {
	buf := make([]*model.ImFunctionalConfig, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i imFunctionalConfigDo) FindInBatches(result *[]*model.ImFunctionalConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i imFunctionalConfigDo) Attrs(attrs ...field.AssignExpr) *imFunctionalConfigDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i imFunctionalConfigDo) Assign(attrs ...field.AssignExpr) *imFunctionalConfigDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i imFunctionalConfigDo) Joins(fields ...field.RelationField) *imFunctionalConfigDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i imFunctionalConfigDo) Preload(fields ...field.RelationField) *imFunctionalConfigDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i imFunctionalConfigDo) FirstOrInit() (*model.ImFunctionalConfig, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFunctionalConfig), nil
	}
}

func (i imFunctionalConfigDo) FirstOrCreate() (*model.ImFunctionalConfig, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ImFunctionalConfig), nil
	}
}

func (i imFunctionalConfigDo) FindByPage(offset int, limit int) (result []*model.ImFunctionalConfig, count int64, err error) {
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

func (i imFunctionalConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i imFunctionalConfigDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i imFunctionalConfigDo) Delete(models ...*model.ImFunctionalConfig) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *imFunctionalConfigDo) withDO(do gen.Dao) *imFunctionalConfigDo {
	i.DO = *do.(*gen.DO)
	return i
}
