package public


//Create an interface
type House interface {
    Location() string
	Price() int
}

//Implement the interface using a struct Familyhouse
type FamilyHouse struct {
    Address string
    Cost int
}
//Implement the interface using a struct Mansion
type Mansion struct {
    Address string
    Cost int
}
//Implement the interface using a struct Apartment
type Apartment struct {
    Address string
    Cost int
}

//Implement the Location() and Price() methods for each struct
func (f FamilyHouse) Location() string {
    return f.Address
}
func (f FamilyHouse) Price() int {
    return f.Cost
}
func (m Mansion) Location() string {
    return m.Address
}
func (m Mansion) Price() int {
    return m.Cost
}
func (a Apartment) Location() string {
    return a.Address
}
func (a Apartment) Price() int {
    return a.Cost
}