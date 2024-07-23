package service

// La ventana de bienvenida debería mostrarse cuando el usuario no ha registrado un fraccionamiento (como el que paga $$)
// O no se ha unido a un fraccionamiento existente.
func (s *Service) ShouldShowWelcomePageIfNotRegistered(userID int) (bool, error) {
	hasRegistered, err := s.dao.HasRegisteredAFracc(userID)
	if err != nil {
		return true, err
	}

	isPartOf, err := s.dao.IsPartOfComunidad(userID)
	if err != nil {
		return true, err
	}

	return (hasRegistered || isPartOf), nil
}
