package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/nicholasmarco27/UTS_5027221042_Nicholas-Marco-Weinandra/master/common/genproto/taskmaster"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	//"google.golang.org/protobuf/types/known/emptypb"
	//"google.golang.org/protobuf/types/known/wrapperspb"
)

type Empty struct{}


type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func init() {
	args := os.Args[1:]
	var configname string = "default-config"
	if len(args) > 0 {
		configname = args[0] + "-config"
	}
	log.Printf("loading config file %s.yml \n", configname)

	viper.SetConfigName(configname)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Fatal error config file: " + err.Error())
	}
}

func (s *httpServer) Run() error {
	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/create", s.handleCreate)
	http.HandleFunc("/update", s.handleUpdate)
	http.HandleFunc("/delete", s.handleDelete)
	http.HandleFunc("/list", s.handleList)
	log.Printf("Starting HTTP server on localhost%s/list\n", s.addr)
	return http.ListenAndServe(s.addr, nil)
}

func (s *httpServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index").Parse(tasksTemplate))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *httpServer) handleCreate(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan data dari form HTML
	title := r.FormValue("title")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")
	quantityStr := r.FormValue("quantity")

	// Convert price string to integer
	price, err := strconv.Atoi(priceStr)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}

	// Convert quantity string to integer
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	// Menginisialisasi koneksi gRPC
	port := ":" + viper.GetString("app.grpc.port")
	client, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Could not connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Membuat klien API tugas
	taskClient := taskmaster.NewTaskApiClient(client)

	// Membuat task baru
	_, err = taskClient.CreateTask(context.Background(), &taskmaster.Task{
		Title:       title,
		Description: description,
		Price:       int32(price),
		Quantity:    int32(quantity),
	})
	if err != nil {
		http.Error(w, "Failed to add transaction", http.StatusInternalServerError)
		return
	}

	// Tetapkan pengalihan arahan ke halaman list
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}

func (s *httpServer) handleUpdate(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan data dari form HTML
	id := r.FormValue("id")
	title := r.FormValue("title")
	description := r.FormValue("description")
	priceStr := r.FormValue("price")
	quantityStr := r.FormValue("quantity")

	// Konversi string harga dan kuantitas ke int32 dan int
	price, err := strconv.ParseInt(priceStr, 10,32)
	if err != nil {
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return
	}
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		http.Error(w, "Invalid quantity", http.StatusBadRequest)
		return
	}

	// Menginisialisasi koneksi gRPC
	port := ":" + viper.GetString("app.grpc.port")
	client, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Could not connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Membuat klien API tugas
	taskClient := taskmaster.NewTaskApiClient(client)

	// Update task
	_, err = taskClient.UpdateTask(context.Background(), &taskmaster.Task{
		Id:          id,
		Title:       title,
		Description: description,
		Price:       int32(price),
		Quantity:    int32(quantity),
	})
	if err != nil {
		http.Error(w, "Failed to update transaction", http.StatusInternalServerError)
		return
	}

	// Tetapkan pengalihan arahan ke halaman list
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}

func (s *httpServer) handleDelete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		http.Error(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	// Menginisialisasi koneksi gRPC
	port := ":" + viper.GetString("app.grpc.port")
	client, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Could not connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Membuat klien API tugas
	taskClient := taskmaster.NewTaskApiClient(client)

	_, err = taskClient.DeleteTask(context.Background(), &taskmaster.TaskId{Id: id})
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/list", http.StatusSeeOther)
}

func (s *httpServer) handleList(w http.ResponseWriter, r *http.Request) {
	// Menginisialisasi koneksi gRPC
	port := ":" + viper.GetString("app.grpc.port")
	client, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		http.Error(w, "Could not connect to gRPC server", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	// Membuat klien API tugas
	taskClient := taskmaster.NewTaskApiClient(client)

	resp, err := taskClient.ListTasks(context.Background(), &taskmaster.Empty{})
	if err != nil {
		http.Error(w, "Failed to retrieve task list", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.New("list").Parse(tasksTemplate))
	if err := tmpl.Execute(w, resp.Tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	httpServer := NewHttpServer(":1000")
	httpServer.Run()
}

type Task struct {
	Id          string
	Title       string
	Description string
	Price       int32
	Quantity    int32
}

var tasksTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Habit Tracker</title>
    <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.3.0/font/bootstrap-icons.css">
    <style>
        body {
            background-color: #EEF7FF;
        }
    </style>
</head>
<body>
    <nav class="navbar navbar-expand-lg navbar-dark bg-info">
        <a class="navbar-brand" href="#">Track Your Transaction</a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarNav">
            <ul class="navbar-nav">
                <li class="nav-item">
                    <a class="nav-link" href="#">Home</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="#">Transactions</a>
                </li>
                <li class="nav-item">
                    <a class="nav-link" href="#">About</a>
                </li>
            </ul>
        </div>
    </nav>

    <div class="container-fluid">
        <h1>Transaction Tracker</h1>
        <form action="/create" method="post">
            <label for="title">Transaction:</label><br>
            <input type="text" id="title" name="title"><br>
            <label for="description">Description:</label><br>
            <textarea id="description" name="description"></textarea><br>
            <label for="price">Price:</label><br> <!-- Added Price Field -->
            <input type="text" id="price" name="price"><br>
            <label for="quantity">Quantity:</label><br> <!-- Added Quantity Field -->
            <input type="text" id="quantity" name="quantity"><br><br>
            <input type="submit" value="Add Transaction">
        </form>
        <hr>
        <h2>Transaction History</h2>
        <a href="/list" class="refresh-btn">Refresh Transaction History</a>
        <ul>
            {{if not (eq (len .Tasks) 0)}}
                {{range .Tasks}}
                <li class="list-group-item">
                    <div class="card-body">
                        <span>{{.Title}} - {{.Description}}</span>
                        <a href="#" onclick="showUpdateForm('{{.Id}}')">Update</a> 
                        <a href="/delete?id={{.Id}}">Delete</a>
                    </div>
                    <form id="updateForm{{.Id}}" class="update-form mt-3" action="/update" method="post" style="display: none;">
                        <input type="hidden" name="id" value="{{.Id}}">
                        <div class="form-group">
                            <label for="title{{.Id}}">New Title:</label><br>
                            <input type="text" id="title{{.Id}}" name="title" value="{{.Title}}"><br>
                        </div>
                        <div>
                            <label for="description{{.Id}}">New Description:</label><br>
                            <textarea id="description{{.Id}}" name="description">{{.Description}}</textarea><br><br>
                        </div>
                        <label for="price{{.Id}}">New Price:</label><br> <!-- Added New Price Field -->
                        <input type="text" id="price{{.Id}}" name="price" value="{{.Price}}"><br>
                        <label for="quantity{{.Id}}">New Quantity:</label><br> <!-- Added New Quantity Field -->
                        <input type="text" id="quantity{{.Id}}" name="quantity" value="{{.Quantity}}"><br><br>
                        <input type="submit" class="btn btn-primary"value="Update Transaction">
                        <a href="/list" class="btn btn-success ml-2">Back</a>
                    </form>
                </li>
                {{end}}
            {{else}}
                <li>No transaction available</li>
            {{end}}
        </ul>
    </div>

    <script>
        // Fungsi untuk menampilkan formulir update task
        function showUpdateForm(taskId) {
            var formId = 'updateForm' + taskId;
            var form = document.getElementById(formId);
            if (form.style.display === 'none') {
                form.style.display = 'block';
            } else {
                form.style.display = 'none';
            }
        }
    </script>
</body>
</html>
`
