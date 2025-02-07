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
    define(['ApiClient', 'model/V1HostAccess', 'model/V1ListHostAccessesResponse'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/V1HostAccess'), require('../model/V1ListHostAccessesResponse'));
  } else {
    // Browser globals (root is window)
    if (!root.PolyaxonSdk) {
      root.PolyaxonSdk = {};
    }
    root.PolyaxonSdk.GitAccessesV1Api = factory(root.PolyaxonSdk.ApiClient, root.PolyaxonSdk.V1HostAccess, root.PolyaxonSdk.V1ListHostAccessesResponse);
  }
}(this, function(ApiClient, V1HostAccess, V1ListHostAccessesResponse) {
  'use strict';

  /**
   * GitAccessesV1 service.
   * @module api/GitAccessesV1Api
   * @version 1.14.4
   */

  /**
   * Constructs a new GitAccessesV1Api. 
   * @alias module:api/GitAccessesV1Api
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the createGitAccess operation.
     * @callback module:api/GitAccessesV1Api~createGitAccessCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HostAccess} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List runs
     * @param {String} owner Owner of the namespace
     * @param {module:model/V1HostAccess} body Artifact store body
     * @param {module:api/GitAccessesV1Api~createGitAccessCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HostAccess}
     */
    this.createGitAccess = function(owner, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling createGitAccess");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling createGitAccess");
      }


      var pathParams = {
        'owner': owner
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
      var returnType = V1HostAccess;

      return this.apiClient.callApi(
        '/api/v1/{owner}/git_accesses', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the deleteGitAccess operation.
     * @callback module:api/GitAccessesV1Api~deleteGitAccessCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Patch run
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Unique integer identifier of the entity
     * @param {module:api/GitAccessesV1Api~deleteGitAccessCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteGitAccess = function(owner, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling deleteGitAccess");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling deleteGitAccess");
      }


      var pathParams = {
        'owner': owner,
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
        '/api/v1/{owner}/git_accesses/{uuid}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getGitAccess operation.
     * @callback module:api/GitAccessesV1Api~getGitAccessCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HostAccess} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Create new run
     * @param {String} owner Owner of the namespace
     * @param {String} uuid Unique integer identifier of the entity
     * @param {module:api/GitAccessesV1Api~getGitAccessCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HostAccess}
     */
    this.getGitAccess = function(owner, uuid, callback) {
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling getGitAccess");
      }

      // verify the required parameter 'uuid' is set
      if (uuid === undefined || uuid === null) {
        throw new Error("Missing the required parameter 'uuid' when calling getGitAccess");
      }


      var pathParams = {
        'owner': owner,
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
      var returnType = V1HostAccess;

      return this.apiClient.callApi(
        '/api/v1/{owner}/git_accesses/{uuid}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the listGitAccessNames operation.
     * @callback module:api/GitAccessesV1Api~listGitAccessNamesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListHostAccessesResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List bookmarked runs for user
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/GitAccessesV1Api~listGitAccessNamesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListHostAccessesResponse}
     */
    this.listGitAccessNames = function(owner, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listGitAccessNames");
      }


      var pathParams = {
        'owner': owner
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
      var returnType = V1ListHostAccessesResponse;

      return this.apiClient.callApi(
        '/api/v1/{owner}/git_accesses/names', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the listGitAccesses operation.
     * @callback module:api/GitAccessesV1Api~listGitAccessesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1ListHostAccessesResponse} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * List archived runs for user
     * @param {String} owner Owner of the namespace
     * @param {Object} opts Optional parameters
     * @param {Number} opts.offset Pagination offset.
     * @param {Number} opts.limit Limit size.
     * @param {String} opts.sort Sort to order the search.
     * @param {String} opts.query Query filter the search search.
     * @param {module:api/GitAccessesV1Api~listGitAccessesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1ListHostAccessesResponse}
     */
    this.listGitAccesses = function(owner, opts, callback) {
      opts = opts || {};
      var postBody = null;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling listGitAccesses");
      }


      var pathParams = {
        'owner': owner
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
      var returnType = V1ListHostAccessesResponse;

      return this.apiClient.callApi(
        '/api/v1/{owner}/git_accesses', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the patchGitAccess operation.
     * @callback module:api/GitAccessesV1Api~patchGitAccessCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HostAccess} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Update run
     * @param {String} owner Owner of the namespace
     * @param {String} host_access_uuid UUID
     * @param {module:model/V1HostAccess} body Artifact store body
     * @param {module:api/GitAccessesV1Api~patchGitAccessCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HostAccess}
     */
    this.patchGitAccess = function(owner, host_access_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling patchGitAccess");
      }

      // verify the required parameter 'host_access_uuid' is set
      if (host_access_uuid === undefined || host_access_uuid === null) {
        throw new Error("Missing the required parameter 'host_access_uuid' when calling patchGitAccess");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling patchGitAccess");
      }


      var pathParams = {
        'owner': owner,
        'host_access.uuid': host_access_uuid
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
      var returnType = V1HostAccess;

      return this.apiClient.callApi(
        '/api/v1/{owner}/git_accesses/{host_access.uuid}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the updateGitAccess operation.
     * @callback module:api/GitAccessesV1Api~updateGitAccessCallback
     * @param {String} error Error message, if any.
     * @param {module:model/V1HostAccess} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Get run
     * @param {String} owner Owner of the namespace
     * @param {String} host_access_uuid UUID
     * @param {module:model/V1HostAccess} body Artifact store body
     * @param {module:api/GitAccessesV1Api~updateGitAccessCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/V1HostAccess}
     */
    this.updateGitAccess = function(owner, host_access_uuid, body, callback) {
      var postBody = body;

      // verify the required parameter 'owner' is set
      if (owner === undefined || owner === null) {
        throw new Error("Missing the required parameter 'owner' when calling updateGitAccess");
      }

      // verify the required parameter 'host_access_uuid' is set
      if (host_access_uuid === undefined || host_access_uuid === null) {
        throw new Error("Missing the required parameter 'host_access_uuid' when calling updateGitAccess");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling updateGitAccess");
      }


      var pathParams = {
        'owner': owner,
        'host_access.uuid': host_access_uuid
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
      var returnType = V1HostAccess;

      return this.apiClient.callApi(
        '/api/v1/{owner}/git_accesses/{host_access.uuid}', 'PUT',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
