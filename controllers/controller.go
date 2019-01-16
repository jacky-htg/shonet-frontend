package controllers

type Page struct {
	Title       string
	Description string
	Data        interface{}
}

type Fashion struct {

}

type Beauty struct {

}

type Journal struct {

}

type Menu struct {
	Fashion Fashion
	Beauty  Beauty
	Journal Journal
}


func GetMenu() (Menu, error) {

	return Menu{}, nil
}
