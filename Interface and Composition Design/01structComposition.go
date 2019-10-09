package main
// ขั้นตอนที่ 1 Structure composition
// ในการพัฒนาระบบนั้นเริ่มต้น
// จากการแบ่งงานใหญ่ออกเป็นงานย่อย ๆ

// โดย code ตัวอย่างนั้น
// ต้องการดึงข้อมูลจากอุปกรณ์ต่าง ๆ ออกมา
// ซึ่งข้อมูลมีมากกว่า 1 บรรทัด
// จากนั้นทำการจัดเก็บข้อมูลเหล่านั้น
// มาดูกันว่า ทำการแยกงานย่อย ๆ ได้อะไรบ้าง

// 1. สร้างโครงสร้างข้อมูลชื่อว่า Data
// โดยสร้างด้วย struct ดังนี้

type Data struct {
	Line string
}

// 2. สร้างส่วนการดึงข้อมูลจากอุปกรณ์
// โดยข้อมูลของอุปกรณ์ประกอบไปด้วย ชื่อเครื่องกับ Timeout
// สร้างด้วย struct และมี method Pull() สำหรับดึงข้อมูล ดังนี้

type Device struct {
	Host string
	Timeout time.Duration
}

func (*Device) Pull(d *Data) error {
	// TODO
	return nil
}

// 3. สร้างส่วนการจัดเก็บข้อมูลจากอุปกรณ์
// มีโครงสร้างเหมือนกัน แต่สิ่งที่ต่างคือมี method Store() สำหรับจัดเก็บข้อมูลดังนี้

type Storage struct {
	Host string
	Timeout time.Duration
}

func (*Storage) Store(d *Data) error {
	// TODO
	return nil
}

// 4. สร้างส่วน logic ของการ pull ข้อมูลจากอุปกรณ์จริง ๆ
// ด้วยการสร้าง method pull() ขึ้นมา
// จากนั้น extract ข้อมูลแต่ละบรรทัดออกมาดังนี้

func pull(d *Device, data []Data) (int, error) {
	for i := range data {
		if err := d.Pull(&data[i]); err != nil {
			return i, err // Found error
		}
	}
	return len(data), nil // Success
}

// 5. สร้างส่วน logic ของการ store ข้อมูล
// ด้วยการสร้าง method store() ขึ้นมาดังนี้

func store(s *Storage, data []Data) (int, error) {
	for i:= range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err //Found error
		}
	}
	return len(data), nil // Success
}

// 6. สิ่งที่ยังขาดไปคือ ส่วนการทำงานหลัก สำหรับดึงและบันทึกข้อมูล

type System struct {
	Device
	Storage
}

func Copy( sys *System, batch int) error {
	data := make([]Data, batch)
	for {
		i, _ := pull(&sys.Device, data)
		store(&sys.Storage, data[:i])
	}
}

// จะเห็นได้ว่าเราทำงานแบ่งปัญหาใหญ่ ๆ ออกเป็นปัญหาเล็ก ๆ
// เพื่อทำการแก้ไขไปเรื่อย ๆ
// จากนั้นทำการรวมหรือ composition เข้าด้วยกัน
// จนสามารถแก้ไขปัญหาที่ต้องการได้
// นี่คือ ความสามารถที่สำคัญของนักพัฒนา software

// แต่ว่านี่เป็นเพียงการเริ่มต้นเท่านั้น
// สำหรับ Coding for Today
// มันยังมีต่ออีกยาว !!

// จากตัวอย่าง code นั้นพบว่า ทำงานได้อย่างดี
// แต่ปัญหาที่น่าสนใจ หรือ Code Smell คือ
// ในตอนนี้ Data กับ Behavior ของ Device และ Storage รวมกันอยู่
// นั่นคือผูกมัดกันอย่างมาก

// ลองคิดดูสิว่า ถ้ามีจำนวน Device และ Storage จำนวนมาก
// ซึ่งมีพฤติกรรมการทำงานที่แตกต่างกัน
// ดังนั้นจึงต้องหาวิธีแก้ไขปัญหา ?