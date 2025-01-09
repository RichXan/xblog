package designpattern

// Person 定义人员结构
type Person struct {
	name   string
	gender string
	age    int
	salary float64
}

func NewPerson(name string, gender string, age int, salary float64) *Person {
	return &Person{
		name:   name,
		gender: gender,
		age:    age,
		salary: salary,
	}
}

// Criteria 定义过滤器接口
type Criteria interface {
	MeetCriteria(persons []*Person) []*Person
}

// MaleCriteria 男性过滤器
type MaleCriteria struct{}

func (c *MaleCriteria) MeetCriteria(persons []*Person) []*Person {
	filtered := make([]*Person, 0)
	for _, person := range persons {
		if person.gender == "Male" {
			filtered = append(filtered, person)
		}
	}
	return filtered
}

// FemaleCriteria 女性过滤器
type FemaleCriteria struct{}

func (c *FemaleCriteria) MeetCriteria(persons []*Person) []*Person {
	filtered := make([]*Person, 0)
	for _, person := range persons {
		if person.gender == "Female" {
			filtered = append(filtered, person)
		}
	}
	return filtered
}

// AgeCriteria 年龄过滤器
type AgeCriteria struct {
	minAge int
	maxAge int
}

func NewAgeCriteria(minAge, maxAge int) *AgeCriteria {
	return &AgeCriteria{
		minAge: minAge,
		maxAge: maxAge,
	}
}

func (c *AgeCriteria) MeetCriteria(persons []*Person) []*Person {
	filtered := make([]*Person, 0)
	for _, person := range persons {
		if person.age >= c.minAge && person.age <= c.maxAge {
			filtered = append(filtered, person)
		}
	}
	return filtered
}

// SalaryCriteria 薪资过滤器
type SalaryCriteria struct {
	minSalary float64
}

func NewSalaryCriteria(minSalary float64) *SalaryCriteria {
	return &SalaryCriteria{minSalary: minSalary}
}

func (c *SalaryCriteria) MeetCriteria(persons []*Person) []*Person {
	filtered := make([]*Person, 0)
	for _, person := range persons {
		if person.salary >= c.minSalary {
			filtered = append(filtered, person)
		}
	}
	return filtered
}

// AndCriteria 组合过滤器（与）
type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func NewAndCriteria(criteria Criteria, otherCriteria Criteria) *AndCriteria {
	return &AndCriteria{
		criteria:      criteria,
		otherCriteria: otherCriteria,
	}
}

func (c *AndCriteria) MeetCriteria(persons []*Person) []*Person {
	firstCriteriaPersons := c.criteria.MeetCriteria(persons)
	return c.otherCriteria.MeetCriteria(firstCriteriaPersons)
}

// OrCriteria 组合过滤器（或）
type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

func NewOrCriteria(criteria Criteria, otherCriteria Criteria) *OrCriteria {
	return &OrCriteria{
		criteria:      criteria,
		otherCriteria: otherCriteria,
	}
}

func (c *OrCriteria) MeetCriteria(persons []*Person) []*Person {
	firstCriteriaItems := c.criteria.MeetCriteria(persons)
	otherCriteriaItems := c.otherCriteria.MeetCriteria(persons)

	for _, person := range otherCriteriaItems {
		if !containsPerson(firstCriteriaItems, person) {
			firstCriteriaItems = append(firstCriteriaItems, person)
		}
	}
	return firstCriteriaItems
}

// 辅助函数：检查列表是否包含特定人员
func containsPerson(persons []*Person, person *Person) bool {
	for _, p := range persons {
		if p == person {
			return true
		}
	}
	return false
}
