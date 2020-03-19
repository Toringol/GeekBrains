package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

/*
* You will be supplied with two data files in CSV format .
* The first file contains statistics about various dinosaurs. The second file contains additional data.
* Given the following formula, speed = ((STRIDE_LENGTH / LEG_LENGTH) - 1) * SQRT(LEG_LENGTH * g)
* Where g = 9.8 m/s^2 (gravitational constant)
*
* Write a program to read in the data files from disk, it must then print the names of only the bipedal dinosaurs from fastest to slowest.
* Do not print any other information.
*
* $ cat dataset1.csv
* NAME,LEG_LENGTH,DIET
* Hadrosaurus,1.4,herbivore
* Struthiomimus,0.72,omnivore
* Velociraptor,1.8,carnivore
* Deinonychus,0.47,herbivore
* Euoplocephalus,2.6,herbivore
* Stegosaurus,1.50,herbivore
* Tyrannosaurus Rex,6.5,carnivore
*
* $ cat dataset2.csv
* NAME,STRIDE_LENGTH,STANCE
* Euoplocephalus,1.97,quadrupedal
* Stegosaurus,1.70,quadrupedal
* Tyrannosaurus Rex,4.76,bipedal
* Hadrosaurus,1.3,bipedal
* Deinonychus,1.11,bipedal
* Struthiomimus,1.24,bipedal
* Velociraptor,2.62,bipedal
 */

/*
* Идея:
* Считываем из второго файла всех bipedal динозавров и
* записываем в скорость пока длину ноги пока не знаем всех остальных данных,
* Потом открываем первый файл и ищем по именам информацию, которую мы получили
* из второго файла, когда находим соответствие пересчитываем скорость,
* когда считали первый файл, сортируем по скорости динозавров и выводим топ.
 */

const (
	g = 9.8
)

type BipedalDinosaur struct {
	Name  string
	Speed float64
}

type DinosaursTop []BipedalDinosaur

func (d DinosaursTop) Len() int {
	return len(d)
}

func (d DinosaursTop) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d DinosaursTop) Less(i, j int) bool {
	return d[i].Speed > d[j].Speed
}

func (d DinosaursTop) String() string {
	topNames := ""
	for iter, item := range d {
		topNames += strconv.Itoa(iter+1) + " - " + item.Name
		if iter != len(d)-1 {
			topNames += "\n"
		}
	}
	return topNames
}

func (d *DinosaursTop) FindBipedalDinosaur(pathFile string) {
	csvFile, err := os.Open(pathFile)
	if err != nil {
		log.Fatalf("Error openning file")
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	bipedalDinosaur := &BipedalDinosaur{}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading file")
		}

		if record[2] == "bipedal" {
			bipedalDinosaur.Name = record[0]
			bipedalDinosaur.Speed, err = strconv.ParseFloat(record[1], 64)
			if err != nil {
				log.Fatalf("Error convert string to float")
			}

			*d = append(*d, *bipedalDinosaur)
		}
	}
}

func (d *DinosaursTop) CalculateSpeed(pathFile string) {
	csvFile, err := os.Open(pathFile)
	if err != nil {
		log.Fatalf("Error openning file")
	}
	defer csvFile.Close()

	r := csv.NewReader(csvFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading file")
		}

		for i := 0; i < len(*d); i++ {

			if record[0] == (*d)[i].Name {
				legLength, err := strconv.ParseFloat(record[1], 64)
				if err != nil {
					log.Fatalf("Error convert string to float")
				}

				(*d)[i].Speed = (((*d)[i].Speed / legLength) - 1) * math.Sqrt(legLength*g)
			}

		}
	}
}

func main() {
	dinosaursTop := &DinosaursTop{}
	dinosaursTop.FindBipedalDinosaur("dataset2.csv")
	dinosaursTop.CalculateSpeed("dataset1.csv")
	sort.Sort(dinosaursTop)
	fmt.Println(dinosaursTop)
}
