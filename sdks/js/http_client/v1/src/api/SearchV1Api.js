// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon sdk
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.14.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.9
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/V1ListSearchesResponse', 'model/V1Search'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/V1ListSearchesResponse'), require('../model/V1Search'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.SearchV1Api = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.V1ListSearchesResponse, root.PolyaxonSdk.V1Search);
  }
}(this, function(ApiClient, V1ListSearchesResponse, V1Search) {
  'use strict';

  /**
   * SearchV1 service.
   * @module api/SearchV1Api
   * @version 1.14.4
   */

  /**
   * Constructs a new SearchV1Api. 
   * @alias module:api/SearchV1Api
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the createSearch operation.
     * @callback module:api/SearchV1Api~createSearchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Search} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List archived runs for user
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {module:model/V1Search} body Artifact store body
     * @param {module:api/SearchV1Api~createSearchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Search}
     */
    this.createSearch = function(owner, project, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling createSearch");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling createSearch");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createSearch");
      }


      var pathParams = {
        'owner': owner,
        'project': project
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Search;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/searches', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteSearch operation.
     * @callback module:api/SearchV1Api~deleteSearchCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update run
     * @param {String} owner Owner of the namespace
     * @param {String} project Project where the experiement will be assigned
     * @param {String} uuid Unique integer identifier of the entity
     * @param {module:api/SearchV1Api~deleteSearchCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteSearch = function(owner, project, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling deleteSearch");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling deleteSearch");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling deleteSearch");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
        'uuid': uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/searches/{uuid}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getSearch operation.
     * @callback module:api/SearchV1Api~getSearchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Search} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List runs
     * @param {String} owner Owner of the namespace
     * @param {String} project Project where the experiement will be assigned
     * @param {String} uuid Unique integer identifier of the entity
     * @param {module:api/SearchV1Api~getSearchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Search}
     */
    this.getSearch = function(owner, project, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getSearch");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling getSearch");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling getSearch");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
        'uuid': uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Search;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/searches/{uuid}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the listSearches operation.
     * @callback module:api/SearchV1Api~listSearchesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListSearchesResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List bookmarked runs for user
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/SearchV1Api~listSearchesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListSearchesResponse}
     */
    this.listSearches = function(owner, project, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listSearches");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling listSearches");
      }


      var pathParams = {
        'owner': owner,
        'project': project
      };
      var queryParams = {
        'offset': opts['offset'],
        'limit': opts['limit'],
        'sort': opts['sort'],
        'query': opts['query'],
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1ListSearchesResponse;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/searches', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the patchSearch operation.
     * @callback module:api/SearchV1Api~patchSearchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Search} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get run
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {String} search_uuid UUID
     * @param {module:model/V1Search} body Artifact store body
     * @param {module:api/SearchV1Api~patchSearchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Search}
     */
    this.patchSearch = function(owner, project, search_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling patchSearch");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling patchSearch");
      }

      // verify the required parameter 'search_uuid' is set
      if (search_uuid === undefined || search_uuid === null) {
        throw new Error("Missing the required parameter 'search_uuid' when calling patchSearch");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling patchSearch");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
        'search.uuid': search_uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Search;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/searches/{search.uuid}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the updateSearch operation.
     * @callback module:api/SearchV1Api~updateSearchCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1Search} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create new run
     * @param {String} owner Owner of the namespace
     * @param {String} project Project under namesapce
     * @param {String} search_uuid UUID
     * @param {module:model/V1Search} body Artifact store body
     * @param {module:api/SearchV1Api~updateSearchCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1Search}
     */
    this.updateSearch = function(owner, project, search_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateSearch");
      }

      // verify the required parameter 'project' is set
      if (project === undefined || project === null) {
        throw new Error("Missing the required parameter 'project' when calling updateSearch");
      }

      // verify the required parameter 'search_uuid' is set
      if (search_uuid === undefined || search_uuid === null) {
        throw new Error("Missing the required parameter 'search_uuid' when calling updateSearch");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateSearch");
      }


      var pathParams = {
        'owner': owner,
        'project': project,
        'search.uuid': search_uuid
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['ApiKey'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = V1Search;

      return this.apiClient.callApi(
        '/api/v1/{owner}/{project}/searches/{search.uuid}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
