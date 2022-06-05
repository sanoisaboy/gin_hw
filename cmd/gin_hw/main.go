package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sanoisaboy/gin_hw/pkg/http_status"
	re "github.com/sanoisaboy/gin_hw/pkg/repository"
)

func GetStudent(r *gin.Engine) {

	//connect to the database
	a := re.NewCrdbRepository(connect_string)

	//set http routing
	r.GET("/v1/users/id", func(ctx *gin.Context) {

		id := ctx.Query("id")
		id_int, err := strconv.Atoi(id)

		//get the data from the database
		student_get, err := a.GetStudent(ctx, id_int)
		if err != nil {
			log.Println(err)
		}

		//check the get data success or not
		http_status_code, err := http_status.Get_http_status(len(student_get))
		if err != nil {
			log.Println(err)
		}

		//select the reponse message type
		if http_status_code != 200 {
			ctx.JSON(http_status_code, gin.H{
				"message": "request error",
			})
		} else {
			ctx.JSON(http_status_code, gin.H{
				"student": student_get,
			})
		}
	})
}

func ListStudent(r *gin.Engine) {

	//connect to the Database
	a := re.NewCrdbRepository(connect_string)

	//set http routing
	r.GET("/v1/users", func(ctx *gin.Context) {
		//get the data from the database
		student_get, err := a.ListStudent(ctx)
		if err != nil {
			log.Println(err)
		}
		//reponse the data
		ctx.JSON(200, gin.H{
			"student": student_get,
		})

	})

}

func CreateStudent(r *gin.Engine) {

	//connect to the Database
	a := re.NewCrdbRepository(connect_string)

	//set http routing
	r.POST("/v1/users", func(ctx *gin.Context) {

		student_name := ctx.Query("student_name")
		id := ctx.Query("id")
		id_int, _ := strconv.Atoi(id)
		point := ctx.Query("point")
		point_int, _ := strconv.Atoi(point)

		//check data need query success or not
		http_status_code, err := http_status.Create_http_status(student_name, id, point)
		if err != nil {
			log.Println(err)
		}

		//insert data to the database success or not
		http_status_code, err = a.CreateStudent(ctx, student_name, id_int, point_int, http_status_code)
		if err != nil {
			log.Println(err)
		}

		//select the reponse message type
		if http_status_code != 201 {
			ctx.JSON(http_status_code, gin.H{
				"message": "create error",
			})
		} else {
			ctx.JSON(http_status_code, gin.H{
				"message":      "create success",
				"student_name": student_name,
				"id":           id,
				"point":        point,
			})
		}
	})

}

func UpdateStudent(r *gin.Engine) {

	//connect to the database
	a := re.NewCrdbRepository(connect_string)

	//set http routing
	r.PATCH("/v1/users", func(ctx *gin.Context) {

		id := ctx.Query("id")
		id_int, _ := strconv.Atoi(id)
		point := ctx.Query("point")
		point_int, _ := strconv.Atoi(point)

		//check data need query success or not
		http_status_code, err := http_status.Update_http_status(id, point)
		if err != nil {
			log.Println(err)
		}

		//Update data to the database success or not
		http_status_code, student_update, err := a.UpdateStudent(ctx, id_int, point_int, http_status_code)
		if err != nil {
			log.Println(err)
		}

		//select the reponse message type
		if http_status_code != 200 {
			ctx.JSON(http_status_code, gin.H{
				"message": "update error",
			})

		} else {
			ctx.JSON(http_status_code, gin.H{
				"message": "update success",
				"":        student_update,
			})
		}

	})
}

func DeleteStudent(r *gin.Engine) {

	//connect to the database
	a := re.NewCrdbRepository(connect_string)

	//set http routing
	r.DELETE("/v1/users", func(ctx *gin.Context) {
		id := ctx.Query("id")
		id_int, err := strconv.Atoi(id)

		//check data need query success or not
		http_status_code, err := http_status.Delete_http_status(id)
		if err != nil {
			log.Println(err)
		}

		//Delete data to the database success or not
		http_status_code, err = a.DeleteStudent(ctx, id_int)
		if err != nil {
			log.Println(err)
		}

		//select the reponse message type
		if http_status_code != 204 {
			ctx.JSON(http_status_code, gin.H{
				"message": "delete fail",
			})
		} else {
			ctx.JSON(http_status_code, gin.H{})
		}

	})
}

var connect_string string

//use flag to connect to database
func init() {
	flag.StringVar(&connect_string, "connect_string", "Default string", "Set String")
}

func main() {
	flag.Parse()
	r := gin.Default()
	ListStudent(r)
	GetStudent(r)
	CreateStudent(r)
	UpdateStudent(r)
	DeleteStudent(r)

	//run GIN server
	r.Run(":8002")
}
