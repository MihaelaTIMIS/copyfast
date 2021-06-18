package tuto

import "fmt"

// nil je n'ai pas de valeur ni de pointer
var leNomDeMaVariablePrivateNil string
var LeNomDeMaVariablePublicNil string

var leNomDeMaVariableTypePrivateValue string = "VersionType"
var LeNomDeMaVariableTypePublicValue string = "VersionType"

var leNomDeMaVariableNonTypePrivateValue = "VersionNonType"
var LeNomDeMaVariableNonTypePublicValue = "VersionNonType"

// Constante value obligatoire
const maConstantePriveTypePrivateValue string = "VersionType"
const MaConstantePriveTypePublicValue string = "VersionType"

const maConstantePriveNonTypePrivateValue = "VersionNonType"
const MaConstantePriveNonTypePublicValue = "VersionNonType"

func sampleFunction() {

	// nil je n'ai pas de valeur ni de pointer
	var leNomDeMaVariablePrivateNil string
	var LeNomDeMaVariablePublicNil string

	fmt.Println(leNomDeMaVariablePrivateNil, LeNomDeMaVariablePublicNil)

	var leNomDeMaVariableTypePrivateValue string = "VersionType"
	var LeNomDeMaVariableTypePublicValue string = "VersionType"

	fmt.Println(leNomDeMaVariableTypePrivateValue, LeNomDeMaVariableTypePublicValue)

	var leNomDeMaVariableNonTypePrivateValue = "VersionNonType"
	var LeNomDeMaVariableNonTypePublicValue = "VersionNonType"

	fmt.Println(leNomDeMaVariableNonTypePrivateValue, LeNomDeMaVariableNonTypePublicValue)

	// Constante value obligatoire
	const maConstantePriveTypePrivateValue string = "VersionType"
	const MaConstantePriveTypePublicValue string = "VersionType"

	const maConstantePriveNonTypePrivateValue = "VersionNonType"
	const MaConstantePriveNonTypePublicValue = "VersionNonType"

	variableLiterral := "Literral"
	fmt.Println(variableLiterral)
}
