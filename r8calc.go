package main

import (
	"log"
	"math"
	"os"
	"strings"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

type trainCar struct {
	model       string
	emptyWeight float64
	cargoWeight float64
	isEngine    bool
}

type train struct {
	cars []trainCar
}

func (t *train) engineGroups() [][]trainCar {
	var currentGroup []trainCar
	var engineGroups [][]trainCar

	currentGroup = make([]trainCar, 0)
	engineGroups = make([][]trainCar, 0)

	for _, car := range t.cars {
		if car.isEngine {
			currentGroup = append(currentGroup, car)
		} else if len(currentGroup) > 0 {
			engineGroups = append(engineGroups, currentGroup)
			currentGroup = make([]trainCar, 0)
		}
	}

	if len(currentGroup) > 0 {
		engineGroups = append(engineGroups, currentGroup)
	}
	return engineGroups
}

func (t *train) tonnage() int {
	var tonnage float64
	for _, car := range t.cars {
		tonnage += car.cargoWeight + car.emptyWeight
	}
	return int(math.Ceil(tonnage))
}

func setCell(sheet *ole.IDispatch, row int, col int, val int) {
	cell := oleutil.MustGetProperty(sheet, "Cells", row, col).ToIDispatch()
	oleutil.PutProperty(cell, "Value", val)
	cell.Release()
}

func getCell(sheet *ole.IDispatch, row int, col int) int64 {
	cell := oleutil.MustGetProperty(sheet, "Cells", row, col).ToIDispatch()
	val := oleutil.MustGetProperty(cell, "Value")
	cell.Release()
	return val.Val
}

func incrementCell(sheet *ole.IDispatch, row int, col int) {
	cell := oleutil.MustGetProperty(sheet, "Cells", row, col).ToIDispatch()
	val := oleutil.MustGetProperty(cell, "Value")
	oleutil.PutProperty(cell, "Value", int(val.Value().(float64))+1)
	cell.Release()
}

func loadSheet(filename string, train train) error {
	ole.CoInitialize(0)
	oleobj, err := oleutil.CreateObject("Excel.Application")
	if err != nil {
		return err
	}

	excel, err := oleobj.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return err
	}
	defer excel.Release()

	workbooks_, err := oleutil.GetProperty(excel, "Workbooks")
	if err != nil {
		return err
	}
	workbooks := workbooks_.ToIDispatch()
	defer workbooks.Release()

	workbook_, err := oleutil.CallMethod(workbooks, "Open", filename)
	if err != nil {
		return err
	}
	workbook := workbook_.ToIDispatch()
	defer workbook.Release()

	bnsfSheet_, err := oleutil.GetProperty(workbook, "Worksheets", 1)
	if err != nil {
		return err
	}
	bnsfSheet := bnsfSheet_.ToIDispatch()
	defer bnsfSheet.Release()

	upSheet_, err := oleutil.GetProperty(workbook, "Worksheets", 2)
	if err != nil {
		return err
	}
	upSheet := upSheet_.ToIDispatch()
	defer upSheet.Release()

	tonnage := train.tonnage()

	setCell(bnsfSheet, 3, 7, tonnage)
	setCell(upSheet, 3, 7, tonnage)

	var headEngines, midEngines, rearEngines []trainCar

	engineGroups := train.engineGroups()
	headEngines = engineGroups[0]

	if len(engineGroups) == 2 {
		if train.cars[len(train.cars)-1].isEngine {
			midEngines = make([]trainCar, 0)
			rearEngines = engineGroups[1]
		} else {
			midEngines = engineGroups[1]
			rearEngines = make([]trainCar, 0)
		}
	} else if len(engineGroups) == 3 {
		midEngines = engineGroups[1]
		rearEngines = engineGroups[2]
	}

	for _, engine := range headEngines {
		code := strings.Split(engine.model, "_")[1]
		incrementCell(bnsfSheet, bnsfCells(code)[0], 2)
		incrementCell(bnsfSheet, bnsfCells(code)[3], 2)
		incrementCell(upSheet, upCells(code)[0], 2)
		incrementCell(upSheet, upCells(code)[3], 2)
	}

	for _, engine := range midEngines {
		code := strings.Split(engine.model, "_")[1]
		incrementCell(bnsfSheet, bnsfCells(code)[1], 2)
		incrementCell(bnsfSheet, bnsfCells(code)[4], 2)
		incrementCell(upSheet, upCells(code)[1], 2)
		incrementCell(upSheet, upCells(code)[4], 2)
	}

	for _, engine := range rearEngines {
		code := strings.Split(engine.model, "_")[1]
		incrementCell(bnsfSheet, bnsfCells(code)[2], 2)
		incrementCell(bnsfSheet, bnsfCells(code)[5], 2)
		incrementCell(upSheet, upCells(code)[2], 2)
		incrementCell(upSheet, upCells(code)[5], 2)
	}

	oleutil.PutProperty(excel, "Visible", true)
	ole.CoUninitialize()

	return nil
}

func main() {
	var train train
	var cargoWeight, emptyWeight float64
	var isEngine bool

	if len(os.Args) != 3 {
		log.Fatalf("%s <xml> <xsl>", os.Args[0])
	}

	cars := parseTrain(os.Args[1])

	for _, car := range cars {
		carFile := car.FileName
		carType := carFile[:strings.IndexByte(carFile, '.')]
		emptyWeight = carWeight(carType)
		cargoWeight = car.Weight
		if car.UnitType == "US_DieselEngine" {
			isEngine = true
		} else {
			isEngine = false
		}

		train.cars = append(train.cars, trainCar{carType, emptyWeight, cargoWeight, isEngine})
	}

	cwd, _ := os.Getwd()
	err := loadSheet(cwd+"\\"+os.Args[2], train)
	if err != nil {
		log.Fatal(err)
	}
}
