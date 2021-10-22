package containerregistry

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// PasswordName enumerates the values for password name.
type PasswordName string

const (
	// Password ...
	Password PasswordName = "password"
	// Password2 ...
	Password2 PasswordName = "password2"
)

// PossiblePasswordNameValues returns an array of possible values for the PasswordName const type.
func PossiblePasswordNameValues() []PasswordName {
	return []PasswordName{Password, Password2}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Creating, Succeeded}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// Basic ...
	Basic SkuTier = "Basic"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{Basic}
}
