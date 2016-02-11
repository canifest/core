var canifestApp = angular.module('canifestApp', [
    'ngRoute',
    'operationsServices',
    'canifestControllers'
    ]);

canifestApp.config(['$routeProvider', '$httpProvider',
    function($routeProvider, $httpProvider) {
        $routeProvider.
            when('/operations', {
                templateUrl: 'operations.html',
                controller: 'OperationsCtrl'
            }).
            otherwise({
                redirectTo: '/operations'
            });

//        $httpProvider.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
    }]);