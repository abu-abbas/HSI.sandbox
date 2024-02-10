# HSI Sandbox [Go] - Tugas Level 4

1. Buatkan REST API untuk CRUD tabel pada tugas level 3

    Dependensi yang saya gunakan
    ```
    [database]
    github.com/go-chi/chi v1.5.5
    github.com/jmoiron/sqlx v1.3.5
	github.com/mattn/go-sqlite3 v1.14.22

	[router]
	github.com/go-chi/chi v1.5.5

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
        │  └── tugas_4.db
        ├── entity
        │  ├── detail.go
        │  └── item.go
        ├── go.mod
        ├── go.sum
        ├── handler
        │  └── response.go
        ├── main.go
        ├── model
        │  ├── detail.go
        │  └── item.go
        ├── readme.md
        └── utils
           └── error.go
    ```

    Untuk inisialisai awal (entry point) yaitu `main.go`
    ```go
    package main

    import "github.com/abu-abbas/level_4/server/api"

    func main() {
    	api.CreateApp().Mount()
    }
    ```

    `func CreateApp()` yaitu untuk inisialisai config server dan router pada library `api.go` adalah sebagai berikut:
    ```go
    func CreateApp() *AppServer {
    	cfg := config.GetYamlValue().ServerConfig
    	svr := &AppServer{
    		config: cfg,
    		router: chi.NewRouter(),
    	}

    	svr.routes()

    	return svr
    }
    ```

    Kemudian setelah persiapan selesai, proses selanjutnya adalah mounting http server agar dapat membalas setiap request yang datang beserta event handle pada saat unmounting http server
    ```go
     func (app *AppServer) Mount() {
    	ctx := context.Background()

    	server := http.Server{
    		Addr:    fmt.Sprintf(":%s", app.config.Port),
    		Handler: app.router,
    	}

    	shutdownComplete := handleShutdown(func() {
    		if err := server.Shutdown(ctx); err != nil {
    			log.Printf("server shutdown failed: %v\n", err)
    		}
    	})

    	log.Printf("http listen and serve on port: %s", app.config.Port)
    	if err := server.ListenAndServe(); err == http.ErrServerClosed {
    		<-shutdownComplete
    	} else {
    		log.Printf("http listen and serve failed: %v\n", err)
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
    func (app *AppServer) routes() {
    	app.router.Use(middleware.RequestID)
    	app.router.Use(middleware.Logger)
    	app.router.Use(middleware.Recoverer)
    	app.router.Use(middleware.URLFormat)
    	app.router.Use(render.SetContentType(render.ContentTypeJSON))
    	app.router.Route(
    		"/api",
    		func(r chi.Router) {
    			r.Mount("/v1", v1Handler(r))
    		},
    	)

    	// base path
    	app.router.Get(
    		"/",
    		func(w http.ResponseWriter, r *http.Request) {
    			w.Write([]byte("root."))
    		},
    	)
    }

    func v1Handler(r chi.Router) http.Handler {
    	r.Mount("/items", itemServices(r))

    	return r
    }

    func itemServices(r chi.Router) http.Handler {
    	controllerItem := controllers.Item{}

    	r.Get("/", controllerItem.Index)
    	r.Post("/", controllerItem.Create)

    	return r
    }
    ```

    Setelah semua sudah siap, berikut ini adalah detail untuk proses CRUD dengan protokol REST API
    * Proses Read Item (Fetching All Item)
    ![Get All Items](https://raw.githubusercontent.com/abu-abbas/HSI.sandbox/main/go/tugas/level_4/snapshot/getItems.png)

    Berikut ini `controller` untuk items

    ```go
    func (i *Item) Index(w http.ResponseWriter, r *http.Request) {
    	res := handler.JsonBody{}
    	items, err := i.model.Get()

    	if err != nil {
    		log.Errorf("gagal mendapatkan items, error: %v\n", err)
    		res.HttpStatus = "error"
    		res.HttpCode = http.StatusInternalServerError
    		res.Message = "Terjadi kesalahan pada server. Silakan hubungi Admin"
    	} else {
    		res.HttpStatus = "success"
    		res.HttpCode = http.StatusOK
    		res.Payload = items
    	}

    	handler.JsonResponse(w, res)
    }

    ```

    Handler untuk setiap `response` baik sukses maupun gagal dalam bentuk `json`
    ```go
    type JsonBody struct {
    	HttpStatus string      `json:"status"`
    	HttpCode   int         `json:"code,omitempty"`
    	Payload    interface{} `json:"data,omitempty"`
    	Message    string      `json:"message,omitempty"`
    }

    func JsonResponse(w http.ResponseWriter, body JsonBody) {
    	response, _ := json.Marshal(body)

    	w.Header().Set("Content-Type", "application/json")
    	w.WriteHeader(body.HttpCode)
    	w.Write(response)
    }
    ```

    Proses fetching data pada `model` seperti dilaporkan pada tugas level 3
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
    			return items, err
    		}
    	}

    	return items, nil
    }
    ```

    * Proses Create Item

2. Buatkan client untuk menembak service lain dengan protocol REST API

    `N/A`

