angular.module('canifestApp.sections.operations', []).
controller('OperationsCtrl',
    ['$scope', '$http','operations',
    function($scope, $http,operations){
      var vm = this;
      operations.getOperations().then(function(resp){
        vm.operations = resp.data;
      });
    }
]);
