package main

import (
	"database/sql"
	"fmt"

	"log"
	"net/http"
	"strconv"

	"forum/internal/config"
	"forum/internal/handlers"
	"forum/internal/middleware"
	"forum/internal/repositories"
	"forum/internal/services"
)

func registerHandlers(db *sql.DB) {
	// Create the repositories
	postRepo := repositories.NewSQLitePostRepository(db)
	userRepo := repositories.NewSQLiteUserRepository(db)
	commentRepo := repositories.NewSqliteCommentRepository(db)
	// Create the services
	postService := services.NewPostService(*postRepo)
	authService := services.NewAuthService(*userRepo)
	commentService := services.NewCommentService(*commentRepo)
	//userService := services.NewRegistrationService(*userRepo)

	registrationService := services.NewRegistrationService(*userRepo)

	// Create the handlers
	postHandler := handlers.NewPostHandler(*postService, "templates/Home.html")
	authHandler := handlers.NewAuthHandler(*authService)
	commentHandler := handlers.NewCommentHandler(*commentService)
	loginHandler, _ := handlers.NewLoginHandler("templates/Login.html")
	resetpasswordHnadler, _ := handlers.NewPasswordReset("templates/Resetpassword.html")
	//userHandler := handlers.NewRegistrationHandler(*userService)
	registrationHandler := handlers.NewRegistrationHandler(*registrationService)
	homeHandler, err := handlers.NewHomeHandler("templates/Home.html", *postService)
	if err != nil {
		log.Fatal(err)
	}

	// Register the handlers
	// public endpoints
	http.HandleFunc("/posts/all", postHandler.GetAllPosts)
	http.HandleFunc("/auth/login", authHandler.Login)
	http.HandleFunc("/auth/register", registrationHandler.RegisterUser)
	http.HandleFunc("/comments/display", commentHandler.GetAllCommentsById)
	http.HandleFunc("/posts/category", postHandler.GetPostByCategory)
	http.HandleFunc("/", homeHandler.HomePage)
	http.HandleFunc("/login", loginHandler.LoginPage)
	http.HandleFunc("/resetpassword", resetpasswordHnadler.PasswordResetPage)
	http.HandleFunc("/auth/logout", authHandler.Logout)

	// Apply authentication middleware to the "/restricted" endpoint
	http.Handle("/posts/create", middleware.AuthenticationMiddleware(http.HandlerFunc(postHandler.CreatePost)))

	http.Handle("/posts/liked_post", middleware.AuthenticationMiddleware(http.HandlerFunc(postHandler.GetLikedPostByUser)))

	http.Handle("/posts/created_post", middleware.AuthenticationMiddleware(http.HandlerFunc(postHandler.GetPostByUserID)))

	http.Handle("/postlike/", middleware.AuthenticationMiddleware(http.HandlerFunc(postHandler.LikePost)))

	http.Handle("/postdislike/", middleware.AuthenticationMiddleware(http.HandlerFunc(postHandler.DislikePost)))

	http.Handle("/comments/add", middleware.AuthenticationMiddleware(http.HandlerFunc(commentHandler.CreateComment)))

	http.Handle("/comments/likes", middleware.AuthenticationMiddleware(http.HandlerFunc(commentHandler.LikesComment)))

	http.Handle("/comments/dislikes", middleware.AuthenticationMiddleware(http.HandlerFunc(commentHandler.DislikesComment)))
}

func main() {
	// Create the configuration
	cfg := config.NewConfig()

	// Create the SQLite database connection
	db, err := sql.Open("sqlite3", cfg.DatabasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	config.CreateTables(db) // Call the function to create tables

	// Register the handlers with SQLite repositories
	registerHandlers(db)

	// Serve static files (CSS, JavaScript, images, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Start the server
	addr := ":" + strconv.Itoa(cfg.Port)
	fmt.Println(addr)
	log.Printf("Server listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
