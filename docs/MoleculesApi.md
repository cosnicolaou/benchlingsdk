# \MoleculesApi

All URIs are relative to */api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveMolecules**](MoleculesApi.md#ArchiveMolecules) | **Post** /molecules:archive | Archive Molecules
[**BulkCreateMolecules**](MoleculesApi.md#BulkCreateMolecules) | **Post** /molecules:bulk-create | Bulk Create Molecules
[**BulkUpdateMolecules**](MoleculesApi.md#BulkUpdateMolecules) | **Post** /molecules:bulk-update | Bulk Update Molecules
[**CreateMolecule**](MoleculesApi.md#CreateMolecule) | **Post** /molecules | Create a Molecule
[**GetMolecule**](MoleculesApi.md#GetMolecule) | **Get** /molecules/{molecule_id} | Get a Molecule
[**ListMolecules**](MoleculesApi.md#ListMolecules) | **Get** /molecules | List Molecules
[**UnarchiveMolecules**](MoleculesApi.md#UnarchiveMolecules) | **Post** /molecules:unarchive | Unarchive Molecules
[**UpdateMolecule**](MoleculesApi.md#UpdateMolecule) | **Patch** /molecules/{molecule_id} | Update a Molecule



## ArchiveMolecules

> MoleculesArchivalChange ArchiveMolecules(ctx).MoleculesArchive(moleculesArchive).Execute()

Archive Molecules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moleculesArchive := *openapiclient.NewMoleculesArchive([]string{"MoleculeIds_example"}, "Reason_example") // MoleculesArchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoleculesApi.ArchiveMolecules(context.Background()).MoleculesArchive(moleculesArchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoleculesApi.ArchiveMolecules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ArchiveMolecules`: MoleculesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `MoleculesApi.ArchiveMolecules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveMoleculesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moleculesArchive** | [**MoleculesArchive**](MoleculesArchive.md) |  | 

### Return type

[**MoleculesArchivalChange**](MoleculesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateMolecules

> AsyncTaskLink BulkCreateMolecules(ctx).MoleculesBulkCreateRequest(moleculesBulkCreateRequest).Execute()

Bulk Create Molecules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moleculesBulkCreateRequest := *openapiclient.NewMoleculesBulkCreateRequest() // MoleculesBulkCreateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoleculesApi.BulkCreateMolecules(context.Background()).MoleculesBulkCreateRequest(moleculesBulkCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoleculesApi.BulkCreateMolecules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateMolecules`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `MoleculesApi.BulkCreateMolecules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateMoleculesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moleculesBulkCreateRequest** | [**MoleculesBulkCreateRequest**](MoleculesBulkCreateRequest.md) |  | 

### Return type

[**AsyncTaskLink**](AsyncTaskLink.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateMolecules

> AsyncTaskLink BulkUpdateMolecules(ctx).MoleculesBulkUpdateRequest(moleculesBulkUpdateRequest).Execute()

Bulk Update Molecules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moleculesBulkUpdateRequest := *openapiclient.NewMoleculesBulkUpdateRequest() // MoleculesBulkUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoleculesApi.BulkUpdateMolecules(context.Background()).MoleculesBulkUpdateRequest(moleculesBulkUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoleculesApi.BulkUpdateMolecules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateMolecules`: AsyncTaskLink
    fmt.Fprintf(os.Stdout, "Response from `MoleculesApi.BulkUpdateMolecules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateMoleculesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moleculesBulkUpdateRequest** | [**MoleculesBulkUpdateRequest**](MoleculesBulkUpdateRequest.md) |  | 

### Return type

[**AsyncTaskLink**](AsyncTaskLink.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMolecule

> Molecule CreateMolecule(ctx).MoleculeCreate(moleculeCreate).Execute()

Create a Molecule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moleculeCreate := *openapiclient.NewMoleculeCreate(*openapiclient.NewMoleculeBaseRequestChemicalStructure(), "Name_example") // MoleculeCreate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoleculesApi.CreateMolecule(context.Background()).MoleculeCreate(moleculeCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoleculesApi.CreateMolecule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMolecule`: Molecule
    fmt.Fprintf(os.Stdout, "Response from `MoleculesApi.CreateMolecule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMoleculeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moleculeCreate** | [**MoleculeCreate**](MoleculeCreate.md) |  | 

### Return type

[**Molecule**](Molecule.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMolecule

> Molecule GetMolecule(ctx, moleculeId).Execute()

Get a Molecule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moleculeId := "moleculeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoleculesApi.GetMolecule(context.Background(), moleculeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoleculesApi.GetMolecule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMolecule`: Molecule
    fmt.Fprintf(os.Stdout, "Response from `MoleculesApi.GetMolecule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moleculeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMoleculeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Molecule**](Molecule.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMolecules

> MoleculesPaginatedList ListMolecules(ctx).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).AuthorIdsAnyOf(authorIdsAnyOf).ChemicalSubstructureMol(chemicalSubstructureMol).ChemicalSubstructureSmiles(chemicalSubstructureSmiles).Execute()

List Molecules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    pageSize := int32(56) // int32 | Number of results to return. Defaults to 50, maximum of 100.  (optional) (default to 50)
    nextToken := "nextToken_example" // string | Token for pagination (optional)
    sort := "sort_example" // string |  (optional) (default to "modifiedAt:desc")
    modifiedAt := "modifiedAt_example" // string | Datetime, in RFC 3339 format. Supports the > and < operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. > 2017-04-30.  (optional)
    name := "name_example" // string | Name of a Molecule. Restricts results to those with the specified name, alias, or entity registry ID. (optional)
    nameIncludes := "nameIncludes_example" // string | Name substring of a Molecule. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. (optional)
    folderId := "folderId_example" // string | ID of a folder. Restricts results to those in the folder. (optional)
    mentionedIn := []string{"Inner_example"} // []string | Comma-separated list of entry IDs. Restricts results to Molecules mentioned in those entries.  (optional)
    projectId := "projectId_example" // string | ID of a project. Restricts results to those in the project. (optional)
    registryId := "registryId_example" // string | ID of a registry. Restricts results to those registered in this registry. Specifying \"null\" returns unregistered items.  (optional)
    schemaId := "schemaId_example" // string | ID of a schema. Restricts results to those of the specified schema.  (optional)
    schemaFields := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | Schema field value. For Integer, Float, and Date type fields, supports the >= and <= operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  (optional)
    archiveReason := "ANY_ARCHIVED" // string | Archive reason. Restricts results to those with the specified archive reason. Use “NOT_ARCHIVED” to filter for unarchived Molecules. Use \"ANY_ARCHIVED\" to filter for archived Molecules regardless of reason. Use \"ANY_ARCHIVED_OR_NOT_ARCHIVED\" to return items for both archived and unarchived.  (optional)
    mentions := []string{"Inner_example"} // []string | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  (optional)
    ids := "mol_yWs5X7lv,mol_RhYGVnHF" // string | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  (optional)
    entityRegistryIdsAnyOf := "TP001,TP002" // string | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  (optional)
    namesAnyOf := "MyName1,MyName2" // string | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs.  (optional)
    authorIdsAnyOf := "ent_a0SApq3z,ent_b4AApz9b" // string | Comma separated list of user or app IDs (optional)
    chemicalSubstructureMol := "Format described at https://en.wikipedia.org/wiki/Chemical_table_file#Molfile As an example, ethanol is represented as follows: ChEBI
  Marvin  10060515352D

  3  2  0  0  0  0            999 V2000
    4.8667   -3.3230    0.0000 C   0  0  0  0  0  0  0  0  0  0  0  0
    5.5812   -2.9105    0.0000 C   0  0  0  0  0  0  0  0  0  0  0  0
    6.2956   -3.3230    0.0000 O   0  0  0  0  0  0  0  0  0  0  0  0
  1  2  1  0  0  0  0
  2  3  1  0  0  0  0
M  END
" // string | mol-formatted string for a chemical substructure to search by (optional)
    chemicalSubstructureSmiles := "CCO,C(C1C(C(C(C(O1)O)O)O)O)O" // string | SMILES string for a chemical substructure to search by (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoleculesApi.ListMolecules(context.Background()).PageSize(pageSize).NextToken(nextToken).Sort(sort).ModifiedAt(modifiedAt).Name(name).NameIncludes(nameIncludes).FolderId(folderId).MentionedIn(mentionedIn).ProjectId(projectId).RegistryId(registryId).SchemaId(schemaId).SchemaFields(schemaFields).ArchiveReason(archiveReason).Mentions(mentions).Ids(ids).EntityRegistryIdsAnyOf(entityRegistryIdsAnyOf).NamesAnyOf(namesAnyOf).AuthorIdsAnyOf(authorIdsAnyOf).ChemicalSubstructureMol(chemicalSubstructureMol).ChemicalSubstructureSmiles(chemicalSubstructureSmiles).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoleculesApi.ListMolecules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMolecules`: MoleculesPaginatedList
    fmt.Fprintf(os.Stdout, "Response from `MoleculesApi.ListMolecules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMoleculesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Number of results to return. Defaults to 50, maximum of 100.  | [default to 50]
 **nextToken** | **string** | Token for pagination | 
 **sort** | **string** |  | [default to &quot;modifiedAt:desc&quot;]
 **modifiedAt** | **string** | Datetime, in RFC 3339 format. Supports the &gt; and &lt; operators. Time zone defaults to UTC. Restricts results to those modified in the specified range. e.g. &gt; 2017-04-30.  | 
 **name** | **string** | Name of a Molecule. Restricts results to those with the specified name, alias, or entity registry ID. | 
 **nameIncludes** | **string** | Name substring of a Molecule. Restricts results to those with names, aliases, or entity registry IDs that include the provided substring. | 
 **folderId** | **string** | ID of a folder. Restricts results to those in the folder. | 
 **mentionedIn** | **[]string** | Comma-separated list of entry IDs. Restricts results to Molecules mentioned in those entries.  | 
 **projectId** | **string** | ID of a project. Restricts results to those in the project. | 
 **registryId** | **string** | ID of a registry. Restricts results to those registered in this registry. Specifying \&quot;null\&quot; returns unregistered items.  | 
 **schemaId** | **string** | ID of a schema. Restricts results to those of the specified schema.  | 
 **schemaFields** | **map[string]interface{}** | Schema field value. For Integer, Float, and Date type fields, supports the &gt;&#x3D; and &lt;&#x3D; operators. If present, the schemaId param must also be present. Restricts results to those with a field of whose value matches the filter.  | 
 **archiveReason** | **string** | Archive reason. Restricts results to those with the specified archive reason. Use “NOT_ARCHIVED” to filter for unarchived Molecules. Use \&quot;ANY_ARCHIVED\&quot; to filter for archived Molecules regardless of reason. Use \&quot;ANY_ARCHIVED_OR_NOT_ARCHIVED\&quot; to return items for both archived and unarchived.  | 
 **mentions** | **[]string** | Comma-separated list of item IDs. Restricts results to those that mention the given items in the description.  | 
 **ids** | **string** | Comma-separated list of ids. Matches all of the provided IDs, or returns a 400 error that includes a list of which IDs are invalid.  | 
 **entityRegistryIdsAnyOf** | **string** | Comma-separated list of entity registry IDs. Restricts results to those that match any of the specified registry IDs.  | 
 **namesAnyOf** | **string** | Comma-separated list of names. Restricts results to those that match any of the specified names, aliases, or entity registry IDs.  | 
 **authorIdsAnyOf** | **string** | Comma separated list of user or app IDs | 
 **chemicalSubstructureMol** | **string** | mol-formatted string for a chemical substructure to search by | 
 **chemicalSubstructureSmiles** | **string** | SMILES string for a chemical substructure to search by | 

### Return type

[**MoleculesPaginatedList**](MoleculesPaginatedList.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnarchiveMolecules

> MoleculesArchivalChange UnarchiveMolecules(ctx).MoleculesUnarchive(moleculesUnarchive).Execute()

Unarchive Molecules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moleculesUnarchive := *openapiclient.NewMoleculesUnarchive([]string{"MoleculeIds_example"}) // MoleculesUnarchive |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoleculesApi.UnarchiveMolecules(context.Background()).MoleculesUnarchive(moleculesUnarchive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoleculesApi.UnarchiveMolecules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnarchiveMolecules`: MoleculesArchivalChange
    fmt.Fprintf(os.Stdout, "Response from `MoleculesApi.UnarchiveMolecules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveMoleculesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moleculesUnarchive** | [**MoleculesUnarchive**](MoleculesUnarchive.md) |  | 

### Return type

[**MoleculesArchivalChange**](MoleculesArchivalChange.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMolecule

> Molecule UpdateMolecule(ctx, moleculeId).MoleculeUpdate(moleculeUpdate).Execute()

Update a Molecule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    moleculeId := "moleculeId_example" // string | 
    moleculeUpdate := *openapiclient.NewMoleculeUpdate() // MoleculeUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MoleculesApi.UpdateMolecule(context.Background(), moleculeId).MoleculeUpdate(moleculeUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MoleculesApi.UpdateMolecule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMolecule`: Molecule
    fmt.Fprintf(os.Stdout, "Response from `MoleculesApi.UpdateMolecule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**moleculeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMoleculeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moleculeUpdate** | [**MoleculeUpdate**](MoleculeUpdate.md) |  | 

### Return type

[**Molecule**](Molecule.md)

### Authorization

[basicApiKeyAuth](../README.md#basicApiKeyAuth), [oAuth](../README.md#oAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

