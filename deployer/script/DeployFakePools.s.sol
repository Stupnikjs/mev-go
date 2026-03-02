// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script} from "forge-std/Script.sol";
import {stdJson} from "forge-std/StdJson.sol";
import "v3-core/contracts/interfaces/IUniswapV3Factory.sol";
import "v3-periphery/contracts/interfaces/INonfungiblePositionManager.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MockToken is ERC20 {
    constructor(string memory name, string memory symbol) ERC20(name, symbol) {
        _mint(msg.sender, 1_000_000 ether);
    }
}

contract DeployFakePools is Script {
    using stdJson for string;
    address factoryAddress;
    PoolConfig config;  
    struct PoolConfig {
        address tokenA;
        address tokenB;
        uint256 fee; // On utilise uint256 ici car le JSON parse souvent en 256
    }

    function setUp() public {
       
        string memory root = vm.projectRoot();
        string memory path = string.concat(root, "/config.json");
        string memory json = vm.readFile(path);
        factoryAddress = vm.parseJsonAddress(json, ".sepolia_v3_factory");
        managerAddress = vm.parseJsonAddress(json, ".sepolia_v3_position_manager");
        bytes memory data = json.parseRaw(".pool1");
        config = abi.decode(data, (PoolConfig));
        
    }

    function run() public {
        vm.startBroadcast();
        // Utilisation de l'interface officielle
        IUniswapV3Factory factory = IUniswapV3Factory(factoryAddress);
        address pool = factory.createPool(config.tokenA, config.tokenB, uint24(3000));
        IUniswapV3Pool(pool).initialize(79228162514264337593543950336);

        MockToken token0 = new MockToken("Fake Tether", "fUSDT");
        MockToken token1 = new MockToken("Fake Wrapped ETH", "fWETH");

        IERC20(token1).approve(managerAddress, 1000 ether);
        IERC20(token0).approve(managerAddress, 1000 ether);

        // 4. Ajout de liquidité (Mint)
        INonfungiblePositionManager pm = INonfungiblePositionManager(managerAddress);
        
        // Tri des tokens (Obligatoire : token0 < token1)
        (address t0, address t1) = token0 < token1 
            ? (token0, token1) 
            : (token1, token0);

        pm.mint(INonfungiblePositionManager.MintParams({
            token0: t0,
            token1: t1,
            fee: uint24(config.fee),
            tickLower: -887220,
            tickUpper: 887220,
            amount0Desired: 10 ether,
            amount1Desired: 10 ether,
            amount0Min: 0,
            amount1Min: 0,
            recipient: vm.addr(deployerPrivateKey),
            deadline: block.timestamp
        }));

        vm.stopBroadcast();
    }
}
