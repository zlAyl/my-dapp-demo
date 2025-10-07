// SPDX-License-Identifier: SEE LICENSE IN LICENSE
pragma solidity ^0.8;

contract Count {
    uint256 public count;

    constructor() {
        count = 1;
    }

    function AddCount() external  returns(uint256){
        count++;
        return count;
    }
}