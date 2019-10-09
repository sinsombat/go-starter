package main

// ขั้นตอนที่ 3 Interface Composition
// จะเห็นได้ว่าการใช้งาน interface Puller และ Storer นั้นจะทำงานร่วมกันเสมอ
// ดังจะเห็นได้จาก method Copy()
// แสดงดัง code

func Copy( sys *System, batch int) error {
	data := make([]Data, batch)
	for {
		i, _ := pull(&sys.Device, data)
		store(&sys.Storage, data[:i])
	}
}

// สังเกตุได้ว่าสิ่งที่ส่งมายัง method Copy() คือ System
// ซึ่งเป็น struct ที่รวมเอา Device และ Storage ไว้ด้วยกัน
// แต่สิ่งที่ควรส่งมาคือ พฤติกรรมของการทำงานมากกว่า !!

// คำถามคือ จะส่งอะไรมาระหว่าง Puller และ Storer ?
// คำตอบคือทั้งสองตัว !!

// สิ่งที่ต้องทำคือ สร้าง interface ใหม่ชื่อว่า PullStorer
// เพื่อทำงานรวมทั้งสอง interface เข้าด้วยกัน
// แสดงดัง code

type PullStorer interface {
	Puller
	Storer
}

func Copy( ps PullStorer, batch int) error {
	data := make([]Data, batch)
	for {
		i, _ := pull(ps, data)
		store(ps, data[:i])
	}
}

// ยังไม่พอนะ ยังไม่จบ