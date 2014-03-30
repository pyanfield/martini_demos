// inspired by http://blog.gopheracademy.com/day-11-martini

package main

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/cznic/kv"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"io"
	"os"
)

var dbFile string = "kv_demo.kv"

type KVData struct {
	Key   string `form:"key"`
	Value string `form:"value"`
}

// 返回 KVData 类型得数组
func AllDatas(db *kv.DB) []KVData {
	var datas []KVData
	alldatas := getAll(db)
	for key, value := range alldatas {
		data := KVData{}
		data.Key = key
		data.Value = value
		datas = append(datas, data)
	}
	return datas
}

func getBuf(n int) []byte {
	return make([]byte, n)
}

// 打开 kv 数据库
func openDB(dbFile string, opts *kv.Options) (*kv.DB, error) {
	createOpen := kv.Open
	status := "opening"

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		createOpen = kv.Create
		status = "creating"
	}

	if opts == nil {
		opts = &kv.Options{}
	}

	db, err := createOpen(dbFile, opts)

	if err != nil {
		return nil, fmt.Errorf("error %s %s: %v \n", status, dbFile, err)
	}
	return db, nil
}

// 返回数据库中得所有key/value值
func getAll(db *kv.DB) map[string]string {
	m := map[string]string{}
	enum, err := db.SeekFirst()
	if err == io.EOF {
		fmt.Println("Empty Database")
		return m
	}
	for {
		k, v, err := enum.Next()
		if err == io.EOF {
			fmt.Printf("End of Database \n")
			return m
		}
		m[string(k)] = string(v)
		// fmt.Printf("%s = %s \n", string(k), string(v))
	}
	return m
}

// 自定义一个中间件，返回 martini.Handler，这样就可以再每个request中都会调用该中间件的方法
func DB() martini.Handler {
	return func(c martini.Context) {
		db, err := openDB(dbFile, nil)
		if err != nil {
			panic(err)
		}
		// map 一个kv.DB得实例到request context中，这样就可以再 HTTP 方法中定义一个 kv.DB 参数，从而将它注入到我们得代码中
		c.Map(db)
		defer db.Close()
		c.Next()
	}
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Use(DB())

	m.Get("/kv", func(r render.Render, db *kv.DB) {
		r.HTML(200, "list", AllDatas(db))
	})
	// binding.Form 将表单数据解析到 KVData结构体中，然后map到 request context中
	m.Post("/kv", binding.Form(KVData{}), func(data KVData, r render.Render, db *kv.DB) {
		err := db.Set([]byte(data.Key), []byte(data.Value))
		if err != nil {
			panic(err)
		}
		// fmt.Printf(">>>>>>> %s = %s \n", data.Key, data.Value)
		r.HTML(200, "list", AllDatas(db))
	})

	m.Run()
}
