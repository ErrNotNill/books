package models

type Crud interface {
	Create()
	Read()
	Update()
	Delete()
}

func (b *Books)Create(){
	//POST books
	//make id
}
func (a *Authors)Create(){
	//POST authors
	//make id
}
func (b *Books)Read(){
	//GET books
	//get by id
}
func (a *Authors)Read(){
	//GET authors
	//get by id
}
func (b *Books)Update(){
	//PUT books
	//put by id
}
func (a *Authors)Update(){
	//PUT authors
	//put by id
}
func (b *Books)Delete(){
	//DELETE books
	//delete by id
}
func (a *Authors)Delete(){
	//DELETE authors
	//delete by id
}