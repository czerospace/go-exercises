package main

import (
	"github.com/gin-gonic/gin"
)

// 创建任务的 API 处理函数
func createTaskHandler(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	taskRegistry := TaskRegistry{}
	err := taskRegistry.CreateTask(task)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(201, gin.H{"message": "Task created successfully"})
}

// 获取任务列表的 API 处理函数
func listTaskHandler(c *gin.Context) {
	taskRegistry := TaskRegistry{}
	tasks, err := taskRegistry.ListTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list tasks"})
	}

	c.JSON(200, tasks)
}

func listServicesHandler(c *gin.Context) {
	c.JSON(200, nil)
}

func createServicesHandler(c *gin.Context) {
	c.JSON(200, nil)
}

func main() {
	// 初始化 Gin Web 框架
	r := gin.Default()

	// 注册 API 路由
	r.POST("/tasks", createTaskHandler)
	r.GET("/tasks", listTaskHandler)
	r.GET("/services", listServicesHandler)
	r.POST("/services", createServicesHandler)
	r.Any("resource/:type", restHandler)

	// 启动 Web 服务器
	r.Run(":8080")
}

// Task 结构体
type Task struct {
	Name        string
	Description string
}

// 定义一个 Registry 接口
type Registry interface {
	ListTasks() ([]Task, error)
	CreateTask(task Task) error
}

// 新建一个 TaskRegistry结构体 实现 Registry 结构
type TaskRegistry struct {
}

func (t *TaskRegistry) ListTasks() ([]Task, error) {
	return []Task{{
		Name:        "Task",
		Description: "TaskRegistry",
	}}, nil
}

func (t *TaskRegistry) CreateTask(task Task) error {
	return nil
}

// 新建一个 MysqlRegistry 结构体 实现 Registry 结构
type MysqlTaskRegistry struct {
}

func (m *MysqlTaskRegistry) ListTasks() ([]Task, error) {
	return []Task{{
		Name:        "mysql",
		Description: "MysqlTaskRegistry",
	}}, nil
}

func (m *MysqlTaskRegistry) CreateTask(task Task) error {
	return nil
}

// 模仿 k8s apiserver 源码中的 restStorage
// 自定义 handerStorage 接口
type handerStorage interface {
	List(c *gin.Context)
	Create(c *gin.Context)
}

// 定义 TaskStorage 结构体去实现 handerStorage 中的方法
type TaskStorage struct {
}

func (t *TaskStorage) List(c *gin.Context) {
	taskRegistry := TaskRegistry{}
	tasks, err := taskRegistry.ListTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list tasks"})
	}

	c.JSON(200, tasks)
}

func (t *TaskStorage) Create(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	taskRegistry := TaskRegistry{}
	err := taskRegistry.CreateTask(task)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(201, gin.H{"message": "Task created successfully"})
}

// 定义 ServiceStorage 结构体去实现 handerStorage 中的方法
type ServiceStorage struct {
}

func (t *ServiceStorage) List(c *gin.Context) {
	taskRegistry := MysqlTaskRegistry{}
	tasks, err := taskRegistry.ListTasks()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to list tasks"})
	}

	c.JSON(200, tasks)
}

func (t *ServiceStorage) Create(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": "Bad Request"})
		return
	}
	taskRegistry := MysqlTaskRegistry{}
	err := taskRegistry.CreateTask(task)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(201, gin.H{"message": "Task created successfully"})
}
func restHandler(c *gin.Context) {
	// eg: http://127.0.0.1:8080/resource/task
	// 只需要把 task 放入 map 中，由 handerStorage 处理
	// m := map[string]handerStorage{"task": &TaskStorage{}}

	// 现有又要增加一个 reousrce 类型为  service
	// 只需 service 实现 handerStorage 接口中的方法，然后加入 map 即可
	m := map[string]handerStorage{
		"task":    &TaskStorage{},
		"service": &ServiceStorage{},
	}
	resourceType := c.Param("type")
	switch c.Request.Method {
	case "GET":
		m[resourceType].List(c)
	case "POST":
		m[resourceType].Create(c)
	}
}
