package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"jtapi/model"
	"log"
	"net/http"
	"strconv"
)

// API => 新增
// POST
func Add(c *gin.Context) {
	var tt model.Test
	err := c.BindJSON(&tt)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg":err.Error()})
	}
	id,err := tt.Add()
	fmt.Println("add =========> " + strconv.Itoa(int(id)))
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg":"", "data":id})
}

// API => 列表查询
// GET
func Page(c *gin.Context) {
	var tt model.Test
	tests, err := tt.Page()
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg":err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg":"", "data":tests})
}

// API => 修改
// PUT
func Update(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)
	var tt model.Test
	err := c.BindJSON(&tt)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg":err.Error()})
	}
	// 赋值
	tt.Id = id
	tt.Update()
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg":"", "data":""})
}

// API => 查看详情
// GET
func Get(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)
	tt := model.Test{Id: id}
	t, err := tt.Get()
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg":err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg":"", "data":t})
}

// API => 修改
// DELETE
func Delete(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)
	tt := model.Test{Id: id}
	nums, err := tt.Delete()
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{"code": 1, "msg":err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg":"", "data":nums})
}