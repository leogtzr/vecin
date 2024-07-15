package email

import (
	"context"
	"fmt"
	"github.com/mailersend/mailersend-go"
	"log"
	"time"
	"vecin/internal/config"
)

type EmailSender interface {
	Send(user, email, token string) error
}

// MailerSend -> provider: mailersend.com
type MailerSend struct {
	Config config.Mailing
}

func (m MailerSend) Send(user, email, token string) error {
	log.Printf("debug:x send stuff, user=(%s), token=(%s)", user, token)

	ms := mailersend.NewMailersend(m.Config.ApiKey)

	log.Printf("debug:x api=(%s), user=(%s), token=(%s)", m.Config.ApiKey, user, token)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// TODO:put this in an env variable or a configuration:
	subject := "Confirma tu cuenta - Vecin"

	// Variables para el mensaje
	recipientName := user
	confirmationLink := m.Config.ConfirmationLink // "https://example.com/confirm?token=abc123"
	// imageURL := "http://localhost:8180/assets/images/logo.png"

	// Contenido en texto plano
	text := fmt.Sprintf(`Hola, %s,

Gracias por registrarte en Vecin. Por favor, haz clic en el siguiente enlace para confirmar tu cuenta:

%s/%s

Si no te registraste en Vecin, ignora este correo.

Saludos,
El equipo de Vecin`, recipientName, confirmationLink, token)

	html := fmt.Sprintf(`Hola, %s,<br><br>
Gracias por registrarte en Vecin. Por favor, haz clic en el siguiente enlace para confirmar tu cuenta:<br><br>
<a href="%s/%s">Confirmar cuenta</a><br><br>
🏡<br><br>
Si no te registraste en Vecin, ignora este correo.<br><br>
Saludos,<br>
El equipo de Vecin`, recipientName, confirmationLink, token)

	from := mailersend.From{
		Name:  "Vecin",
		Email: "vecin@trial-yzkq3401v96ld796.mlsender.net",
	}

	recipients := []mailersend.Recipient{
		{
			Name:  recipientName,
			Email: email,
		},
	}

	variables := []mailersend.Variables{
		{
			Email: email,
			Substitutions: []mailersend.Substitution{
				{
					Var:   "name",
					Value: recipientName,
				},
				{
					Var:   "link",
					Value: confirmationLink,
				},
			},
		},
	}

	tags := []string{"foo", "bar"}

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)
	message.SetSubstitutions(variables)
	message.SetTags(tags)

	res, err := ms.Email.Send(ctx, message)
	if err != nil {
		log.Printf("Error al enviar el correo: %v", err)
		return err
	}

	messageID := res.Header.Get("X-Message-Id")
	if messageID != "" {
		log.Printf("Correo enviado con éxito. ID del mensaje: %s", messageID)
	} else {
		log.Println("Correo enviado, pero no se recibió el ID del mensaje.")
	}

	return nil
}
