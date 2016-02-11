var operationsServices = angular.module('operationsServices', ['ngResource']);

// This isn't fully baked right now...
operationsServices.factory('Operations', ['$resource',
    function($resource){
        return $resource('api/list');
    }
]);

