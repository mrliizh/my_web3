## Solidity 编码规范、高级特性

- 所有人都可以存钱
  - ETH
- 只有合约 owner 才可以取钱
- 只要取钱，合约就销毁掉 selfdestruct
- 扩展：支持主币以外的资产
  - ERC20
  - ERC721

```js
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;
contract Bank {
    // 状态变量
    address public immutable owner;
    // 事件
    event Deposit(address _ads, uint256 amount);
    event Withdraw(uint256 amount);
    // receive
    receive() external payable {
        emit Deposit(msg.sender, msg.value);
    }
    // 构造函数
    constructor() payable {
        owner = msg.sender;
    }
    // 方法
    function withdraw() external {
        require(msg.sender == owner, "Not Owner");
        emit Withdraw(address(this).balance);
        selfdestruct(payable(msg.sender));
    }
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
}
```



WETH 是包装 ETH 主币，作为 ERC20 的合约。 标准的 ERC20 合约包括如下几个

- 3 个查询
  - balanceOf: 查询指定地址的 Token 数量
  - allowance: 查询指定地址对另外一个地址的剩余授权额度
  - totalSupply: 查询当前合约的 Token 总量
- 2 个交易
  - transfer: 从当前调用者地址发送指定数量的 Token 到指定地址。
    - 这是一个写入方法，所以还会抛出一个 Transfer 事件。
  - transferFrom: 当向另外一个合约地址存款时，对方合约必须调用 transferFrom 才可以把 Token 拿到它自己的合约中。
- 2 个事件
  - Transfer
  - Approval
- 1 个授权
  - approve: 授权指定地址可以操作调用者的最大 Token 数量。

```js
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;
contract WETH {
    string public name = "Wrapped Ether";
    string public symbol = "WETH";
    uint8 public decimals = 18;
    event Approval(address indexed src, address indexed delegateAds, uint256 amount);
    event Transfer(address indexed src, address indexed toAds, uint256 amount);
    event Deposit(address indexed toAds, uint256 amount);
    event Withdraw(address indexed src, uint256 amount);
    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;
    function deposit() public payable {
        balanceOf[msg.sender] += msg.value;
        emit Deposit(msg.sender, msg.value);
    }
    function withdraw(uint256 amount) public {
        require(balanceOf[msg.sender] >= amount);
        balanceOf[msg.sender] -= amount;
        payable(msg.sender).transfer(amount);
        emit Withdraw(msg.sender, amount);
    }
    function totalSupply() public view returns (uint256) {
        return address(this).balance;
    }
    function approve(address delegateAds, uint256 amount) public returns (bool) {
        allowance[msg.sender][delegateAds] = amount;
        emit Approval(msg.sender, delegateAds, amount);
        return true;
    }
    function transfer(address toAds, uint256 amount) public returns (bool) {
        return transferFrom(msg.sender, toAds, amount);
    }
    function transferFrom(
        address src,
        address toAds,
        uint256 amount
    ) public returns (bool) {
        require(balanceOf[src] >= amount);
        if (src != msg.sender) {
            require(allowance[src][msg.sender] >= amount);
            allowance[src][msg.sender] -= amount;
        }
        balanceOf[src] -= amount;
        balanceOf[toAds] += amount;
        emit Transfer(src, toAds, amount);
        return true;
    }
    fallback() external payable {
        deposit();
    }
    receive() external payable {
        deposit();
    }
}
```

TodoList: 是类似便签一样功能的东西，记录我们需要做的事情，以及完成状态。 1.需要完成的功能

- 创建任务
- 修改任务名称
  - 任务名写错的时候
- 修改完成状态：
  - 手动指定完成或者未完成
  - 自动切换
    - 如果未完成状态下，改为完成
    - 如果完成状态，改为未完成
- 获取任务 2.思考代码内状态变量怎么安排？ 思考 1：思考任务 ID 的来源？ 我们在传统业务里，这里的任务都会有一个任务 ID，在区块链里怎么实现？？ 答：传统业务里，ID 可以是数据库自动生成的，也可以用算法来计算出来的，比如使用雪花算法计算出 ID 等。在区块链里我们使用数组的 index 索引作为任务的 ID，也可以使用自增的整型数据来表示。 思考 2: 我们使用什么数据类型比较好？ 答：因为需要任务 ID，如果使用数组 index 作为任务 ID。则数据的元素内需要记录任务名称，任务完成状态，所以元素使用 struct 比较好。 如果使用自增的整型作为任务 ID，则整型 ID 对应任务，使用 mapping 类型比较符合。 3.演示代码

```js
contract Demo {
    struct Todo {
        string name;
        bool isCompleted;
    }
    Todo[] public list; // 29414
    // 创建任务
    function create(string memory name_) external {
        list.push(
            Todo({
                name:name_, // ,
                isCompleted:false
            })
        );
    }
    // 修改任务名称
    function modiName1(uint256 index_,string memory name_) external {
        // 方法1: 直接修改，修改一个属性时候比较省 gas
        list[index_].name = name_;
    }
    function modiName2(uint256 index_,string memory name_) external {
        // 方法2: 先获取储存到 storage，在修改，在修改多个属性的时候比较省 gas
        Todo storage temp = list[index_];
        temp.name = name_;
    }
    // 修改完成状态1:手动指定完成或者未完成
    function modiStatus1(uint256 index_,bool status_) external {
        list[index_].isCompleted = status_;
    }
    // 修改完成状态2:自动切换 toggle
    function modiStatus2(uint256 index_) external {
        list[index_].isCompleted = !list[index_].isCompleted;
    }
    // 获取任务1: memory : 2次拷贝
    // 29448 gas
    function get1(uint256 index_) external view
        returns(string memory name_,bool status_){
        Todo memory temp = list[index_];
        return (temp.name,temp.isCompleted);
    }
    // 获取任务2: storage : 1次拷贝
    // 预期：get2 的 gas 费用比较低（相对 get1）
    // 29388 gas
    function get2(uint256 index_) external view
        returns(string memory name_,bool status_){
        Todo storage temp = list[index_];
        return (temp.name,temp.isCompleted);
    }
}
```

众筹合约是一个募集资金的合约，在区块链上，我们是募集以太币，类似互联网业务的水滴筹。区块链早起的 ICO 就是类似业务。

### 1.需求分析



众筹合约分为两种角色：一个是受益人，一个是资助者。

```
// 两种角色:
//      受益人   beneficiary => address         => address 类型
//      资助者   funders     => address:amount  => mapping 类型 或者 struct 类型
```



```
状态变量按照众筹的业务：
// 状态变量
//      筹资目标数量    fundingGoal
//      当前募集数量    fundingAmount
//      资助者列表      funders
//      资助者人数      fundersKey
```



```
需要部署时候传入的数据:
//      受益人
//      筹资目标数量
```



### 2.演示代码

```
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;
contract CrowdFunding {
    address public immutable beneficiary;   // 受益人
    uint256 public immutable fundingGoal;   // 筹资目标数量
    uint256 public fundingAmount;       // 当前的 金额
    mapping(address=>uint256) public funders;
    // 可迭代的映射
    mapping(address=>bool) private fundersInserted;
    address[] public fundersKey; // length
    // 不用自销毁方法，使用变量来控制
    bool public AVAILABLED = true; // 状态
    // 部署的时候，写入受益人+筹资目标数量
    constructor(address beneficiary_,uint256 goal_){
        beneficiary = beneficiary_;
        fundingGoal = goal_;
    }
    // 资助
    //      可用的时候才可以捐
    //      合约关闭之后，就不能在操作了
 function contribute() external payable {
        require(AVAILABLED, "CrowdFunding is closed");

        // 检查捐赠金额是否会超过目标金额
        uint256 potentialFundingAmount = fundingAmount + msg.value;
        uint256 refundAmount = 0;

        if (potentialFundingAmount > fundingGoal) {
            refundAmount = potentialFundingAmount - fundingGoal;
            funders[msg.sender] += (msg.value - refundAmount);
            fundingAmount += (msg.value - refundAmount);
        } else {
            funders[msg.sender] += msg.value;
            fundingAmount += msg.value;
        }

        // 更新捐赠者信息
        if (!fundersInserted[msg.sender]) {
            fundersInserted[msg.sender] = true;
            fundersKey.push(msg.sender);
        }

        // 退还多余的金额
        if (refundAmount > 0) {
            payable(msg.sender).transfer(refundAmount);
        }
    }
    // 关闭
    function close() external returns(bool){
        // 1.检查
        if(fundingAmount<fundingGoal){
            return false;
        }
        uint256 amount = fundingAmount;
        // 2.修改
        fundingAmount = 0;
        AVAILABLED = false;
        // 3. 操作
        payable(beneficiary).transfer(amount);
        return true;
    }
    function fundersLenght() public view returns(uint256){
        return fundersKey.length;
    }
}
```



上面的合约只是一个简化版的 众筹合约，但它已经足以让我们理解本节介绍的类型概念。

这一个实战主要是加深大家对 3 个取钱方法的使用。

- 任何人都可以发送金额到合约
- 只有 owner 可以取款
- 3 种取钱方式

```js
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;
contract EtherWallet {
    address payable public immutable owner;
    event Log(string funName, address from, uint256 value, bytes data);
    constructor() {
        owner = payable(msg.sender);
    }
    receive() external payable {
        emit Log("receive", msg.sender, msg.value, "");
    }
    function withdraw1() external {
        require(msg.sender == owner, "Not owner");
        // owner.transfer 相比 msg.sender 更消耗Gas
        // owner.transfer(address(this).balance);
        payable(msg.sender).transfer(100);
    }
    function withdraw2() external {
        require(msg.sender == owner, "Not owner");
        bool success = payable(msg.sender).send(200);
        require(success, "Send Failed");
    }
    function withdraw3() external {
        require(msg.sender == owner, "Not owner");
        (bool success, ) = msg.sender.call{value: address(this).balance}("");
        require(success, "Call Failed");
    }
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
}
```

多签钱包的功能: 合约有多个 owner，一笔交易发出后，需要多个 owner 确认，确认数达到最低要求数之后，才可以真正的执行。
### 1.原理
- 部署时候传入地址参数和需要的签名数
  - 多个 owner 地址
  - 发起交易的最低签名数
- 有接受 ETH 主币的方法，
- 除了存款外，其他所有方法都需要 owner 地址才可以触发
- 发送前需要检测是否获得了足够的签名数
- 使用发出的交易数量值作为签名的凭据 ID（类似上么）
- 每次修改状态变量都需要抛出事件
- 允许批准的交易，在没有真正执行前取消。
- 足够数量的 approve 后，才允许真正执行。
### 2.代码
```js
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;
contract MultiSigWallet {
    // 状态变量
    address[] public owners;
    mapping(address => bool) public isOwner;
    uint256 public required;
    struct Transaction {
        address to;
        uint256 value;
        bytes data;
        bool exected;
    }
    Transaction[] public transactions;
    mapping(uint256 => mapping(address => bool)) public approved;
    // 事件
    event Deposit(address indexed sender, uint256 amount);
    event Submit(uint256 indexed txId);
    event Approve(address indexed owner, uint256 indexed txId);
    event Revoke(address indexed owner, uint256 indexed txId);
    event Execute(uint256 indexed txId);
    // receive
    receive() external payable {
        emit Deposit(msg.sender, msg.value);
    }
    // 函数修改器
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
        require(!transactions[_txId].exected, "tx is exected");
        _;
    }
    // 构造函数
    constructor(address[] memory _owners, uint256 _required) {
        require(_owners.length > 0, "owner required");
        require(
            _required > 0 && _required <= _owners.length,
            "invalid required number of owners"
        );
        for (uint256 index = 0; index < _owners.length; index++) {
            address owner = _owners[index];
            require(owner != address(0), "invalid owner");
            require(!isOwner[owner], "owner is not unique"); // 如果重复会抛出错误
            isOwner[owner] = true;
            owners.push(owner);
        }
        required = _required;
    }
    // 函数
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
    function submit(
        address _to,
        uint256 _value,
        bytes calldata _data
    ) external onlyOwner returns(uint256){
        transactions.push(
            Transaction({to: _to, value: _value, data: _data, exected: false})
        );
        emit Submit(transactions.length - 1);
        return transactions.length - 1;
    }
    function approv(uint256 _txId)
        external
        onlyOwner
        txExists(_txId)
        notApproved(_txId)
        notExecuted(_txId)
    {
        approved[_txId][msg.sender] = true;
        emit Approve(msg.sender, _txId);
    }
    function execute(uint256 _txId)
        external
        onlyOwner
        txExists(_txId)
        notExecuted(_txId)
    {
        require(getApprovalCount(_txId) >= required, "approvals < required");
        Transaction storage transaction = transactions[_txId];
        transaction.exected = true;
        (bool sucess, ) = transaction.to.call{value: transaction.value}(
            transaction.data
        );
        require(sucess, "tx failed");
        emit Execute(_txId);
    }
    function getApprovalCount(uint256 _txId)
        public
        view
        returns (uint256 count)
    {
        for (uint256 index = 0; index < owners.length; index++) {
            if (approved[_txId][owners[index]]) {
                count += 1;
            }
        }
    }
    function revoke(uint256 _txId)
        external
        onlyOwner
        txExists(_txId)
        notExecuted(_txId)
    {
        require(approved[_txId][msg.sender], "tx not approved");
        approved[_txId][msg.sender] = false;
        emit Revoke(msg.sender, _txId);
    }
}
```