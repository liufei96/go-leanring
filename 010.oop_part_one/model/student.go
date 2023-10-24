package model

type student struct {
	Name  string
	Score float64
}

// 因为student首字母是小写，因此是只能通过当前包使用
// 通过工厂模式解决
func NewStudent(name string, score float64) *student {
	return &student{
		Name:  name,
		Score: score,
	}
}

// 如果score的首字母小写，则外部不能直接访问，我们可以提供一个方法
func (s *student) getScore() float64 {
	return s.Score
}
