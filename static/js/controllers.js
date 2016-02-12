var canifestControllers = angular.module('canifestControllers', []);

canifestControllers.controller('OperationsCtrl', ['$scope', '$http',


    function($scope, $http){

          $http.get('api/list' )
            .success(function(data) {
                $scope.operations = data;
            });

    }

]);