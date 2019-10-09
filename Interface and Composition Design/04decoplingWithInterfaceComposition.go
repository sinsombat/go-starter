package main

// ขั้นตอนที่ 4 Decoupling with Interface composition
// จะเห็นได้ว่า struct System นั้น
// ยังผูกมัดอยู่กับ struct Device และ Storage คือข้อมูลนั่นเอง
// ดังนั้นเราควรที่จะแยกออกมา
// ด้วยการใช้งาน interface composition แทน
// นั่นคือ

type System struct {
	Puller
	Storer
}

func main() {
	sys := System{
		Puller: &Device{},
		Storer: &Storage{},
	}

	Copy(&sys, 3)
}

// มาถึงตรงนี้จะเห็นได้ว่า
// การทำงานแต่ละส่วนแยกออกจากกันอย่างชัดเจนด้วย interface
// ซึ่งทำให้ง่ายต่อการเปลี่ยนแปลง
// และลดผลกระทบจากการเปลี่ยนแปลงอีกด้วย