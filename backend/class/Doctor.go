package class

type Doctor struct {
	Id int `json:"Id"`
	Nume string `json:"nume"`
	Prenume string `json:"prenume"`
	Mail string `json:"mail"`
	Specialitate string `json:"specialitate"`
	Cnp string `json:"cnp"`
}
func (r Doctor) GetId() int {
	return r.Id
}

func (r Doctor) GetNume() string {
	return r.Nume
}

func (r Doctor) GetPrenume() string {
	return r.Prenume
}

func (r Doctor) GetMail() string {
	return r.Mail
}

func (r Doctor) GetSpecialitate() string {
	return r.Specialitate
}

func (r Doctor) GetCnp() string {
	return r.Cnp
}

