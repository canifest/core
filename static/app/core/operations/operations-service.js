angular.module('canifestApp.core.operations', [
])

/**
 * handles operation services
 */
.factory('operations', function operations($http,$q) {

  /**
   * Retrieves the list of available operations
   * @returns {promise} promise that will resolve  with the list of available operations
   */
  function getOperations() {
    var deferred = $q.defer();

    $http.get('api/list' )
    .then(function(data){
      deferred.resolve(data);
    })
    .catch(function(err){
      deferred.reject(err);
    });

    return deferred.promise;
  }

  return {
    /**
    * Retrieves the list of available operations
     */
    getOperations: getOperations
  };
});
