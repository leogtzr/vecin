package database

import (
	"database/sql"
	"errors"
	"log"
	"vecin/internal/model"
)

// func (database *postgresBookDAO) AddAll(books []book.BookInfo) error {
// 	for _, book := range books {
// 		log.Printf("Reading: (%s)", book)
// 		bookInfo, err := database.GetBookByID(book.ID)
// 		if err == nil && bookInfo.ID == book.ID {
// 			log.Printf("Book with ID: %d already exists, skipping", book.ID)
// 			continue
// 		}

// 		var bookID int
// 		stmt, err := database.db.Prepare("INSERT INTO books(id, title, author, description, read, added_on, goodreads_link) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id")
// 		if err != nil {
// 			return err
// 		}

// 		err = stmt.QueryRow(book.ID, book.Title, book.Author, book.Description, book.HasBeenRead, book.AddedOn, book.GoodreadsLink).Scan(&bookID)
// 		if err != nil {
// 			return err
// 		}

// 		for _, imageName := range book.ImageNames {
// 			imgBytes, err := os.ReadFile(filepath.Join("images", imageName))
// 			if err != nil {
// 				return err
// 			}

// 			imgStmt, err := database.db.Prepare("INSERT INTO book_images(book_id, image) VALUES($1, $2)")
// 			if err != nil {
// 				return err
// 			}

// 			_, err = imgStmt.Exec(bookID, imgBytes)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

// func (database *postgresBookDAO) AddImageToBook(bookID int, imageData []byte) error {
// 	return addImageToBook(bookID, imageData, database.db)
// }

// func (database *postgresBookDAO) AddUser(userID, email, name, oauthIdentifier string) error {
// 	return addUser(database.db, userID, email, name, oauthIdentifier)
// }

func (dao *daoImpl) Close() error {
	return dao.db.Close()
}

func (dao *daoImpl) GetUserByUsername(username string) (*model.Usuario, error) {
	query := `SELECT usuario_id, username, nombre, apellido, telefono, email, password_hash, activo 
              FROM usuario WHERE nombre_usuario = $1`
	row := dao.db.QueryRow(query, username)

	var user model.Usuario
	err := row.Scan(&user.ID, &user.Username, &user.Nombre, &user.Apellido, &user.Telefono, &user.Email, &user.HashContrasena, &user.Activo)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err // Usuario no encontrado
		}
		return nil, err
	}

	return &user, nil
}

func (dao *daoImpl) GetUserByEmail(email string) (*model.Usuario, error) {
	query := `SELECT usuario_id, username, nombre, apellido, telefono, email, password_hash, activo 
              FROM usuario WHERE email = $1`
	row := dao.db.QueryRow(query, email)

	var user model.Usuario
	err := row.Scan(&user.ID, &user.Username, &user.Nombre, &user.Apellido, &user.Telefono, &user.Email, &user.HashContrasena, &user.Activo)
	if err != nil {
		log.Printf("debug:x error=(%v)", err)
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("debug:x error 1")
			return nil, err
		}

		log.Printf("debug:x error 2")
		return nil, err
	}

	return &user, nil
}

func (dao *daoImpl) UserExistsByEmail(email string) (bool, error) {
	query := `SELECT 1 FROM usuario WHERE email = $1`
	row := dao.db.QueryRow(query, email)

	var exists int
	if err := row.Scan(&exists); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// SaveCommunity saves a comunity into the database.
func (dao *daoImpl) SaveCommunity(data model.RegisterFormData) (int, error) {
	var comunidadID int
	var err error

	err = dao.db.QueryRow(`
        INSERT INTO comunidad (nombre, direccion_calle, direccion_numero, direccion_colonia, direccion_cp, direccion_ciudad, direccion_estado, direccion_pais, tipo, modelo_suscripcion)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING comunidad_id
    `,
		data.NombreComunidad,
		data.DireccionCalle,
		data.DireccionNumero,
		data.DireccionColonia,
		data.DireccionCodigoPostal,
		data.DireccionCiudad,
		data.DireccionEstado,
		data.DireccionPais,
		data.TipoComunidad,
		data.ModeloSuscripcion).Scan(&comunidadID)
	if err != nil {
		return -1, err
	}

	log.Printf("debug:x done saving to comunidad, id=%d", comunidadID)

	return comunidadID, nil
}

// TODO: check what is going on here.
func (dao *daoImpl) SaveUser(data model.SignUpFormData) (int, error) {
	return -1, nil
}

func (dao *daoImpl) HasRegisteredAFracc(userID int) (bool, error) {
	query := `SELECT COUNT(*) FROM comunidad WHERE usuario_registrante_id = $1`

	var count int
	if err := dao.db.QueryRow(query, userID).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (dao *daoImpl) IsPartOfComunidad(userID int) (bool, error) {
	query := `
        SELECT COUNT(*)
        FROM habitante
        WHERE email = (
            SELECT email
            FROM usuario
            WHERE usuario_id = $1
        )
    `

	var count int
	if err := dao.db.QueryRow(query, userID).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func (dao *daoImpl) DB() *sql.DB {
	return dao.db
}

/*
func (dao *daoImpl) SaveCommunity(data model.RegisterFormData) (int, error) {
	var comunidadID int
	var suscripcionID int
	var pagoID int

	tx, err := dao.db.Begin()
	if err != nil {
		return -1, err
	}

	// Insertar en la tabla comunidad
	err = tx.QueryRow(`
        INSERT INTO comunidad (nombre, direccion_calle, direccion_numero, direccion_colonia, direccion_cp, direccion_ciudad, direccion_estado, direccion_pais, tipo, modelo_suscripcion)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING comunidad_id
    `,
		data.NombreComunidad,
		data.DireccionCalle,
		data.DireccionNumero,
		data.DireccionColonia,
		data.DireccionCodigoPostal,
		data.DireccionCiudad,
		data.DireccionEstado,
		data.DireccionPais,
		data.TipoComunidad,
		data.ModeloSuscripcion).Scan(&comunidadID)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	log.Printf("debug:x done saving to comunidad, id=%d", comunidadID)

	// Insertar en la tabla suscripcion
	err = tx.QueryRow(`
        INSERT INTO suscripcion (usuario_id, comunidad_id, modelo_suscripcion, fecha_inicio, fecha_fin, monto)
        VALUES ($1, $2, $3, $4, $5, $6) RETURNING suscripcion_id
    `,
		data.UsuarioID, // Debes asegurarte de tener UsuarioID en el data
		comunidadID,
		data.ModeloSuscripcion,
		data.FechaInicioSuscripcion, // Asegúrate de tener estos campos en el data
		data.FechaFinSuscripcion,
		data.MontoSuscripcion).Scan(&suscripcionID)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	log.Printf("debug:x done saving to suscripcion, id=%d", suscripcionID)

	// Insertar en la tabla pago
	err = tx.QueryRow(`
        INSERT INTO pago (suscripcion_id, fecha_pago, monto)
        VALUES ($1, $2, $3) RETURNING pago_id
    `,
		suscripcionID,
		data.FechaPago, // Asegúrate de tener estos campos en el data
		data.MontoPago).Scan(&pagoID)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	log.Printf("debug:x done saving to pago, id=%d", pagoID)

	// Confirmar la transacción
	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return comunidadID, nil
}

*/

// func (database *postgresBookDAO) CreateBook(book book.BookInfo) error {
// 	stmt, err := database.db.Prepare("INSERT INTO books (title, author, description, read, goodreads_link) VALUES ($1, $2, $3, $4, $5)")
// 	if err != nil {
// 		return err
// 	}
// 	defer stmt.Close()

// 	insertedBookIDResult, err := stmt.Exec(book.Title, book.Author, book.Description, book.HasBeenRead, book.GoodreadsLink)
// 	if err != nil {
// 		return err
// 	}

// 	insertedBookID, err := insertedBookIDResult.LastInsertId()
// 	if err != nil {
// 		return err
// 	}

// 	imgStmt, err := database.db.Prepare("INSERT INTO book_images(book_id, image) VALUES($1, $2)")
// 	if err != nil {
// 		return err
// 	}

// 	_, err = imgStmt.Exec(insertedBookID, book.Image)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (database *postgresBookDAO) GetAllAuthors() ([]string, error) {
// 	return getAllAuthors(database.db)
// }

// func (database *postgresBookDAO) GetBookByID(id int) (book.BookInfo, error) {
// 	var err error
// 	var queryStr = `SELECT b.id, b.title, b.author, b.description, b.read, b.added_on, b.goodreads_link FROM books b WHERE b.id=$1`

// 	bookRows, err := database.db.Query(queryStr, id)
// 	if err != nil {
// 		return book.BookInfo{}, err
// 	}

// 	defer func() {
// 		_ = bookRows.Close()
// 	}()

// 	var bookInfo book.BookInfo
// 	var bookID int
// 	var title string
// 	var author string
// 	var description string
// 	var hasBeenRead bool
// 	var addedOn time.Time
// 	var goodreadsLink sql.NullString
// 	if bookRows.Next() {
// 		if err := bookRows.Scan(&bookID, &title, &author, &description, &hasBeenRead, &addedOn, &goodreadsLink); err != nil {
// 			return book.BookInfo{}, err
// 		}

// 		bookInfo.ID = bookID
// 		bookInfo.Title = title
// 		bookInfo.Author = author
// 		bookInfo.Description = description
// 		bookInfo.HasBeenRead = hasBeenRead
// 		bookInfo.AddedOn = addedOn.Format("2006-01-02")
// 		if goodreadsLink.Valid {
// 			bookInfo.GoodreadsLink = goodreadsLink.String
// 		} else {
// 			bookInfo.GoodreadsLink = ""
// 		}
// 	}

// 	bookImages, err := database.GetImagesByBookID(id)
// 	if err != nil {
// 		return book.BookInfo{}, err
// 	}

// 	bookInfo.Base64Images = bookImages

// 	return bookInfo, nil
// }

// func (database *postgresBookDAO) GetBookCount() (int, error) {
// 	return getBookCount(database.db)
// }

// func (database *postgresBookDAO) GetBooksWithPagination(offset, limit int) ([]book.BookInfo, error) {
// 	return getBooksWithPagination(offset, limit, database.db)
// }

// func (database *postgresBookDAO) GetBooksBySearchTypeCoincidence(titleSearchText string, bookSearchType book.BookSearchType) ([]book.BookInfo, error) {
// 	var err error
// 	queryStr := `SELECT b.id, b.title, b.author, b.description, b.read, b.added_on, b.goodreads_link FROM books b WHERE LOWER(b.title) LIKE '%' || LOWER($1) || '%' ORDER BY b.title`

// 	if bookSearchType == book.ByAuthor {
// 		queryStr = `SELECT b.id, b.title, b.author, b.description, b.read, b.added_on, b.goodreads_link FROM books b WHERE LOWER(b.author) LIKE '%' || LOWER($1) || '%' ORDER BY b.title`
// 	}

// 	booksByTitleRows, err := database.db.Query(queryStr, "%"+titleSearchText+"%")
// 	if err != nil {
// 		return []book.BookInfo{}, err
// 	}

// 	defer booksByTitleRows.Close()

// 	var books []book.BookInfo
// 	var id int
// 	var title string
// 	var author string
// 	var description string
// 	var hasBeenRead bool
// 	var addedOn time.Time
// 	var goodreadsLink string
// 	for booksByTitleRows.Next() {
// 		var bookInfo book.BookInfo
// 		if err := booksByTitleRows.Scan(&id, &title, &author, &description, &hasBeenRead, &addedOn, &goodreadsLink); err != nil {
// 			return []book.BookInfo{}, err
// 		}

// 		bookInfo.ID = id
// 		bookInfo.Title = title
// 		bookInfo.Author = author
// 		bookImages, err := database.GetImagesByBookID(id)
// 		if err != nil {
// 			return []book.BookInfo{}, err
// 		}

// 		bookInfo.Base64Images = bookImages
// 		bookInfo.Description = description
// 		bookInfo.HasBeenRead = hasBeenRead
// 		bookInfo.AddedOn = addedOn.Format("2006-01-02")
// 		books = append(books, bookInfo)
// 	}

// 	return books, nil
// }

// func (database *postgresBookDAO) GetImagesByBookID(bookID int) ([]book.BookImageInfo, error) {
// 	return getImagesByBookID(bookID, database.db)
// }

// func (database *postgresBookDAO) GetUserInfoByID(id string) (user.UserInfo, error) {
// 	var err error
// 	var queryStr = `SELECT u.user_id, u.email, u.name FROM users u WHERE u.user_id=$1`

// 	userRow, err := database.db.Query(queryStr, id)
// 	if err != nil {
// 		return user.UserInfo{}, err
// 	}

// 	defer func() {
// 		_ = userRow.Close()
// 	}()

// 	var userInfo user.UserInfo
// 	var userID string
// 	var email string
// 	var name string
// 	if userRow.Next() {
// 		if err := userRow.Scan(&userID, &email, &name); err != nil {
// 			return user.UserInfo{}, err
// 		}

// 		userInfo.Sub = userID
// 		userInfo.Email = email
// 		userInfo.Name = name
// 	}

// 	return userInfo, nil
// }

// func (database *postgresBookDAO) LikedBy(bookID, userID string) (bool, error) {
// 	queryStr := "SELECT EXISTS(SELECT 1 FROM book_likes WHERE book_id=$1 AND user_id=$2)"

// 	rows, err := database.db.Query(queryStr, bookID, userID)
// 	if err != nil {
// 		return false, err
// 	}
// 	defer rows.Close()

// 	var exists bool

// 	if rows.Next() {
// 		if err := rows.Scan(&exists); err != nil {
// 			return false, err
// 		}
// 	}

// 	if err != nil {
// 		return false, err
// 	}

// 	return exists, nil
// }

// func (database *postgresBookDAO) LikeBook(bookID, userID string) error {
// 	_, err := database.db.Exec("INSERT INTO book_likes(book_id, user_id) VALUES($1, $2) ON CONFLICT(book_id, user_id) DO NOTHING", bookID, userID)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (database *postgresBookDAO) LikesCount(bookID int) (int, error) {
// 	var count int
// 	if err := database.db.QueryRow("SELECT COUNT(*) FROM book_likes WHERE book_id = $1", bookID).Scan(&count); err != nil {
// 		return -1, err
// 	}

// 	return count, nil
// }

func (dao *daoImpl) Ping() error {
	return dao.db.Ping()
}

// func (database *postgresBookDAO) RemoveImage(imageID int) error {
// 	if _, err := database.db.Exec("DELETE FROM book_images WHERE image_id=$1", imageID); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (database *postgresBookDAO) UnlikeBook(bookID, userID string) error {
// 	if _, err := database.db.Exec("DELETE FROM book_likes WHERE book_id=$1 AND user_id=$2", bookID, userID); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (database *postgresBookDAO) UpdateBook(title string, author string, description string, read bool, goodreadsLink string, id int) error {
// 	return updateBook(title, author, description, read, goodreadsLink, id, database.db)
// }

// func (database *postgresBookDAO) GetWishListBooks() ([]book.WishListBook, error) {
// 	return database.wishListBooks, nil
// }
