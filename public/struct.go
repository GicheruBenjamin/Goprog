package public


type Car struct {
    Name string
    Year int
}
func (c Car) GetName() string {
    return c.Name
}
func (c Car) GetYear() int {
    return c.Year
}
