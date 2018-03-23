package wrong

func TestFilmsListOwnershipMapping(t *testing.T) {
	// preparing
	ownershipMapping := s.FilmsListOwnershipMapper(user, ownerships)
	// checks
}

func TestFilmsListOwnershipMappingSet(t *testing.T) {
	// preparing
	mappedFilms := s.FilmsListOwnershipMappingSet(films, ownershipMapping, mapperSettings)
	// checks
}

func TestFilmsListOwnershipMappingSort(t *testing.T) {
	// preparing
	sortedFilms := s.FilmsListOwnershipMappingSort(films, ownershipMapping, "id")
	// checks
}

// ... many tests
