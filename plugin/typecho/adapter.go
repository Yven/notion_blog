package typecho

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"notion_blog/notion"
)

type Typecho struct {
	Db      *gorm.DB
	Content *Contents
	Meta    []*Metas
}

type DbOptions struct {
	Host     string
	Port     string
	User     string
	Passwd   string
	Database string
	Charset  string
}

func NewDb(opt DbOptions) *Typecho {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", opt.User, opt.Passwd, opt.Host, opt.Port, opt.Database, opt.Charset)
	// dsn := "root:123456@tcp(127.0.0.1:3306)/typecho?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "typecho_", SingularTable: false},
	})
	if err != nil {
		log.Fatalf("数据库连接失败：%s", err)
	}

	return &Typecho{Db: db}
}

func (c *Typecho) Writer(list *notion.List) error {
	// 将数据转为表结构
	c.Content = c.TransformContent(list)
	c.Meta = c.TransformMeta(list)
	// 查询文章是否存在
	existContent := &Contents{}
	contentRes := c.Db.Where(Contents{Slug: c.Content.Slug}).Find(existContent)
	if contentRes.Error != nil {
		return contentRes.Error
	}

	// 开启事务
	tx := c.Db.Begin()

	if contentRes.RowsAffected == 0 {
		// 新增文章
		if res := tx.Create(c.Content); res.Error != nil {
			tx.Rollback()
			return res.Error
		}
	} else {
		// 修改文章
		c.Content.Cid = existContent.Cid
		if res := tx.Model(existContent).Updates(c.Content); res.Error != nil {
			tx.Rollback()
			return res.Error
		}
	}

	// 遍历关联标签表
	for _, meta := range c.Meta {
		// 查询标签是否存在
		existMeta := &Metas{}
		if metaRes := tx.Where(Metas{Type: meta.Type, Slug: meta.Slug}).First(existMeta); metaRes.Error == nil || errors.Is(metaRes.Error, gorm.ErrRecordNotFound) {
			if metaRes.RowsAffected == 0 {
				// 新增标签
				if res := tx.Create(meta); res.Error != nil {
					tx.Rollback()
					return res.Error
				}
			} else {
				// 记录标签
				meta = existMeta
			}
		} else {
			tx.Rollback()
			return metaRes.Error
		}

		// 查询是否存在绑定关系
		if relatRes := tx.Where(Relationships{Cid: c.Content.Cid, Mid: meta.Mid}).First(&Relationships{}); relatRes.Error == nil || errors.Is(relatRes.Error, gorm.ErrRecordNotFound) {
			if relatRes.RowsAffected == 0 {
				// 新增关联关系
				if res := tx.Create(Relationships{Cid: c.Content.Cid, Mid: meta.Mid}); res.Error != nil {
					tx.Rollback()
					return res.Error
				}

				// 更新标签
				meta.Count = existMeta.Count + 1
				if res := tx.Model(meta).Updates(*meta); res.Error != nil {
					tx.Rollback()
					return res.Error
				}
			}
		} else {
			fmt.Println(relatRes.Error)
			tx.Rollback()
			return relatRes.Error
		}
	}

	// 文章存在时
	if contentRes.RowsAffected != 0 {
		// 查询是否存在关联关系
		nowRelat := &[]Relationships{}
		if relatRes := tx.Where(Relationships{Cid: c.Content.Cid}).Find(nowRelat); relatRes.Error == nil || errors.Is(relatRes.Error, gorm.ErrRecordNotFound) {
			// 如果存在关联关系，寻找需要删除的关联关系
			if relatRes.RowsAffected != 0 {
				var delRow []Relationships
				for _, relationships := range *nowRelat {
					isExist := false
					for _, meta := range c.Meta {
						if meta.Mid == relationships.Mid {
							isExist = true
							break
						}
					}
					if isExist {
						delRow = append(delRow, relationships)
					}
				}
				// 删除关联关系
				for _, relationships := range delRow {
					if res := tx.Delete(&Relationships{Cid: c.Content.Cid, Mid: relationships.Mid}); res.Error != nil {
						tx.Rollback()
						return res.Error
					}
				}
			}
		} else {
			tx.Rollback()
			return relatRes.Error
		}
	}

	tx.Commit()
	return nil
}

func (c *Typecho) Delete() {}
func (c *Typecho) Close()  {}
