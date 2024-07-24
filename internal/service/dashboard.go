package service

import "log"

// La ventana de bienvenida debería mostrarse cuando el usuario no ha registrado un fraccionamiento (como el que paga $$)
// O no se ha unido a un fraccionamiento existente.
func (s *Service) ShouldShowWelcomePageIfNotRegistered(userID int) (bool, error) {
	hasRegistered, err := s.dao.HasRegisteredAFracc(userID)
	if err != nil {
		log.Printf("debug:x error checking if the user has registered a community: %v", err)
		return true, err
	}

	isPartOf, err := s.dao.IsPartOfComunidad(userID)
	if err != nil {
		log.Printf("debug:x error checking if the user is part of a community (habitant): %v", err)
		return true, err
	}

	log.Printf("debug:x HasRegisteredAFracc=%v, IsPartOfComunidad=%v, userID=%d", hasRegistered, isPartOf, userID)

	return hasRegistered || isPartOf, nil
}
