//AttributeGenerator creates physical and mental attributes for a character
package main

import (
	"PythonToGolangConversion/DiceRoller"
	"fmt"
)

//character structure not currently used
/*type character struct {
	head, torso, abdomen, rarm, larm, rleg, lleg int  // hit points
	str, intel, will, charisma, build, dex int  // physical attributes
	mass, throw, load int  //abilities
}*/

//SetAttrib generates a value for an attribute
func SetAttrib() int  {
	attrib := DiceRoller.MultiDie(2, 2)
	return attrib
}

//MakeAttrib assigns values to attributes
func MakeAttrib() map[string]int {
	attribMap := make(map[string]int)

	//Get values of core attributes
	attribMap["str"] = SetAttrib()  //strength
	attribMap["intel"] = SetAttrib()  //intelligence
	attribMap["will"] = SetAttrib()  //willpower
	attribMap["charisma"] = SetAttrib()
	attribMap["build"] = SetAttrib()
	attribMap["dex"] = SetAttrib()  //dexterity

	//Get values of secondary attributes
	CharBuild, ok := attribMap["build"]
	if ok == true {
		attribMap["mass"] = (CharBuild * 5) + 15 //kilograms
	}

	charStr, ok := attribMap["str"]
	if ok == true {
		attribMap["throw"] = charStr * 2              //meters
		attribMap["load"] = (charStr + CharBuild) * 2 //kilograms
	}
	return attribMap
}

//HitPoints generates the hit points for each body part
func HitPoints(build, bodypart int) int {
	var hp int

	switch bodypart {
	case 1:
		hp = build  //head
	case 2:
		hp = build * 4  //torso
	case 3:
		hp = build * 2  //arms, legs, abdomen
	}
	return hp
}

func MakeHP(build int) map[string]int {
	hpMap := make(map[string]int)

	//Get hit points for body parts
	hpMap["head"] = HitPoints(build, 1)
	hpMap["torso"] = HitPoints(build, 2)
	hpMap["rarm"] = HitPoints(build, 3)
	hpMap["larm"] = HitPoints(build, 3)
	hpMap["rleg"] = HitPoints(build, 3)
	hpMap["lleg"] = HitPoints(build, 3)
	hpMap["abdomen"] = HitPoints(build, 3)

	return hpMap
}

func main() {
	attributes := MakeAttrib()
	globalBuild := attributes["build"]
	hitPoints := MakeHP(globalBuild)
	//fmt.Println(attributes)
	//fmt.Println(hitPoints)
	for key, value := range attributes {
		fmt.Printf("%s:%d\n", key, value)
	}
	for key, value := range hitPoints {
		fmt.Printf("%s:%d\n", key, value)
	}
}
