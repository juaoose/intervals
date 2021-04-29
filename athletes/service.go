package athlete

type AthleteService struct {
	athletesRepository *AthletesRepository
}

func (a *AthleteService) createAthlete(athlete Athlete) (*Athlete, error) {
	return a.athletesRepository.CreateUser(athlete)
}

func CreateService(athletesRepository *AthletesRepository) *AthleteService {
	return &AthleteService{athletesRepository: athletesRepository}
}
