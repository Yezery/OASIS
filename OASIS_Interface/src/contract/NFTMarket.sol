// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// 0x5b2D36d73D974f4b4f223EB8Cab82DAc69AD2EcD

//goerli: 0x27A0bf96E92EA917dD7f72797Ca93A9457b7375A
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "./NFT.sol";

contract NFTMarket is IERC721Receiver, Ownable {
    using SafeMath for uint256;
    using Address for address payable;
    address payable public marketplaceOwner;

    //  NFT商品结构体
    struct Sale {
        uint256 saleId;
        address seller;
        address nftContract;
        uint256 tokenId;
        uint256 price;
        bool isActive;
        string tokenURI;
    }

    //  手续费
    uint256 public feePercentage;

    //  NFT商品列表
    Sale[] private sales;

    //  当前商品Id
    uint256 public SaleId = sales.length;

    //  记录交易表
    mapping(uint256 => address) public tokenOwners;

    //====================  事件集
    event SaleCreated(
        uint256 saleId,
        address indexed seller,
        address indexed nftContract,
        uint256 tokenId,
        string tokenURI,
        uint256 price
    );
    event SaleCancelled(
        uint256 saleId,
        address indexed seller,
        address indexed nftContract,
        uint256 tokenId
    );
    event SaleCompleted(
        uint256 saleId,
        address indexed buyer,
        address indexed seller,
        address indexed nftContract,
        uint256 tokenId,
        uint256 price
    );
    event NFTReceived(
        address indexed operator,
        address indexed from,
        uint256 tokenId,
        string metadata
    );

    //====================  初始化构造：给市场合约设置交易手续费
    constructor(uint256 _feePercentage) payable {
        feePercentage = _feePercentage;
        marketplaceOwner = payable(msg.sender);
    }

    //==================== 修改手续费
    function changeFeePercentage(uint256 _newFeePercentage) public onlyOwner {
        feePercentage = _newFeePercentage;
    }

    //====================  safeTransferFrom回调
    function onERC721Received(
        address operator,
        address from,
        uint256 tokenId,
        bytes memory data
    ) public override returns (bytes4) {
        // 记录交易
        tokenOwners[tokenId] = from;
        string memory metadata = string(data);
        // 触发事件
        emit NFTReceived(operator, from, tokenId, metadata);

        // 更新合约状态
        // Sale memory sale = sales[tokenId];
        // sale.isActive = sale.isActive ;
        // sales[tokenId] = sale;

        // 返回ERC721Receiver接口的标准返回值
        return
            bytes4(
                keccak256("onERC721Received(address,address,uint256,bytes)")
            );
    }

    //====================  给市场合约充值gas
    function fund() public payable {
        require(
            msg.value > 0,
            "Marketplace: Funding amount must be greater than zero"
        );
    }

    //====================  上架
    function createSale(
        address nftContract,
        uint256 tokenId,
        uint256 price,
        string memory metadata
    ) external {
        require(
            price > 0 wei,
            "Marketplace: Price must be greater than zero wei"
        );
        IERC721 nft = IERC721(nftContract);
        NFT NFTcontract = NFT(nftContract);
        require(
            nft.ownerOf(tokenId) == msg.sender,
            "Marketplace: Only the owner can sell the NFT"
        );
        require(
            nft.getApproved(tokenId) == address(this),
            "Marketplace: Contract not approved for NFT"
        );

        uint256 saleId = sales.length;
        string memory tokenURI = NFTcontract.tokenURI(tokenId);
        bytes memory data = bytes(metadata);
        sales.push(
            Sale(
                saleId,
                msg.sender,
                nftContract,
                tokenId,
                price,
                true,
                tokenURI
            )
        );
        nft.safeTransferFrom(msg.sender, address(this), tokenId, data);
        emit SaleCreated(
            saleId,
            msg.sender,
            nftContract,
            tokenId,
            tokenURI,
            price
        );
    }

    //====================  下架
    function cancelSale(uint256 saleId) external {
        Sale storage sale = sales[saleId];
        require(sale.isActive, "Marketplace: Sale is inactive");
        require(
            msg.sender == sale.seller,
            "Marketplace: Only the seller can cancel the sale"
        );

        IERC721 nft = IERC721(sale.nftContract);
        sale.isActive = false;
        nft.safeTransferFrom(sale.seller, address(this), sale.tokenId);

        emit SaleCancelled(saleId, sale.seller, sale.nftContract, sale.tokenId);
    }

    //====================  购买
    modifier onlyUnsoldSales(uint256 saleId) {
        require(sales[saleId].isActive, "Marketplace: Sale is not active");
        _;
    }

    function buy(uint256 saleId) public payable onlyUnsoldSales(saleId) {
        Sale storage sale = sales[saleId];

        uint256 fee = (sale.price * feePercentage) / 1000;
        uint256 totalPrice = sale.price;
        require(msg.value >= totalPrice, "Marketplace: Insufficient funds");

        sale.isActive = false;
        require(
            msg.sender != sale.seller,
            "Marketplace: Seller cannot buy own NFT"
        );

        IERC721 nft = IERC721(sale.nftContract);
        require(
            nft.ownerOf(sale.tokenId) == address(this),
            "Marketplace: NFT to seller must Market"
        );
        nft.safeTransferFrom(address(this), msg.sender, sale.tokenId);
        // nft.safeTransferFrom(sale.seller, msg.sender, sale.tokenId);

        address payable seller = payable(sale.seller);
        seller.transfer(totalPrice - fee);

        if (msg.value > totalPrice) {
            payable(msg.sender).transfer(msg.value - totalPrice);
        }

        emit SaleCompleted(
            saleId,
            msg.sender,
            sale.seller,
            sale.nftContract,
            sale.tokenId,
            sale.price
        );
    }

    //====================  查看当前市场合约余额
    function getMarketBalance() public view returns (uint256) {
        return address(this).balance;
    }

    //====================  查看上架NFT列表信息
    function getSales() external view returns (Sale[] memory) {
        return sales;
    }

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

    //==================== 市场合约拥有者取出手续费
    function withdrawBalance() public onlyOwner {
        uint256 contractBalance = address(this).balance;
        require(contractBalance > 0, "No balance to withdraw");
        payable(marketplaceOwner).transfer(contractBalance);
    }

    //===================== 查询个人上架
    function getOwnerSales(address Owner) public view returns (Sale[] memory) {
        uint256 count = 0;
        Sale[] memory result;
        for (uint256 i = 0; i < SaleId; i++) {
            IERC721Enumerable nft = IERC721Enumerable(sales[i].nftContract);
            if (
                sales[i].seller == Owner &&
                sales[i].isActive &&
                nft.ownerOf(sales[i].tokenId) == Owner
            ) {
                count++;
                result[count - 1] = sales[i];
            }
        }
        return result;
    }
}
