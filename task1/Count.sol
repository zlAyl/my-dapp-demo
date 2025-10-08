// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

contract Count {
    uint256 public count;
    event AddCount(address indexed  sender,uint256 nextCount);
    constructor() {
        count = 1;
    }

    function addCount() external  returns(uint256){
        count++;
        emit AddCount(msg.sender,count);
        return count;
    }

}