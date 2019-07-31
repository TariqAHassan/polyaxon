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

/**
 * Experiment service
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.7
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.ExperimentService);
  }
}(this, function(expect, ExperimentService) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new ExperimentService.ExperimentServiceApi();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('ExperimentServiceApi', function() {
    describe('archiveExperiment', function() {
      it('should call archiveExperiment successfully', function(done) {
        //uncomment below and update the code to test archiveExperiment
        //instance.archiveExperiment(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('bookmarkExperiment', function() {
      it('should call bookmarkExperiment successfully', function(done) {
        //uncomment below and update the code to test bookmarkExperiment
        //instance.bookmarkExperiment(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('createExperiment', function() {
      it('should call createExperiment successfully', function(done) {
        //uncomment below and update the code to test createExperiment
        //instance.createExperiment(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('createExperimentStatus', function() {
      it('should call createExperimentStatus successfully', function(done) {
        //uncomment below and update the code to test createExperimentStatus
        //instance.createExperimentStatus(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteExperiment', function() {
      it('should call deleteExperiment successfully', function(done) {
        //uncomment below and update the code to test deleteExperiment
        //instance.deleteExperiment(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteExperiments', function() {
      it('should call deleteExperiments successfully', function(done) {
        //uncomment below and update the code to test deleteExperiments
        //instance.deleteExperiments(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getExperiment', function() {
      it('should call getExperiment successfully', function(done) {
        //uncomment below and update the code to test getExperiment
        //instance.getExperiment(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getExperimentCodeRef', function() {
      it('should call getExperimentCodeRef successfully', function(done) {
        //uncomment below and update the code to test getExperimentCodeRef
        //instance.getExperimentCodeRef(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('greateExperimentCodeRef', function() {
      it('should call greateExperimentCodeRef successfully', function(done) {
        //uncomment below and update the code to test greateExperimentCodeRef
        //instance.greateExperimentCodeRef(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listArchivedExperiments', function() {
      it('should call listArchivedExperiments successfully', function(done) {
        //uncomment below and update the code to test listArchivedExperiments
        //instance.listArchivedExperiments(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listBookmarkedExperiments', function() {
      it('should call listBookmarkedExperiments successfully', function(done) {
        //uncomment below and update the code to test listBookmarkedExperiments
        //instance.listBookmarkedExperiments(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listExperimentStatuses', function() {
      it('should call listExperimentStatuses successfully', function(done) {
        //uncomment below and update the code to test listExperimentStatuses
        //instance.listExperimentStatuses(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listExperiments', function() {
      it('should call listExperiments successfully', function(done) {
        //uncomment below and update the code to test listExperiments
        //instance.listExperiments(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('restartExperiment', function() {
      it('should call restartExperiment successfully', function(done) {
        //uncomment below and update the code to test restartExperiment
        //instance.restartExperiment(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('restoreExperiment', function() {
      it('should call restoreExperiment successfully', function(done) {
        //uncomment below and update the code to test restoreExperiment
        //instance.restoreExperiment(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('resumeExperiment', function() {
      it('should call resumeExperiment successfully', function(done) {
        //uncomment below and update the code to test resumeExperiment
        //instance.resumeExperiment(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('startExperimentTensorboard', function() {
      it('should call startExperimentTensorboard successfully', function(done) {
        //uncomment below and update the code to test startExperimentTensorboard
        //instance.startExperimentTensorboard(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('stopExperiment', function() {
      it('should call stopExperiment successfully', function(done) {
        //uncomment below and update the code to test stopExperiment
        //instance.stopExperiment(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('stopExperimentTensorboard', function() {
      it('should call stopExperimentTensorboard successfully', function(done) {
        //uncomment below and update the code to test stopExperimentTensorboard
        //instance.stopExperimentTensorboard(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('stopExperiments', function() {
      it('should call stopExperiments successfully', function(done) {
        //uncomment below and update the code to test stopExperiments
        //instance.stopExperiments(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('unBookmarkExperiment', function() {
      it('should call unBookmarkExperiment successfully', function(done) {
        //uncomment below and update the code to test unBookmarkExperiment
        //instance.unBookmarkExperiment(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateExperiment2', function() {
      it('should call updateExperiment2 successfully', function(done) {
        //uncomment below and update the code to test updateExperiment2
        //instance.updateExperiment2(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
