import request from "@/utils/axiosRequest"
import store from '@/store';

import { create } from "ipfs-http-client"
import Web3 from 'web3'
import { updateSaleactive, insertSale, deleteSale, } from "./Sale"
import { postOwnerContractList, getOwnerUpSaleNFTs, updateNFTOwnerListAfterBuy } from "./ownerContractLIst";
import { mintNFTContractABI, mintNFTContractBytecode, marketContractAddress, marketContractABI, ipfsPublicGatewayUrl, rpcUrl } from "@/contract/Contract"
//  IPFS的根URL
// const baseURL = ipfsBaseURI();
//  铸币合约ABI
const MintNFTContractABI = mintNFTContractABI()
//  铸币合约ByteCode
const MintNFTContractBytecode = mintNFTContractBytecode()
//  市场合约地址
const MarketContractAddress = marketContractAddress();
//  市场合约ABI
const MarketContractABI = marketContractABI();
// 指定其他节点的 IP 和端口
// const peerAddress = ipfsPeerAddress()
//  IPFS网关
const publicGatewayUrl = ipfsPublicGatewayUrl();
// 服务器IP
const IP = store.state.IP

//====================  部署铸造NFT合约，返回NFT合约实例 
async function deployNFTContract(Name, Symbol, Maxmums) {
  let contract = new store.state.Web3.eth.Contract(MintNFTContractABI);
  //  部署合约
  let NewNFTContractAddress = await contract.deploy({ data: MintNFTContractBytecode, arguments: [Name, Symbol, publicGatewayUrl, Maxmums] })
    .send({
      from: store.state.currentAddress
    });
  console.log(NewNFTContractAddress._address);
  //  通过部署合约返回的合约地址，与合约ABI产生合约实例
  let NFTContract = await new store.state.Web3.eth.Contract(MintNFTContractABI, NewNFTContractAddress._address);
  return NFTContract;
}
//====================  保存到IPFS，返回IPFS哈希值
async function savetoIPFS(uploadFiles) {
  try {
    if (uploadFiles.length != 0) {
      //ipfs的add方法是将东西添加到ipfs网络上
      let results = await create({
        host: `${IP}`,
        port: "5001",
        protocol: "http",
      }).add(
        uploadFiles[0].raw,
        { remote: true, pin: true, }
      );
      return results.path;
    }
  } catch (err) {
    console.error(err);
    return null
  }
}

// 
async function mintNFT(NFTContract, name, symbol, maxmums, NFTName, description, uploadFiles) {
  let nftCount = 0
  let ipfsHash = await savetoIPFS(uploadFiles)
  let ipfsPath = `ipfs/${ipfsHash}?filename=${NFTName}`
  if (ipfsHash != null) {
    try {
      await NFTContract.methods._currentId().call().then(re => { nftCount = re })
    } catch (error) {
      console.log(error);
      console.log(NFTContract.methods);
    }
    await NFTContract.methods.mint(NFTName, description, ipfsHash).send({ from: store.state.currentAddress })
    // // ====================  信息存储到数据库中
    await request({
      url: '/createNFT',
      method: 'post',
      data: {
        "ownerAddress": store.state.currentAddress,
        "currentOwner": store.state.currentAddress,
        "nftAddress": NFTContract._address,
        "ipfsPath": `${publicGatewayUrl}${ipfsPath}`,
        "isActive": false,
        "price": "0",
        "seriesName": name,
        "symbol": symbol,
        "nftName": NFTName,
        "maxmums": Number(maxmums),
        "description": description,
        "tokenId": Number(nftCount)
      }
    }).then(res => {
      if (res.status == 200) {
        postOwnerContractList({ ownerAddress: store.state.currentAddress }).then((re) => {
          store.commit("setOwnerNFTList", re.data.data);
        });
        return
      } else {
        return
      }
    })
  }
}
export async function addMint(to, nftAddress, NFTName, description, uploadFiles, seriesName, symbol, maxmums) {
  let nftCount = 0
  let ipfsHash = await savetoIPFS(uploadFiles)
  let ipfsPath = `ipfs/${ipfsHash}?filename=${NFTName}`
  let contract = await getNFTStruct(nftAddress)
  if (ipfsHash != null) {
    try {
      await contract.methods._currentId().call().then(re => { nftCount = re })
    } catch (error) {
      console.log(error);
    }
    await contract.methods.giveMint(to, NFTName, description, ipfsHash).send({ from: store.state.currentAddress })
    // // ====================  信息存储到数据库中
    await request({
      url: '/createNFT',
      method: 'post',
      data: {
        "ownerAddress": to,
        "currentOwner":store.state.currentAddress,
        "nftAddress": contract._address,
        "ipfsPath": `${publicGatewayUrl}${ipfsPath}`,
        "isActive": false,
        "price": "0",
        "seriesName": seriesName,
        "symbol": symbol,
        "nftName": NFTName,
        "maxmums": Number(maxmums),
        "description": description,
        "tokenId": Number(nftCount)
      }
    }).then(res => {
      if (res.status == 200) {
        postOwnerContractList({ ownerAddress: store.state.currentAddress }).then((re) => {
          store.commit("setOwnerNFTList", re.data.data);
        });
        return
      } else {
        return
      }
    })
  }
}

//====================  生成NFT合约实例
export async function getNFTStruct(nftAddress) {
  return await new store.state.Web3.eth.Contract(MintNFTContractABI, nftAddress);
}

//==================== NFT总控
export async function MakeNFT(Name, Symbol, uploadFiles, Maxmums, FirstNFTName, Description) {
  console.log("部署合约");
  const NFTContract = await deployNFTContract(Name, Symbol, Maxmums);
  console.log("铸造");
  await mintNFT(NFTContract, Name, Symbol, Maxmums, FirstNFTName, Description, uploadFiles)
}
//====================  市场合约
let MarketContract;
let marketUseWeb3;
// 实例化市场合约

console.log("初始化WEB3......");
if (window.ethereum != undefined) {
  marketUseWeb3 = new Web3(window.ethereum);
  store.commit("setWEB3", new Web3(window.ethereum));
  console.log("初始化WEB3(MetaMask)完毕");
} else {
  marketUseWeb3 = new Web3(rpcUrl());
  store.commit("setWEB3", new Web3(rpcUrl()));
  console.log("初始化WEB3(RPC)完毕");
}

MarketContract = new marketUseWeb3.eth.Contract(
  MarketContractABI,
  MarketContractAddress
);
//====================  上架NFT
export async function getSaleList() {
  let SaleList = []
  try {
    await MarketContract.methods.getSales().call().then(re => {
      SaleList = re
    });
  } catch (error) {
    console.log(error);
  }
  return SaleList
}

export async function UpSale(NFT) {
  let SaleId;
  try {
    //授权
    let NFTContract = await new store.state.Web3.eth.Contract(
      MintNFTContractABI,
      NFT.nftAddress
    );
    await NFTContract.methods
      .approve(MarketContractAddress, NFT.tokenId)
      .send({ from: store.state.currentAddress });

    await MarketContract.methods.createSale(NFT.nftAddress, NFT.tokenId, store.state.Web3.utils.toWei(NFT.price, 'ether')
      , "").send({ from: store.state.currentAddress })
    await MarketContract.methods.getSalesId().call().then(re => {
      SaleId = re
    })
  } catch (error) {
    console.log(error);
  }
  var Sale = {
    saleId: Number(SaleId),
    NFTOwnerList: NFT
  }
  try {
    await insertSale(Sale)
    NFT.isActive = true
    NFT.price = store.state.Web3.utils.toWei(NFT.price, 'ether')
    await updateSaleactive(NFT)
    // await postOwnerContractList({ ownerAddress: store.state.currentAddress }).then((re) => {
    //   store.commit("setOwnerNFTList", re.data.data);
    // });
  } catch (error) {
    console.log(error);
  }
}

export async function Buy(NFT) {

  try {
    await MarketContract.methods.buy(NFT.saleId - 1)
      .send({
        from: store.state.currentAddress,
        value: store.state.Web3.utils.toWei(NFT.price, 'ether'),
      });
    NFT.isActive = false
    NFT.currentOwner = store.state.currentAddress
    await updateNFTOwnerListAfterBuy(NFT)
    await deleteSale(NFT)
    // await postOwnerContractList({ ownerAddress: store.state.currentAddress }).then((re) => {
    //   store.commit("setOwnerNFTList", re.data.data);
    // });

  } catch (error) {
    console.log(error);
    return
  }

}

export async function getFeePercentage() {
  let FeePercentage;
  await MarketContract.methods.feePercentage().call().then(re => { FeePercentage = re })
  return FeePercentage;
}

export async function DownSale(NFT) {
  let SaleId;
  try {
    await getOwnerUpSaleNFTs(NFT).then(re => { SaleId = re.data.data.saleId; })
    if (SaleId != undefined) {
      await MarketContract.methods.cancelSale(SaleId)
        .send({
          from: store.state.currentAddress
        });
      await deleteSale(NFT)
      // await postOwnerContractList({ ownerAddress: store.state.currentAddress }).then((re) => {
      //   store.commit("setOwnerNFTList", re.data.data);
      // });
    } else {
      return
    }
  } catch (error) {
    console.log(error);
    return
  }
}