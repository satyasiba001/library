package models

type Member struct{
	member_id int
	name string
	email string
	phone string
	date string
	unique_id string
}

type Book struct{
	book_id uint16
	title string
	author string
	category string
	count int
}

type BookTransactions struct{
	borrow_id int
	member Member
	book Book
	borrow_date string
	due_date string
	return_date string
}