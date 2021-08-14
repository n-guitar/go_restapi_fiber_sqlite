module github.com/n-guitar/go_restapi_fiber_sqlite

go 1.16

require (
	github.com/gofiber/fiber/v2 v2.17.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.20.7
)

replace (
	github.com/n-guitar/go_restapi_fiber_sqlite/pkg/book => ./pkg/book
	github.com/n-guitar/go_restapi_fiber_sqlite/pkg/database => ./pkg/database
)
