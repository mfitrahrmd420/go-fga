package user

type Student struct {
	No           int
	Name         string
	AbsentReason string
}

// constructor method
func NewStudent(name string, num int) User {
	return &Teacher{name: name, num: num}
}

// implementation
func (t *Student) Absent() {
	t.AbsentReason = "BOLOS BOS"
}

func (t *Student) Attend() {
}
