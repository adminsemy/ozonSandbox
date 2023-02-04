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

//Добавляем процессор в занятые
func appendToBusyProcessors(p *processor) {
	err := appendFastToBusyProcessors(p)
	if err != nil {
		appendBinaryToBusyProcessors(p)
	}
}

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
func appendBinaryToBusyProcessors(p *processor) {
	var iStart = 0
	var iEnd = (len(busyProcessors) - 1)
	var iCenter int
	for {
		if iStart+1 == iEnd {
			busyProcessors = append(busyProcessors[:iEnd], busyProcessors[iStart:]...)
			busyProcessors[iEnd] = p
			break
		}
		iCenter = (iEnd-iStart)/2 + iStart
		center := busyProcessors[iCenter]
		if center.FreeTime < p.FreeTime {
			iStart = iCenter
			continue
		}
		if center.FreeTime > p.FreeTime {
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
				busyProcessors = append(busyProcessors[:iEnd], busyProcessors[iStart:]...)
				busyProcessors[iEnd] = p
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
	var index int = 0
	for i, b := range busyProcessors {
		if firstElement.Energy > b.Energy {
			firstElement = b
			index = i
		}
		if len(busyProcessors)-1 == i ||
			busyProcessors[i+1].FreeTime > timeIn {
			busyProcessors = append(busyProcessors[:index], busyProcessors[index+1:]...)
			break
		}
	}
	firstElement.FreeTime = timeIn + duration
	appendToBusyProcessors(firstElement)
	return firstElement.Energy * duration, nil
}

func generalProcessorsTime(timeIn, duration uint64) uint64 {
	res, err := freeInBusyProcessors(timeIn, duration)
	if err == nil {
		return res
	}
	if len(processors) == 0 {
		return 0
	}
	proc := processors[0]
	proc.FreeTime = duration + timeIn
	appendToBusyProcessors(proc)
	processors = processors[1:]
	return proc.Energy * duration
}

//Сканирование данных из консоли
func scanNumbers() uint64 {
	testFile, err := os.Open("./tests/06")
	if err != nil {
		fmt.Println("Not found file", err)
		return 0
	}
	defer testFile.Close()
	input := bufio.NewReader(testFile)
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
