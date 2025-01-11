//SPDX-License-Identifier: MIT
pragma solidity  ^0.8.26;
contract Example{
    function get_balance ()public view returns(uint){
        return address(this).balance;
    }

    function get_code()public view returns(bytes memory){
        return address(this).code;
    }

    function get_codehash()public view returns(bytes32){
        return address(this).codehash;
    }

    function get_contract_address() public view returns (address) {
        return address(this);
    }
}
