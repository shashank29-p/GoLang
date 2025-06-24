type Person struct {
	//Structural data
	Name  string
	Age   int
	Email string
	//Nested struct
	Address Address
	//Anonymous struct - Use only meant to be used once
	Phone  struct{
		CountryCode string
		Number      string
	}

}

type Address struct {
	Street  string
	City    string
	State   string
}

type Government struct {
	//Embedded struct
	Person
}