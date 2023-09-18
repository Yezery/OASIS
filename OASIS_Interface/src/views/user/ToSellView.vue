<template>
  <div class="ToSellMain animate__animated animate__fadeInRight">
    <div class="ToSellHeaderBox">
      <div class="ToSellHeader">
        <router-link :to="{ name: 'MarketShop' }">
          <i class="el-icon-arrow-left" />
        </router-link>
      </div>
    </div>
    <div class="ToSellMain">
      <div
        class="NFTInf"
        v-for="inf in UserNFTListInf"
        :key="inf.ipfsPath"
      >
        <div class="imageBox">
          <img
            class="NFTImage"
            :src="inf.ipfsPath"
            alt=""
          >
        </div>
        <div class="Inf">
          <template v-if="inf.isActive">
            <div class="NFTName">
              {{ inf.nftName }}
            </div>
            <div class="ownerAndToSell">
              <div class="ToSellBox">
                <div
                  v-if="!inf.isActive"
                  class="ToSellinnerBox animate__animated animate__fadeInUp"
                >
                  <i class="el-icon-sort" />
                  <span
                    class="apporve "
                    @click="inf.isActive = !inf.isActive"
                  > 上架 </span>
                </div>
                <div
                  v-else
                  class="downSale animate__animated animate__fadeInUp"
                  @click="downSale(inf)"
                >
                  <i class="el-icon-sort" />
                  <span class="apporve "> 下架 </span>
                </div>
              </div>
            </div>
          </template>
          <template v-else>
            <div class="NFTName animate__animated animate__fadeIn">
              <el-input
                v-model="inf.price"
                placeholder="请输入内容"
              />
              ETH
            </div>
            <div class="ownerAndToSell animate__animated animate__fadeInDown">
              <div class="ToSellBox">
                <div class="ToSellinnerBox">
                  <i class="el-icon-sort" />
                  <span
                    class="apporve"
                    @click="UpSale(inf)"
                  > 确认 </span>
                </div>
                <div
                  class="ToSellinnerBox "
                  style="background-color: rgba(225, 44, 44, 0.7);"
                >
                  <i
                    class="el-icon-sort"
                    style="color: white;"
                  />
                  <span
                    class="apporve"
                    style="color: white;"
                    @click="inf.isActive = !inf.isActive"
                  > 取消 </span>
                </div>
              </div>
            </div>
          </template>
        </div>
      </div>
      <!-- absolute -->
    </div>
  </div>
</template>

<script>
  import { UpSale ,DownSale} from "@/api/axios/contract";
  export default {
    data() {
      return {
        UserNFTList: [],
        UserNFTListInf: [],
        MintNFTContractAbi: this.$store.state.MintNFTContractAbi,
        contractAddress: "",
        NTFList: [],
        contract: null,
        value: 1,
       
      };
    },
  mounted() {
    this.UserNFTListInf = this.$store.state.ownerNFTList;
    },
  methods: {
    async downSale(NFT) {
      await DownSale(NFT)
      this.UserNFTListInf = this.$store.state.ownerNFTList;
    },
      // async GETNFTContract(contractAddress) {
      //   let NewContract = await new this.$store.state.Web3.eth.Contract(
      //     this.MintNFTContractAbi,
      //     contractAddress
      //   );
      //   this.contractAddress = contractAddress;
      //   await NewContract.methods
      //     .getNFTsByOwner(this.$store.state.currentAddress)
      //     .call()
      //     .then((re) => {
      //       this.UserNFTList = re;
      //     });
      //   return NewContract;
      // },
      // async GetNFTInf() {
      //   let Contract = await this.GETNFTContract(this.contractAddress);
      //   for (let index = 0; index < this.UserNFTList.length; index++) {
      //     var NFTURI;
      //     var NFTNAME;
      //     await Contract.methods
      //       .tokenURI(Number(this.UserNFTList[index]))
      //       .call()
      //       .then((res) => {
      //         NFTURI = res;
      //       });
      //     await Contract.methods
      //       .name()
      //       .call()
      //       .then((res) => {
      //         NFTNAME = res;
      //       });
      //     let web3 = new this.Web3(window.ethereum);
      //     this.contract = new web3.eth.Contract(
      //       this.$store.state.MarketContractAbi,
      //       this.$store.state.MarketContractAddress
      //     );
      //     await this.contract.methods
      //       .fetchMyNFTs()
      //       .call()
      //       .then((res) => {
      //         this.NTFList = res;
      //       });
      //     console.log(await this.contract.methods.fetchMyNFTs().call());
      //     var NFT = {
      //       NFTURI: NFTURI,
      //       NFTNAME: NFTNAME,
      //       NFTAddress: Contract._address,
      //       NFTTokenId: this.UserNFTList[index],
      //     };

      //     this.UserNFTListInf.push(NFT);
      //   }
      // },
      // async upSale(NFT) {
      //   this.$prompt("请输入价格", "上架提醒", {
      //     confirmButtonText: "确定",
      //     cancelButtonText: "取消",
      //     inputPattern: /^[1-9]\d*$/,
      //     inputErrorMessage: "不能为0",
      //   })
      //     .then(async ({ value }) => {

      //       let NFTContract = await this.GETNFTContract(NFT.NFTAddress);
      //       let MarketContract = this.$store.state.MarketContract;
      //       //授权
      //       NFTContract.methods
      //         .approve(this.$store.state.MarketContractAddress, NFT.NFTTokenId)
      //         .send({ from: this.$store.state.currentAddress });

      //       // 上架
      //       let Value;
      //       await MarketContract.methods
      //         .getListingPrice()
      //         .call()
      //         .then((res) => {
      //           Value = res;
      //         });
      //       await MarketContract.methods
      //         .createMarketplaceItem(NFT.NFTAddress, 1, 1)
      //         .send({
      //           from: this.$store.state.currentAddress,
      //           // value: Value
      //           value: this.$store.state.Web3.utils.toWei(Value, "wei"),
      //         });
      //       this.$message({
      //         type: "success",
      //         message: "你的邮箱是: " + value,
      //       });
      //     })
      //     .catch((res) => {
      //       this.$message({
      //         type: "info",
      //         message: "取消输入",
      //       });
      //       console.log(res);
      //     });
      // },
      async UpSale(NFT) {
        if (Number(NFT.price) == 0 || NFT.price == "") {
          return;
        }
        await UpSale(NFT);
        this.UserNFTListInf = await this.$store.state.ownerNFTList;
      },
      
    },
  };
</script>

<style lang="scss" scoped>
.ToSellMain {
  height: 100%;
  text-align: center;
  width: 100%;
  font-family: Arial, Helvetica, sans-serif;

  .NFTInf {
    background-color: var(--White--);
    border-radius: 30px;
    width: 315px;
    height: 390px;
    display: inline-block;
    overflow: hidden;
    transition: all 0.3s ease-in-out;
    &:hover {
      box-shadow: rgba(14, 30, 37, 0.12) 0px 2px 4px 0px,
        rgba(14, 30, 37, 0.32) 0px 2px 16px 0px;
      transition: all 0.3s ease-in-out;
    }
  }
  @media screen and (max-width: 1600px) and (min-width: 1600px) {
    .ToSellMain {
      width: 100%;
      justify-content: center;
      align-items: center;
      flex-direction: column;
    }
  }
  .imageBox {
    width: 100%;
    height: 63%;
    border-radius: 30px;
    overflow: hidden;
  }
  .Inf {
    width: 100%;
    height: 100%;
    position: relative;
  }
  .NFTImage {
    object-fit: cover;
    width: 100%;
    transition: all 0.6s;
    cursor: pointer;

    overflow: hidden;
    &:hover {
      transform: scale(1.1);

      transition: all 0.6s;
    }
  }
  .NFTName {
    display: inline-block;
    color: var(--Dark--);
    font-weight: 800;
    text-align: left;
    margin-top: 20px;
    font-size: 13px;
  }
  .ownerAndToSell {
    width: 100%;
    display: flex;
    margin-top: 18px;
    justify-content: center;
    align-items: center;
  }
  .ToSellBox {
    display: flex;
    justify-content: flex-start;
    align-items: center;
  }

  .ToSellinnerBox {
    margin: 2%;
    width: 82.79px;
    padding: 8px 0px 8px 0px;
    color: #55c960;
    background-color: rgba(85, 201, 96, 0.12);
    border-radius: 10px;
    cursor: pointer;
    transition: all 0.3s ease-in-out;
    &:hover {
      background-color: rgba(85, 201, 96, 0.24);
      transition: all 0.3s ease-in-out;
    }
  }
  .downSale{
    @extend .ToSellinnerBox;
    background-color: #d63131e5;
    color: white;
    &:hover {
      background-color: rgb(255, 0, 0);
      transition: all 0.3s ease-in-out;
    }
  }
  .apporve {
    font-size: 12px;
    font-weight: 800;
  }
  .ToSellHeader {
    display: flex;
    justify-content: space-between;
    width: 95%;
    height: 10%;
    font-size: 25px;
    margin-top: 3%;
    border-radius: 30px;
    padding: 3.5% 0;
    font-weight: 800;
    color: var(--Dark--);
  }
  .ToSellHeaderBox {
    width: 100%;
    text-align: center;
    display: flex;
    justify-content: center;
    align-items: center;
  }
}
</style>
