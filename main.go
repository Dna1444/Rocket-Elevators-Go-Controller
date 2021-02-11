package main

import (
	"fmt"
	"math"
	"sort"
)

func senario1() {
	battery := Battery{
		ID:                         1,
		status:                     "online",
		amountOfColumns:            4,
		amountOfFloors:             60,
		amountOfBasements:          6,
		amountOfElevatorsPerColumn: 5,
		columnsList:                []Column{},
		FloorRequestList:           []FloorRequestButton{},
	}
	battery.createFloorRequestbutton()
	battery.createColumns()

	battery.columnsList[1].elevatorsList[0].currentFloor = 20
	battery.columnsList[1].elevatorsList[0].status = "moving"
	battery.columnsList[1].elevatorsList[0].direction = "down"
	battery.columnsList[1].elevatorsList[0].floorRequestList = append(battery.columnsList[1].elevatorsList[0].floorRequestList, 5)
	battery.columnsList[1].elevatorsList[1].currentFloor = 3
	battery.columnsList[1].elevatorsList[1].status = "moving"
	battery.columnsList[1].elevatorsList[1].direction = "up"
	battery.columnsList[1].elevatorsList[1].floorRequestList = append(battery.columnsList[1].elevatorsList[1].floorRequestList, 15)
	battery.columnsList[1].elevatorsList[2].currentFloor = 13
	battery.columnsList[1].elevatorsList[2].status = "moving"
	battery.columnsList[1].elevatorsList[2].direction = "down"
	battery.columnsList[1].elevatorsList[2].floorRequestList = append(battery.columnsList[1].elevatorsList[2].floorRequestList, 1)
	battery.columnsList[1].elevatorsList[3].currentFloor = 15
	battery.columnsList[1].elevatorsList[3].status = "moving"
	battery.columnsList[1].elevatorsList[3].direction = "down"
	battery.columnsList[1].elevatorsList[3].floorRequestList = append(battery.columnsList[1].elevatorsList[3].floorRequestList, 2)
	battery.columnsList[1].elevatorsList[4].currentFloor = 6
	battery.columnsList[1].elevatorsList[4].status = "moving"
	battery.columnsList[1].elevatorsList[4].direction = "down"
	battery.columnsList[1].elevatorsList[4].floorRequestList = append(battery.columnsList[1].elevatorsList[4].floorRequestList, 1)

	battery.assignElevator(20, "up")

}

func senario2() {
	battery := Battery{
		ID:                         1,
		status:                     "online",
		amountOfColumns:            4,
		amountOfFloors:             60,
		amountOfBasements:          6,
		amountOfElevatorsPerColumn: 5,
		columnsList:                []Column{},
		FloorRequestList:           []FloorRequestButton{},
	}
	battery.createFloorRequestbutton()
	battery.createColumns()

	battery.columnsList[2].elevatorsList[0].currentFloor = 1
	battery.columnsList[2].elevatorsList[0].status = "idle"
	battery.columnsList[2].elevatorsList[0].direction = "up"
	battery.columnsList[2].elevatorsList[0].floorRequestList = append(battery.columnsList[2].elevatorsList[0].floorRequestList, 21)
	battery.columnsList[2].elevatorsList[1].currentFloor = 21
	battery.columnsList[2].elevatorsList[1].status = "moving"
	battery.columnsList[2].elevatorsList[1].direction = "up"
	battery.columnsList[2].elevatorsList[1].floorRequestList = append(battery.columnsList[2].elevatorsList[1].floorRequestList, 28)
	battery.columnsList[2].elevatorsList[2].currentFloor = 33
	battery.columnsList[2].elevatorsList[2].status = "moving"
	battery.columnsList[2].elevatorsList[2].direction = "down"
	battery.columnsList[2].elevatorsList[2].floorRequestList = append(battery.columnsList[2].elevatorsList[2].floorRequestList, 1)
	battery.columnsList[2].elevatorsList[3].currentFloor = 40
	battery.columnsList[2].elevatorsList[3].status = "moving"
	battery.columnsList[2].elevatorsList[3].direction = "down"
	battery.columnsList[2].elevatorsList[3].floorRequestList = append(battery.columnsList[2].elevatorsList[3].floorRequestList, 24)
	battery.columnsList[2].elevatorsList[4].currentFloor = 39
	battery.columnsList[2].elevatorsList[4].status = "moving"
	battery.columnsList[2].elevatorsList[4].direction = "down"
	battery.columnsList[2].elevatorsList[4].floorRequestList = append(battery.columnsList[2].elevatorsList[4].floorRequestList, 1)

	battery.assignElevator(36, "up")

}

func senario3() {
	battery := Battery{
		ID:                         1,
		status:                     "online",
		amountOfColumns:            4,
		amountOfFloors:             60,
		amountOfBasements:          6,
		amountOfElevatorsPerColumn: 5,
		columnsList:                []Column{},
		FloorRequestList:           []FloorRequestButton{},
	}
	battery.createFloorRequestbutton()
	battery.createColumns()

	battery.columnsList[3].elevatorsList[0].currentFloor = 58
	battery.columnsList[3].elevatorsList[0].status = "move"
	battery.columnsList[3].elevatorsList[0].direction = "down"
	battery.columnsList[3].elevatorsList[0].floorRequestList = append(battery.columnsList[3].elevatorsList[0].floorRequestList, 1)
	battery.columnsList[3].elevatorsList[1].currentFloor = 50
	battery.columnsList[3].elevatorsList[1].status = "moving"
	battery.columnsList[3].elevatorsList[1].direction = "up"
	battery.columnsList[3].elevatorsList[1].floorRequestList = append(battery.columnsList[3].elevatorsList[1].floorRequestList, 60)
	battery.columnsList[3].elevatorsList[2].currentFloor = 46
	battery.columnsList[3].elevatorsList[2].status = "moving"
	battery.columnsList[3].elevatorsList[2].direction = "up"
	battery.columnsList[3].elevatorsList[2].floorRequestList = append(battery.columnsList[3].elevatorsList[2].floorRequestList, 58)
	battery.columnsList[3].elevatorsList[3].currentFloor = 1
	battery.columnsList[3].elevatorsList[3].status = "moving"
	battery.columnsList[3].elevatorsList[3].direction = "up"
	battery.columnsList[3].elevatorsList[3].floorRequestList = append(battery.columnsList[3].elevatorsList[3].floorRequestList, 54)
	battery.columnsList[3].elevatorsList[4].currentFloor = 60
	battery.columnsList[3].elevatorsList[4].status = "moving"
	battery.columnsList[3].elevatorsList[4].direction = "down"
	battery.columnsList[3].elevatorsList[4].floorRequestList = append(battery.columnsList[3].elevatorsList[4].floorRequestList, 1)

	battery.columnsList[3].requestElevator(51, "down")

}

func senario4() {
	battery := Battery{
		ID:                         1,
		status:                     "online",
		amountOfColumns:            4,
		amountOfFloors:             60,
		amountOfBasements:          6,
		amountOfElevatorsPerColumn: 5,
		columnsList:                []Column{},
		FloorRequestList:           []FloorRequestButton{},
	}
	battery.createFloorRequestbutton()
	battery.createColumns()

	battery.columnsList[0].elevatorsList[0].currentFloor = -4
	battery.columnsList[0].elevatorsList[0].status = "idle"
	battery.columnsList[0].elevatorsList[1].currentFloor = 1
	battery.columnsList[0].elevatorsList[1].status = "idle"
	battery.columnsList[0].elevatorsList[2].currentFloor = -3
	battery.columnsList[0].elevatorsList[2].status = "moving"
	battery.columnsList[0].elevatorsList[2].direction = "down"
	battery.columnsList[0].elevatorsList[2].floorRequestList = append(battery.columnsList[0].elevatorsList[2].floorRequestList, -5)
	battery.columnsList[0].elevatorsList[3].currentFloor = -6
	battery.columnsList[0].elevatorsList[3].status = "moving"
	battery.columnsList[0].elevatorsList[3].direction = "up"
	battery.columnsList[0].elevatorsList[3].floorRequestList = append(battery.columnsList[0].elevatorsList[3].floorRequestList, 1)
	battery.columnsList[0].elevatorsList[4].currentFloor = -1
	battery.columnsList[0].elevatorsList[4].status = "moving"
	battery.columnsList[0].elevatorsList[4].direction = "down"
	battery.columnsList[0].elevatorsList[4].floorRequestList = append(battery.columnsList[0].elevatorsList[4].floorRequestList, -6)

	battery.columnsList[0].requestElevator(-3, "up")

}

func main() {

	senario4()

}

// Battery struct
type Battery struct {
	ID                         int
	status                     string
	amountOfColumns            int
	amountOfFloors             int
	amountOfBasements          int
	amountOfElevatorsPerColumn int
	columnsList                []Column
	FloorRequestList           []FloorRequestButton
}

// Column struct
type Column struct {
	ID                int
	status            string
	amountOfFloors    int
	amountOfElevators int
	servedFloors      []int
	isBasement        bool
	elevatorsList     []Elevator
	callButtonsList   []CallButton
}

// Elevator struct
type Elevator struct {
	ID               int
	status           string
	amountOfFloors   int
	currentFloor     int
	direction        string
	floorRequestList []int
	door             Door
}

// Door struct
type Door struct {
	ID     int
	status string
}

// CallButton struct
type CallButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

// FloorRequestButton struct
type FloorRequestButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

// making my columns
func (b *Battery) createColumns() {
	columnID := 1
	floor := 1

	// BASEMENT
	// Check to see if there is a basement in the building
	if b.amountOfBasements > 0 {
		// What floors are being serve?
		servedFloors := []int{}
		b.amountOfColumns--

		for i := 0; i < b.amountOfBasements; i++ {
			//We want the lobby to be floor 1 not 0
			if floor == 1 {
				servedFloors = append(servedFloors, floor)
				floor = -1 //Skip 0 in the next iteration
			} else {
				servedFloors = append(servedFloors, floor)
				floor--
			}
		}

		// Creating a basement column
		// createElevators returns a list of elevators
		// createButtons returns a list of hallway call buttons
		b.columnsList = append(b.columnsList, Column{ID: columnID, status: "online", amountOfFloors: b.amountOfBasements, amountOfElevators: b.amountOfElevatorsPerColumn, servedFloors: servedFloors, isBasement: true, elevatorsList: createElevators(b.amountOfBasements, b.amountOfElevatorsPerColumn), callButtonsList: createCallButtons(b.amountOfBasements, true)})
		columnID++

	}

	// NON-BASEMENT
	// Determine the number of floors per column
	amountOfFloorsPerColumn := int(math.Ceil(float64(b.amountOfFloors / b.amountOfColumns)))

	floor = 1
	for i := 0; i < b.amountOfColumns; i++ {
		// What floors are being serve?
		servedFloors := []int{}

		for j := 0; j < amountOfFloorsPerColumn; j++ {

			if floor <= b.amountOfFloors {
				if floor != 1 {
					servedFloors = append(servedFloors, 1)
				}
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}

		// Creating non-basement columns
		// createElevators returns a list of elevators
		// createButtons returns a list of hallway call buttons
		b.columnsList = append(b.columnsList, Column{ID: columnID, status: "online", amountOfFloors: amountOfFloorsPerColumn, amountOfElevators: b.amountOfElevatorsPerColumn, servedFloors: servedFloors, isBasement: false, elevatorsList: createElevators(amountOfFloorsPerColumn, b.amountOfElevatorsPerColumn), callButtonsList: createCallButtons(amountOfFloorsPerColumn, false)})
		columnID++
	}
}

// maing my floor request button
func (b *Battery) createFloorRequestbutton() {

	floorRequestID := 1
	// making the button for basement
	if b.amountOfBasements > 0 {
		buttonFloor := -1

		for i := 0; i < b.amountOfBasements; i++ {
			b.FloorRequestList = append(b.FloorRequestList, FloorRequestButton{ID: floorRequestID, status: "off", floor: buttonFloor, direction: "down"})
			buttonFloor--
			floorRequestID++

		}

	} else {
		buttonFloor := 1

		for i := 0; i < b.amountOfFloors; i++ {
			b.FloorRequestList = append(b.FloorRequestList, FloorRequestButton{ID: floorRequestID, status: "off", floor: buttonFloor, direction: "up"})
			buttonFloor++
			floorRequestID++

		}

	}

}

// making my elevators
func createElevators(amountOfFloors int, amountOfElevators int) []Elevator {
	elevatorID := 1
	elevatorsList := []Elevator{}

	for i := 0; i < amountOfElevators; i++ {
		elevatorsList = append(elevatorsList, Elevator{ID: elevatorID, status: "idle", amountOfFloors: amountOfFloors, direction: "null", currentFloor: 0, door: Door{elevatorID, "closed"}})

		elevatorID++
	}

	return elevatorsList
}

// making my callbuttons
func createCallButtons(amountOfFloors int, isBasement bool) []CallButton {
	buttonsList := []CallButton{}
	callButtonID := 1
	if isBasement {
		buttonFloor := -1

		for i := 0; i < amountOfFloors; i++ {

			buttonsList = append(buttonsList, CallButton{callButtonID, "off", buttonFloor, "up"})
			buttonFloor--
			callButtonID++
		}
	} else {
		buttonFloor := 1
		for i := 0; i < amountOfFloors; i++ {

			buttonsList = append(buttonsList, CallButton{callButtonID, "off", buttonFloor, "up"})

			buttonFloor++
			callButtonID++
		}
	}

	return buttonsList
}

// find the column that serve that floor
func (b Battery) findBestColumn(requestFloor int) Column {
	bestColumn := Column{}
	for i, column := range b.columnsList {
		if find(requestFloor, b.columnsList[i].servedFloors) {
			bestColumn = column

		}

	}
	return bestColumn
}

// function for contains
func find(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (c Column) findElevator(requestedFloor int, requestedDirection string) Elevator {
	bestElevator := Elevator{}
	bestScore := 5
	referenceGap := 10000000

	if requestedFloor == 1 {

		for _, elevator := range c.elevatorsList {
			fmt.Println(elevator.ID, "checking the id of the elevator")
			//The elevator is at my floor and going in the direction I want
			if requestedFloor == elevator.currentFloor && elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is lower than me, is coming up and I want to go up
			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" {
				bestElevator, bestScore, referenceGap = checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is higher than me, is coming down and I want to go down
			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" {
				bestElevator, bestScore, referenceGap = checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is idle
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is not available, but still could take the call if nothing better is found
			} else {
				bestElevator, bestScore, referenceGap = checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			}
		}

	} else {

		for _, elevator := range c.elevatorsList {
			//The elevator is at my floor and going in the direction I want
			if requestedFloor == elevator.currentFloor && elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = checkIfElevatorIsBetter(1, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is lower than me, is coming up and I want to go up
			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" && requestedDirection == elevator.direction {
				bestElevator, bestScore, referenceGap = checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is higher than me, is coming down and I want to go down
			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == elevator.direction {
				bestElevator, bestScore, referenceGap = checkIfElevatorIsBetter(2, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is idle
			} else if elevator.status == "idle" {
				bestElevator, bestScore, referenceGap = checkIfElevatorIsBetter(3, elevator, bestScore, referenceGap, bestElevator, requestedFloor)

				//The elevator is not available, but still could take the call if nothing better is found
			} else {
				bestElevator, bestScore, referenceGap = checkIfElevatorIsBetter(4, elevator, bestScore, referenceGap, bestElevator, requestedFloor)
			}
		}

	}
	return bestElevator
}

func checkIfElevatorIsBetter(scoreToCheck int, newElevator Elevator, bestScore int, referenceGap int, bestElevator Elevator, requestedFloor int) (Elevator, int, int) {
	if scoreToCheck < bestScore {
		bestScore = scoreToCheck
		bestElevator = newElevator
		referenceGap = int(math.Abs(float64(newElevator.currentFloor - requestedFloor)))
	} else if bestScore == scoreToCheck {
		gap := int(math.Abs(float64(newElevator.currentFloor - requestedFloor)))
		if referenceGap > gap {
			bestElevator = newElevator
			referenceGap = gap
		}
	}

	return bestElevator, bestScore, referenceGap
}

func (e *Elevator) move() {

	for len(e.floorRequestList) != 0 {
		destination := e.floorRequestList[0]
		e.status = "moving"

		if e.currentFloor < destination {
			e.direction = "up"
			for e.currentFloor < destination {
				fmt.Println("Elevator", e.ID, "is currently at floor", e.currentFloor)
				e.currentFloor++
			}
		} else if e.currentFloor > destination {
			e.direction = "down"
			for e.currentFloor > destination {
				fmt.Println("Elevator", e.ID, "is currently at floor", e.currentFloor)
				e.currentFloor--
			}
		}

		e.status = "idle"
		e.direction = "null"
		e.floorRequestList = e.floorRequestList[1:]
		fmt.Println("Elevator", e.ID, "stopped on floor", e.currentFloor)
		e.operateDoors()
	}
}

func (e *Elevator) sortFloorList(checkForLobby string) {

	if checkForLobby == "not" {
		if e.direction == "up" {
			sort.Ints(e.floorRequestList)
		} else {
			sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestList)))
		}
	} else if checkForLobby == "lobby" {
		if e.direction == "down" {
			sort.Ints(e.floorRequestList)
		} else {
			sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestList)))
		}
	}

}

func (c Column) requestElevator(requestedFloor int, direction string) {
	fmt.Println("taking column", c.ID)
	fromLobby := "not"
	// Find the best elevator
	elevator := c.findElevator(requestedFloor, direction)
	// Add floor to floorRequestList
	elevator.floorRequestList = append(elevator.floorRequestList, requestedFloor)
	// Sort the floorRequestList
	elevator.sortFloorList(fromLobby)
	// Make the elevator move to the user
	elevator.move()

}

func (e *Elevator) operateDoors() {
	e.door.status = "opened"
	fmt.Println("Doors opening...")
	fmt.Println("please wait 5")
	e.door.status = "closed"
	fmt.Println("Doors closing...")
}

func (b Battery) assignElevator(requestFloor int, direction string) {
	fromLobby := "yes"
	column := b.findBestColumn(requestFloor)
	fmt.Println("this column id", column.ID)
	elevator := column.findElevator(1, direction)
	fmt.Println("this elevator id", elevator.ID)
	elevator.floorRequestList = append(elevator.floorRequestList, requestFloor)
	elevator.sortFloorList(fromLobby)
	elevator.move()

}
