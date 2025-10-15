pragma solidity ^0.8.26;

contract Store {
  event ItemSet(bytes32 key, bytes32 value);

  string public version;
  mapping (bytes32 => bytes32) public items;

  constructor(string memory _version) {
    version = _version;
  }

  function setItem(bytes32 key, bytes32 value) external {
    items[key] = value;
    emit ItemSet(key, value);
  }
}

//已调用setItem("0x0000000000000000000000000000000000000000000000000000000000000012", "0x00000000000000000000000000000000000000000000000000000000000000ab")