package main

// ยังไม่จบนะ จะสังเกตุได้ว่า มี composition interface ที่หน้าตาเหมือนกันเลย
// นั่นคือเกิด duplication code ขึ้นมาแล้ว (Remove interface pollution)
// ดังนั้นให้ทำการลบ code ส่วนนั้นทิ้งไปซะ
// นั่นคือ interface PullStorer นั่นเอง
// จะได้ code ทั้งหมดดังนี้

type Data struct {
	Line string
}

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

type System struct {
	Puller
	Storer
}

type Device struct {
	Host string
	Timeout time.Duration
}


func (*Device) Pull(d *Data) error {
	// TODO
	return nil
}

type Storage struct {
	Host string
	Timeout time.Duration
}

func (*Storage) Store(d *Data) error {
	// TODO
	return nil
}

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


func Copy( s *System, batch int) error {
	data := make([]Data, batch)
	for {
		i, _ := pull(s, data)
		store(s, data[:i])
	}
}

func main() {
	sys := System{
		Puller: &Device{},
		Storer: &Storage{},
	}

	Copy(&sys, 3)
}

// มาถึงจุดนี้กันได้อย่างไร ?
// รู้สึกอย่างไรบ้างสำหรับการออกแบบ Interface และ Composition

// สุดท้ายแล้วยังมีเรื่องที่น่าสนใจและมึนงงอีกมากมาย
// ทั้งเรื่องการทำงานของ Garbase Collection
// ทั้งเรื่องการทำงานของ Slice
// ทั้งเรื่องการทำงานของ Goroutine
// ทั้งเรื่องการทำงานของ Concurrency
// ทั้งเรื่องของ Profiling และ Optimization

// แน่นอนว่าต้องใช้เวลาในการย่อยอีกสักพัก
// แล้วจะนำมาเขียนสรุปต่อไปครับ

// ขอให้สนุกกับการเขียน code

// ที่สำคัญต้องเข้าใจด้วยว่าปัญหาที่คุณกำลังแก้ไขคืออะไร
// นั่นคือคุณเข้าใจ domain และ data นั่นเอง