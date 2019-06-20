package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"time"
)

//表的设计

//用户
type User struct {
	Id int	`orm:"pk;auto"`
	IsActive bool `orm:"default(false)"`	//激活
	Email string
	Name string
	PassWord string
	Articles []*Article `orm:"reverse(many)"`
	Adminrivilege int `orm:"default(0);size(1)"` //0普通1管理员2超级管理员
}
//文章标签
type ArticLabel struct {
	Id int
	Name string
	Articles []*Article `orm:"reverse(many)"`
}
//文章
type Article struct {
	Id int `orm:"pk;auto"`
	ArtiName string `orm:"size(20)"`//标题
	Atime time.Time `orm:"auto_now"`
	Acount int `orm:"default(0);null"`
	Acontent string	`orm:"type(text)"`
	ArticLabels []*ArticLabel `orm:"rel(m2m)"`
	TwoArticleType *TwoArticleType `orm:"rel(fk);on_delete(set_null);null"`
	User *User `orm:"rel(fk)"`
	ArticIndexBanner []*ArticIndexBanner `orm:"reverse(many)"`
	ArticBanner []*ArticBanner `orm:"reverse(many)"`
	Top int `orm:"default(9);size(1)"`    //置顶
	Key string	`orm:"unique"`
	IsDel bool	`orm:"default(false)"`
}
//点赞表
type Praise struct {
	Id int `orm:"pk;auto"`
	ArticleId int
	UserId int
}
//1级文章类型表
type ArticleType struct {
	Id int
	TypeName string `orm:"size(20)"`
	//TwoArticleType *TwoArticleType `orm:"rel(fk);on_delete(set_null);null"`
	TwoArticleType []*TwoArticleType `orm:"reverse(many)"`
}
//2级文章类型表
type TwoArticleType struct {
	Id int
	TypeName string `orm:"size(20)"`
	ArticleType *ArticleType `orm:"rel(fk)"`
	Article []*Article `orm:"reverse(many)"`
}


//首页滚动展示
type ArticIndexBanner struct {
	Id int
	Img string
	Article *Article `orm:"rel(fk)"`
}
//首页文章固定展示
type ArticBanner struct {
	Id int
	Img string
	Article *Article `orm:"rel(fk)"`
}
//图片
type ArticIMG struct {
	Id int
	Key string
	Addr string
}
func init(){
//ORM操作数据库
//获取连接对象
orm.RegisterDataBase("default","mysql","root:rootroot@tcp(127.0.0.1:3306)/test?charset=utf8")

//创建表
orm.RegisterModel(new(Praise),new(ArticIMG),new(User),new(Article),new(ArticleType),new(TwoArticleType),new(ArticLabel),new(ArticIndexBanner),new(ArticBanner))

//生成表
//第一个参数是数据库别名，第二个参数是是否强制更新
orm.RunSyncdb("default",false,true)

}