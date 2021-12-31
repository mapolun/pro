package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func Export(c *gin.Context) {
	//type params struct {
	//	Id   int    `json:"id"   binding:"required"`
	//	Name string `json:"name" binding:"required"`
	//}
	//if err := c.ShouldBind(&params{}); err != nil {
	//	fmt.Println(err.Error())
	//	response.Error(c, "参数错误")
	//}
	f := excelize.NewFile()
	// 这里设置表头
	f.SetCellValue("Sheet1", "A1", "序号")
	f.SetCellValue("Sheet1", "B1", "名称")
	line := 1

	fruits := getFruits()
	// 循环写入数据
	for _, v := range fruits {
		line++
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", line), v.ID)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", line), v.Name)
	}

	// 保存文件
	if err := f.SaveAs("fruits.xlsx"); err != nil {
		fmt.Println(err)
	}
}

type fruit struct {
	ID    int
	Name  string
	Price float64
}

func getFruits() []fruit {

	var data []fruit

	for i := 1; i <= 10000; i++ {
		data = append(data, fruit{
			ID:    i,
			Name:  "测试",
			Price: 8,
		})
	}

	return data
}
