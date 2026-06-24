package annalyn


// CanFastAttack can be executed only when the knight is sleeping.
func CanFastAttack(knightIsAwake bool) bool {
    if knightIsAwake != true {
return true
    } 
    return false
}

// CanSpy can be executed if at least one of the characters is awake.
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
       if knightIsAwake != true && archerIsAwake != true && prisonerIsAwake != true {
           return false
    } else if knightIsAwake == true && archerIsAwake == true && prisonerIsAwake == true {
		return true
    } else if knightIsAwake == true || (archerIsAwake != true && prisonerIsAwake != true) {
        return true     
    } else  if archerIsAwake == true || (knightIsAwake != true && prisonerIsAwake != true) {
        return true
    } else   if prisonerIsAwake == true || (knightIsAwake != true && archerIsAwake != true) {
        return true
    } 
    return false
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
    if archerIsAwake != true && prisonerIsAwake != true {
        return false
    } else if archerIsAwake == true && prisonerIsAwake != true {
        return false
    }  else if archerIsAwake == true && prisonerIsAwake == true {
        return false 
        }
    return true
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping.
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
    if prisonerIsAwake == true && (knightIsAwake != true && archerIsAwake != true) {
        return true
        } 
    if petDogIsPresent == true && archerIsAwake != true {
        return true
        } 
    return false
}
