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

// GetTask200Response - struct for GetTask200Response
type GetTask200Response struct {
	AIGGenerateInputAsyncTask *AIGGenerateInputAsyncTask
	AOPProcessOutputAsyncTask *AOPProcessOutputAsyncTask
	AsyncTask *AsyncTask
	AutofillPartsAsyncTask *AutofillPartsAsyncTask
	AutofillTranslationsAsyncTask *AutofillTranslationsAsyncTask
	BulkCreateAaSequencesAsyncTask *BulkCreateAaSequencesAsyncTask
	BulkCreateContainersAsyncTask *BulkCreateContainersAsyncTask
	BulkCreateCustomEntitiesAsyncTask *BulkCreateCustomEntitiesAsyncTask
	BulkCreateDnaOligosAsyncTask *BulkCreateDnaOligosAsyncTask
	BulkCreateDnaSequencesAsyncTask *BulkCreateDnaSequencesAsyncTask
	BulkCreateFeaturesAsyncTask *BulkCreateFeaturesAsyncTask
	BulkCreateRnaOligosAsyncTask *BulkCreateRnaOligosAsyncTask
	BulkCreateRnaSequencesAsyncTask *BulkCreateRnaSequencesAsyncTask
	BulkRegisterEntitiesAsyncTask *BulkRegisterEntitiesAsyncTask
	BulkUpdateAaSequencesAsyncTask *BulkUpdateAaSequencesAsyncTask
	BulkUpdateContainersAsyncTask *BulkUpdateContainersAsyncTask
	BulkUpdateCustomEntitiesAsyncTask *BulkUpdateCustomEntitiesAsyncTask
	BulkUpdateDnaOligosAsyncTask *BulkUpdateDnaOligosAsyncTask
	BulkUpdateDnaSequencesAsyncTask *BulkUpdateDnaSequencesAsyncTask
	BulkUpdateRnaOligosAsyncTask *BulkUpdateRnaOligosAsyncTask
	BulkUpdateRnaSequencesAsyncTask *BulkUpdateRnaSequencesAsyncTask
	CreateConsensusAlignmentAsyncTask *CreateConsensusAlignmentAsyncTask
	CreateNucleotideConsensusAlignmentAsyncTask *CreateNucleotideConsensusAlignmentAsyncTask
	CreateNucleotideTemplateAlignmentAsyncTask *CreateNucleotideTemplateAlignmentAsyncTask
	CreateTemplateAlignmentAsyncTask *CreateTemplateAlignmentAsyncTask
	ExportsAsyncTask *ExportsAsyncTask
	FindMatchingRegionsAsyncTask *FindMatchingRegionsAsyncTask
	TransfersAsyncTask *TransfersAsyncTask
}

// AIGGenerateInputAsyncTaskAsGetTask200Response is a convenience function that returns AIGGenerateInputAsyncTask wrapped in GetTask200Response
func AIGGenerateInputAsyncTaskAsGetTask200Response(v *AIGGenerateInputAsyncTask) GetTask200Response {
	return GetTask200Response{
		AIGGenerateInputAsyncTask: v,
	}
}

// AOPProcessOutputAsyncTaskAsGetTask200Response is a convenience function that returns AOPProcessOutputAsyncTask wrapped in GetTask200Response
func AOPProcessOutputAsyncTaskAsGetTask200Response(v *AOPProcessOutputAsyncTask) GetTask200Response {
	return GetTask200Response{
		AOPProcessOutputAsyncTask: v,
	}
}

// AsyncTaskAsGetTask200Response is a convenience function that returns AsyncTask wrapped in GetTask200Response
func AsyncTaskAsGetTask200Response(v *AsyncTask) GetTask200Response {
	return GetTask200Response{
		AsyncTask: v,
	}
}

// AutofillPartsAsyncTaskAsGetTask200Response is a convenience function that returns AutofillPartsAsyncTask wrapped in GetTask200Response
func AutofillPartsAsyncTaskAsGetTask200Response(v *AutofillPartsAsyncTask) GetTask200Response {
	return GetTask200Response{
		AutofillPartsAsyncTask: v,
	}
}

// AutofillTranslationsAsyncTaskAsGetTask200Response is a convenience function that returns AutofillTranslationsAsyncTask wrapped in GetTask200Response
func AutofillTranslationsAsyncTaskAsGetTask200Response(v *AutofillTranslationsAsyncTask) GetTask200Response {
	return GetTask200Response{
		AutofillTranslationsAsyncTask: v,
	}
}

// BulkCreateAaSequencesAsyncTaskAsGetTask200Response is a convenience function that returns BulkCreateAaSequencesAsyncTask wrapped in GetTask200Response
func BulkCreateAaSequencesAsyncTaskAsGetTask200Response(v *BulkCreateAaSequencesAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkCreateAaSequencesAsyncTask: v,
	}
}

// BulkCreateContainersAsyncTaskAsGetTask200Response is a convenience function that returns BulkCreateContainersAsyncTask wrapped in GetTask200Response
func BulkCreateContainersAsyncTaskAsGetTask200Response(v *BulkCreateContainersAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkCreateContainersAsyncTask: v,
	}
}

// BulkCreateCustomEntitiesAsyncTaskAsGetTask200Response is a convenience function that returns BulkCreateCustomEntitiesAsyncTask wrapped in GetTask200Response
func BulkCreateCustomEntitiesAsyncTaskAsGetTask200Response(v *BulkCreateCustomEntitiesAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkCreateCustomEntitiesAsyncTask: v,
	}
}

// BulkCreateDnaOligosAsyncTaskAsGetTask200Response is a convenience function that returns BulkCreateDnaOligosAsyncTask wrapped in GetTask200Response
func BulkCreateDnaOligosAsyncTaskAsGetTask200Response(v *BulkCreateDnaOligosAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkCreateDnaOligosAsyncTask: v,
	}
}

// BulkCreateDnaSequencesAsyncTaskAsGetTask200Response is a convenience function that returns BulkCreateDnaSequencesAsyncTask wrapped in GetTask200Response
func BulkCreateDnaSequencesAsyncTaskAsGetTask200Response(v *BulkCreateDnaSequencesAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkCreateDnaSequencesAsyncTask: v,
	}
}

// BulkCreateFeaturesAsyncTaskAsGetTask200Response is a convenience function that returns BulkCreateFeaturesAsyncTask wrapped in GetTask200Response
func BulkCreateFeaturesAsyncTaskAsGetTask200Response(v *BulkCreateFeaturesAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkCreateFeaturesAsyncTask: v,
	}
}

// BulkCreateRnaOligosAsyncTaskAsGetTask200Response is a convenience function that returns BulkCreateRnaOligosAsyncTask wrapped in GetTask200Response
func BulkCreateRnaOligosAsyncTaskAsGetTask200Response(v *BulkCreateRnaOligosAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkCreateRnaOligosAsyncTask: v,
	}
}

// BulkCreateRnaSequencesAsyncTaskAsGetTask200Response is a convenience function that returns BulkCreateRnaSequencesAsyncTask wrapped in GetTask200Response
func BulkCreateRnaSequencesAsyncTaskAsGetTask200Response(v *BulkCreateRnaSequencesAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkCreateRnaSequencesAsyncTask: v,
	}
}

// BulkRegisterEntitiesAsyncTaskAsGetTask200Response is a convenience function that returns BulkRegisterEntitiesAsyncTask wrapped in GetTask200Response
func BulkRegisterEntitiesAsyncTaskAsGetTask200Response(v *BulkRegisterEntitiesAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkRegisterEntitiesAsyncTask: v,
	}
}

// BulkUpdateAaSequencesAsyncTaskAsGetTask200Response is a convenience function that returns BulkUpdateAaSequencesAsyncTask wrapped in GetTask200Response
func BulkUpdateAaSequencesAsyncTaskAsGetTask200Response(v *BulkUpdateAaSequencesAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkUpdateAaSequencesAsyncTask: v,
	}
}

// BulkUpdateContainersAsyncTaskAsGetTask200Response is a convenience function that returns BulkUpdateContainersAsyncTask wrapped in GetTask200Response
func BulkUpdateContainersAsyncTaskAsGetTask200Response(v *BulkUpdateContainersAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkUpdateContainersAsyncTask: v,
	}
}

// BulkUpdateCustomEntitiesAsyncTaskAsGetTask200Response is a convenience function that returns BulkUpdateCustomEntitiesAsyncTask wrapped in GetTask200Response
func BulkUpdateCustomEntitiesAsyncTaskAsGetTask200Response(v *BulkUpdateCustomEntitiesAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkUpdateCustomEntitiesAsyncTask: v,
	}
}

// BulkUpdateDnaOligosAsyncTaskAsGetTask200Response is a convenience function that returns BulkUpdateDnaOligosAsyncTask wrapped in GetTask200Response
func BulkUpdateDnaOligosAsyncTaskAsGetTask200Response(v *BulkUpdateDnaOligosAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkUpdateDnaOligosAsyncTask: v,
	}
}

// BulkUpdateDnaSequencesAsyncTaskAsGetTask200Response is a convenience function that returns BulkUpdateDnaSequencesAsyncTask wrapped in GetTask200Response
func BulkUpdateDnaSequencesAsyncTaskAsGetTask200Response(v *BulkUpdateDnaSequencesAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkUpdateDnaSequencesAsyncTask: v,
	}
}

// BulkUpdateRnaOligosAsyncTaskAsGetTask200Response is a convenience function that returns BulkUpdateRnaOligosAsyncTask wrapped in GetTask200Response
func BulkUpdateRnaOligosAsyncTaskAsGetTask200Response(v *BulkUpdateRnaOligosAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkUpdateRnaOligosAsyncTask: v,
	}
}

// BulkUpdateRnaSequencesAsyncTaskAsGetTask200Response is a convenience function that returns BulkUpdateRnaSequencesAsyncTask wrapped in GetTask200Response
func BulkUpdateRnaSequencesAsyncTaskAsGetTask200Response(v *BulkUpdateRnaSequencesAsyncTask) GetTask200Response {
	return GetTask200Response{
		BulkUpdateRnaSequencesAsyncTask: v,
	}
}

// CreateConsensusAlignmentAsyncTaskAsGetTask200Response is a convenience function that returns CreateConsensusAlignmentAsyncTask wrapped in GetTask200Response
func CreateConsensusAlignmentAsyncTaskAsGetTask200Response(v *CreateConsensusAlignmentAsyncTask) GetTask200Response {
	return GetTask200Response{
		CreateConsensusAlignmentAsyncTask: v,
	}
}

// CreateNucleotideConsensusAlignmentAsyncTaskAsGetTask200Response is a convenience function that returns CreateNucleotideConsensusAlignmentAsyncTask wrapped in GetTask200Response
func CreateNucleotideConsensusAlignmentAsyncTaskAsGetTask200Response(v *CreateNucleotideConsensusAlignmentAsyncTask) GetTask200Response {
	return GetTask200Response{
		CreateNucleotideConsensusAlignmentAsyncTask: v,
	}
}

// CreateNucleotideTemplateAlignmentAsyncTaskAsGetTask200Response is a convenience function that returns CreateNucleotideTemplateAlignmentAsyncTask wrapped in GetTask200Response
func CreateNucleotideTemplateAlignmentAsyncTaskAsGetTask200Response(v *CreateNucleotideTemplateAlignmentAsyncTask) GetTask200Response {
	return GetTask200Response{
		CreateNucleotideTemplateAlignmentAsyncTask: v,
	}
}

// CreateTemplateAlignmentAsyncTaskAsGetTask200Response is a convenience function that returns CreateTemplateAlignmentAsyncTask wrapped in GetTask200Response
func CreateTemplateAlignmentAsyncTaskAsGetTask200Response(v *CreateTemplateAlignmentAsyncTask) GetTask200Response {
	return GetTask200Response{
		CreateTemplateAlignmentAsyncTask: v,
	}
}

// ExportsAsyncTaskAsGetTask200Response is a convenience function that returns ExportsAsyncTask wrapped in GetTask200Response
func ExportsAsyncTaskAsGetTask200Response(v *ExportsAsyncTask) GetTask200Response {
	return GetTask200Response{
		ExportsAsyncTask: v,
	}
}

// FindMatchingRegionsAsyncTaskAsGetTask200Response is a convenience function that returns FindMatchingRegionsAsyncTask wrapped in GetTask200Response
func FindMatchingRegionsAsyncTaskAsGetTask200Response(v *FindMatchingRegionsAsyncTask) GetTask200Response {
	return GetTask200Response{
		FindMatchingRegionsAsyncTask: v,
	}
}

// TransfersAsyncTaskAsGetTask200Response is a convenience function that returns TransfersAsyncTask wrapped in GetTask200Response
func TransfersAsyncTaskAsGetTask200Response(v *TransfersAsyncTask) GetTask200Response {
	return GetTask200Response{
		TransfersAsyncTask: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetTask200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AIGGenerateInputAsyncTask
	err = newStrictDecoder(data).Decode(&dst.AIGGenerateInputAsyncTask)
	if err == nil {
		jsonAIGGenerateInputAsyncTask, _ := json.Marshal(dst.AIGGenerateInputAsyncTask)
		if string(jsonAIGGenerateInputAsyncTask) == "{}" { // empty struct
			dst.AIGGenerateInputAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.AIGGenerateInputAsyncTask = nil
	}

	// try to unmarshal data into AOPProcessOutputAsyncTask
	err = newStrictDecoder(data).Decode(&dst.AOPProcessOutputAsyncTask)
	if err == nil {
		jsonAOPProcessOutputAsyncTask, _ := json.Marshal(dst.AOPProcessOutputAsyncTask)
		if string(jsonAOPProcessOutputAsyncTask) == "{}" { // empty struct
			dst.AOPProcessOutputAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.AOPProcessOutputAsyncTask = nil
	}

	// try to unmarshal data into AsyncTask
	err = newStrictDecoder(data).Decode(&dst.AsyncTask)
	if err == nil {
		jsonAsyncTask, _ := json.Marshal(dst.AsyncTask)
		if string(jsonAsyncTask) == "{}" { // empty struct
			dst.AsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.AsyncTask = nil
	}

	// try to unmarshal data into AutofillPartsAsyncTask
	err = newStrictDecoder(data).Decode(&dst.AutofillPartsAsyncTask)
	if err == nil {
		jsonAutofillPartsAsyncTask, _ := json.Marshal(dst.AutofillPartsAsyncTask)
		if string(jsonAutofillPartsAsyncTask) == "{}" { // empty struct
			dst.AutofillPartsAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.AutofillPartsAsyncTask = nil
	}

	// try to unmarshal data into AutofillTranslationsAsyncTask
	err = newStrictDecoder(data).Decode(&dst.AutofillTranslationsAsyncTask)
	if err == nil {
		jsonAutofillTranslationsAsyncTask, _ := json.Marshal(dst.AutofillTranslationsAsyncTask)
		if string(jsonAutofillTranslationsAsyncTask) == "{}" { // empty struct
			dst.AutofillTranslationsAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.AutofillTranslationsAsyncTask = nil
	}

	// try to unmarshal data into BulkCreateAaSequencesAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkCreateAaSequencesAsyncTask)
	if err == nil {
		jsonBulkCreateAaSequencesAsyncTask, _ := json.Marshal(dst.BulkCreateAaSequencesAsyncTask)
		if string(jsonBulkCreateAaSequencesAsyncTask) == "{}" { // empty struct
			dst.BulkCreateAaSequencesAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkCreateAaSequencesAsyncTask = nil
	}

	// try to unmarshal data into BulkCreateContainersAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkCreateContainersAsyncTask)
	if err == nil {
		jsonBulkCreateContainersAsyncTask, _ := json.Marshal(dst.BulkCreateContainersAsyncTask)
		if string(jsonBulkCreateContainersAsyncTask) == "{}" { // empty struct
			dst.BulkCreateContainersAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkCreateContainersAsyncTask = nil
	}

	// try to unmarshal data into BulkCreateCustomEntitiesAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkCreateCustomEntitiesAsyncTask)
	if err == nil {
		jsonBulkCreateCustomEntitiesAsyncTask, _ := json.Marshal(dst.BulkCreateCustomEntitiesAsyncTask)
		if string(jsonBulkCreateCustomEntitiesAsyncTask) == "{}" { // empty struct
			dst.BulkCreateCustomEntitiesAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkCreateCustomEntitiesAsyncTask = nil
	}

	// try to unmarshal data into BulkCreateDnaOligosAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkCreateDnaOligosAsyncTask)
	if err == nil {
		jsonBulkCreateDnaOligosAsyncTask, _ := json.Marshal(dst.BulkCreateDnaOligosAsyncTask)
		if string(jsonBulkCreateDnaOligosAsyncTask) == "{}" { // empty struct
			dst.BulkCreateDnaOligosAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkCreateDnaOligosAsyncTask = nil
	}

	// try to unmarshal data into BulkCreateDnaSequencesAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkCreateDnaSequencesAsyncTask)
	if err == nil {
		jsonBulkCreateDnaSequencesAsyncTask, _ := json.Marshal(dst.BulkCreateDnaSequencesAsyncTask)
		if string(jsonBulkCreateDnaSequencesAsyncTask) == "{}" { // empty struct
			dst.BulkCreateDnaSequencesAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkCreateDnaSequencesAsyncTask = nil
	}

	// try to unmarshal data into BulkCreateFeaturesAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkCreateFeaturesAsyncTask)
	if err == nil {
		jsonBulkCreateFeaturesAsyncTask, _ := json.Marshal(dst.BulkCreateFeaturesAsyncTask)
		if string(jsonBulkCreateFeaturesAsyncTask) == "{}" { // empty struct
			dst.BulkCreateFeaturesAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkCreateFeaturesAsyncTask = nil
	}

	// try to unmarshal data into BulkCreateRnaOligosAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkCreateRnaOligosAsyncTask)
	if err == nil {
		jsonBulkCreateRnaOligosAsyncTask, _ := json.Marshal(dst.BulkCreateRnaOligosAsyncTask)
		if string(jsonBulkCreateRnaOligosAsyncTask) == "{}" { // empty struct
			dst.BulkCreateRnaOligosAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkCreateRnaOligosAsyncTask = nil
	}

	// try to unmarshal data into BulkCreateRnaSequencesAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkCreateRnaSequencesAsyncTask)
	if err == nil {
		jsonBulkCreateRnaSequencesAsyncTask, _ := json.Marshal(dst.BulkCreateRnaSequencesAsyncTask)
		if string(jsonBulkCreateRnaSequencesAsyncTask) == "{}" { // empty struct
			dst.BulkCreateRnaSequencesAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkCreateRnaSequencesAsyncTask = nil
	}

	// try to unmarshal data into BulkRegisterEntitiesAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkRegisterEntitiesAsyncTask)
	if err == nil {
		jsonBulkRegisterEntitiesAsyncTask, _ := json.Marshal(dst.BulkRegisterEntitiesAsyncTask)
		if string(jsonBulkRegisterEntitiesAsyncTask) == "{}" { // empty struct
			dst.BulkRegisterEntitiesAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkRegisterEntitiesAsyncTask = nil
	}

	// try to unmarshal data into BulkUpdateAaSequencesAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkUpdateAaSequencesAsyncTask)
	if err == nil {
		jsonBulkUpdateAaSequencesAsyncTask, _ := json.Marshal(dst.BulkUpdateAaSequencesAsyncTask)
		if string(jsonBulkUpdateAaSequencesAsyncTask) == "{}" { // empty struct
			dst.BulkUpdateAaSequencesAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkUpdateAaSequencesAsyncTask = nil
	}

	// try to unmarshal data into BulkUpdateContainersAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkUpdateContainersAsyncTask)
	if err == nil {
		jsonBulkUpdateContainersAsyncTask, _ := json.Marshal(dst.BulkUpdateContainersAsyncTask)
		if string(jsonBulkUpdateContainersAsyncTask) == "{}" { // empty struct
			dst.BulkUpdateContainersAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkUpdateContainersAsyncTask = nil
	}

	// try to unmarshal data into BulkUpdateCustomEntitiesAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkUpdateCustomEntitiesAsyncTask)
	if err == nil {
		jsonBulkUpdateCustomEntitiesAsyncTask, _ := json.Marshal(dst.BulkUpdateCustomEntitiesAsyncTask)
		if string(jsonBulkUpdateCustomEntitiesAsyncTask) == "{}" { // empty struct
			dst.BulkUpdateCustomEntitiesAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkUpdateCustomEntitiesAsyncTask = nil
	}

	// try to unmarshal data into BulkUpdateDnaOligosAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkUpdateDnaOligosAsyncTask)
	if err == nil {
		jsonBulkUpdateDnaOligosAsyncTask, _ := json.Marshal(dst.BulkUpdateDnaOligosAsyncTask)
		if string(jsonBulkUpdateDnaOligosAsyncTask) == "{}" { // empty struct
			dst.BulkUpdateDnaOligosAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkUpdateDnaOligosAsyncTask = nil
	}

	// try to unmarshal data into BulkUpdateDnaSequencesAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkUpdateDnaSequencesAsyncTask)
	if err == nil {
		jsonBulkUpdateDnaSequencesAsyncTask, _ := json.Marshal(dst.BulkUpdateDnaSequencesAsyncTask)
		if string(jsonBulkUpdateDnaSequencesAsyncTask) == "{}" { // empty struct
			dst.BulkUpdateDnaSequencesAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkUpdateDnaSequencesAsyncTask = nil
	}

	// try to unmarshal data into BulkUpdateRnaOligosAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkUpdateRnaOligosAsyncTask)
	if err == nil {
		jsonBulkUpdateRnaOligosAsyncTask, _ := json.Marshal(dst.BulkUpdateRnaOligosAsyncTask)
		if string(jsonBulkUpdateRnaOligosAsyncTask) == "{}" { // empty struct
			dst.BulkUpdateRnaOligosAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkUpdateRnaOligosAsyncTask = nil
	}

	// try to unmarshal data into BulkUpdateRnaSequencesAsyncTask
	err = newStrictDecoder(data).Decode(&dst.BulkUpdateRnaSequencesAsyncTask)
	if err == nil {
		jsonBulkUpdateRnaSequencesAsyncTask, _ := json.Marshal(dst.BulkUpdateRnaSequencesAsyncTask)
		if string(jsonBulkUpdateRnaSequencesAsyncTask) == "{}" { // empty struct
			dst.BulkUpdateRnaSequencesAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.BulkUpdateRnaSequencesAsyncTask = nil
	}

	// try to unmarshal data into CreateConsensusAlignmentAsyncTask
	err = newStrictDecoder(data).Decode(&dst.CreateConsensusAlignmentAsyncTask)
	if err == nil {
		jsonCreateConsensusAlignmentAsyncTask, _ := json.Marshal(dst.CreateConsensusAlignmentAsyncTask)
		if string(jsonCreateConsensusAlignmentAsyncTask) == "{}" { // empty struct
			dst.CreateConsensusAlignmentAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.CreateConsensusAlignmentAsyncTask = nil
	}

	// try to unmarshal data into CreateNucleotideConsensusAlignmentAsyncTask
	err = newStrictDecoder(data).Decode(&dst.CreateNucleotideConsensusAlignmentAsyncTask)
	if err == nil {
		jsonCreateNucleotideConsensusAlignmentAsyncTask, _ := json.Marshal(dst.CreateNucleotideConsensusAlignmentAsyncTask)
		if string(jsonCreateNucleotideConsensusAlignmentAsyncTask) == "{}" { // empty struct
			dst.CreateNucleotideConsensusAlignmentAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.CreateNucleotideConsensusAlignmentAsyncTask = nil
	}

	// try to unmarshal data into CreateNucleotideTemplateAlignmentAsyncTask
	err = newStrictDecoder(data).Decode(&dst.CreateNucleotideTemplateAlignmentAsyncTask)
	if err == nil {
		jsonCreateNucleotideTemplateAlignmentAsyncTask, _ := json.Marshal(dst.CreateNucleotideTemplateAlignmentAsyncTask)
		if string(jsonCreateNucleotideTemplateAlignmentAsyncTask) == "{}" { // empty struct
			dst.CreateNucleotideTemplateAlignmentAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.CreateNucleotideTemplateAlignmentAsyncTask = nil
	}

	// try to unmarshal data into CreateTemplateAlignmentAsyncTask
	err = newStrictDecoder(data).Decode(&dst.CreateTemplateAlignmentAsyncTask)
	if err == nil {
		jsonCreateTemplateAlignmentAsyncTask, _ := json.Marshal(dst.CreateTemplateAlignmentAsyncTask)
		if string(jsonCreateTemplateAlignmentAsyncTask) == "{}" { // empty struct
			dst.CreateTemplateAlignmentAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.CreateTemplateAlignmentAsyncTask = nil
	}

	// try to unmarshal data into ExportsAsyncTask
	err = newStrictDecoder(data).Decode(&dst.ExportsAsyncTask)
	if err == nil {
		jsonExportsAsyncTask, _ := json.Marshal(dst.ExportsAsyncTask)
		if string(jsonExportsAsyncTask) == "{}" { // empty struct
			dst.ExportsAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.ExportsAsyncTask = nil
	}

	// try to unmarshal data into FindMatchingRegionsAsyncTask
	err = newStrictDecoder(data).Decode(&dst.FindMatchingRegionsAsyncTask)
	if err == nil {
		jsonFindMatchingRegionsAsyncTask, _ := json.Marshal(dst.FindMatchingRegionsAsyncTask)
		if string(jsonFindMatchingRegionsAsyncTask) == "{}" { // empty struct
			dst.FindMatchingRegionsAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.FindMatchingRegionsAsyncTask = nil
	}

	// try to unmarshal data into TransfersAsyncTask
	err = newStrictDecoder(data).Decode(&dst.TransfersAsyncTask)
	if err == nil {
		jsonTransfersAsyncTask, _ := json.Marshal(dst.TransfersAsyncTask)
		if string(jsonTransfersAsyncTask) == "{}" { // empty struct
			dst.TransfersAsyncTask = nil
		} else {
			match++
		}
	} else {
		dst.TransfersAsyncTask = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AIGGenerateInputAsyncTask = nil
		dst.AOPProcessOutputAsyncTask = nil
		dst.AsyncTask = nil
		dst.AutofillPartsAsyncTask = nil
		dst.AutofillTranslationsAsyncTask = nil
		dst.BulkCreateAaSequencesAsyncTask = nil
		dst.BulkCreateContainersAsyncTask = nil
		dst.BulkCreateCustomEntitiesAsyncTask = nil
		dst.BulkCreateDnaOligosAsyncTask = nil
		dst.BulkCreateDnaSequencesAsyncTask = nil
		dst.BulkCreateFeaturesAsyncTask = nil
		dst.BulkCreateRnaOligosAsyncTask = nil
		dst.BulkCreateRnaSequencesAsyncTask = nil
		dst.BulkRegisterEntitiesAsyncTask = nil
		dst.BulkUpdateAaSequencesAsyncTask = nil
		dst.BulkUpdateContainersAsyncTask = nil
		dst.BulkUpdateCustomEntitiesAsyncTask = nil
		dst.BulkUpdateDnaOligosAsyncTask = nil
		dst.BulkUpdateDnaSequencesAsyncTask = nil
		dst.BulkUpdateRnaOligosAsyncTask = nil
		dst.BulkUpdateRnaSequencesAsyncTask = nil
		dst.CreateConsensusAlignmentAsyncTask = nil
		dst.CreateNucleotideConsensusAlignmentAsyncTask = nil
		dst.CreateNucleotideTemplateAlignmentAsyncTask = nil
		dst.CreateTemplateAlignmentAsyncTask = nil
		dst.ExportsAsyncTask = nil
		dst.FindMatchingRegionsAsyncTask = nil
		dst.TransfersAsyncTask = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetTask200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetTask200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetTask200Response) MarshalJSON() ([]byte, error) {
	if src.AIGGenerateInputAsyncTask != nil {
		return json.Marshal(&src.AIGGenerateInputAsyncTask)
	}

	if src.AOPProcessOutputAsyncTask != nil {
		return json.Marshal(&src.AOPProcessOutputAsyncTask)
	}

	if src.AsyncTask != nil {
		return json.Marshal(&src.AsyncTask)
	}

	if src.AutofillPartsAsyncTask != nil {
		return json.Marshal(&src.AutofillPartsAsyncTask)
	}

	if src.AutofillTranslationsAsyncTask != nil {
		return json.Marshal(&src.AutofillTranslationsAsyncTask)
	}

	if src.BulkCreateAaSequencesAsyncTask != nil {
		return json.Marshal(&src.BulkCreateAaSequencesAsyncTask)
	}

	if src.BulkCreateContainersAsyncTask != nil {
		return json.Marshal(&src.BulkCreateContainersAsyncTask)
	}

	if src.BulkCreateCustomEntitiesAsyncTask != nil {
		return json.Marshal(&src.BulkCreateCustomEntitiesAsyncTask)
	}

	if src.BulkCreateDnaOligosAsyncTask != nil {
		return json.Marshal(&src.BulkCreateDnaOligosAsyncTask)
	}

	if src.BulkCreateDnaSequencesAsyncTask != nil {
		return json.Marshal(&src.BulkCreateDnaSequencesAsyncTask)
	}

	if src.BulkCreateFeaturesAsyncTask != nil {
		return json.Marshal(&src.BulkCreateFeaturesAsyncTask)
	}

	if src.BulkCreateRnaOligosAsyncTask != nil {
		return json.Marshal(&src.BulkCreateRnaOligosAsyncTask)
	}

	if src.BulkCreateRnaSequencesAsyncTask != nil {
		return json.Marshal(&src.BulkCreateRnaSequencesAsyncTask)
	}

	if src.BulkRegisterEntitiesAsyncTask != nil {
		return json.Marshal(&src.BulkRegisterEntitiesAsyncTask)
	}

	if src.BulkUpdateAaSequencesAsyncTask != nil {
		return json.Marshal(&src.BulkUpdateAaSequencesAsyncTask)
	}

	if src.BulkUpdateContainersAsyncTask != nil {
		return json.Marshal(&src.BulkUpdateContainersAsyncTask)
	}

	if src.BulkUpdateCustomEntitiesAsyncTask != nil {
		return json.Marshal(&src.BulkUpdateCustomEntitiesAsyncTask)
	}

	if src.BulkUpdateDnaOligosAsyncTask != nil {
		return json.Marshal(&src.BulkUpdateDnaOligosAsyncTask)
	}

	if src.BulkUpdateDnaSequencesAsyncTask != nil {
		return json.Marshal(&src.BulkUpdateDnaSequencesAsyncTask)
	}

	if src.BulkUpdateRnaOligosAsyncTask != nil {
		return json.Marshal(&src.BulkUpdateRnaOligosAsyncTask)
	}

	if src.BulkUpdateRnaSequencesAsyncTask != nil {
		return json.Marshal(&src.BulkUpdateRnaSequencesAsyncTask)
	}

	if src.CreateConsensusAlignmentAsyncTask != nil {
		return json.Marshal(&src.CreateConsensusAlignmentAsyncTask)
	}

	if src.CreateNucleotideConsensusAlignmentAsyncTask != nil {
		return json.Marshal(&src.CreateNucleotideConsensusAlignmentAsyncTask)
	}

	if src.CreateNucleotideTemplateAlignmentAsyncTask != nil {
		return json.Marshal(&src.CreateNucleotideTemplateAlignmentAsyncTask)
	}

	if src.CreateTemplateAlignmentAsyncTask != nil {
		return json.Marshal(&src.CreateTemplateAlignmentAsyncTask)
	}

	if src.ExportsAsyncTask != nil {
		return json.Marshal(&src.ExportsAsyncTask)
	}

	if src.FindMatchingRegionsAsyncTask != nil {
		return json.Marshal(&src.FindMatchingRegionsAsyncTask)
	}

	if src.TransfersAsyncTask != nil {
		return json.Marshal(&src.TransfersAsyncTask)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetTask200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AIGGenerateInputAsyncTask != nil {
		return obj.AIGGenerateInputAsyncTask
	}

	if obj.AOPProcessOutputAsyncTask != nil {
		return obj.AOPProcessOutputAsyncTask
	}

	if obj.AsyncTask != nil {
		return obj.AsyncTask
	}

	if obj.AutofillPartsAsyncTask != nil {
		return obj.AutofillPartsAsyncTask
	}

	if obj.AutofillTranslationsAsyncTask != nil {
		return obj.AutofillTranslationsAsyncTask
	}

	if obj.BulkCreateAaSequencesAsyncTask != nil {
		return obj.BulkCreateAaSequencesAsyncTask
	}

	if obj.BulkCreateContainersAsyncTask != nil {
		return obj.BulkCreateContainersAsyncTask
	}

	if obj.BulkCreateCustomEntitiesAsyncTask != nil {
		return obj.BulkCreateCustomEntitiesAsyncTask
	}

	if obj.BulkCreateDnaOligosAsyncTask != nil {
		return obj.BulkCreateDnaOligosAsyncTask
	}

	if obj.BulkCreateDnaSequencesAsyncTask != nil {
		return obj.BulkCreateDnaSequencesAsyncTask
	}

	if obj.BulkCreateFeaturesAsyncTask != nil {
		return obj.BulkCreateFeaturesAsyncTask
	}

	if obj.BulkCreateRnaOligosAsyncTask != nil {
		return obj.BulkCreateRnaOligosAsyncTask
	}

	if obj.BulkCreateRnaSequencesAsyncTask != nil {
		return obj.BulkCreateRnaSequencesAsyncTask
	}

	if obj.BulkRegisterEntitiesAsyncTask != nil {
		return obj.BulkRegisterEntitiesAsyncTask
	}

	if obj.BulkUpdateAaSequencesAsyncTask != nil {
		return obj.BulkUpdateAaSequencesAsyncTask
	}

	if obj.BulkUpdateContainersAsyncTask != nil {
		return obj.BulkUpdateContainersAsyncTask
	}

	if obj.BulkUpdateCustomEntitiesAsyncTask != nil {
		return obj.BulkUpdateCustomEntitiesAsyncTask
	}

	if obj.BulkUpdateDnaOligosAsyncTask != nil {
		return obj.BulkUpdateDnaOligosAsyncTask
	}

	if obj.BulkUpdateDnaSequencesAsyncTask != nil {
		return obj.BulkUpdateDnaSequencesAsyncTask
	}

	if obj.BulkUpdateRnaOligosAsyncTask != nil {
		return obj.BulkUpdateRnaOligosAsyncTask
	}

	if obj.BulkUpdateRnaSequencesAsyncTask != nil {
		return obj.BulkUpdateRnaSequencesAsyncTask
	}

	if obj.CreateConsensusAlignmentAsyncTask != nil {
		return obj.CreateConsensusAlignmentAsyncTask
	}

	if obj.CreateNucleotideConsensusAlignmentAsyncTask != nil {
		return obj.CreateNucleotideConsensusAlignmentAsyncTask
	}

	if obj.CreateNucleotideTemplateAlignmentAsyncTask != nil {
		return obj.CreateNucleotideTemplateAlignmentAsyncTask
	}

	if obj.CreateTemplateAlignmentAsyncTask != nil {
		return obj.CreateTemplateAlignmentAsyncTask
	}

	if obj.ExportsAsyncTask != nil {
		return obj.ExportsAsyncTask
	}

	if obj.FindMatchingRegionsAsyncTask != nil {
		return obj.FindMatchingRegionsAsyncTask
	}

	if obj.TransfersAsyncTask != nil {
		return obj.TransfersAsyncTask
	}

	// all schemas are nil
	return nil
}

type NullableGetTask200Response struct {
	value *GetTask200Response
	isSet bool
}

func (v NullableGetTask200Response) Get() *GetTask200Response {
	return v.value
}

func (v *NullableGetTask200Response) Set(val *GetTask200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTask200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTask200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTask200Response(val *GetTask200Response) *NullableGetTask200Response {
	return &NullableGetTask200Response{value: val, isSet: true}
}

func (v NullableGetTask200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTask200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

