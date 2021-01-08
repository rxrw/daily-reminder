package services

import (
	"strconv"
	"strings"
)

//ILottery 抽象接口
type ILottery interface {
	AmIWin() int
}

//Lottery 抽象类
type Lottery struct {
	correct string
	current string
}

func (l *Lottery) sameNum(correct []string, current []string) int {
	sameNum := 0
	i, j := 0, 0
	for {
		if i > len(correct)-1 || j > len(current)-1 {
			break
		}
		intCorrect, _ := strconv.Atoi(correct[i])
		intCurrent, _ := strconv.Atoi(current[j])
		if intCorrect == intCurrent {
			sameNum++
			i++
			j++
		} else if intCorrect < intCurrent {
			i++
		} else {
			j++
		}
	}
	return sameNum
}

//NewLottery creates a new Lottery
func NewLottery(correct string, current string, name string) ILottery {
	switch name {
	case "ssq":
		return &SsqLottery{lottery: &Lottery{correct: correct, current: current}}
	case "dlt":
		return &DltLottery{lottery: &Lottery{correct: correct, current: current}}
	case "fcsd":
		return &FcsdLottery{lottery: &Lottery{correct: correct, current: current}}
	}
	return nil
}

//SsqLottery 双色球
type SsqLottery struct {
	lottery *Lottery
}

//FcsdLottery 3d
type FcsdLottery struct {
	lottery *Lottery
}

// DltLottery 大乐透
type DltLottery struct {
	lottery *Lottery
}

//AmIWin int
func (s *SsqLottery) AmIWin() int {
	//最后一位是篮球 拿出来
	current := strings.Split(s.lottery.current, ",")
	currentBlue, _ := strconv.Atoi(current[len(current)-1])
	current = append(current[:len(current)-1])
	correct := strings.Split(s.lottery.correct, ",")
	correctBlue, _ := strconv.Atoi(correct[len(correct)-1])
	correct = append(correct[:len(correct)-1])
	sameNum := s.lottery.sameNum(correct, current)
	var level int
	if currentBlue != correctBlue {
		//蓝球没中的情况
		if sameNum < 2 {
			level = 7
		} else if sameNum == 4 {
			level = 5
		} else if sameNum == 5 {
			level = 4
		} else if sameNum == 6 {
			level = 2
		}
	} else {
		//蓝球中了的情况
		if sameNum < 3 {
			level = 6
		} else if sameNum == 3 {
			level = 5
		} else if sameNum == 4 {
			level = 4
		} else if sameNum == 5 {
			level = 3
		} else if sameNum == 6 {
			level = 1
		}
	}
	switch level {
	case 1:
		return -1
	case 2:
		return -1
	case 3:
		return 3000
	case 4:
		return 200
	case 5:
		return 10
	case 6:
		return 5
	default:
		return 0
	}
}

//AmIWin int
func (s *FcsdLottery) AmIWin() int {
	if s.lottery.correct == s.lottery.current {
		return 1000
	}
	return 0
}

//AmIWin int
func (s *DltLottery) AmIWin() int {
	current := strings.Split(s.lottery.current, ",")
	currentBlue := append(current[len(current)-2:])
	current = append(current[:len(current)-2])
	correct := strings.Split(s.lottery.correct, ",")
	correctBlue := append(correct[len(correct)-2:])
	correct = append(correct[:len(correct)-2])
	sameNum := s.lottery.sameNum(current, correct)
	sameBlue := s.lottery.sameNum(currentBlue, correctBlue) //蓝球相同数量
	var level int
	if sameBlue == 0 {
		switch sameNum {
		case 3:
			level = 9
			break
		case 4:
			level = 7
			break
		case 5:
			level = 3
			break
		default:
			level = 0
		}
	} else if sameBlue == 1 {
		switch sameNum {
		case 2:
			level = 9
			break
		case 3:
			level = 8
			break
		case 4:
			level = 5
			break
		case 5:
			level = 2
			break
		default:
			level = 0
		}
	} else if sameNum == 2 {
		switch sameNum {
		case 0:
		case 1:
			level = 9
			break
		case 2:
			level = 8
			break
		case 3:
			level = 6
			break
		case 4:
			level = 4
			break
		case 5:
			level = 1
			break
		default:
			level = 0
		}
	}
	switch level {
	case 1:
		return -1
	case 2:
		return -1
	case 3:
		return 10000
	case 4:
		return 3000
	case 5:
		return 300
	case 6:
		return 200
	case 7:
		return 100
	case 8:
		return 15
	case 9:
		return 5
	default:
		return 0
	}
}
