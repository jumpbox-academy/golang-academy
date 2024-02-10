
# Mini API Instruction with TODO 

Start with [gin framework](https://gin-gonic.com/docs/quickstart/)

1. Download and install it
```bash
go get -u github.com/gin-gonic/gin
```

2. Import it in your code:
```go
import "github.com/gin-gonic/gin"
```

3. (Optional) Import net/http. This is required for example if using constants such as http.StatusOK.
```go
import "net/http"
```

4. Copy this to your `main.go` project
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
```

## Create TODO feature
1. Create `todo` Directory
```
mkdir todo && cd todo
```

2. Create a file name `todo`
```
touch todo.go
```

3. Fill in the information 
```go
package todo
// สร้าง structure ที่ represent table (entity) รับส่ง data กับ frontend ได้
type Todo struct {
    // ฝั่ง json ส่งเข้ามา key ชื่อว่า text แต่เรา mapping เป็น Title
    Title string `json:"text"`
    // Embedded gorm.Model
    gorm.Model
}

// ป้องกัน gorm ตั้งชื่อไม่ได้เป็นไปตามที่เราต้องการ ด้วยการกำหนดด้วยตัวเอง
func (Todo) TableName() string {
    return "todos"
}

// สร้าง type TodoHandler ที่มี db เป็น gorm.DB อยู่ด้านใน
type TodoHandler struct {
    db *gorm.DB
}

// สร้าง Structure เพื่อให้ฝั่งคนใช้งานนำไปใช้ รับ db ที่เป็น gorm.DB
func NewTodoHandler (db *gorm.DB) *TodoHandler {
    return &TodoHandler{db: db}
}

// สร้าง Handler
func (t *TodoHandler) NewTask(c *gin.Context) {
    var todo Todo

    // function ที่ bind JSON มีให้เลือกใช้งาน 2 แบบ
    // 1. BindJSON(...) หมายถึง Must Bind เมื่อเกิดปัญหาจะทำการคืน Status 400 ไปอัตโนมัติ
    // 2. ShouldBindJSON หมายถึง Should Bind เมื่อเกิดปัญหาเราต้อง Handle error นั้นเอง
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // สร้าง Data เพื่อไปลง Database
    // function Create จะคืน result(r) ออกมา
    r := t.db.Create(&todo)
    if err := r.Error; err !=nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }
    
    // เราจะทำการคืนของออกไปคือ ID ที่อยู่ใน gormModel
    c.JSON(http.StatusCreated, gin.H{
        "ID": todo.Model.ID,
    })
}
```

4. Integrate in `main.go` 
```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

```go
    // วางไว้บน gin framework
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

    // gin Framework
    ...


    //วางไว้ใต้ gin Framework
    handler := todo.NewTodoHandler(db)
    r.POST("/todos", handler.NewTask)


    // Run ไว้ด้านล่างสุดเสมอ
    r.Run() // listen and serve on 0.0.0.0:8080
```
5. Test API via VSCode Extenstion `REST Client` and prepare test api

```
mkdir test && cd test && touch create_new_todo.http
```

```go
@base_url=http://localhost:8080

POST {{base_url}}/todos
Content-Type: application/json

{
    "text": "test 1"
}
```

6. Add Auto migrate
```go
// ใส่ไว้ข้างบน gin framework
db.AutoMigrate(&todo.Todo{}) 
```

7. Install `Database Client` Extension for connect sqlite
