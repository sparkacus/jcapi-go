# \ActiveDirectoryApi

All URIs are relative to *https://console.jumpcloud.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivedirectoriesDelete**](ActiveDirectoryApi.md#ActivedirectoriesDelete) | **Delete** /activedirectories/{id} | Delete an Active Directory
[**ActivedirectoriesGet**](ActiveDirectoryApi.md#ActivedirectoriesGet) | **Get** /activedirectories/{id} | Get an Active Directory
[**ActivedirectoriesList**](ActiveDirectoryApi.md#ActivedirectoriesList) | **Get** /activedirectories | List Active Directories
[**ActivedirectoriesPost**](ActiveDirectoryApi.md#ActivedirectoriesPost) | **Post** /activedirectories | Create a new Active Directory
[**GraphActiveDirectoryAssociationsList**](ActiveDirectoryApi.md#GraphActiveDirectoryAssociationsList) | **Get** /activedirectories/{activedirectory_id}/associations | List the associations of an Active Directory instance
[**GraphActiveDirectoryAssociationsPost**](ActiveDirectoryApi.md#GraphActiveDirectoryAssociationsPost) | **Post** /activedirectories/{activedirectory_id}/associations | Manage the associations of an Active Directory instance
[**GraphActiveDirectoryTraverseUserGroup**](ActiveDirectoryApi.md#GraphActiveDirectoryTraverseUserGroup) | **Get** /activedirectories/{activedirectory_id}/usergroups | List the User Groups associated with an Active Directory instance


# **ActivedirectoriesDelete**
> ActivedirectoriesDelete(ctx, id, contentType, accept)
Delete an Active Directory

This endpoint allows you to delete an Active Directory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| ObjectID of this Active Directory instance. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivedirectoriesGet**
> ActiveDirectoryOutput ActivedirectoriesGet(ctx, id, contentType, accept)
Get an Active Directory

This endpoint returns a specific Active Directory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **id** | **string**| ObjectID of this Active Directory instance. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]

### Return type

[**ActiveDirectoryOutput**](active-directory-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivedirectoriesList**
> []ActiveDirectoryOutput ActivedirectoriesList(ctx, contentType, accept, optional)
List Active Directories

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **fields** | **string**| The comma separated fields included in the returned records. If omitted the default list of fields will be returned.  | [default to ]
 **filter** | **string**| Supported operators are: eq, ne, gt, ge, lt, le, between, search | [default to ]
 **limit** | **int32**| The number of records to return at once. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]
 **sort** | **string**| The comma separated fields used to sort the collection. Default sort is ascending, prefix with &#x60;-&#x60; to sort descending.  | [default to ]

### Return type

[**[]ActiveDirectoryOutput**](active-directory-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivedirectoriesPost**
> ActiveDirectoryOutput ActivedirectoriesPost(ctx, contentType, accept, optional)
Create a new Active Directory

This endpoint allows you to create a new Active Directory.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**ActiveDirectoryInput**](ActiveDirectoryInput.md)|  | 

### Return type

[**ActiveDirectoryOutput**](active-directory-output.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphActiveDirectoryAssociationsList**
> []GraphConnection GraphActiveDirectoryAssociationsList(ctx, activedirectoryId, targets, contentType, accept, optional)
List the associations of an Active Directory instance

This endpoint returns the direct associations of this Active Directory instance.  A direct association can be a non-homogenous relationship between 2 different objects. For example Active Directory and Users.   #### Sample Request ``` https://console.jumpcloud.com/api/v2/activedirectories/{activedirectory_id}/associations?targets=user ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **activedirectoryId** | **string**|  | 
  **targets** | [**[]string**](string.md)|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activedirectoryId** | **string**|  | 
 **targets** | [**[]string**](string.md)|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]

### Return type

[**[]GraphConnection**](GraphConnection.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphActiveDirectoryAssociationsPost**
> GraphActiveDirectoryAssociationsPost(ctx, activedirectoryId, contentType, accept, optional)
Manage the associations of an Active Directory instance

This endpoint allows you to manage the _direct_ associations of an Active Directory instance.  A direct association can be a non-homogenous relationship between 2 different objects. For example Active Directory and Users.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/activedirectories/{activedirectory_id}/associations ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **activedirectoryId** | **string**|  | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activedirectoryId** | **string**|  | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **body** | [**GraphManagementReq**](GraphManagementReq.md)|  | 

### Return type

 (empty response body)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GraphActiveDirectoryTraverseUserGroup**
> []GraphObjectWithPaths GraphActiveDirectoryTraverseUserGroup(ctx, activedirectoryId, contentType, accept, optional)
List the User Groups associated with an Active Directory instance

This endpoint will return User Groups associated with an Active Directory instance. Each element will contain the group's type, id, attributes and paths.  The `attributes` object is a key/value hash of attributes specifically set for this group.  The `paths` array enumerates each path from this Active Directory instance to the corresponding User Group; this array represents all grouping and/or associations that would have to be removed to deprovision the User Group from this Active Directory instance.  See `/members` and `/associations` endpoints to manage those collections.  #### Sample Request ``` https://console.jumpcloud.com/api/v2/activedirectories/{activedirectory_id}/usersgroups ```

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context containing the authentication | nil if no authentication
  **activedirectoryId** | **string**| ObjectID of the Active Directory instance. | 
  **contentType** | **string**|  | [default to application/json]
  **accept** | **string**|  | [default to application/json]
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activedirectoryId** | **string**| ObjectID of the Active Directory instance. | 
 **contentType** | **string**|  | [default to application/json]
 **accept** | **string**|  | [default to application/json]
 **limit** | **int32**| The number of records to return at once. | [default to 10]
 **skip** | **int32**| The offset into the records to return. | [default to 0]

### Return type

[**[]GraphObjectWithPaths**](GraphObjectWithPaths.md)

### Authorization

[x-api-key](../README.md#x-api-key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
