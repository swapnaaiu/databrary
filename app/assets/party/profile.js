'use strict';

app.controller('party/profile', [
  '$scope', 'displayService', 'party', 'pageService','constantService',
  function ($scope, display, party, page, constants) {

    var getUsers = function(volumes){

      var users = {};
      
      users.sponsors = [];
      users.labGroupMembers = [];
      users.nonGroupAffiliates = [];
      users.otherCollaborators = [];

      _(volumes).pluck('access').flatten().each(function(v){
        if(v.children > 0 ){
          users.sponsors.push(v);
        } else if(v.parents && v.parents.length > 0 && v.party.members && v.party.members.length > 0) {
          users.labGroupMembers.push(v);
        } else if(v.parents && v.parents.length > 0){
          users.nonGroupAffiliates.push(v);
        } else {
          users.otherCollaborators.push(v);
        }

      }).value();

      return users;
    };

    // This is a quick helper function to make sure that the isSelected
    // class is set to empty and to avoid repeating code. 
    var unSetSelected = function(v){
      v.isSelected = '';
      return v; 
    };
      
    
    
    $scope.clickVolume = function(volume) {

      $scope.volumes.individual = _.map($scope.volumes.individual, unSetSelected);
      $scope.volumes.collaborator = _.map($scope.volumes.collaborator, unSetSelected);
      $scope.volumes.inherited = _.map($scope.volumes.inherited, unSetSelected);

      for(var i = 0; i < volume.access.length; i += 1 ){
        for (var j = 0; j < $scope.users.sponsors.length; j += 1){
          if($scope.users.sponsors[j].id === volume.access[i].party.id){
            $scope.users.sponsors[j].isSelected = 'userSelected';
            return; 
          }
        }
        for (var j = 0; j < $scope.users.labGroupMembers.length; j += 1){
          if($scope.users.labGroupMembers[j].id === volume.access[i].party.id){
            $scope.users.labGroupMembers[j].isSelected = 'userSelected';
            return; 
          }
        }
        for (var j = 0; j < $scope.users.nonGroupAffiliates.length; j += 1){
          if($scope.users.nonGroupAffiliates[j].id === volume.access[i].party.id){
            $scope.users.nonGroupAffiliates[j].isSelected = 'userSelected';
            return; 
          }
        }
        for (var j = 0; j < $scope.users.otherCollaborators.length; j += 1){
          if($scope.users.otherCollaborators[j].id === volume.access[i].party.id){
            $scope.users.otherCollaborators[j].isSelected = 'userSelected';
            return; 
          }
        }        
      }
    };

    // This should take in a user, then select volumes on each thing. 
    $scope.clickUser = function(user){

      // Make sure that all the selections have been cleared out. 
      $scope.volumes.individual = _.map($scope.volumes.individual, unSetSelected);
      $scope.volumes.collaborator = _.map($scope.volumes.collaborator, unSetSelected);
      $scope.volumes.inherited = _.map($scope.volumes.inherited, unSetSelected);
      var iterrateVolume = function(_item, i, volumeArray){
        for(var j = 0; j < volumeArray[i].access.length; j += 1){
          if(volumeArray[i].access[j] == user){
            volumeArray[i].access[j].isSelected = 'volumeSelected';
            return false; 
          }
        }
      };

      _.each($scope.volumes.individual, iterrateVolume);
      _.each($scope.volumes.collaborator, iterrateVolume);
      _.each($scope.volumes.inherited, iterrateVolume);

    };

    var getVolumes = function(volumes) {
      var tempVolumes = {};
      tempVolumes.individual = [];
      tempVolumes.collaborator = [];
      tempVolumes.inherited = [];

      _.each(volumes, function(v){
        if(v.isIndividual){
          tempVolumes.individual.push(v); 
	} else if(tempVolumes.isCollaborator){
          tempVolumes.collaborator.push(v); 
	} else{
          tempVolumes.inherited.push(v); 
	}
      });
      return tempVolumes;
    };

    $scope.party = party;
    $scope.volumes = getVolumes(party.volumes);
    console.log(party.volumes);
    $scope.users = getUsers(party.volumes);  
    console.log($scope.users);
    $scope.page = page;
    $scope.profile = page.$location.path() === '/profile';
    display.title = party.name;
  }
]);
