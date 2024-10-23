//SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;
 
interface INativeMinter {
  event NativeCoinMinted(address indexed sender, address indexed recipient, uint256 amount);
 
  function mintNativeCoin(address addr, uint256 amount) external;
}

contract NativeCoinMint {
    function mint(address addr, uint256 amount) external {
        INativeMinter(0x0200000000000000000000000000000000000001).mintNativeCoin(addr,amount * 10 ** 18);
    }
}