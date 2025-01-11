//SPDX-License-Identifier: MIT
pragma solidity  ^0.8.23;

contract Coutner{
    uint public count;

    function store(uint num) public{
        count = num;
    } 

    function get() public view returns (uint){
        return count;
    }
}