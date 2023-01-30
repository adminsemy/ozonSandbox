package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
)

//Характеристики процессоров - количество затрат энергии и
//время, когда освободится процессор
type processor struct {
	Energy   uint64
	FreeTime uint64
}

//Занятые процессоры
var busyProcessors []*processor

//Рабочие процессоры
var processors []*processor

//Ближайшее время освобождения процессора
var timeFreeProcessor uint64

//Есть ли свободные процессоры
var freeProcessor bool = true

//Индекс следующего свободного процессора
var freeIndexProcessor uint64

//Добавляем в занятые процессоры путем быстрой вставки
func appendFastToBusyProcessors(proc *processor) error {
	if len(busyProcessors) == 0 {
		busyProcessors = append(busyProcessors, proc)
		return nil
	}
	lastIndex := len(busyProcessors) - 1
	lastElement := busyProcessors[lastIndex]
	if lastElement.FreeTime < proc.FreeTime {
		busyProcessors = append(busyProcessors, proc)
		return nil
	}
	firstElement := busyProcessors[0]
	if firstElement.FreeTime > proc.FreeTime {
		busyProcessors = append([]*processor{proc}, busyProcessors...)
		return nil
	}
	return errors.New("не удалось добавить простым добавлением")
}

//Вставляем по алгоритму бинарного дерева
func appendToBusyProcessors(p *processor) {
	var iStart uint64 = 1
	var iEnd uint64 = uint64(len(busyProcessors) - 2)
	var iCenter uint64
	for {
		iCenter = (iEnd-iStart)%2 + iStart
		if iCenter+1 == iEnd {
			start := busyProcessors[:iCenter]
			start = append(start, p)
			busyProcessors = append(start, busyProcessors[iCenter:]...)
			break
		}
		center := busyProcessors[iCenter]
		if center.FreeTime > p.FreeTime {
			iStart = iCenter
			continue
		}
		if center.FreeTime < p.FreeTime {
			iEnd = iCenter
			continue
		}
		if center.FreeTime == p.FreeTime {
			if center.Energy > p.Energy {
				iEnd = iCenter
				continue
			}
			if center.Energy < p.Energy {
				iStart = iCenter
				continue
			}
			if center.Energy == p.Energy {
				start := busyProcessors[:iCenter]
				start = append(start, p)
				busyProcessors = append(start, busyProcessors[iCenter:]...)
				break
			}
		}
	}
}

func freeInBusyProcessors(timeIn, duration uint64) (uint64, error) {
	if len(busyProcessors) == 0 {
		return 0, errors.New("нет элементов")
	}
	firstElement := busyProcessors[0]
	if firstElement.FreeTime > timeIn {
		return 0, errors.New("нет свободных процессоров")
	}
	busyProcessors = busyProcessors[1:]
	firstElement.FreeTime = timeIn + duration
	err := appendFastToBusyProcessors(firstElement)
	if err != nil {
		appendToBusyProcessors(firstElement)
	}
	return firstElement.Energy * duration, nil
}

func generalProcessorsTime(timeIn, duration uint64) uint64 {
	if timeFreeProcessor > timeIn && freeProcessor == false {
		return 0
	}
	for _, p := range processors {
		if p.FreeTime <= timeIn {
			freeProcessor = true
			p.FreeTime = duration + timeIn
			return p.Energy * duration
		}
		if timeFreeProcessor > p.FreeTime || timeFreeProcessor == 0 {
			timeFreeProcessor = p.FreeTime
		}
	}
	freeProcessor = false
	return 0
}

//Сканирование данных из консоли
func scanNumbers() uint64 {
	input := bufio.NewReader(os.Stdin)
	var countProcessors int
	var tasks int
	var result uint64

	fmt.Fscan(input, &countProcessors, &tasks)

	processors = make([]*processor, countProcessors)

	for i := 0; i < countProcessors; i++ {
		var time uint64
		fmt.Fscan(input, &time)
		proc := &processor{Energy: time, FreeTime: 0}
		processors[i] = proc
	}

	sort.Slice(processors, func(i, j int) bool {
		return processors[i].Energy < processors[j].Energy
	})
	for i := 0; i < tasks; i++ {
		var timeIn uint64   //количество строк
		var duration uint64 //количество символов

		fmt.Fscan(input, &timeIn, &duration)

		result += generalProcessorsTime(timeIn, duration)
	}
	return result
}

func main() {
	result := scanNumbers()
	fmt.Println(result)
	//test()
}
