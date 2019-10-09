package main

// ขั้นตอนที่ 2 Decoupling with Interface
// โดย interface นั้นใช้สำหรับการกำหนดพฤติกรรมการทำงาน
// ซึ่งถูกใช้งานจากผู้เรียกใช้งานนั่นเอง
// เป็นแนวทางที่มาก ๆ ใน Go

// จากตัวอย่างจะแยกพฤติกรรมออกมาจาก Device และ Storage ดังนี้
// สำหรับ Device จะแยกออกมาเป็น Puller
// สำหรับ Storage จะแยกออกมาเป็น Storer
// เขียน code ได้ดังนี้

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

// ผลที่ตามมาคือ
// ต้องทำการแก้ไข code ในส่วนของ method ที่ทำงาน pull() และ store() ข้อมูล
// จากเดิมที่ส่ง struct เข้าไป ก็เปลี่ยนเป็น interface ได้ดังนี้

func pull(p Puller, data []Data) (int, error) {
	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err // Found error
		}
	}
	return len(data), nil // Success
}

func store(s Storer, data []Data) (int, error) {
	for i:= range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err //Found error
		}
	}
	return len(data), nil // Success
}

// มาถึงตรงนี้เราสามารถแยกส่วนของข้อมูลกับพฤติกรรมการทำงานออกจากกัน
// ทำให้ตอบโจทย์เรื่องการเปลี่ยนแปลงได้ง่ายมากขึ้น
// และลดผลกระทบจากการเปลี่ยนแปลงอีกด้วย

// ในงาน GopherCon มักจะพูดว่า interface คือ first-class citizen ด้วยนะ

// โดยการทำงานร่วมกันนั้น
// เราไม่ต้องรู้หรอกว่า คุณเป็นใคร สนใจเพียงว่าทำอะไรร่วมกันได้บ้าง
// นั่นคือที่มาของ interface นั่นเอง