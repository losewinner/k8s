package maintenance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/preview/maintenance/mgmt/2018-06-01-preview/maintenance"

// ApplyUpdate apply Update request
type ApplyUpdate struct {
	autorest.Response `json:"-"`
	// ApplyUpdateProperties - Properties of the apply update
	*ApplyUpdateProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Fully qualified identifier of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of the resource
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ApplyUpdate.
func (au ApplyUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if au.ApplyUpdateProperties != nil {
		objectMap["properties"] = au.ApplyUpdateProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ApplyUpdate struct.
func (au *ApplyUpdate) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var applyUpdateProperties ApplyUpdateProperties
				err = json.Unmarshal(*v, &applyUpdateProperties)
				if err != nil {
					return err
				}
				au.ApplyUpdateProperties = &applyUpdateProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				au.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				au.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				au.Type = &typeVar
			}
		}
	}

	return nil
}

// ApplyUpdateProperties properties for apply update
type ApplyUpdateProperties struct {
	// Status - The status. Possible values include: 'Pending', 'InProgress', 'Completed', 'RetryNow', 'RetryLater'
	Status UpdateStatus `json:"status,omitempty"`
	// ResourceID - The resourceId
	ResourceID *string `json:"resourceId,omitempty"`
	// LastUpdateTime - Last Update time
	LastUpdateTime *date.Time `json:"lastUpdateTime,omitempty"`
}

// Configuration maintenance configuration record type
type Configuration struct {
	autorest.Response `json:"-"`
	// Location - Gets or sets location of the resource
	Location *string `json:"location,omitempty"`
	// Tags - Gets or sets tags of the resource
	Tags map[string]*string `json:"tags"`
	// ConfigurationProperties - Gets or sets properties of the resource
	*ConfigurationProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Fully qualified identifier of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of the resource
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Configuration.
func (c Configuration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if c.Location != nil {
		objectMap["location"] = c.Location
	}
	if c.Tags != nil {
		objectMap["tags"] = c.Tags
	}
	if c.ConfigurationProperties != nil {
		objectMap["properties"] = c.ConfigurationProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Configuration struct.
func (c *Configuration) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				c.Location = &location
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				c.Tags = tags
			}
		case "properties":
			if v != nil {
				var configurationProperties ConfigurationProperties
				err = json.Unmarshal(*v, &configurationProperties)
				if err != nil {
					return err
				}
				c.ConfigurationProperties = &configurationProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				c.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				c.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				c.Type = &typeVar
			}
		}
	}

	return nil
}

// ConfigurationAssignment configuration Assignment
type ConfigurationAssignment struct {
	autorest.Response `json:"-"`
	// Location - Location of the resource
	Location *string `json:"location,omitempty"`
	// ConfigurationAssignmentProperties - Properties of the configuration assignment
	*ConfigurationAssignmentProperties `json:"properties,omitempty"`
	// ID - READ-ONLY; Fully qualified identifier of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of the resource
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for ConfigurationAssignment.
func (ca ConfigurationAssignment) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if ca.Location != nil {
		objectMap["location"] = ca.Location
	}
	if ca.ConfigurationAssignmentProperties != nil {
		objectMap["properties"] = ca.ConfigurationAssignmentProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for ConfigurationAssignment struct.
func (ca *ConfigurationAssignment) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				ca.Location = &location
			}
		case "properties":
			if v != nil {
				var configurationAssignmentProperties ConfigurationAssignmentProperties
				err = json.Unmarshal(*v, &configurationAssignmentProperties)
				if err != nil {
					return err
				}
				ca.ConfigurationAssignmentProperties = &configurationAssignmentProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				ca.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				ca.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				ca.Type = &typeVar
			}
		}
	}

	return nil
}

// ConfigurationAssignmentProperties properties for configuration assignment
type ConfigurationAssignmentProperties struct {
	// MaintenanceConfigurationID - The maintenance configuration Id
	MaintenanceConfigurationID *string `json:"maintenanceConfigurationId,omitempty"`
	// ResourceID - The unique resourceId
	ResourceID *string `json:"resourceId,omitempty"`
}

// ConfigurationProperties properties for maintenance configuration
type ConfigurationProperties struct {
	// Namespace - Gets or sets namespace of the resource
	Namespace *string `json:"namespace,omitempty"`
	// ExtensionProperties - Gets or sets extensionProperties of the maintenanceConfiguration
	ExtensionProperties map[string]*string `json:"extensionProperties"`
	// MaintenanceScope - Gets or sets maintenanceScope of the configuration. Possible values include: 'ScopeAll', 'ScopeHost', 'ScopeResource', 'ScopeInResource'
	MaintenanceScope Scope `json:"maintenanceScope,omitempty"`
}

// MarshalJSON is the custom marshaler for ConfigurationProperties.
func (cp ConfigurationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cp.Namespace != nil {
		objectMap["namespace"] = cp.Namespace
	}
	if cp.ExtensionProperties != nil {
		objectMap["extensionProperties"] = cp.ExtensionProperties
	}
	if cp.MaintenanceScope != "" {
		objectMap["maintenanceScope"] = cp.MaintenanceScope
	}
	return json.Marshal(objectMap)
}

// Error an error response received from the Azure Maintenance service.
type Error struct {
	// Error - Details of the error
	Error *ErrorDetails `json:"error,omitempty"`
}

// ErrorDetails an error response details received from the Azure Maintenance service.
type ErrorDetails struct {
	// Code - Service-defined error code. This code serves as a sub-status for the HTTP error code specified in the response.
	Code *string `json:"code,omitempty"`
	// Message - Human-readable representation of the error.
	Message *string `json:"message,omitempty"`
}

// ListConfigurationAssignmentsResult response for ConfigurationAssignments list
type ListConfigurationAssignmentsResult struct {
	autorest.Response `json:"-"`
	// Value - The list of configuration Assignments
	Value *[]ConfigurationAssignment `json:"value,omitempty"`
}

// ListMaintenanceConfigurationsResult response for MaintenanceConfigurations list
type ListMaintenanceConfigurationsResult struct {
	autorest.Response `json:"-"`
	// Value - The list of maintenance Configurations
	Value *[]Configuration `json:"value,omitempty"`
}

// ListUpdatesResult response for Updates list
type ListUpdatesResult struct {
	autorest.Response `json:"-"`
	// Value - The pending updates
	Value *[]Update `json:"value,omitempty"`
}

// Operation represents an operation returned by the GetOperations request
type Operation struct {
	// Name - Name of the operation
	Name *string `json:"name,omitempty"`
	// Display - Display name of the operation
	Display *OperationInfo `json:"display,omitempty"`
	// Origin - Origin of the operation
	Origin *string `json:"origin,omitempty"`
	// Properties - Properties of the operation
	Properties interface{} `json:"properties,omitempty"`
}

// OperationInfo information about an operation
type OperationInfo struct {
	// Provider - Name of the provider
	Provider *string `json:"provider,omitempty"`
	// Resource - Name of the resource type
	Resource *string `json:"resource,omitempty"`
	// Operation - Name of the operation
	Operation *string `json:"operation,omitempty"`
	// Description - Description of the operation
	Description *string `json:"description,omitempty"`
}

// OperationsListResult result of the List Operations operation
type OperationsListResult struct {
	autorest.Response `json:"-"`
	// Value - A collection of operations
	Value *[]Operation `json:"value,omitempty"`
}

// Resource definition of a Resource
type Resource struct {
	// ID - READ-ONLY; Fully qualified identifier of the resource
	ID *string `json:"id,omitempty"`
	// Name - READ-ONLY; Name of the resource
	Name *string `json:"name,omitempty"`
	// Type - READ-ONLY; Type of the resource
	Type *string `json:"type,omitempty"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	return json.Marshal(objectMap)
}

// Update maintenance update on a resource
type Update struct {
	// MaintenanceScope - The impact area. Possible values include: 'ScopeAll', 'ScopeHost', 'ScopeResource', 'ScopeInResource'
	MaintenanceScope Scope `json:"maintenanceScope,omitempty"`
	// ImpactType - The impact type. Possible values include: 'None', 'Freeze', 'Restart', 'Redeploy'
	ImpactType ImpactType `json:"impactType,omitempty"`
	// Status - The status. Possible values include: 'Pending', 'InProgress', 'Completed', 'RetryNow', 'RetryLater'
	Status UpdateStatus `json:"status,omitempty"`
	// ImpactDurationInSec - Duration of impact in seconds
	ImpactDurationInSec *int32 `json:"impactDurationInSec,omitempty"`
	// NotBefore - Time when Azure will start force updates if not self-updated by customer before this time
	NotBefore *date.Time `json:"notBefore,omitempty"`
	// UpdateProperties - Properties of the apply update
	*UpdateProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for Update.
func (u Update) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if u.MaintenanceScope != "" {
		objectMap["maintenanceScope"] = u.MaintenanceScope
	}
	if u.ImpactType != "" {
		objectMap["impactType"] = u.ImpactType
	}
	if u.Status != "" {
		objectMap["status"] = u.Status
	}
	if u.ImpactDurationInSec != nil {
		objectMap["impactDurationInSec"] = u.ImpactDurationInSec
	}
	if u.NotBefore != nil {
		objectMap["notBefore"] = u.NotBefore
	}
	if u.UpdateProperties != nil {
		objectMap["properties"] = u.UpdateProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for Update struct.
func (u *Update) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "maintenanceScope":
			if v != nil {
				var maintenanceScope Scope
				err = json.Unmarshal(*v, &maintenanceScope)
				if err != nil {
					return err
				}
				u.MaintenanceScope = maintenanceScope
			}
		case "impactType":
			if v != nil {
				var impactType ImpactType
				err = json.Unmarshal(*v, &impactType)
				if err != nil {
					return err
				}
				u.ImpactType = impactType
			}
		case "status":
			if v != nil {
				var status UpdateStatus
				err = json.Unmarshal(*v, &status)
				if err != nil {
					return err
				}
				u.Status = status
			}
		case "impactDurationInSec":
			if v != nil {
				var impactDurationInSec int32
				err = json.Unmarshal(*v, &impactDurationInSec)
				if err != nil {
					return err
				}
				u.ImpactDurationInSec = &impactDurationInSec
			}
		case "notBefore":
			if v != nil {
				var notBefore date.Time
				err = json.Unmarshal(*v, &notBefore)
				if err != nil {
					return err
				}
				u.NotBefore = &notBefore
			}
		case "properties":
			if v != nil {
				var updateProperties UpdateProperties
				err = json.Unmarshal(*v, &updateProperties)
				if err != nil {
					return err
				}
				u.UpdateProperties = &updateProperties
			}
		}
	}

	return nil
}

// UpdateProperties properties for update
type UpdateProperties struct {
	// ResourceID - The resourceId
	ResourceID *string `json:"resourceId,omitempty"`
}
