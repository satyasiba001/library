package handler

import(
	"github.com/gin-gonic/gin"
	// "github.com/satyasiba/library/db"
)

func routes(router *gin.Engine){

	router := gin.Default()

	//member functions
	router.POST("/members", getBodyDate())
	// router.GET("/members/:id", GetUserByID(MemberServices))
	// router.DELETE("/members/:id", DeleteMemberByID(MemberServices))

	// // bookRoutes
	// router.POST("/books", CreateNewBook(BookServices))
	// router.DELETE("/books/:id", DeleteBookByID(BookServices))
	// router.GET("/books", Filter(BookServices))

	// // transaction routes
	// router.POST("/borrow", BorrowBook(TransactionServices))
	// router.PATCH("/return", ReturnBook(TransactionServices))
}

func StartApp() {

	router := gin.Default()
	Handlers(router)
	router.Run()
}

