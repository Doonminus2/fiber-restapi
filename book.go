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

// Handler functions

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

// function createbook ไว้ใช้สร้าง post methon
func createBook(c *fiber.Ctx) error {
	book := new(Book)
	// มีการ ทำดัก error ไว้ในกรณีที่ ส่งข้อมูลเป็น json เเล้วไม่ถูก type ตามที่ struct กำหนดไว้ ก็จะขึ้น json: cannot unmarshal string into Go struct field Book.id of type int
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)
}

// updateBook = แก้ไขข้อมูลหนังสือเดิมผ่าน id
func updataBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())

	}

	for i, book := range books {
		if book.ID == id {
			book.Title = bookUpdate.Title
			book.Author = bookUpdate.Author
			books[i] = book
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)

}

func deleteBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.SendStatus(fiber.StatusNotFound)
}
