package gorm

import (
	"encoding/json"
	"fctl/pkg/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
)

var DB *gorm.DB

func ConnectDB(dsn string) (db *gorm.DB) {
	var err error

	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %w", err))
	}

	return db
}

type Relationship struct {
	Type              string `json:"type"`               //关联类型：belongs_to、has_one、has_many、many2many
	Table             string `json:"table"`              //关联表名
	Alias             string `json:"alias"`              //别名（可不声明，默认用表名）
	ForeignKey        string `json:"foreign_key"`        //外键（可不声明，默认为'id'或'表名_id'）
	ReferenceKey      string `json:"reference_key"`      //引用键（可不声明，默认为'id'或'表名_id'）
	RelationshipTable string `json:"relationship_table"` //当many2many时，连接表名
	JoinForeignKey    string `json:"join_foreign_key"`   //当many2many时，本表在连接表的外键
	JoinTargetKey     string `json:"join_target_key"`    //当many2many时，关联表在连接表的外键
}

// genModels is gorm/gen generated models
func genModels(g *gen.Generator, db *gorm.DB, tables []string) (models []interface{}, err error) {
	var tablesList []string

	if len(tables) == 0 || isEmpty(tables) {
		// Execute tasks for all tables in the database
		tablesList, err = db.Migrator().GetTables()

		if err != nil {
			return nil, fmt.Errorf("GORM migrator get all tables fail: %w", err)
		}
	} else {
		tablesList = tables
	}
	// Execute some data table tasks
	models = make([]interface{}, len(tablesList))
	for i, tableName := range tablesList {
		models[i] = g.GenerateModel(tableName)
	}
	g = relation(g)
	return models, nil
}

// 处理关系模型
func relation(g *gen.Generator) *gen.Generator {
	// 判断是否设置了关系模型
	root, _ := os.Getwd()
	file, err := os.ReadFile(root + "/config/database/config.json")
	var relationship map[string][]*Relationship
	if err == nil {
		err = json.Unmarshal(file, &relationship)
		if err != nil {
			panic("表关系JSON文件配置出错：" + err.Error())
		}
	}
	for key, val := range relationship {
		for _, v := range val {
			temp := g.GenerateModel(v.Table)
			var fieldType field.RelationshipType
			switch v.Type {
			case "belongs_to":
				fieldType = field.BelongsTo
				break
			case "has_one":
				fieldType = field.HasOne
				break
			case "has_many":
				fieldType = field.HasMany
				break
			case "many2many":
				fieldType = field.Many2Many
				break
			default:
				panic("关系类型错误，只支持belongs_to、has_one、has_many、many2many")
			}
			// 初始化
			tag := make(field.GormTag, 0)
			tag.Set("foreignKey", v.ForeignKey)
			relateConfig := field.RelateConfig{
				JSONTag: v.ForeignKey,
				GORMTag: tag,
			}
			if v.Type == "many2many" {
				if v.RelationshipTable == "" {
					panic("表" + v.Table + "的many2many必须声明连接表")
				}
				tag.Set("joinForeignKey", v.JoinForeignKey)
				tag.Set("joinReferences", v.JoinTargetKey)
			} else {
				if v.ForeignKey == "" {
					panic("表" + v.Table + "必须声明外键")
				}
			}
			if v.ReferenceKey != "" {
				tag.Set("references", v.ReferenceKey)
			}
			var up = g.GenerateModel(key, gen.FieldRelate(fieldType, v.Table, temp, &relateConfig))
			g.ApplyBasic(temp, up)
		}
	}
	return g
}

func isEmpty(arr []string) bool {
	for _, item := range arr {
		if item != "" {
			return false
		}
	}
	return true
}

func Run(opts ...Option) error {
	Init(opts...)
	DB = ConnectDB(global.CFG.MySQL.DSNURL()).Debug()
	g := gen.NewGenerator(gen.Config{
		FieldNullable:  false, // gorm pointer when field is nullable
		FieldCoverable: false, // gorm pointer when field has default value
		//FieldWithIndexTag: true, // gorm with gorm index tag
		FieldWithTypeTag: true, // gorm with gorm column type tag
		OutPath:          global.CFG.Dal.OutPath,
		OutFile:          global.CFG.Dal.OutFile,
		ModelPkgPath:     global.CFG.Dal.ModelPkgPath,
	})
	g.UseDB(DB)

	tables := strings.Split(global.CFG.Dal.Table, ",")
	models, err := genModels(g, DB, tables)
	if err != nil {
		log.Fatalln("get tables info fail:", err)
	}
	g.ApplyBasic(models...)
	g.Execute()
	return nil
}
