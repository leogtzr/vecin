package handler

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"vecin/internal/database"
	"vecin/internal/middleware"
	"vecin/internal/model"
	"vecin/internal/service"
)

var templateHTMLFiles = []string{
	"internal/template/head.html",
	"internal/template/nav.html",
	"internal/template/footer.html",
	"internal/template/scripts.html",
}

type RequestData struct {
	BookID string `json:"book_id"`
}

type PageVariables struct {
	Year         string
	SiteKey      string
	LoggedIn     bool
	UseAnalytics bool
	AppName      string
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}

func getTemplatePath(templateFileName string) string {
	templateDir := os.Getenv("TEMPLATE_DIR")
	if templateDir == "" {
		templateDir = "internal/template" // default value for local development
	}
	templatePath := filepath.Join(templateDir, templateFileName)

	return templatePath
}

func addTemplateFiles(additionalFiles ...string) []string {
	return append(templateHTMLFiles, additionalFiles...)
}

// ToDo: finish this...
func isLoggedIn(r *http.Request) bool {
	return !true
}

// IndexPage renders the home or index page.
// path: "/"
func IndexPage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Year:     time.Now().Format("2006"),
		AppName:  "Vecin",
		LoggedIn: isLoggedIn(r),
	}

	tmpl, err := template.ParseFiles(
		addTemplateFiles("internal/template/index.html", "internal/template/login.html")...,
	)
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", pageVariables)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// path: "/landing"
func LandingPage(w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	pageVariables := PageVariables{
		Year:    now.Format("2006"),
		AppName: "Vecin",
	}

	tmpl, err := template.ParseFiles(
		addTemplateFiles("internal/template/landing.html", "internal/template/landing_page.html")...,
	)
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", pageVariables)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func Login(dao *database.DAO, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := (*dao).GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashContrasena), []byte(password)); err != nil {
		http.Error(w, "Usuario o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	store := middleware.GetSessionStore()
	session, _ := store.Get(r, "session")
	session.Values["user_id"] = user.ID
	_ = session.Save(r, w)

	http.Redirect(w, r, "/landing", http.StatusSeeOther)
}

func redirectLoginPage(w http.ResponseWriter) {
	templatePath := getTemplatePath("register-login.html")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)

	pageVariables := PageVariables{
		Year:     time.Now().Format("2006"),
		AppName:  "Vecin",
		LoggedIn: false,
	}

	err = t.Execute(w, pageVariables)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
}

// RegisterFracc handles the rendering to register a fraccionamiento.
// path: "/registrar-fraccionamiento"
func RegisterFracc(w http.ResponseWriter, r *http.Request) {
	// TODO:
	// Verificar si el usuario no ha hecho login, si no mandar a hacer una cuenta.
	loggedIn := isLoggedIn(r)
	if !loggedIn {
		redirectLoginPage(w)

		return
	}

	pageVariables := PageVariables{
		Year:    time.Now().Format("2006"),
		AppName: "Vecin",
	}

	tmpl, err := template.ParseFiles(
		addTemplateFiles("internal/template/registrar_fraccionamiento.html")...,
	)
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", pageVariables)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func getDatabaseEmailFromSessionID(db *sql.DB, userID string) (string, error) {
// 	queryStr := "SELECT u.email FROM users u WHERE u.user_id=$1"

// 	rows, err := db.Query(queryStr, userID)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer rows.Close()

// 	var email string

// 	if rows.Next() {
// 		if err := rows.Scan(&email); err != nil {
// 			return "", err
// 		}
// 	}

// 	return email, nil
// }

// func getUserInfoFromAuth0(accessToken string) (*user.UserInfo, error) {
// 	userInfoEndpoint := fmt.Sprintf("https://%s/userinfo", os.Getenv("AUTH0_DOMAIN"))

// 	req, err := http.NewRequest("GET", userInfoEndpoint, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creando la solicitud: %v", err)
// 	}

// 	req.Header.Add("Authorization", "Bearer "+accessToken)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("error al realizar la solicitud: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("error en la respuesta de Auth0: %s", body)
// 	}

// 	var userInfo user.UserInfo
// 	err = json.Unmarshal(body, &userInfo)
// 	if err != nil {
// 		return nil, fmt.Errorf("error al decodificar la respuesta JSON: %v", err)
// 	}

// 	return &userInfo, nil
// }

// func redirectToErrorPage(w http.ResponseWriter, r *http.Request) {
// 	http.Redirect(w, r, "/error", http.StatusSeeOther)
// }

func redirectToErrorPageWithMessageAndStatusCode(w http.ResponseWriter, errorMessage string, httpStatusCode int) {
	type ErrorVariables struct {
		AppName      string
		Year         string
		ErrorMessage string
	}

	pageVariables := ErrorVariables{
		Year:         time.Now().Format("2006"),
		ErrorMessage: errorMessage,
		AppName:      "Vecin",
	}

	w.WriteHeader(httpStatusCode)

	tmpl, err := template.ParseFiles(
		addTemplateFiles("internal/template/error5xx.html")...,
	)
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", pageVariables)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// func redirectToErrorLoginPage(w http.ResponseWriter) {
// 	templatePath := getTemplatePath("errorLogin.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		log.Printf("error: %v", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	type ErrorVariables struct {
// 		Year         string
// 		ErrorMessage string
// 	}

// 	w.WriteHeader(http.StatusUnauthorized)

// 	err = t.Execute(w, nil)
// 	if err != nil {
// 		log.Printf("error: %v", err)
// 		return
// 	}
// }

// func writeErrorGeneralStatus(w http.ResponseWriter, err error) {
// 	log.Printf("error: %v", err)

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]string{
// 		"status": "error",
// 	})
// }

// func writeUnauthenticated(w http.ResponseWriter) {
// 	w.Header().Set("Content-Type", "application/json")
// Note: might be a good idea to add the status code to the header (404? 200? 401? 402? 5XX?)
// 	json.NewEncoder(w).Encode(map[string]string{"status": "unauthenticated"})
// }

// func getCurrentUserID(r *http.Request) (string, error) {
// 	session, err := auth.SessionStore.Get(r, "user-session")
// 	if err != nil {
// 		return "", err
// 	}

// 	userID, ok := session.Values["user_id"].(string)
// 	if !ok {
// 		return "", errors.New("user_id not found in session")
// 	}

// 	for k, v := range session.Values {
// 		fmt.Printf("k=(%v), v=(%v)\n", k, v)
// 	}

// 	fmt.Println("--------")
// 	fmt.Println(session)
// 	fmt.Println(userID)
// 	fmt.Println("----- end")

// 	return userID, nil
// }

// func BooksByAuthorPage(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	now := time.Now()
// 	pageVariables := PageVariablesForAuthors{
// 		Year:         now.Format("2006"),
// 		SiteKey:      captcha.SiteKey,
// 		UseAnalytics: useAnalytics,
// 	}

// 	authors, err := (*database).GetAllAuthors()
// 	if err != nil {
// 		log.Printf("Error getting authors: %v", err)
// 		redirectToErrorPageWithMessageAndStatusCode(w, "error getting information from the database", http.StatusInternalServerError)
// 		return
// 	}

// 	_, err = getCurrentUserID(r)
// 	if err != nil {
// 		log.Printf("(BooksByAuthorPage) User is not logged in: %v", err)
// 		pageVariables.LoggedIn = false
// 	} else {
// 		log.Println("User is logged in")
// 		pageVariables.LoggedIn = true
// 	}

// 	pageVariables.Authors = authors

// 	templatePath := getTemplatePath("books_by_author.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Printf("Error parsing template: %v", err)
// 		return
// 	}

// 	err = t.Execute(w, pageVariables)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Printf("Error executing template: %v", err)
// 	}
// }

// 	var pages []int
// 	for i := start; i <= end; i++ {
// 		pages = append(pages, i)
// 	}

// 	pageVariables.Pages = pages

// 	pageVariables.StartPage = start
// 	pageVariables.EndPage = end

// 	return nil
// }

// 	pageVariables := PageResultsVariables{}
// 	pageVariables.UseAnalytics = useAnalytics
// 	pageVariables.Results = books

// 	err = setUpPaginationFor(pageInt, database, &pageVariables)
// 	if err != nil {
// 		log.Printf("Error setting up pagination: %v", err)
// 		redirectToErrorPageWithMessageAndStatusCode(w, "error getting information from the database", http.StatusInternalServerError)
// 		return
// 	}

// 	setAuthenticationForPageResults(r, &pageVariables, database)

// 	templateDir := os.Getenv("TEMPLATE_DIR")
// 	if templateDir == "" {
// 		templateDir = "internal/template" // default value for local development
// 	}
// 	templatePath := getTemplatePath("allbooks.html")

// 	t := template.New("").Funcs(sprig.TxtFuncMap())

// 	tmpl, err := t.ParseFiles(templatePath)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Printf("Error parsing template: %v", err)
// 		return
// 	}

// 	err = tmpl.ExecuteTemplate(w, "allbooks.html", pageVariables)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Printf("Error executing template: %v", err)
// 	}
// }

// //func Autocomplete(db *sql.DB, w http.ResponseWriter, r *http.Request) {
// //	query := r.URL.Query().Get("q")
// //
// //	searchTypesStr := r.URL.Query().Get("searchType")
// //	searchTypes := strings.Split(searchTypesStr, ",")
// //
// //	var suggestions []string
// //
// //	var queryStr string
// //	var rows *sql.Rows
// //	var err error
// //
// //	// Perform DB query based on queryParam("q")
// //
// //	w.Header().Set("Content-Type", "application/json")
// //	json.NewEncoder(w).Encode(map[string][]string{
// //		"suggestions": suggestions,
// //	})
// //}

// func BooksList(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	authorParam := r.URL.Query().Get("start_with")

// 	booksByAuthor, err := (*database).GetBooksBySearchTypeCoincidence(authorParam, book.ByAuthor)
// 	if err != nil {
// 		log.Printf("error: %v", err)
// 		http.Error(w, "Database error", http.StatusInternalServerError)
// 		return
// 	}

// 	type BookDetail struct {
// 		ID           int                  `json:"id"`
// 		Title        string               `json:"title"`
// 		Author       string               `json:"author"`
// 		Description  string               `json:"description"`
// 		Base64Images []book.BookImageInfo `json:"images"`
// 	}

// 	var results []BookDetail

// 	for _, book := range booksByAuthor {
// 		bookDetail := BookDetail{}
// 		bookDetail.ID = book.ID
// 		bookDetail.Title = book.Title
// 		bookDetail.Author = book.Author
// 		bookDetail.Description = book.Description
// 		bookDetail.Base64Images = book.Base64Images

// 		results = append(results, bookDetail)
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	_ = json.NewEncoder(w).Encode(results)
// }

// func getTotalBooks(db *sql.DB) (int, error) {
// 	queryStr := `SELECT count(*) FROM books`
// 	rows, err := db.Query(queryStr)
// 	if err != nil {
// 		return -1, err
// 	}

// 	var count int

// 	for rows.Next() {
// 		err := rows.Scan(&count)
// 		if err != nil {
// 			return -1, err
// 		}
// 	}

// 	return count, nil
// }

// func BooksCount(database *database.DAO, w http.ResponseWriter) {
// 	count, err := (*database).GetBookCount()
// 	if err != nil {
// 		http.Error(w, "Database error", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(map[string]int{
// 		"booksCount": count,
// 	})
// }

// func SearchBooksPage(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	bookQuery := r.URL.Query().Get("textSearch")
// 	searchTypesStr := r.URL.Query().Get("searchType")
// 	searchTypesParams := uniqueSearchTypes(strings.Split(searchTypesStr, ","))

// 	if len(searchTypesParams) == 0 || (len(searchTypesParams) == 1 && searchTypesParams[0] == "") {
// 		searchTypesParams = []string{"byTitle"}
// 	}

// 	var results []book.BookInfo
// 	var err error

// 	for _, searchTypeParam := range searchTypesParams {
// 		searchType := parseBookSearchType(searchTypeParam)
// 		switch searchType {
// 		case book.ByTitle:
// 			booksByTitle, err := (*database).GetBooksBySearchTypeCoincidence(bookQuery, book.ByTitle)
// 			if err != nil {
// 				log.Printf("error: %v", err)
// 				redirectToErrorPageWithMessageAndStatusCode(w, "Error getting information from the database", http.StatusInternalServerError)

// 				return
// 			}
// 			results = append(results, booksByTitle...)

// 		case book.ByAuthor:
// 			booksByAuthor, err := (*database).GetBooksBySearchTypeCoincidence(bookQuery, book.ByAuthor)
// 			if err != nil {
// 				log.Printf("error getting info from the database: %v", err)
// 				redirectToErrorPageWithMessageAndStatusCode(w, "error getting info from the database", http.StatusInternalServerError)
// 				return
// 			}
// 			results = append(results, booksByAuthor...)

// 		case book.Unknown:
// 			log.Printf("Tipo de búsqueda en libros desconocido.")
// 			redirectToErrorPageWithMessageAndStatusCode(w, "Wrong search", http.StatusInternalServerError)

// 			return
// 		}
// 	}

// 	now := time.Now()
// 	pageVariables := PageResultsVariables{
// 		Year:         now.Format("2006"),
// 		SiteKey:      captcha.SiteKey,
// 		Results:      results,
// 		UseAnalytics: useAnalytics,
// 	}

// 	templatePath := getTemplatePath("search_books.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		log.Printf("template error: %v", err)
// 		redirectToErrorPageWithMessageAndStatusCode(w, "template error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = t.Execute(w, pageVariables)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("error: %v", err)
// 		return
// 	}
// }

// func ErrorPage(w http.ResponseWriter, _ *http.Request) {
// 	templateDir := os.Getenv("TEMPLATE_DIR")
// 	if templateDir == "" {
// 		templateDir = "internal/template"
// 	}
// 	templatePath := getTemplatePath("error5xx.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	err = t.Execute(w, nil)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// }

// func IngresarPage(w http.ResponseWriter, r *http.Request) {
// 	oauthState := generateRandomString(32)

// 	session, _ := auth.SessionStore.Get(r, "user-session")
// 	session.Values["oauth_state"] = oauthState
// 	session.Save(r, w)

// 	if isDevMode() {
// 		http.Redirect(w, r, "/auth/callback", http.StatusSeeOther)

// 		return
// 	}

// 	//url := auth.GoogleOauthConfig.AuthCodeURL(oauthState)
// 	url := auth.Config.AuthCodeURL(oauthState)
// 	http.Redirect(w, r, url, http.StatusSeeOther)
// }

// func setDevCredentialsAndRedirect(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	userInfo, err := getDevUserInfo()
// 	if err != nil {
// 		log.Printf("error: cannot get user info from Auth0: %v", err)
// 		http.Error(w, "cannot get user info from Auth0", http.StatusInternalServerError)

// 		return
// 	}

// 	err = (*database).AddUser(userInfo.Sub, userInfo.Email, userInfo.Name, "Google")
// 	if err != nil {
// 		http.Error(w, "Error al guardar el usuario en la base de datos", http.StatusInternalServerError)
// 		return
// 	}

// 	session, _ := auth.SessionStore.Get(r, "user-session")
// 	session.Values["user_id"] = userInfo.Sub
// 	session.Save(r, w)

// 	now := time.Now()

// 	pageVariables := PageVariables{
// 		Year:    now.Format("2006"),
// 		SiteKey: captcha.SiteKey,
// 	}

// 	_, err = getCurrentUserID(r)
// 	if err != nil {
// 		pageVariables.LoggedIn = false
// 	} else {
// 		pageVariables.LoggedIn = true
// 	}

// 	templateDir := os.Getenv("TEMPLATE_DIR")
// 	if templateDir == "" {
// 		templateDir = "internal/template"
// 	}
// 	templatePath := filepath.Join(templateDir, "index.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Printf("Error al analizar la plantilla: %v", err)
// 		return
// 	}

// 	err = t.Execute(w, pageVariables)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Printf("Error al ejecutar la plantilla: %v", err)
// 	}
// }

// func Auth0Callback(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	if isDevMode() {
// 		setDevCredentialsAndRedirect(database, w, r)

// 		return
// 	}

// 	code := r.URL.Query().Get("code")

// 	token, err := auth.Config.Exchange(r.Context(), code)
// 	if err != nil {
// 		log.Printf("Error: %v", err)
// 		http.Error(w, "Cannot get Auth0 token", http.StatusInternalServerError)
// 		return
// 	}

// 	userInfo, err := getUserInfoFromAuth0(token.AccessToken)
// 	if err != nil {
// 		log.Printf("error: cannot get user info from Auth0: %v", err)
// 		http.Error(w, "cannot get user info from Auth0", http.StatusInternalServerError)
// 		return
// 	}

// 	err = (*database).AddUser(userInfo.Sub, userInfo.Email, userInfo.Name, "Google")
// 	if err != nil {
// 		http.Error(w, "Error al guardar el usuario en la base de datos", http.StatusInternalServerError)
// 		return
// 	}

// 	session, _ := auth.SessionStore.Get(r, "user-session")
// 	session.Values["user_id"] = userInfo.Sub
// 	session.Save(r, w)

// 	now := time.Now()

// 	pageVariables := PageVariables{
// 		Year:    now.Format("2006"),
// 		SiteKey: captcha.SiteKey,
// 	}

// 	_, err = getCurrentUserID(r)
// 	if err != nil {
// 		pageVariables.LoggedIn = false
// 	} else {
// 		pageVariables.LoggedIn = true
// 	}

// 	templateDir := os.Getenv("TEMPLATE_DIR")
// 	if templateDir == "" {
// 		templateDir = "internal/template"
// 	}
// 	templatePath := filepath.Join(templateDir, "index.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Printf("Error al analizar la plantilla: %v", err)
// 		return
// 	}

// 	err = t.Execute(w, pageVariables)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		log.Printf("Error al ejecutar la plantilla: %v", err)
// 	}
// }

// func setAuthenticationForPageResults(r *http.Request, pageResultsVariables *PageResultsVariables, database *database.DAO) {
// 	dbID, err := getCurrentUserID(r)
// 	if err != nil {
// 		log.Printf("error: checking authentication information for user '%v'", err)
// 		pageResultsVariables.LoggedIn = false
// 	} else {
// 		pageResultsVariables.LoggedIn = true
// 		if isDevMode() {
// 			pageResultsVariables.IsAdmin = true
// 			return
// 		}
// 		userInfo, err := (*database).GetUserInfoByID(dbID)
// 		if err != nil {
// 			log.Printf("error: %v", err)
// 		}
// 		pageResultsVariables.IsAdmin = userInfo.Email == os.Getenv("LEONLIB_MAINAPP_USER")
// 	}
// }

// func CheckLikeStatus(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	userID, err := getCurrentUserID(r)
// 	if err != nil {
// 		writeUnauthenticated(w)

// 		return
// 	}

// 	vars := mux.Vars(r)
// 	bookID := vars["book_id"]

// 	exists, err := (*database).LikedBy(bookID, userID)
// 	if err != nil {
// 		writeErrorLikeStatus(w, err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	if exists {
// 		json.NewEncoder(w).Encode(map[string]string{"status": "liked"})
// 	} else {
// 		json.NewEncoder(w).Encode(map[string]string{"status": "not-liked"})
// 	}
// }

// func LikeBook(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	userID, err := getCurrentUserID(r)
// 	if err != nil {
// 		http.Error(w, "Error al obtener información de la sesión", http.StatusInternalServerError)
// 		return
// 	}

// 	err = r.ParseForm()
// 	if err != nil {
// 		w.Write([]byte(fmt.Sprintf("error like book: %v", err.Error())))
// 	}
// 	bookID := r.PostFormValue("book_id")

// 	err = (*database).LikeBook(bookID, userID)

// 	if err != nil {
// 		http.Error(w, "Error al dar like en la base de datos", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte("Liked successfully"))
// }

// func AddBook(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseMultipartForm(2 << 20)
// 	if err != nil {
// 		log.Printf("error adding book: %v", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	book := book.BookInfo{}
// 	book.Title = r.FormValue("title")
// 	book.Author = r.FormValue("author")
// 	book.Description = r.FormValue("description")
// 	book.HasBeenRead = r.FormValue("read") == "on"
// 	book.GoodreadsLink = r.FormValue("goodreadsLink")

// 	var imageData []byte
// 	file, _, err := r.FormFile("image")
// 	if err == nil {
// 		defer file.Close()
// 		imageData, err = io.ReadAll(file)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	} else if !errors.Is(err, http.ErrMissingFile) {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	book.Image = imageData

// 	err = (*database).CreateBook(book)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte("Libro agregado con éxito"))
// }

// func UnlikeBook(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	userID, err := getCurrentUserID(r)
// 	if err != nil {
// 		http.Error(w, "Error al obtener información de la sesión", http.StatusInternalServerError)
// 		return
// 	}

// 	var requestData RequestData

// 	err = json.NewDecoder(r.Body).Decode(&requestData)
// 	if err != nil {
// 		http.Error(w, "Error al decodificar el cuerpo de la solicitud", http.StatusInternalServerError)
// 		return
// 	}

// 	bookID := requestData.BookID

// 	fmt.Printf("debug:x trying to unlike book_id=(%s), user_id=(%s)\n", bookID, userID)

// 	err = (*database).UnlikeBook(bookID, userID)
// 	if err != nil {
// 		http.Error(w, "Error al quitar el like en la base de datos", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte("Unliked successfully"))
// }

// func LikesCount(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	bookID := r.URL.Query().Get("book_id")
// 	if bookID == "" {
// 		http.Error(w, "book_id is required", http.StatusBadRequest)
// 		return
// 	}

// 	id, err := strconv.Atoi(bookID)
// 	if err != nil {
// 		http.Error(w, "Invalid book_id", http.StatusBadRequest)
// 		return
// 	}

// 	var count int
// 	count, err = (*database).LikesCount(id)
// 	if err != nil {
// 		http.Error(w, "Error querying the database", http.StatusInternalServerError)
// 		return
// 	}

// 	resp := map[string]int{
// 		"count": count,
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(resp)
// }

// func CreateDBFromFile(database *database.DAO, w http.ResponseWriter) {
// 	libraryDir := "library"
// 	libraryDirPath := filepath.Join(libraryDir, "books_db.toml")

// 	var library book.Library

// 	if _, err := toml.DecodeFile(libraryDirPath, &library); err != nil {
// 		writeErrorGeneralStatus(w, err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")

// 	startTime := time.Now()

// 	err := (*database).AddAll(library.Book)
// 	if err != nil {
// 		writeErrorGeneralStatus(w, err)
// 		return
// 	}

// 	elapsedTime := time.Since(startTime)

// 	log.Printf("Books loaded in: %.2f seconds\n", elapsedTime.Seconds())

// 	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
// }

// func InfoBook(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	idQueryParam := r.URL.Query().Get("id")

// 	id, err := strconv.Atoi(idQueryParam)
// 	if err != nil {
// 		redirectToErrorPage(w, r)
// 		return
// 	}

// 	bookByID, err := (*database).GetBookByID(id)
// 	if err != nil {
// 		log.Printf("error: getting information from the database")
// 		redirectToErrorPageWithMessageAndStatusCode(w, "error getting information from the database", http.StatusInternalServerError)
// 		return
// 	}

// 	now := time.Now()

// 	pageVariables := &PageResultsVariables{
// 		Year:         now.Format("2006"),
// 		SiteKey:      captcha.SiteKey,
// 		Results:      []book.BookInfo{bookByID},
// 		UseAnalytics: useAnalytics,
// 	}

// 	setAuthenticationForPageResults(r, pageVariables, database)

// 	templatePath := getTemplatePath("book_info.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		redirectToErrorPage(w, r)
// 		return
// 	}

// 	err = t.Execute(w, pageVariables)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("template error: %v", err)
// 		return
// 	}
// }

// func ModifyBook(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	// TODO: check auth here
// 	err := r.ParseMultipartForm(2 << 20)
// 	if err != nil {
// 		writeErrorGeneralStatus(w, err)

// 		return
// 	}

// 	bookIDParam := r.FormValue("book_id")
// 	title := r.FormValue("title")
// 	author := r.FormValue("author")
// 	description := r.FormValue("description")
// 	read := r.FormValue("read") == "on"
// 	goodreadsLink := r.FormValue("goodreadsLink")

// 	id, err := strconv.Atoi(bookIDParam)
// 	if err != nil {
// 		writeErrorGeneralStatus(w, err)

// 		return
// 	}

// 	err = addImageToBook(database, id, r)
// 	if err != nil {
// 		writeErrorGeneralStatus(w, err)

// 		return
// 	}

// 	err = (*database).UpdateBook(title, author, description, read, goodreadsLink, id)
// 	if err != nil {
// 		writeErrorGeneralStatus(w, err)

// 		return
// 	}

// 	w.Write([]byte("Libro modificado con exito"))
// }

// func addImageToBook(database *database.DAO, id int, r *http.Request) error {
// 	var imageData []byte
// 	file, _, err := r.FormFile("image")
// 	if err == nil {
// 		defer file.Close()
// 		imageData, err = io.ReadAll(file)
// 		if err != nil {
// 			return err
// 		}
// 	} else if !errors.Is(err, http.ErrMissingFile) {
// 		return err
// 	}

// 	if len(imageData) == 0 {
// 		return nil
// 	}

// 	err = (*database).AddImageToBook(id, imageData)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func ModifyBookPage(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	idQueryParam := r.URL.Query().Get("book_id")

// 	id, err := strconv.Atoi(idQueryParam)
// 	if err != nil {
// 		redirectToErrorPageWithMessageAndStatusCode(w, "wrong ID", http.StatusInternalServerError)
// 		return
// 	}

// 	bookByID, err := (*database).GetBookByID(id)
// 	if err != nil {
// 		redirectToErrorPageWithMessageAndStatusCode(w, "error getting information from the database", http.StatusInternalServerError)
// 		return
// 	}

// 	now := time.Now()

// 	type BookToModifyVariables struct {
// 		Year          string
// 		SiteKey       string
// 		Book          book.BookInfo
// 		LoggedIn      bool
// 		GoodreadsLink template.URL
// 	}

// 	pageVariables := BookToModifyVariables{
// 		Year:          now.Format("2006"),
// 		SiteKey:       captcha.SiteKey,
// 		Book:          bookByID,
// 		GoodreadsLink: template.URL(bookByID.GoodreadsLink),
// 	}

// 	//_, err = getCurrentUserID(r)
// 	//if err != nil {
// 	//	redirectToErrorPageWithMessageAndStatusCode(w, "Error al obtener información de la sesión", http.StatusInternalServerError)
// 	//
// 	//	return
// 	//}
// 	pageVariables.LoggedIn = true

// 	templatePath := getTemplatePath("modify.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		redirectToErrorPageWithMessageAndStatusCode(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = t.Execute(w, pageVariables)
// 	if err != nil {
// 		redirectToErrorPageWithMessageAndStatusCode(w, err.Error(), http.StatusInternalServerError)

// 		return
// 	}
// }

// func AboutPage(w http.ResponseWriter, r *http.Request) {
// 	templatePath := getTemplatePath("about.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		redirectToErrorPageWithMessageAndStatusCode(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	now := time.Now()

// 	pageVariables := PageVariables{
// 		Year:         now.Format("2006"),
// 		SiteKey:      captcha.SiteKey,
// 		LoggedIn:     false,
// 		UseAnalytics: useAnalytics,
// 	}

// 	err = t.Execute(w, pageVariables)
// 	if err != nil {
// 		redirectToErrorPageWithMessageAndStatusCode(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func ContactPage(w http.ResponseWriter, _ *http.Request) {
// 	templatePath := getTemplatePath("contact.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		redirectToErrorPageWithMessageAndStatusCode(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	now := time.Now()

// 	pageVariables := PageVariables{
// 		Year:         now.Format("2006"),
// 		SiteKey:      captcha.SiteKey,
// 		LoggedIn:     false,
// 		UseAnalytics: useAnalytics,
// 	}

// 	err = t.Execute(w, pageVariables)
// 	if err != nil {
// 		redirectToErrorPageWithMessageAndStatusCode(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func AddBookPage(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	dbID, err := getCurrentUserID(r)
// 	if err != nil {
// 		redirectToErrorLoginPage(w)
// 		return
// 	}

// 	if !isDevMode() {
// 		userInfo, err := (*database).GetUserInfoByID(dbID)
// 		if err != nil {
// 			log.Printf("error: %v", err)
// 			redirectToErrorLoginPage(w)

// 			return
// 		}
// 		if userInfo.Email != os.Getenv("LEONLIB_MAINAPP_USER") {
// 			log.Printf("error: %s is not %s", userInfo.Email, os.Getenv("LEONLIB_MAINAPP_USER"))

// 			redirectToErrorLoginPage(w)
// 			return
// 		}
// 	}
// 	templatePath := getTemplatePath("add_book.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		log.Printf("template error: %v", err)
// 		redirectToErrorPageWithMessageAndStatusCode(w, fmt.Sprintf("template error: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	now := time.Now()

// 	pageVariables := PageVariables{
// 		Year:     now.Format("2006"),
// 		SiteKey:  captcha.SiteKey,
// 		LoggedIn: false,
// 	}

// 	err = t.Execute(w, pageVariables)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		log.Printf("template error: %v", err)
// 		return
// 	}
// }

// func RemoveImage(database *database.DAO, w http.ResponseWriter, r *http.Request) {
// 	// TODO: check auth
// 	r.ParseForm()
// 	imageIDParam := r.PostFormValue("image_id")

// 	imageID, err := strconv.Atoi(imageIDParam)
// 	if err != nil {
// 		http.Error(w, "Error removing image", http.StatusInternalServerError)
// 		return
// 	}

// 	err = (*database).RemoveImage(imageID)
// 	if err != nil {
// 		http.Error(w, "Error removing image", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte("Image removed OK..."))
// }

// func WishListBooksPage(database *database.DAO, w http.ResponseWriter, _ *http.Request) {
// 	templatePath := getTemplatePath("wishlistbooks.html")

// 	t, err := template.ParseFiles(templatePath)
// 	if err != nil {
// 		redirectToErrorPageWithMessageAndStatusCode(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	now := time.Now()

// 	var results []book.WishListBook

// 	results, err = (*database).GetWishListBooks()
// 	if err != nil {
// 		redirectToErrorPageWithMessageAndStatusCode(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	pageVariables := PageResultsVariablesForWishList{
// 		Year:         now.Format("2006"),
// 		SiteKey:      captcha.SiteKey,
// 		IsAdmin:      false, // TODO: pending
// 		LoggedIn:     false,
// 		Results:      results,
// 		UseAnalytics: useAnalytics,
// 	}

// 	err = t.Execute(w, pageVariables)
// 	if err != nil {
// 		redirectToErrorPageWithMessageAndStatusCode(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func isDevMode() bool {
// 	runMode := os.Getenv("RUN_MODE")

// 	return runMode == "dev"
// }

func FormRegisterFracc(dao *database.DAO, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var formData model.RegisterFormData
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		redirectToErrorPageWithMessageAndStatusCode(w, "Unable to process input data", http.StatusInternalServerError)

		return
	}

	log.Printf("Form Data: %v", formData)

	// Save the data:
	comunidadID, err := (*dao).SaveCommunity(formData)
	// TODO: fix this...

	w.WriteHeader(http.StatusOK)
	resp := map[string]string{
		"message":           "RegisterPage OK",
		"fraccionamientoID": fmt.Sprintf("%d", comunidadID),
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func GenError(dao *database.DAO, w http.ResponseWriter, r *http.Request) {
	redirectToErrorPageWithMessageAndStatusCode(w, "error: just testing...", http.StatusInternalServerError)
}

func ViewFraccsPage(dao *database.DAO, w http.ResponseWriter, r *http.Request) {
	now := time.Now()

	pageVariables := PageVariables{
		Year:    now.Format("2006"),
		AppName: "Vecin",
	}

	tmpl, err := template.ParseFiles(
		addTemplateFiles("internal/template/view-fraccs.html")...,
	)
	if err != nil {
		log.Printf("Error parsing templates: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", pageVariables)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func RegisterPage(dao *database.DAO, w http.ResponseWriter, r *http.Request) {
	//now := time.Now()
	//
	//pageVariables := PageVariables{
	//	Year:    now.Format("2006"),
	//	AppName: "Vecin",
	//}
	//
	//tmpl, err := template.ParseFiles(
	//	addTemplateFiles("internal/template/register.html")...,
	//)
	//if err != nil {
	//	log.Printf("Error parsing templates: %v", err)
	//	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	return
	//}
	//
	//err = tmpl.ExecuteTemplate(w, "base", pageVariables)
	//if err != nil {
	//	log.Printf("Error executing template: %v", err)
	//	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	//	return
	//}

	templatePath := getTemplatePath("register.html")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		redirectToErrorPageWithMessageAndStatusCode(w, err.Error(), http.StatusInternalServerError)
		return
	}

	now := time.Now()

	pageVariables := PageVariables{
		Year:     now.Format("2006"),
		LoggedIn: false,
		AppName:  "Vecin",
	}

	err = t.Execute(w, pageVariables)
	if err != nil {
		redirectToErrorPageWithMessageAndStatusCode(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateAccountPage(dao *database.DAO, w http.ResponseWriter, r *http.Request) {
	templatePath := getTemplatePath("create-account.html")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type ErrorVariables struct {
		Year         string
		ErrorMessage string
	}

	w.WriteHeader(http.StatusUnauthorized)

	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
}

func SignUp(svc *service.Service, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	var signUpFormData model.SignUpFormData
	err := json.NewDecoder(r.Body).Decode(&signUpFormData)
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		redirectToErrorPageWithMessageAndStatusCode(w, "Unable to process input data", http.StatusInternalServerError)

		return
	}

	log.Printf("Form Data: %v", signUpFormData)

	if signUpFormData.Password != signUpFormData.ConfirmPassword {
		writePasswordDoNotMatchToResponse(w)

		return
	}

	token, err := svc.GenerateToken()
	if err != nil {
		log.Printf("Error generating token: %v", err)
		redirectToErrorPageWithMessageAndStatusCode(w, "Unable to generate confirmation token", http.StatusInternalServerError)
		return
	}

	err = svc.SaveUser(signUpFormData, token)
	if err != nil {
		log.Printf("Error saving user: %v", err)
		redirectToErrorPageWithMessageAndStatusCode(w, "Unable to save user", http.StatusInternalServerError)
		return
	}

	err = svc.SendConfirmationEmail(signUpFormData.Username, signUpFormData.Email, token)
	if err != nil {
		log.Printf("Error sending confirmation email: %v", err)

		redirectToErrorPageWithMessageAndStatusCode(w, "Unable to send confirmation email", http.StatusInternalServerError)
		return
	}

	// TODO: do something useful with this information...
	w.WriteHeader(http.StatusOK)
	resp := map[string]string{
		"message": "SignUp OK",
		"id":      "1",
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}

func writePasswordDoNotMatchToResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"message": "passwords do not match"})
}

func CheckEmail(svc *service.Service, w http.ResponseWriter, r *http.Request) {
	type CheckEmailRequest struct {
		Email string `json:"email"`
	}

	var req CheckEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)

		return
	}

	type CheckEmailResponse struct {
		Exists bool `json:"exists"`
	}

	exists, err := svc.CheckEmail(req.Email)
	if err != nil {
		log.Printf("Error checking email: %v", err)
		http.Error(w, "Error checking email", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(CheckEmailResponse{
		Exists: exists,
	})
}
