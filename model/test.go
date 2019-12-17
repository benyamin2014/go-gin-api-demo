package model

import (
	db "jtapi/repo"
	"log"
)

//定义实体
type Test struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

// 新增方法
func (t *Test) Add()(id int64, err error) {
	rs, err := db.SqlDB.Exec("insert into test (id,name) values(?,?)", t.Id, t.Name)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id)
	return
}

// 列表查询
func (t *Test) Page()(tests []Test, err error) {
	tests = make([]Test, 0)
	rows, err := db.SqlDB.Query("select id,name from test")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var test Test
		rows.Scan(&test.Id, &test.Name)
		tests = append(tests, test)
	}
	if err = rows.Err(); err!=nil {
		return
	}
	return
}

// 详情查询
func (t *Test) Get()(test Test, err error) {
	err = db.SqlDB.QueryRow("select id, name from test where id = ?", t.Id).Scan(
		&test.Id, &test.Name,
		)
	return
}

// 修改
func (t *Test) Update() (ra int64, err error) {
	stmt, err := db.SqlDB.Prepare("update test set name=? where id=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(t.Name, t.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

// 删除
func (t *Test) Delete()(result int64, err error) {
	rs, err := db.SqlDB.Exec("delete from test where id=?", t.Id)
	if err != nil {
		log.Fatal(err)
	}
	result, err = rs.RowsAffected()
	return
}
