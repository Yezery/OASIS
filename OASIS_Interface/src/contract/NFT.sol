// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
contract NFT is ERC721, IERC721Enumerable, Ownable {
    uint256 public _maximums; // 最大铸造量
    uint256 public _currentId; //当前tokenId
    string private _baseTokenURI; // IPFS网关
    mapping(uint256 => string) public  _nftMetaData; // tokenId 映射 NFT元数据
    
    //====================  工厂
    constructor(
        string memory name, // 名称
        string memory symbol, // 代币
        string memory baseURI, // IPFS网关
        uint256 maximums // 最大铸造量
    ) ERC721(name, symbol) {
        _baseTokenURI = baseURI;
        _maximums = maximums;
    }

    //====================  铸币
    function mint(
        string memory nftName,
        string memory description,
        string memory ipfsHash
    ) public onlyOwner {
        require(
            totalSupply() < _maximums,
            "The number of mints exceeds the maximum limit"
        );
        uint256 newTokenId = _currentId;
        _safeMint(msg.sender, newTokenId);
        // 设置NFT图片路径
        string memory image = string(
            abi.encodePacked(
                _baseURI(),
                "ipfs/",
                ipfsHash,
                "?filename=",
                nftName
            )
        );
         // 构造JSON格式元数据
        string memory metadata = string(
            abi.encodePacked(
                 "{\"name\":\"",
                nftName,
                "\",\"description\":\"",
                description,
                "\",\"image\":\"",
                image,
                "\"}"
            )
        );
        // 添加映射
        _nftMetaData[newTokenId] = metadata;
        _currentId++;
    }
    
    //====================  赠送NFT
    function giveMint(
        address to,
        string memory nftName,
        string memory description,
        string memory ipfsHash
    ) public onlyOwner {
        require(
            totalSupply() < _maximums,
            "The number of mints exceeds the maximum limit"
        );
        uint256 newTokenId = _currentId;
        _safeMint(to, newTokenId);
        // 设置NFT图片路径
        string memory image = string(
            abi.encodePacked(
                _baseURI(),
                "ipfs/",
                ipfsHash,
                "?filename=",
                nftName
            )
        );
        // 构造JSON格式元数据
        string memory metadata = string(
            abi.encodePacked(
                 "{\"name\":\"",
                nftName,
                "\",\"description\":\"",
                description,
                "\",\"image\":\"",
                image,
                "\"}"
            )
        );
        // 添加映射
        _nftMetaData[newTokenId] = metadata;
        _currentId++;
    }

    //====================  通过 tokenId 获取 NFT元数据 ERC721 Override
    function tokenURI(uint256 tokenId)
        public
        view
        virtual
        override
        returns (string memory)
    {
        require(_exists(tokenId), "ERC721Metadata: nonexistent token");

        return _nftMetaData[tokenId];
    }

    //====================  ERC721 Override
    function _baseURI() internal view virtual override returns (string memory) {
        return _baseTokenURI;
    }

    //====================  ERC721 Override
    function totalSupply() public view virtual override returns (uint256) {
        return _currentId;
    }

    //====================  ERC721 Override
    function tokenByIndex(uint256 index)
        external
        view
        virtual
        override
        returns (uint256)
    {
        require(index < _currentId, "Index out of bounds");
        return index;
    }

    //====================  ERC721 Override
    function tokenOfOwnerByIndex(address owner, uint256 index)
        external
        view
        virtual
        override
        returns (uint256)
    {
        require(index < balanceOf(owner), "Index out of bounds");
        return index;
    }

    //====== 通用
    //====================  查看个人NFT列表信息
    function getUserNFTs(address userAddress, address _nftAddress)
        external
        view
        returns (uint256[] memory)
    {
        IERC721Enumerable nft = IERC721Enumerable(_nftAddress);
        uint256 balance = nft.balanceOf(userAddress);
        uint256[] memory tokens = new uint256[](balance);
        for (uint256 i = 0; i < balance; i++) {
            tokens[i] = nft.tokenOfOwnerByIndex(userAddress, i);
        }
        return tokens;
    }

    // function ChangeBaseURI(string memory baseURI) public onlyOwner {
    //     _baseTokenURI = baseURI;
    // }
}
