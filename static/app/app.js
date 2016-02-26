angular.module('canifestApp', [
    'ui.router',
    'canifestApp.sections.operations',
    'canifestApp.core.operations',
    'canifestApp.constants'
])

.config(['$stateProvider', '$urlRouterProvider','ROUTES',
    function($stateProvider, $urlRouterProvider, ROUTES) {

      $urlRouterProvider.otherwise(ROUTES.OPERATION_STATE_URL);

      $stateProvider
      .state(ROUTES.OPERATION_STATE, {
        url: ROUTES.OPERATION_STATE_URL,
        templateUrl: 'app/sections/operations/operations.tpl.html',
        controller: 'OperationsCtrl as operationsCtrl'
      });
}]);
