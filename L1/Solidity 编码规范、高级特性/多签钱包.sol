// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract MultiSigWallet {
    address[] public owners;
    mapping(address => bool) public isOwner;
    uint256 public required;

    struct Transaction {
        address to;
        uint256 value;
        bytes data;
        bool executed;
        string description;
    }

    Transaction[] public transactions;
    mapping(uint256 => mapping(address => bool)) public approved;
    mapping(uint256 => uint256) public approvalCount;

    event Deposit(address indexed sender, uint256 amount);
    event Submit(uint256 indexed txId);
    event Approve(address indexed owner, uint256 indexed txId);
    event Revoke(address indexed owner, uint256 indexed txId);
    event Execute(uint256 indexed txId);

    receive() external payable {
        emit Deposit(msg.sender, msg.value);
    }

    modifier onlyOwner() {
        require(isOwner[msg.sender], "not owner");
        _;
    }

    modifier txExists(uint256 _txId) {
        require(_txId < transactions.length, "tx doesn't exist");
        _;
    }

    modifier notApproved(uint256 _txId) {
        require(!approved[_txId][msg.sender], "tx already approved");
        _;
    }

    modifier notExecuted(uint256 _txId) {
        require(!transactions[_txId].executed, "tx is executed");
        _;
    }

    constructor(address[] memory _owners, uint256 _required) {
        require(_owners.length > 0, "owner required");
        require(
            _required > 0 && _required <= _owners.length,
            "invalid required number of owners"
        );
        for (uint256 index = 0; index < _owners.length; index++) {
            address owner = _owners[index];
            require(owner != address(0), "invalid owner");
            require(!isOwner[owner], "owner is not unique");
            isOwner[owner] = true;
            owners.push(owner);
        }
        required = _required;
    }

    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }

    function submit(address _to, uint256 _value, bytes calldata _data, string memory _description) external onlyOwner returns(uint256) {
        transactions.push(
            Transaction({to: _to, value: _value, data: _data, executed: false, description: _description})
        );
        emit Submit(transactions.length - 1);
        return transactions.length - 1;
    }

    function approve(uint256 _txId) external onlyOwner txExists(_txId) notApproved(_txId) notExecuted(_txId) {
        approved[_txId][msg.sender] = true;
        approvalCount[_txId] += 1;
        emit Approve(msg.sender, _txId);
    }

    function execute(uint256 _txId) external onlyOwner txExists(_txId) notExecuted(_txId) {
        require(approvalCount[_txId] >= required, "approvals < required");
        Transaction storage transaction = transactions[_txId];
        transaction.executed = true;
        (bool success, ) = transaction.to.call{value: transaction.value}(transaction.data);
        require(success, "tx failed");
        emit Execute(_txId);
    }

    function revoke(uint256 _txId) external onlyOwner txExists(_txId) notExecuted(_txId) {
        require(approved[_txId][msg.sender], "tx not approved");
        approved[_txId][msg.sender] = false;
        approvalCount[_txId] -= 1;
        emit Revoke(msg.sender, _txId);
    }
}