package models

import (
	"fmt"
	"database/sql"

	"github.com/astaxie/beedb"
	_ "github.com/ziutek/mymysql/godrv"
)

//计算器运算记录
type CalculationRecord struct{
	Id int `PK`
	Expression string
	Result float64
}

//连接数据库
func GetLink() beedb.Model{
	db,err:=sql.Open("mymysql","calculator/huwh/123456")
	if err!=nil{
		fmt.Println(err)
	}
	orm:=beedb.New(db)
	return orm
}

//添加记录
func SaveCalculationRecord(calculationRecord CalculationRecord) (record CalculationRecord){
	db:=GetLink()
	db.Save(&CalculationRecord)
	return record
}

//查询历史记录
func  GetAllCalculationRecord() (records []CalculationRecord){
	db:=GetLink()
	db.FindAll(&records)
	return
}