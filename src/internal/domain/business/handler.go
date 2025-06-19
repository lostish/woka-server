package business

import "github.com/gin-gonic/gin"

const business_by_name = "/:name"
const business_media = business_by_name + "/media"
const business_media_search = business_media + "/search"
const business_media_by_id = business_by_name + "/:mediaId"
const business_tasks = business_by_name + "/tasks"
const business_tasks_by_name = business_tasks + "/:task_name"
const business_reports = business_by_name + "/reports"
const business_reports_by_id = business_reports + "/:id"
const business_members = business_by_name + "/members"
const business_members_by_id = business_members + "/:id"
const business_products = business_by_name + "/products"
const business_products_by_name = business_products + "/:product_name"

// const business_media_download_by_id = business_media_by_id + "/download"

func RegisterRoutes(rg *gin.RouterGroup) {
	rg.GET("/", handleListBusiness)
	rg.GET(business_by_name, handleGetBusiness)
	rg.GET(business_media, handleListMedia)
	rg.GET(business_media_search, handleSearchMedia)
	rg.GET(business_media_by_id, handleGetMedia)
	rg.GET(business_tasks, handleListTask)
	rg.GET(business_tasks_by_name, handleGetTask)
	rg.GET(business_reports, handleListReports)
	rg.GET(business_reports_by_id, handleGetReport)
	rg.GET(business_members, handleListMembers)
	rg.GET(business_members_by_id, handleGetMember)
	rg.GET(business_products, handleListProduct)
	rg.GET(business_products_by_name, handleGetProduct)

	rg.POST("/", handleCreateBusiness)
	rg.POST(business_media, handleCreateMedia)
	rg.POST(business_tasks, handleCreateTask)
	rg.POST(business_reports, handleCreateReport)
	rg.POST(business_members, handleCreateMember)
	rg.POST(business_products, handleCreateProduct)

	rg.PUT(business_by_name, handleUpdateBusinnes)
	rg.PUT(business_media_by_id, handleUpdateMedia)
	rg.PUT(business_tasks_by_name, handleUpdateTask)
	rg.PUT(business_reports_by_id, handleUpdateReport)
	rg.PUT(business_members_by_id, handleUpdateMember)
	rg.PUT(business_products_by_name, handleUpdateProduct)

	rg.DELETE(business_by_name, handleDeleteBusiness)
	rg.DELETE(business_media_by_id, handleDeleteMedia)
	rg.DELETE(business_tasks_by_name, handleDeleteTask)
	rg.DELETE(business_reports_by_id, handleDeleteReport)
	rg.DELETE(business_members_by_id, handleDeleteMember)
	rg.DELETE(business_products_by_name, handleDeleteProduct)
}

func handleListBusiness(c *gin.Context) {}

func handleGetBusiness(c *gin.Context) {}

func handleListMedia(c *gin.Context) {}

func handleSearchMedia(c *gin.Context) {}

func handleGetMedia(c *gin.Context) {}

func handleListTask(c *gin.Context) {}

func handleGetTask(c *gin.Context) {}

func handleListReports(c *gin.Context) {}

func handleGetReport(c *gin.Context) {}

func handleListMembers(c *gin.Context) {}

func handleGetMember(c *gin.Context) {}

func handleListProduct(c *gin.Context) {}

func handleGetProduct(c *gin.Context) {}

func handleCreateBusiness(c *gin.Context) {}

func handleCreateMedia(c *gin.Context) {}

func handleCreateTask(c *gin.Context) {}

func handleCreateReport(c *gin.Context) {}

func handleCreateMember(c *gin.Context) {}

func handleCreateProduct(c *gin.Context) {}

func handleUpdateBusinnes(c *gin.Context) {}

func handleUpdateMedia(c *gin.Context) {}

func handleUpdateTask(c *gin.Context) {}

func handleUpdateReport(c *gin.Context) {}

func handleUpdateMember(c *gin.Context) {}

func handleUpdateProduct(c *gin.Context) {}

func handleDeleteBusiness(c *gin.Context) {}

func handleDeleteMedia(c *gin.Context) {}

func handleDeleteTask(c *gin.Context) {}

func handleDeleteReport(c *gin.Context) {}

func handleDeleteMember(c *gin.Context) {}

func handleDeleteProduct(c *gin.Context) {}
