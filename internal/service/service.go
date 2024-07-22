package service

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"time"
	"vecin/internal/config"
	"vecin/internal/database"
	"vecin/internal/email"
	"vecin/internal/model"
)

type Service struct {
	dao         database.DAO
	Config      *config.Config
	EmailSender email.EmailSender
}

func NewService(dao database.DAO, cfg *config.Config, emailSender email.EmailSender) *Service {
	return &Service{dao: dao, Config: cfg, EmailSender: emailSender}
}

func (s *Service) GenerateToken() (string, error) {
	bytes := make([]byte, s.Config.UserTokenLen)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func (s *Service) CalculateExpiry(duration time.Duration) time.Time {
	return time.Now().Add(duration)
}

func (s *Service) ConfirmAccount(token string) (model.Usuario, error) {
	log.Printf("debug:x token: (%s)", token)
	var userID int
	var expiry time.Time

	tx, err := s.dao.DB().Begin()
	if err != nil {
		log.Printf("debug:x error: (%s), error al iniciar tx", err.Error())
		return model.Usuario{}, err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			// http.Error(w, "Error interno", http.StatusInternalServerError)
			log.Printf("debug:x internal error")
		}
	}()

	err = tx.QueryRow("SELECT usuario_id, fecha_expiracion FROM confirmacion_cuenta WHERE token = $1", token).Scan(&userID, &expiry)
	if err != nil {
		_ = tx.Rollback()
		//http.Error(w, "Token inválido o expirado", http.StatusBadRequest)
		return model.Usuario{}, fmt.Errorf("invalid token or expired: %v", err)
	}

	if time.Now().After(expiry) {
		_ = tx.Rollback()
		//http.Error(w, "El token ha expirado", http.StatusBadRequest)
		return model.Usuario{}, fmt.Errorf("token has expired: %v", err)
	}

	_, err = tx.Exec("UPDATE usuario SET activo = TRUE WHERE usuario_id = $1", userID)
	if err != nil {
		_ = tx.Rollback()
		//http.Error(w, "Error al activar la cuenta", http.StatusInternalServerError)
		return model.Usuario{}, fmt.Errorf("couldn't update usuario: %v", err)
	}

	_, err = tx.Exec("DELETE FROM confirmacion_cuenta WHERE token = $1", token)
	if err != nil {
		_ = tx.Rollback()
		//http.Error(w, "Error al eliminar el token de confirmación", http.StatusInternalServerError)
		return model.Usuario{}, fmt.Errorf("couldn't delete confirmacion: %v", err)
	}

	var userName string
	var nombre string
	var apellido string
	var telefono string
	err = tx.QueryRow("SELECT username, nombre, apellido, telefono FROM usuario WHERE usuario_id = $1", userID).
		Scan(&userName, &nombre, &apellido, &telefono)
	if err != nil {
		_ = tx.Rollback()
		//http.Error(w, "Token inválido o expirado", http.StatusBadRequest)
		return model.Usuario{}, fmt.Errorf("unable to get user info: %v", err)
	}

	if err := tx.Commit(); err != nil {
		//http.Error(w, "Error al confirmar la transacción", http.StatusInternalServerError)
		return model.Usuario{}, fmt.Errorf("couldn't commit transaction: %v", err)
	}

	var userConfirmedAccount model.Usuario
	userConfirmedAccount.Username = userName
	userConfirmedAccount.Nombre = nombre
	userConfirmedAccount.Apellido = apellido
	userConfirmedAccount.Telefono = telefono

	return userConfirmedAccount, nil
}

func (s *Service) SendConfirmationEmail(username, email, token string) error {
	//bypass email sending:
	//return s.EmailSender.Send(username, email, token)
	return nil
}

func (s *Service) SaveUser(signUpFormData model.SignUpFormData, token string) error {
	tx, err := s.dao.DB().Begin()
	if err != nil {
		return err
	}

	// Save the user
	var userID int
	err = tx.QueryRow("INSERT INTO usuario (username, nombre, apellido, telefono, email, password_hash, activo) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING usuario_id",
		signUpFormData.Username, signUpFormData.Nombre, signUpFormData.Apellido, signUpFormData.Telefono, signUpFormData.Email, signUpFormData.Password, false).Scan(&userID)
	if err != nil {
		log.Printf("debug:x error inserting user (table:usuario): %v", err)
		_ = tx.Rollback()
		return err
	}

	expirationDate := time.Now().Add(s.Config.UserTokenExpiryDays)

	_, err = tx.Exec("INSERT INTO confirmacion_cuenta (usuario_id, token, fecha_expiracion) VALUES ($1, $2, $3)",
		userID, token, expirationDate)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CheckEmail(email string) (bool, error) {
	exists, err := s.dao.UserExistsByEmail(email)
	if err != nil {
		log.Printf("debug:x error: (%s)", err)
		return false, err
	}

	return exists, nil
}
