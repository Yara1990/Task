pragma solidity ^0.5.0;

contract Sampel {
    mapping(string => address) userData;
    
    function setUser(string memory _name ) public returns (address){
        userData[_name] = msg.sender;
        return (userData[_name]);
    }
    
    function getUser(string memory _name) public view returns(address){
        return (userData[_name] );
    }
    
    
}
