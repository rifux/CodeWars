package kata

import "fmt"

type Dinglemouse struct {
	name  string
	age   int
	sex   rune
	order []string
	seen  map[string]bool
}

func NewDinglemouse() *Dinglemouse {
	return &Dinglemouse{
		seen:  make(map[string]bool),
		order: make([]string, 0, 3),
	}
}

func (d *Dinglemouse) SetAge(age int) *Dinglemouse {
	d.age = age
	if !d.seen["age"] {
		d.order = append(d.order, "age")
		d.seen["age"] = true
	}
	return d
}

func (d *Dinglemouse) SetSex(sex rune) *Dinglemouse {
	d.sex = sex
	if !d.seen["sex"] {
		d.order = append(d.order, "sex")
		d.seen["sex"] = true
	}
	return d
}

func (d *Dinglemouse) SetName(name string) *Dinglemouse {
	d.name = name
	if !d.seen["name"] {
		d.order = append(d.order, "name")
		d.seen["name"] = true
	}
	return d
}

func (d *Dinglemouse) Hello() (result string) {
	result += "Hello."
	for _, l := range d.order {
		switch l {
		case "age":
			result += fmt.Sprintf(" I am %d.", d.age)
		case "name":
			result += fmt.Sprintf(" My name is %s.", d.name)
		case "sex":
			sex := "female"
			if d.sex == 'M' {
				sex = "male"
			}
			result += fmt.Sprintf(" I am %s.", sex)
		}
	}

	return
}
