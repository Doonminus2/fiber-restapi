package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// use struct to hold book data. sturct is the one of data suructure is go
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// นี้คือ slice คล้าย array เเต่ ไม่ต้องประกาศขนาดทำงานคล้ายๆ list ใน python
var books []Book

func main() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "Chitsanupat", Author: "pi"})
	books = append(books, Book{ID: 2, Title: "MM", Author: "Mike"})
	// app.Get "/books" คือ การเปิด port ขึ้นมาสามรถยิง get ใส่ได้ ละจะ return ค่าทั้งหมดใน ตัวเเปร books ออกมาเป็น json
	app.Get("/books", getBooks)
	// /id คือ การใส่พารามิดเตอร์เข้าไปถ้าลองยิงใน postman เช่น http://127.0.0.1:8080/books/100 ก็จะ return 100 กลับมาใน console ชอง postman
	app.Get("/books/:id", getOneBook)

	app.Listen(":8080")

}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getOneBook(c *fiber.Ctx) error {
	// import strconv เเละ ใช้ strconv.Atoi ใส่ params ในการเเปลง type จาก int เป็น string เเต่ต้องรับ ค่ามา2 ตัวตาม ตัวอย่างการใช้งานใน doc เเละ อีกตัวที่เอามา err คือ เป็นการ handle error
	bookid, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		// statusBedRequset คือ error400 เเละส่งค่าไปเป็น 400 string โดยจะส่ง error เมื่อไม่สามรถเเปลงค่าเป็น stingได้
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// loopเข้าไปอ่าน ข้อมูลใน book จะ loop อ่านทุกตัวใน slice โดยจะ return ค่ามา 2ตัวคือ ค่าตำเเหน่ง index เเละ 8jk value เเต่สามรถ ใา่ _ เพื่อบอกว่าไม่เอาค่า ตำเเหน่งindex ได้
	for _, book := range books {
		if book.ID == bookid {
			return c.JSON(book)
		}
	}
	// รับค่ามาเป็น int จะ return เป็น string ไม่ได้ เเต่สามรถ ใช้ strconv ในการเเปลง type ของ data ได้
	return c.Status(fiber.StatusNotFound).SendString("Not found i dum")
}
