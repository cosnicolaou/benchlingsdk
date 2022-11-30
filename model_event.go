/*
Benchling API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package benchlingsdk

import (
	"encoding/json"
	"fmt"
)

// Event - struct for Event
type Event struct {
	AssayRunCreatedEvent *AssayRunCreatedEvent
	AssayRunUpdatedFieldsEvent *AssayRunUpdatedFieldsEvent
	AutomationInputGeneratorCompletedV2BetaEvent *AutomationInputGeneratorCompletedV2BetaEvent
	AutomationOutputProcessorCompletedV2BetaEvent *AutomationOutputProcessorCompletedV2BetaEvent
	AutomationOutputProcessorUploadedV2BetaEvent *AutomationOutputProcessorUploadedV2BetaEvent
	EntityRegisteredEvent *EntityRegisteredEvent
	EntryCreatedEvent *EntryCreatedEvent
	EntryUpdatedFieldsEvent *EntryUpdatedFieldsEvent
	EntryUpdatedReviewRecordEvent *EntryUpdatedReviewRecordEvent
	RequestCreatedEvent *RequestCreatedEvent
	RequestUpdatedFieldsEvent *RequestUpdatedFieldsEvent
	StageEntryCreatedEvent *StageEntryCreatedEvent
	StageEntryUpdatedFieldsEvent *StageEntryUpdatedFieldsEvent
	StageEntryUpdatedReviewRecordEvent *StageEntryUpdatedReviewRecordEvent
	WorkflowOutputCreatedEvent *WorkflowOutputCreatedEvent
	WorkflowOutputUpdatedFieldsEvent *WorkflowOutputUpdatedFieldsEvent
	WorkflowTaskCreatedEvent *WorkflowTaskCreatedEvent
	WorkflowTaskGroupCreatedEvent *WorkflowTaskGroupCreatedEvent
	WorkflowTaskGroupUpdatedWatchersEvent *WorkflowTaskGroupUpdatedWatchersEvent
	WorkflowTaskUpdatedAssigneeEvent *WorkflowTaskUpdatedAssigneeEvent
	WorkflowTaskUpdatedFieldsEvent *WorkflowTaskUpdatedFieldsEvent
	WorkflowTaskUpdatedScheduledOnEvent *WorkflowTaskUpdatedScheduledOnEvent
	WorkflowTaskUpdatedStatusEvent *WorkflowTaskUpdatedStatusEvent
}

// AssayRunCreatedEventAsEvent is a convenience function that returns AssayRunCreatedEvent wrapped in Event
func AssayRunCreatedEventAsEvent(v *AssayRunCreatedEvent) Event {
	return Event{
		AssayRunCreatedEvent: v,
	}
}

// AssayRunUpdatedFieldsEventAsEvent is a convenience function that returns AssayRunUpdatedFieldsEvent wrapped in Event
func AssayRunUpdatedFieldsEventAsEvent(v *AssayRunUpdatedFieldsEvent) Event {
	return Event{
		AssayRunUpdatedFieldsEvent: v,
	}
}

// AutomationInputGeneratorCompletedV2BetaEventAsEvent is a convenience function that returns AutomationInputGeneratorCompletedV2BetaEvent wrapped in Event
func AutomationInputGeneratorCompletedV2BetaEventAsEvent(v *AutomationInputGeneratorCompletedV2BetaEvent) Event {
	return Event{
		AutomationInputGeneratorCompletedV2BetaEvent: v,
	}
}

// AutomationOutputProcessorCompletedV2BetaEventAsEvent is a convenience function that returns AutomationOutputProcessorCompletedV2BetaEvent wrapped in Event
func AutomationOutputProcessorCompletedV2BetaEventAsEvent(v *AutomationOutputProcessorCompletedV2BetaEvent) Event {
	return Event{
		AutomationOutputProcessorCompletedV2BetaEvent: v,
	}
}

// AutomationOutputProcessorUploadedV2BetaEventAsEvent is a convenience function that returns AutomationOutputProcessorUploadedV2BetaEvent wrapped in Event
func AutomationOutputProcessorUploadedV2BetaEventAsEvent(v *AutomationOutputProcessorUploadedV2BetaEvent) Event {
	return Event{
		AutomationOutputProcessorUploadedV2BetaEvent: v,
	}
}

// EntityRegisteredEventAsEvent is a convenience function that returns EntityRegisteredEvent wrapped in Event
func EntityRegisteredEventAsEvent(v *EntityRegisteredEvent) Event {
	return Event{
		EntityRegisteredEvent: v,
	}
}

// EntryCreatedEventAsEvent is a convenience function that returns EntryCreatedEvent wrapped in Event
func EntryCreatedEventAsEvent(v *EntryCreatedEvent) Event {
	return Event{
		EntryCreatedEvent: v,
	}
}

// EntryUpdatedFieldsEventAsEvent is a convenience function that returns EntryUpdatedFieldsEvent wrapped in Event
func EntryUpdatedFieldsEventAsEvent(v *EntryUpdatedFieldsEvent) Event {
	return Event{
		EntryUpdatedFieldsEvent: v,
	}
}

// EntryUpdatedReviewRecordEventAsEvent is a convenience function that returns EntryUpdatedReviewRecordEvent wrapped in Event
func EntryUpdatedReviewRecordEventAsEvent(v *EntryUpdatedReviewRecordEvent) Event {
	return Event{
		EntryUpdatedReviewRecordEvent: v,
	}
}

// RequestCreatedEventAsEvent is a convenience function that returns RequestCreatedEvent wrapped in Event
func RequestCreatedEventAsEvent(v *RequestCreatedEvent) Event {
	return Event{
		RequestCreatedEvent: v,
	}
}

// RequestUpdatedFieldsEventAsEvent is a convenience function that returns RequestUpdatedFieldsEvent wrapped in Event
func RequestUpdatedFieldsEventAsEvent(v *RequestUpdatedFieldsEvent) Event {
	return Event{
		RequestUpdatedFieldsEvent: v,
	}
}

// StageEntryCreatedEventAsEvent is a convenience function that returns StageEntryCreatedEvent wrapped in Event
func StageEntryCreatedEventAsEvent(v *StageEntryCreatedEvent) Event {
	return Event{
		StageEntryCreatedEvent: v,
	}
}

// StageEntryUpdatedFieldsEventAsEvent is a convenience function that returns StageEntryUpdatedFieldsEvent wrapped in Event
func StageEntryUpdatedFieldsEventAsEvent(v *StageEntryUpdatedFieldsEvent) Event {
	return Event{
		StageEntryUpdatedFieldsEvent: v,
	}
}

// StageEntryUpdatedReviewRecordEventAsEvent is a convenience function that returns StageEntryUpdatedReviewRecordEvent wrapped in Event
func StageEntryUpdatedReviewRecordEventAsEvent(v *StageEntryUpdatedReviewRecordEvent) Event {
	return Event{
		StageEntryUpdatedReviewRecordEvent: v,
	}
}

// WorkflowOutputCreatedEventAsEvent is a convenience function that returns WorkflowOutputCreatedEvent wrapped in Event
func WorkflowOutputCreatedEventAsEvent(v *WorkflowOutputCreatedEvent) Event {
	return Event{
		WorkflowOutputCreatedEvent: v,
	}
}

// WorkflowOutputUpdatedFieldsEventAsEvent is a convenience function that returns WorkflowOutputUpdatedFieldsEvent wrapped in Event
func WorkflowOutputUpdatedFieldsEventAsEvent(v *WorkflowOutputUpdatedFieldsEvent) Event {
	return Event{
		WorkflowOutputUpdatedFieldsEvent: v,
	}
}

// WorkflowTaskCreatedEventAsEvent is a convenience function that returns WorkflowTaskCreatedEvent wrapped in Event
func WorkflowTaskCreatedEventAsEvent(v *WorkflowTaskCreatedEvent) Event {
	return Event{
		WorkflowTaskCreatedEvent: v,
	}
}

// WorkflowTaskGroupCreatedEventAsEvent is a convenience function that returns WorkflowTaskGroupCreatedEvent wrapped in Event
func WorkflowTaskGroupCreatedEventAsEvent(v *WorkflowTaskGroupCreatedEvent) Event {
	return Event{
		WorkflowTaskGroupCreatedEvent: v,
	}
}

// WorkflowTaskGroupUpdatedWatchersEventAsEvent is a convenience function that returns WorkflowTaskGroupUpdatedWatchersEvent wrapped in Event
func WorkflowTaskGroupUpdatedWatchersEventAsEvent(v *WorkflowTaskGroupUpdatedWatchersEvent) Event {
	return Event{
		WorkflowTaskGroupUpdatedWatchersEvent: v,
	}
}

// WorkflowTaskUpdatedAssigneeEventAsEvent is a convenience function that returns WorkflowTaskUpdatedAssigneeEvent wrapped in Event
func WorkflowTaskUpdatedAssigneeEventAsEvent(v *WorkflowTaskUpdatedAssigneeEvent) Event {
	return Event{
		WorkflowTaskUpdatedAssigneeEvent: v,
	}
}

// WorkflowTaskUpdatedFieldsEventAsEvent is a convenience function that returns WorkflowTaskUpdatedFieldsEvent wrapped in Event
func WorkflowTaskUpdatedFieldsEventAsEvent(v *WorkflowTaskUpdatedFieldsEvent) Event {
	return Event{
		WorkflowTaskUpdatedFieldsEvent: v,
	}
}

// WorkflowTaskUpdatedScheduledOnEventAsEvent is a convenience function that returns WorkflowTaskUpdatedScheduledOnEvent wrapped in Event
func WorkflowTaskUpdatedScheduledOnEventAsEvent(v *WorkflowTaskUpdatedScheduledOnEvent) Event {
	return Event{
		WorkflowTaskUpdatedScheduledOnEvent: v,
	}
}

// WorkflowTaskUpdatedStatusEventAsEvent is a convenience function that returns WorkflowTaskUpdatedStatusEvent wrapped in Event
func WorkflowTaskUpdatedStatusEventAsEvent(v *WorkflowTaskUpdatedStatusEvent) Event {
	return Event{
		WorkflowTaskUpdatedStatusEvent: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Event) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AssayRunCreatedEvent
	err = newStrictDecoder(data).Decode(&dst.AssayRunCreatedEvent)
	if err == nil {
		jsonAssayRunCreatedEvent, _ := json.Marshal(dst.AssayRunCreatedEvent)
		if string(jsonAssayRunCreatedEvent) == "{}" { // empty struct
			dst.AssayRunCreatedEvent = nil
		} else {
			match++
		}
	} else {
		dst.AssayRunCreatedEvent = nil
	}

	// try to unmarshal data into AssayRunUpdatedFieldsEvent
	err = newStrictDecoder(data).Decode(&dst.AssayRunUpdatedFieldsEvent)
	if err == nil {
		jsonAssayRunUpdatedFieldsEvent, _ := json.Marshal(dst.AssayRunUpdatedFieldsEvent)
		if string(jsonAssayRunUpdatedFieldsEvent) == "{}" { // empty struct
			dst.AssayRunUpdatedFieldsEvent = nil
		} else {
			match++
		}
	} else {
		dst.AssayRunUpdatedFieldsEvent = nil
	}

	// try to unmarshal data into AutomationInputGeneratorCompletedV2BetaEvent
	err = newStrictDecoder(data).Decode(&dst.AutomationInputGeneratorCompletedV2BetaEvent)
	if err == nil {
		jsonAutomationInputGeneratorCompletedV2BetaEvent, _ := json.Marshal(dst.AutomationInputGeneratorCompletedV2BetaEvent)
		if string(jsonAutomationInputGeneratorCompletedV2BetaEvent) == "{}" { // empty struct
			dst.AutomationInputGeneratorCompletedV2BetaEvent = nil
		} else {
			match++
		}
	} else {
		dst.AutomationInputGeneratorCompletedV2BetaEvent = nil
	}

	// try to unmarshal data into AutomationOutputProcessorCompletedV2BetaEvent
	err = newStrictDecoder(data).Decode(&dst.AutomationOutputProcessorCompletedV2BetaEvent)
	if err == nil {
		jsonAutomationOutputProcessorCompletedV2BetaEvent, _ := json.Marshal(dst.AutomationOutputProcessorCompletedV2BetaEvent)
		if string(jsonAutomationOutputProcessorCompletedV2BetaEvent) == "{}" { // empty struct
			dst.AutomationOutputProcessorCompletedV2BetaEvent = nil
		} else {
			match++
		}
	} else {
		dst.AutomationOutputProcessorCompletedV2BetaEvent = nil
	}

	// try to unmarshal data into AutomationOutputProcessorUploadedV2BetaEvent
	err = newStrictDecoder(data).Decode(&dst.AutomationOutputProcessorUploadedV2BetaEvent)
	if err == nil {
		jsonAutomationOutputProcessorUploadedV2BetaEvent, _ := json.Marshal(dst.AutomationOutputProcessorUploadedV2BetaEvent)
		if string(jsonAutomationOutputProcessorUploadedV2BetaEvent) == "{}" { // empty struct
			dst.AutomationOutputProcessorUploadedV2BetaEvent = nil
		} else {
			match++
		}
	} else {
		dst.AutomationOutputProcessorUploadedV2BetaEvent = nil
	}

	// try to unmarshal data into EntityRegisteredEvent
	err = newStrictDecoder(data).Decode(&dst.EntityRegisteredEvent)
	if err == nil {
		jsonEntityRegisteredEvent, _ := json.Marshal(dst.EntityRegisteredEvent)
		if string(jsonEntityRegisteredEvent) == "{}" { // empty struct
			dst.EntityRegisteredEvent = nil
		} else {
			match++
		}
	} else {
		dst.EntityRegisteredEvent = nil
	}

	// try to unmarshal data into EntryCreatedEvent
	err = newStrictDecoder(data).Decode(&dst.EntryCreatedEvent)
	if err == nil {
		jsonEntryCreatedEvent, _ := json.Marshal(dst.EntryCreatedEvent)
		if string(jsonEntryCreatedEvent) == "{}" { // empty struct
			dst.EntryCreatedEvent = nil
		} else {
			match++
		}
	} else {
		dst.EntryCreatedEvent = nil
	}

	// try to unmarshal data into EntryUpdatedFieldsEvent
	err = newStrictDecoder(data).Decode(&dst.EntryUpdatedFieldsEvent)
	if err == nil {
		jsonEntryUpdatedFieldsEvent, _ := json.Marshal(dst.EntryUpdatedFieldsEvent)
		if string(jsonEntryUpdatedFieldsEvent) == "{}" { // empty struct
			dst.EntryUpdatedFieldsEvent = nil
		} else {
			match++
		}
	} else {
		dst.EntryUpdatedFieldsEvent = nil
	}

	// try to unmarshal data into EntryUpdatedReviewRecordEvent
	err = newStrictDecoder(data).Decode(&dst.EntryUpdatedReviewRecordEvent)
	if err == nil {
		jsonEntryUpdatedReviewRecordEvent, _ := json.Marshal(dst.EntryUpdatedReviewRecordEvent)
		if string(jsonEntryUpdatedReviewRecordEvent) == "{}" { // empty struct
			dst.EntryUpdatedReviewRecordEvent = nil
		} else {
			match++
		}
	} else {
		dst.EntryUpdatedReviewRecordEvent = nil
	}

	// try to unmarshal data into RequestCreatedEvent
	err = newStrictDecoder(data).Decode(&dst.RequestCreatedEvent)
	if err == nil {
		jsonRequestCreatedEvent, _ := json.Marshal(dst.RequestCreatedEvent)
		if string(jsonRequestCreatedEvent) == "{}" { // empty struct
			dst.RequestCreatedEvent = nil
		} else {
			match++
		}
	} else {
		dst.RequestCreatedEvent = nil
	}

	// try to unmarshal data into RequestUpdatedFieldsEvent
	err = newStrictDecoder(data).Decode(&dst.RequestUpdatedFieldsEvent)
	if err == nil {
		jsonRequestUpdatedFieldsEvent, _ := json.Marshal(dst.RequestUpdatedFieldsEvent)
		if string(jsonRequestUpdatedFieldsEvent) == "{}" { // empty struct
			dst.RequestUpdatedFieldsEvent = nil
		} else {
			match++
		}
	} else {
		dst.RequestUpdatedFieldsEvent = nil
	}

	// try to unmarshal data into StageEntryCreatedEvent
	err = newStrictDecoder(data).Decode(&dst.StageEntryCreatedEvent)
	if err == nil {
		jsonStageEntryCreatedEvent, _ := json.Marshal(dst.StageEntryCreatedEvent)
		if string(jsonStageEntryCreatedEvent) == "{}" { // empty struct
			dst.StageEntryCreatedEvent = nil
		} else {
			match++
		}
	} else {
		dst.StageEntryCreatedEvent = nil
	}

	// try to unmarshal data into StageEntryUpdatedFieldsEvent
	err = newStrictDecoder(data).Decode(&dst.StageEntryUpdatedFieldsEvent)
	if err == nil {
		jsonStageEntryUpdatedFieldsEvent, _ := json.Marshal(dst.StageEntryUpdatedFieldsEvent)
		if string(jsonStageEntryUpdatedFieldsEvent) == "{}" { // empty struct
			dst.StageEntryUpdatedFieldsEvent = nil
		} else {
			match++
		}
	} else {
		dst.StageEntryUpdatedFieldsEvent = nil
	}

	// try to unmarshal data into StageEntryUpdatedReviewRecordEvent
	err = newStrictDecoder(data).Decode(&dst.StageEntryUpdatedReviewRecordEvent)
	if err == nil {
		jsonStageEntryUpdatedReviewRecordEvent, _ := json.Marshal(dst.StageEntryUpdatedReviewRecordEvent)
		if string(jsonStageEntryUpdatedReviewRecordEvent) == "{}" { // empty struct
			dst.StageEntryUpdatedReviewRecordEvent = nil
		} else {
			match++
		}
	} else {
		dst.StageEntryUpdatedReviewRecordEvent = nil
	}

	// try to unmarshal data into WorkflowOutputCreatedEvent
	err = newStrictDecoder(data).Decode(&dst.WorkflowOutputCreatedEvent)
	if err == nil {
		jsonWorkflowOutputCreatedEvent, _ := json.Marshal(dst.WorkflowOutputCreatedEvent)
		if string(jsonWorkflowOutputCreatedEvent) == "{}" { // empty struct
			dst.WorkflowOutputCreatedEvent = nil
		} else {
			match++
		}
	} else {
		dst.WorkflowOutputCreatedEvent = nil
	}

	// try to unmarshal data into WorkflowOutputUpdatedFieldsEvent
	err = newStrictDecoder(data).Decode(&dst.WorkflowOutputUpdatedFieldsEvent)
	if err == nil {
		jsonWorkflowOutputUpdatedFieldsEvent, _ := json.Marshal(dst.WorkflowOutputUpdatedFieldsEvent)
		if string(jsonWorkflowOutputUpdatedFieldsEvent) == "{}" { // empty struct
			dst.WorkflowOutputUpdatedFieldsEvent = nil
		} else {
			match++
		}
	} else {
		dst.WorkflowOutputUpdatedFieldsEvent = nil
	}

	// try to unmarshal data into WorkflowTaskCreatedEvent
	err = newStrictDecoder(data).Decode(&dst.WorkflowTaskCreatedEvent)
	if err == nil {
		jsonWorkflowTaskCreatedEvent, _ := json.Marshal(dst.WorkflowTaskCreatedEvent)
		if string(jsonWorkflowTaskCreatedEvent) == "{}" { // empty struct
			dst.WorkflowTaskCreatedEvent = nil
		} else {
			match++
		}
	} else {
		dst.WorkflowTaskCreatedEvent = nil
	}

	// try to unmarshal data into WorkflowTaskGroupCreatedEvent
	err = newStrictDecoder(data).Decode(&dst.WorkflowTaskGroupCreatedEvent)
	if err == nil {
		jsonWorkflowTaskGroupCreatedEvent, _ := json.Marshal(dst.WorkflowTaskGroupCreatedEvent)
		if string(jsonWorkflowTaskGroupCreatedEvent) == "{}" { // empty struct
			dst.WorkflowTaskGroupCreatedEvent = nil
		} else {
			match++
		}
	} else {
		dst.WorkflowTaskGroupCreatedEvent = nil
	}

	// try to unmarshal data into WorkflowTaskGroupUpdatedWatchersEvent
	err = newStrictDecoder(data).Decode(&dst.WorkflowTaskGroupUpdatedWatchersEvent)
	if err == nil {
		jsonWorkflowTaskGroupUpdatedWatchersEvent, _ := json.Marshal(dst.WorkflowTaskGroupUpdatedWatchersEvent)
		if string(jsonWorkflowTaskGroupUpdatedWatchersEvent) == "{}" { // empty struct
			dst.WorkflowTaskGroupUpdatedWatchersEvent = nil
		} else {
			match++
		}
	} else {
		dst.WorkflowTaskGroupUpdatedWatchersEvent = nil
	}

	// try to unmarshal data into WorkflowTaskUpdatedAssigneeEvent
	err = newStrictDecoder(data).Decode(&dst.WorkflowTaskUpdatedAssigneeEvent)
	if err == nil {
		jsonWorkflowTaskUpdatedAssigneeEvent, _ := json.Marshal(dst.WorkflowTaskUpdatedAssigneeEvent)
		if string(jsonWorkflowTaskUpdatedAssigneeEvent) == "{}" { // empty struct
			dst.WorkflowTaskUpdatedAssigneeEvent = nil
		} else {
			match++
		}
	} else {
		dst.WorkflowTaskUpdatedAssigneeEvent = nil
	}

	// try to unmarshal data into WorkflowTaskUpdatedFieldsEvent
	err = newStrictDecoder(data).Decode(&dst.WorkflowTaskUpdatedFieldsEvent)
	if err == nil {
		jsonWorkflowTaskUpdatedFieldsEvent, _ := json.Marshal(dst.WorkflowTaskUpdatedFieldsEvent)
		if string(jsonWorkflowTaskUpdatedFieldsEvent) == "{}" { // empty struct
			dst.WorkflowTaskUpdatedFieldsEvent = nil
		} else {
			match++
		}
	} else {
		dst.WorkflowTaskUpdatedFieldsEvent = nil
	}

	// try to unmarshal data into WorkflowTaskUpdatedScheduledOnEvent
	err = newStrictDecoder(data).Decode(&dst.WorkflowTaskUpdatedScheduledOnEvent)
	if err == nil {
		jsonWorkflowTaskUpdatedScheduledOnEvent, _ := json.Marshal(dst.WorkflowTaskUpdatedScheduledOnEvent)
		if string(jsonWorkflowTaskUpdatedScheduledOnEvent) == "{}" { // empty struct
			dst.WorkflowTaskUpdatedScheduledOnEvent = nil
		} else {
			match++
		}
	} else {
		dst.WorkflowTaskUpdatedScheduledOnEvent = nil
	}

	// try to unmarshal data into WorkflowTaskUpdatedStatusEvent
	err = newStrictDecoder(data).Decode(&dst.WorkflowTaskUpdatedStatusEvent)
	if err == nil {
		jsonWorkflowTaskUpdatedStatusEvent, _ := json.Marshal(dst.WorkflowTaskUpdatedStatusEvent)
		if string(jsonWorkflowTaskUpdatedStatusEvent) == "{}" { // empty struct
			dst.WorkflowTaskUpdatedStatusEvent = nil
		} else {
			match++
		}
	} else {
		dst.WorkflowTaskUpdatedStatusEvent = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AssayRunCreatedEvent = nil
		dst.AssayRunUpdatedFieldsEvent = nil
		dst.AutomationInputGeneratorCompletedV2BetaEvent = nil
		dst.AutomationOutputProcessorCompletedV2BetaEvent = nil
		dst.AutomationOutputProcessorUploadedV2BetaEvent = nil
		dst.EntityRegisteredEvent = nil
		dst.EntryCreatedEvent = nil
		dst.EntryUpdatedFieldsEvent = nil
		dst.EntryUpdatedReviewRecordEvent = nil
		dst.RequestCreatedEvent = nil
		dst.RequestUpdatedFieldsEvent = nil
		dst.StageEntryCreatedEvent = nil
		dst.StageEntryUpdatedFieldsEvent = nil
		dst.StageEntryUpdatedReviewRecordEvent = nil
		dst.WorkflowOutputCreatedEvent = nil
		dst.WorkflowOutputUpdatedFieldsEvent = nil
		dst.WorkflowTaskCreatedEvent = nil
		dst.WorkflowTaskGroupCreatedEvent = nil
		dst.WorkflowTaskGroupUpdatedWatchersEvent = nil
		dst.WorkflowTaskUpdatedAssigneeEvent = nil
		dst.WorkflowTaskUpdatedFieldsEvent = nil
		dst.WorkflowTaskUpdatedScheduledOnEvent = nil
		dst.WorkflowTaskUpdatedStatusEvent = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Event)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Event)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Event) MarshalJSON() ([]byte, error) {
	if src.AssayRunCreatedEvent != nil {
		return json.Marshal(&src.AssayRunCreatedEvent)
	}

	if src.AssayRunUpdatedFieldsEvent != nil {
		return json.Marshal(&src.AssayRunUpdatedFieldsEvent)
	}

	if src.AutomationInputGeneratorCompletedV2BetaEvent != nil {
		return json.Marshal(&src.AutomationInputGeneratorCompletedV2BetaEvent)
	}

	if src.AutomationOutputProcessorCompletedV2BetaEvent != nil {
		return json.Marshal(&src.AutomationOutputProcessorCompletedV2BetaEvent)
	}

	if src.AutomationOutputProcessorUploadedV2BetaEvent != nil {
		return json.Marshal(&src.AutomationOutputProcessorUploadedV2BetaEvent)
	}

	if src.EntityRegisteredEvent != nil {
		return json.Marshal(&src.EntityRegisteredEvent)
	}

	if src.EntryCreatedEvent != nil {
		return json.Marshal(&src.EntryCreatedEvent)
	}

	if src.EntryUpdatedFieldsEvent != nil {
		return json.Marshal(&src.EntryUpdatedFieldsEvent)
	}

	if src.EntryUpdatedReviewRecordEvent != nil {
		return json.Marshal(&src.EntryUpdatedReviewRecordEvent)
	}

	if src.RequestCreatedEvent != nil {
		return json.Marshal(&src.RequestCreatedEvent)
	}

	if src.RequestUpdatedFieldsEvent != nil {
		return json.Marshal(&src.RequestUpdatedFieldsEvent)
	}

	if src.StageEntryCreatedEvent != nil {
		return json.Marshal(&src.StageEntryCreatedEvent)
	}

	if src.StageEntryUpdatedFieldsEvent != nil {
		return json.Marshal(&src.StageEntryUpdatedFieldsEvent)
	}

	if src.StageEntryUpdatedReviewRecordEvent != nil {
		return json.Marshal(&src.StageEntryUpdatedReviewRecordEvent)
	}

	if src.WorkflowOutputCreatedEvent != nil {
		return json.Marshal(&src.WorkflowOutputCreatedEvent)
	}

	if src.WorkflowOutputUpdatedFieldsEvent != nil {
		return json.Marshal(&src.WorkflowOutputUpdatedFieldsEvent)
	}

	if src.WorkflowTaskCreatedEvent != nil {
		return json.Marshal(&src.WorkflowTaskCreatedEvent)
	}

	if src.WorkflowTaskGroupCreatedEvent != nil {
		return json.Marshal(&src.WorkflowTaskGroupCreatedEvent)
	}

	if src.WorkflowTaskGroupUpdatedWatchersEvent != nil {
		return json.Marshal(&src.WorkflowTaskGroupUpdatedWatchersEvent)
	}

	if src.WorkflowTaskUpdatedAssigneeEvent != nil {
		return json.Marshal(&src.WorkflowTaskUpdatedAssigneeEvent)
	}

	if src.WorkflowTaskUpdatedFieldsEvent != nil {
		return json.Marshal(&src.WorkflowTaskUpdatedFieldsEvent)
	}

	if src.WorkflowTaskUpdatedScheduledOnEvent != nil {
		return json.Marshal(&src.WorkflowTaskUpdatedScheduledOnEvent)
	}

	if src.WorkflowTaskUpdatedStatusEvent != nil {
		return json.Marshal(&src.WorkflowTaskUpdatedStatusEvent)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Event) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AssayRunCreatedEvent != nil {
		return obj.AssayRunCreatedEvent
	}

	if obj.AssayRunUpdatedFieldsEvent != nil {
		return obj.AssayRunUpdatedFieldsEvent
	}

	if obj.AutomationInputGeneratorCompletedV2BetaEvent != nil {
		return obj.AutomationInputGeneratorCompletedV2BetaEvent
	}

	if obj.AutomationOutputProcessorCompletedV2BetaEvent != nil {
		return obj.AutomationOutputProcessorCompletedV2BetaEvent
	}

	if obj.AutomationOutputProcessorUploadedV2BetaEvent != nil {
		return obj.AutomationOutputProcessorUploadedV2BetaEvent
	}

	if obj.EntityRegisteredEvent != nil {
		return obj.EntityRegisteredEvent
	}

	if obj.EntryCreatedEvent != nil {
		return obj.EntryCreatedEvent
	}

	if obj.EntryUpdatedFieldsEvent != nil {
		return obj.EntryUpdatedFieldsEvent
	}

	if obj.EntryUpdatedReviewRecordEvent != nil {
		return obj.EntryUpdatedReviewRecordEvent
	}

	if obj.RequestCreatedEvent != nil {
		return obj.RequestCreatedEvent
	}

	if obj.RequestUpdatedFieldsEvent != nil {
		return obj.RequestUpdatedFieldsEvent
	}

	if obj.StageEntryCreatedEvent != nil {
		return obj.StageEntryCreatedEvent
	}

	if obj.StageEntryUpdatedFieldsEvent != nil {
		return obj.StageEntryUpdatedFieldsEvent
	}

	if obj.StageEntryUpdatedReviewRecordEvent != nil {
		return obj.StageEntryUpdatedReviewRecordEvent
	}

	if obj.WorkflowOutputCreatedEvent != nil {
		return obj.WorkflowOutputCreatedEvent
	}

	if obj.WorkflowOutputUpdatedFieldsEvent != nil {
		return obj.WorkflowOutputUpdatedFieldsEvent
	}

	if obj.WorkflowTaskCreatedEvent != nil {
		return obj.WorkflowTaskCreatedEvent
	}

	if obj.WorkflowTaskGroupCreatedEvent != nil {
		return obj.WorkflowTaskGroupCreatedEvent
	}

	if obj.WorkflowTaskGroupUpdatedWatchersEvent != nil {
		return obj.WorkflowTaskGroupUpdatedWatchersEvent
	}

	if obj.WorkflowTaskUpdatedAssigneeEvent != nil {
		return obj.WorkflowTaskUpdatedAssigneeEvent
	}

	if obj.WorkflowTaskUpdatedFieldsEvent != nil {
		return obj.WorkflowTaskUpdatedFieldsEvent
	}

	if obj.WorkflowTaskUpdatedScheduledOnEvent != nil {
		return obj.WorkflowTaskUpdatedScheduledOnEvent
	}

	if obj.WorkflowTaskUpdatedStatusEvent != nil {
		return obj.WorkflowTaskUpdatedStatusEvent
	}

	// all schemas are nil
	return nil
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


