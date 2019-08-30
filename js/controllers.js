'use strict'

var helloWorldControllers = angular.module('helloWorldControllers', [])

.controller('MainCtrl', ['$scope', '$location', '$http', function MainCtrl($scope, $location, $http){
    $scope.message = "Test Main controller";
    $scope.addCustomer = function() {
        $location.path('/addcustomer');
    };
    $scope.showhtml = function(){
        $location.path('/show');
    };

}])

.controller('ShowCtrl', ['$scope', '$location', '$http', function ShowCtrl($scope, $location, $http){
    $scope.showMessage = "Test Show controller";
    $scope.mainBack = function() {
        $location.path('/');
    }
}])

.controller('CustomCtrl', ['$scope', '$location', '$http', function CustomCtrl($scope, $location, $http){
    $scope.cName = "Armin Eminagic";
    $scope.cPhone = "123456";
    
    $scope.changeCustomer = function(){
        $scope.cName = $scope.iName;
        $scope.cPhone = $scope.iPhone;
        
    };
    $scope.mainBack = function() {
        $location.path('/');
    }
}])

.controller('AddCustomerCtrl', ['$scope', '$location', function AddCustomerCtrl($scope, $location){
    $scope.submit = function(){
        $location.path('/addedCustomer/' + $scope.cName + "/" + $scope.cCity);
    };

    $scope.mainBack = function() {
        $location.path('/');
    }
}])

.controller('AddedCustomerCtrl', ['$scope', '$routeParams', '$location', function AddCustomerCtrl($scope, $routeParams, $location){
    $scope.name = $routeParams.customer;
    $scope.city = $routeParams.city;
    $scope.mainBack = function(){
        $location.path('/');
    };
}]);





