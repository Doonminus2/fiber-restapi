package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Chitsanupat", Author: "pi"})
	books = append(books, Book{ID: 2, Title: "MM", Author: "Mike"})
	// app.Get "/books" คือ การเปิด port ขึ้นมาสามรถยิง get ใส่ได้ ละจะ return ค่าทั้งหมดใน ตัวเเปร books ออกมาเป็น json
	app.Get("/books", getBooks)
	// /id คือ การใส่พารามิดเตอร์เข้าไปถ้าลองยิงใน postman เช่น http://127.0.0.1:8080/books/100 ก็จะ return 100 กลับมาใน console ชอง postman
	app.Get("/books/:id", getOneBook)

	// สร้าง methons post เพื่อสร้างข้อมูล
	app.Post("/books", createBook)

	app.Put("/books/:id", updataBook)

	app.Delete("/books/:id", deleteBook)
	app.Listen(":8080")
}
