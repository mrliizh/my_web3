pragma solidity >=0.4.25 <0.9.0; 
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract RccToken is ERC20{
    constructor() ERC20("RccToken", "Rcc"){
        _mint(msg.sender, 20000000*1_000_000_000_000_000_000);
    }
}