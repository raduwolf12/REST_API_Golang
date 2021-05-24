package class


type Pacient struct {
	Id int `json:"Id"`
	Nume string `json:"nume"`
	Prenume string `json:"prenume"`
	Mail string `json:"mail"`
	Cnp string `json:"cnp"`
}

func (r Pacient) GetId() int {
	return r.Id
}

func (r Pacient) GetNume() string {
	return r.Nume
}

func (r Pacient) GetPrenume() string {
	return r.Prenume
}

func (r Pacient) GetMail() string {
	return r.Mail
}

func (r Pacient) GetCnp() string {
	return r.Cnp
}


