//SPDX-License-Identifier: MIT
pragma solidity  ^0.8.23;

contract EtherUnits {
    uint public oneWei = 1 ;
    bool public isoneWei = (1 wei == 1); 

    bool public isOnegwei = (1 gwei == 1e9);

    bool public isOneether = (1 ether == 1e18);

}