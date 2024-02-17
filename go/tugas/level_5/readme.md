# HSI Sandbox [Go] - Tugas Level 5

1. Buatkan CRUD menggunakan framework Fiber

    Dependensi yang saya gunakan
    ```
    [database]
    github.com/jmoiron/sqlx v1.3.5
	github.com/mattn/go-sqlite3 v1.14.22

	[http-framework]
	github.com/gofiber/fiber/v2 v2.52.0

	[utils]
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/viper v1.18.2
    ```

    Struktur folder
    ```
        .
        ├── api
        │  ├── routes.go
        │  └── server.go
        ├── config
        │  ├── config.go
        │  └── config.toml
        ├── controllers
        │  └── item.go
        ├── db
        │  ├── connection.go
        │  └── tugas_5.db
        ├── entity
        │  ├── detail.go
        │  └── item.go
        ├── go.mod
        ├── go.sum
        ├── main.go
        ├── model
        │  ├── detail.go
        │  └── item.go
        ├── routes
        └── utils
           └── error.go

    ```

    Untuk inisialisai awal (entry point) yaitu `main.go`
    ```go
    package main

    import "github.com/abu-abbas/level_5/api"

    func main() {
    	api.CreateApp().Mount()
    }
    ```

    `func CreateApp()` yaitu untuk inisialisai config server dan router pada library `api.go` adalah sebagai berikut:
    ```go
    type AppServer struct {
    	config config.ServerConfig
    	app    *fiber.App
    }

    func CreateApp() *AppServer {
    	svr := &AppServer{
    		config: config.GetYamlValue().ServerConfig,
    		app:    fiber.New(),
    	}

    	svr.routes()
    	return svr
    }
    ```

    Kemudian setelah persiapan selesai, proses selanjutnya adalah mounting http server agar dapat membalas setiap request yang datang beserta event handle pada saat unmounting http server. Mirip dengan level 4 seebelumnya.
    ```go
    func (server *AppServer) Mount() {
    	ctx := context.Background()
    	shutdownComplete := handleShutdown(func() {
    		if err := server.app.ShutdownWithContext(ctx); err != nil {
    			log.Printf("server shutdown failed: %v\n", err)
    		}
    	})

    	log.Printf("http listen and serve on port: %s", server.config.Port)
    	if err := server.app.Listen(fmt.Sprintf(":%s", server.config.Port)); err == http.ErrServerClosed {
    		<-shutdownComplete
    	}

    	log.Println("shutdown gracefully")
    }

    func handleShutdown(onShutdownSignal func()) <-chan struct{} {
    	shutdown := make(chan struct{})
    	go func() {
    		shutdownSignal := make(chan os.Signal, 1)
    		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

    		<-shutdownSignal

    		onShutdownSignal()
    		close(shutdown)
    	}()

    	return shutdown
    }
    ```

    Router yang berfungsi sebagai pengatur request datang dan mengembalikan response
    ```go
    func (server *AppServer) routes() {
    	server.app.Use(requestid.New())
    	server.app.Use(logger.New(
    		logger.Config{
    			Format: "[${ip}]:${port} ${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
    		},
    	))

    	server.app.Get("/", func(c *fiber.Ctx) error {
    		return c.SendString("Hello, World!")
    	})

    	middleware := basicauth.New(basicauth.Config{
    		Users: map[string]string{
    			"admin": "x",
    		},
    	})

    	api := server.app.Group("/api", middleware)
    	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
    		c.Set("Version", "v1")
    		return c.Next()
    	})

    	item := controllers.Item{}

    	v1.Get("/", item.Index)
    	v1.Post("/", item.Create)
    }
    ```

    Setelah semua sudah siap, berikut ini adalah detail untuk proses CRUD dengan protokol REST API
    * Akses root path tanpa auth
    ![Root access tanpa auth](https://github.com/abu-abbas/HSI.sandbox/blob/main/go/tugas/level_5/snapshot/unprotected_route.png?raw=true)


    * Akses `/api/v1` yang dilidungi dengan middleware basic auth (contoh tanpa auth)
    ![Tanpa Auth](https://github.com/abu-abbas/HSI.sandbox/blob/main/go/tugas/level_5/snapshot/protected_route_1.png?raw=true)

        karena route `/api/v1` sudah dilindungi oleh middlewar basic auth maka dari itu akan terjadi error ketika diakses tanpa otentikasi


    * Akses `/api/v1` yang dilidungi dengan middleware basic auth
    ![Get All Items](https://github.com/abu-abbas/HSI.sandbox/blob/main/go/tugas/level_5/snapshot/protected_route_2.png?raw=true)

        `controller` untuk handle request diatas adalah sbb:
        ```go
        func (i *Item) Index(c *fiber.Ctx) error {
        	items, err := i.model.Get()
        	if err != nil {
        		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
        			"status":  "error",
        			"message": "Terjadi kesalahan pada server",
        		})
        	}

        	return c.Status(fiber.StatusOK).JSON(fiber.Map{
        		"status": "success",
        		"data":   items,
        	})
        }
        ```

        `model` untuk `get all item` masih sama dengan level sebelumnya
        ```go
        func (i Item) Get() ([]entity.Item, error) {
        	var items []entity.Item

        	qry := "SELECT * FROM items"
        	con := db.Connect()
        	err := con.Select(&items, qry)

        	if err != nil {
        		if err == sql.ErrNoRows {
        			return items, errors.New("item tidak ditemukan")
        		} else {
        			utils.ErrorCheck(err)
        			return items, err
        		}
        	}

        	return items, nil
        }
        ```

    * Proses insert data yang dimulai dari `controllers` adalah sbb:
        ```go
        func (i *Item) findById(id int64, c *fiber.Ctx) error {
        	fetch, err := i.model.FindById(id)
        	if err != nil {
        		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
        			"status":  "error",
        			"message": "Terjadi kesalahan pada server",
        		})
        	}

        	return c.Status(fiber.StatusOK).JSON(fiber.Map{
        		"status": "success",
        		"data":   fetch,
        	})
        }

        func (i *Item) Create(c *fiber.Ctx) error {
        	entity := entity.Item{}
        	err := c.BodyParser(&entity)
        	if err != nil {
        		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
        			"status":  "error",
        			"message": "Terjadi kesalahan pada server",
        		})
        	}

        	itemId, err := i.model.Create(entity)
        	if err != nil {
        		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
        			"status":  "error",
        			"message": "Terjadi kesalahan pada server",
        		})
        	}

        	return i.findById(itemId, c)
        }

        ```

        Untuk `model` tidak ada perubahan, masih sama dengan level sebelumnya:
        ```go
        func (i Item) FindById(id int64) (entity.Item, error) {
        	var item entity.Item

        	qry := "SELECT * FROM items WHERE id=?"
        	con := db.Connect()
        	err := con.Get(&item, qry, id)

        	return i.resultCheck(item, err)
        }

        func (i Item) Create(item entity.Item) (int64, error) {
        	qry := "INSERT INTO items (name, status, amount) VALUES (:name, :status, :amount)"
        	con := db.Connect()
        	res, err := con.NamedExec(qry, &item)

        	if err != nil {
        		return -1, err
        	}

        	return res.LastInsertId()
        }

        func (i Item) CreateMany(items []entity.Item) int64 {
        	qry := "INSERT INTO items (name, status, amount) VALUES (:name, :status, :amount)"
        	trx := db.Begin()
        	res, err := trx.NamedExec(qry, items)
        	if err != nil {
        		trx.Rollback()
        	}

        	rowAffected, errRowAffected := res.RowsAffected()
        	if errRowAffected != nil {
        		trx.Rollback()
        	}

        	trx.Commit()
        	return rowAffected
        }
        ```

2. Tambahkan middleware Authorization untuk CRUD tersebut

    Middleware untuk auth pada `fiber` sangat fleksibel bisa diletakan pada saat awal inisialisasi `fiber` ataupun pada `route` tertentu. Berikut ini konfigurasi yang saya gunakan:
    ```go
    package api

    import "github.com/gofiber/fiber/v2/middleware/basicauth"

    middleware := basicauth.New(basicauth.Config{
		Users: map[string]string{
			"admin": "x",
		},
	})

	api := server.app.Group("/api", middleware)
    ```
