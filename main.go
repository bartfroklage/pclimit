package main

import (
	"os"
	"log"
	"fmt"
	"time"
	"github.com/gen2brain/beeep"
)

func main() {

	/* Configure logging. */
	f, err := os.OpenFile("pclimit.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	
	/* Start settings. */
	lockedState := false
	currentDayIx := int(time.Now().Weekday())
		
	for {
		ReadSettings()
		timeUsed := ReadTimeUsed()
		timeLeft := settings.DayLimits[currentDayIx].Limit - (timeUsed / 12)

		time.Sleep(5 * time.Second)
		newLockedState := IsWindowsLocked() 

		if (settings.Status == "allow") {
			if lockedState && !newLockedState {
				beeep.Notify("PC-Limit update...", "You're computer is opened up.", "assets/information.png")
				time.Sleep(5 * time.Second)
			}
			lockedState = newLockedState
			continue
		}
		if (settings.Status == "block") {
			if (!newLockedState) {
				beeep.Notify("PC-Limit update...", "You're computer is blocked.", "assets/information.png")
				time.Sleep(5 * time.Second)
			}
			LockWindows()
			lockedState = newLockedState
			continue
		}

		if (lockedState && !newLockedState || (!newLockedState && (timeLeft % 10 == 0) && (timeUsed % 12 == 0))) {
			var message string
			switch {
			case timeLeft < 1:
				message = "Your time is up."
			case timeLeft == 1:
				message = "You have one minute left."
			default:
				message = fmt.Sprintf("You have %d minutes left.", timeLeft)
			}

			title := "PC-Limit update..."
			if (lockedState && !newLockedState) {
				title = "Your time is running again..."
			}
			beeep.Notify(title, message, "assets/information.png")
		} 

		if (!lockedState && timeLeft <= 0) {
			LockWindows()
		}

		if (!lockedState) {
			IncreaseTimeUsed()			
		}		
		lockedState = newLockedState
	}	
}