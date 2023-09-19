<template>
  <div class="userHomeMain">
    <el-container class="userHomeMainBox">
      <el-aside style="width: auto;" class="animate__animated animate__fadeInLeft">
        <ChatMemu ref="ChatMemu" />
      </el-aside>
      <div class="inf">
        <div class="infBox">
          <div class="content ">
            <div class="contentTop animate__animated animate__fadeInDown">
              <div class="userInf ">
                <div class="userInfTop">
                  <div class="userAvatarBox">
                    <img :src="userAvatar" alt="">
                  </div>
                  <div class="userName">
                    {{ userName }}
                  </div>
                </div>
                <div class="userInfBottom">
                  <div class="UserBalance">
                    <span style="font-size: 4vw;font-weight: 800;"> {{ userBalance }}</span> ETH
                  </div>
                </div>
              </div>
              <!-- <div class="userPet" /> -->
            </div>
            <div class="contentBottom animate__animated animate__fadeInUp">
              <div class="NFTListBox">
                <div class="NFTMenu" v-if="NFTSeriesnameList.length !== 0">
                  <el-input placeholder="请输入内容" v-model="SearchVo.key">
                    <el-button slot="append" icon="el-icon-search" @click="SearchNFT" />
                  </el-input>
                </div>
                <div class="NFTList">
                  <el-collapse v-if="NFTSeriesnameList.length !== 0">
                    <el-collapse-item v-for="address,i in nftContractAddressList" :key="i" :name="NFTSeriesnameList[i]">
                      <template slot="title">
                        <el-popover title="合约地址" placement="top-start" width="350" trigger="hover" :content="address">
                          <i class="header-icon el-icon-info" slot="reference" />
                        </el-popover>
                        <span style="margin-left: 1%;font-weight: 800;font-size: 1vw;">
                          {{ NFTSeriesnameList[i] }}
                        </span>
                        <span class="ADDNFT" v-if="isOwnerCheckArray[i]">
                          <el-popover title="为该合约添加新的NFT 🎉" placement="top-start" width="200" trigger="hover" content="">
                            <router-link class="addImit" :to="{ name: 'addImit',query:{nftContract:address} }"><el-button type="success" plain>前往</el-button>
                            </router-link>
                            <i class="el-icon-plus" slot="reference" />
                          </el-popover>
                        </span>
                      </template>
                      <div class="collapseInnerBox">
                        <template v-for="inf in NFTArray">
                          <template v-for="nft,k in inf">
                            <template v-if="nft.nftAddress == address">
                              <div class="NFTInf" :key="k">
                                <div style="height:65%;width: 100%;overflow: hidden;">
                                  <img class="NFTImage" :src="nft.ipfsPath" alt="">
                                </div>
                                <div class="Inf">
                                  <div class="InfTop">
                                    <div class="NFTName">
                                    {{ nft.nftName }}
                                    </div>
                                    <div class="Active">
                                      <span style="font-size: 25px;">#{{ nft.tokenId }}</span>
                                    </div>
                                  </div>
                                  <div v-if="nft.isActive" class="priceBox">
                                    <span class="price">{{ $store.state.Web3.utils.fromWei(nft.price, 'ether') }}</span> ETH
                                  </div>
                                  <div v-else class="priceBox">
                                    <span class="price" />
                                  </div>
                                  <div class="InfBottom " style="background-color: #d63131e6;" @click="OpenMessageBox(nft,2)" v-if="nft.isActive">
                                    <i class="el-icon-sold-out" />
                                  </div>
                                  <div class="InfBottom" v-else @click="OpenMessageBox(nft,1)">
                                    <i class="el-icon-sell" />
                                  </div>
                                </div>
                              </div>
                            </template>
                          </template>
                        </template>
                      </div>
                    </el-collapse-item>
                  </el-collapse>
                  <div v-else>
                    <el-empty>
                      <template slot="description">
                        <div>
                          <span style="font-weight: 800;
                          margin-bottom: 10%;margin-top: 5%;">未查到相关藏品 </span>
                          <el-button type="success" round plain @click="toMint">
                            前往创造
                          </el-button>
                        </div>
                      </template>
                    </el-empty>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </el-container>
    <div class="MessageMask" v-if="MessageShow">
      <div class="Message animate__animated animate__fadeInUp">
        <div class="MessageLeft">
          <div class="imageBox">
            <img :src="changeNFT.ipfsPath" alt="">
          </div>
        </div>
        <div class="MessageRight">
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                系列名 <span class="tipshelp">Series Name</span>
              </div>
              <div class="tipsTitle2">
                {{ changeNFT.seriesName }}
              </div>
            </div>
          </div>
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                代币符号 <span class="tipshelp">Symbol</span>
              </div>
              <div class="tipsTitle2">
                {{ changeNFT.symbol }}
              </div>
            </div>
            <!-- <el-input
                v-model="Symbol"
                placeholder="Please enter the token symbol"
                maxlength="11"
              /> -->
          </div>
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                名称 <span class="tipshelp">NFT Name</span>
              </div>
              <div class="tipsTitle2">
                {{ changeNFT.nftName }}
              </div>
            </div>
            <!-- <el-input
                v-model="FirstNFTName"
                placeholder="Please enter the name of Genesis NFT"
              /> -->
          </div>
          <div class="select">
            <div class="tipsBox">
              <div class="tipsTitle">
                序号 <span class="tipshelp">Token ID</span>
              </div>
              <div class="tipsTitle2">
                #{{ changeNFT.tokenId }}
              </div>
            </div>
          </div>
          <div class="select" v-if="opt==1">
            <div class="tipsBox">
              <div class="tipsTitle2">
                <el-divider />
              </div>
            </div>
          </div>

          <div class="select" v-if="opt==1">
            <div class="tipsBox">
              <div class="tipsTitle">
                请输入 <span class="tipshelp">价格</span>
              </div>
              <div class="tipsTitle2">
                <el-input-number v-model="Price" :precision="3" :step="0.001"></el-input-number>
              </div>
            </div>
          </div>
          <div class="select" v-if="opt==2">
            <div class="tipsBox">
              <div class="tipsTitle">
                价格 <span class="tipshelp">Price</span>
              </div>
              <div class="tipsTitle2">
                {{ $store.state.Web3.utils.fromWei(changeNFT.price, 'ether') }} ETH
              </div>
            </div>
          </div>
          <div class="select">
            <div class="sumbitBox">
              <el-button @click="CloseMessageBox(1)" class="createButton" type="primary" plain>
                取消
              </el-button>
              <el-button @click="upSale" class="createButton" type="success" plain v-if="opt==1">
                上架
              </el-button>
              <el-button @click="downSale" class="createButton" type="success" plain v-if="opt==2">
                下架
              </el-button>

            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import { UpSale, DownSale, getNFTStruct } from "@/api/axios/contract.js";
  import ChatMemu from "@/views/chat/ChatMemu.vue";
  import { getOwnerNFTsByAddress, search } from "@/api/axios/ownerContractLIst";
  export default {
    components: { ChatMemu },
    data() {
      return {
        userBalance: 0,
        userName: "",
        userAvatar: require("@/assets/webAssets/MetaMask.png"),
        UserNFTListInf: [],
        NFTArray: [],
        nftContractAddressList: [],
        NFTSeriesnameList: [],
        isOwnerCheckArray: [],
        SearchVo: {
          key: "",
          isActive: false,
          minPrice: "",
          maxPrice: "",
          minMaxmums: "",
          maxMaxmums: "",
        },

        MessageShow: false,
        changeNFT: {},
        Price: 0,
        opt: 0,
      };
    },
    async mounted() {
      await this.init();
      await this.getNFTSeriesnameList();
      console.log(this.nftContractAddressList);
      await this.GetNFTContractNFT();
    },
    methods: {
      async init() {
        let user = this.$store.state.currentAddress;
        try {
          this.userBalance = this.$store.state.Web3.utils
            .fromWei(await this.$store.state.Web3.eth.getBalance(user), "ether")
            .slice(0, 4);
          this.userName = `${user.slice(0, 5)}...${user.slice(37)}`;
          this.userAvatar =
            "data:image/png;base64," + new this.Identicon(user, 120).toString();
        } catch (error) {
          return;
        }
      },
      async upSale() {
        if (this.Price > 0) {
          try {
            this.changeNFT.price = this.Price.toString();
            if (await UpSale(this.changeNFT)) {
              this.$notify({
                title: "上架成功 🎉",
                type: "success",
                position: "top-left",
                offset: 200,
              });
              this.CloseMessageBox(2);
            }
          } catch (error) {
            this.$notify.error({
              title: "上架异常",
              position: "top-left",
              offset: 200,
            });
          }
        } else {
              this.$notify({
            title: `价格填写有误`,
            type: "warning",
            position: "top-left",
            offset: 200,
              });
        }
        this.Price = 0
        this.UserNFTListInf = await this.$store.state.ownerNFTList;
      },
      async downSale() {
        try {
          if (await DownSale(this.changeNFT)) {
            this.$notify({
              title: "下架成功",
              type: "success",
              position: "top-left",
              offset: 200,
            });
            this.UserNFTListInf = this.$store.state.ownerNFTList;
            this.CloseMessageBox(2);
          }
        } catch (error) {
          this.$notify.error({
            title: "下架失败",
            position: "top-left",
            offset: 200,
          });
        }
      },
      async GetNFTContractNFT() {
        let NFTInfList = [];
        for (const nftaddress of this.nftContractAddressList) {
          let nft = {
            currentowner: this.$store.state.currentAddress,
            nftAddress: nftaddress,
          };
          await getOwnerNFTsByAddress(nft).then((re) => {
            NFTInfList.push(re.data.data);
            console.log(re);
          });
        }
        this.NFTArray = NFTInfList;
      },
      async getSetAddressArray(ContractAddressArray) {
        for (const nft of ContractAddressArray) {
          this.nftContractAddressList.push(nft.nftAddress);
        }
        this.nftContractAddressList = new Set(this.nftContractAddressList);
      },
      async getNFTSeriesnameList() {
        await this.getSetAddressArray(this.$store.state.ownerNFTList);
        for (const key of this.nftContractAddressList) {
          let SeriesName;
          let contract = await getNFTStruct(key);
          await contract.methods
            .name()
            .call()
            .then((re) => {
              SeriesName = re;
            });
          this.NFTSeriesnameList.push(SeriesName);
          let owner;
          await contract.methods
            .owner()
            .call()
            .then((re) => {
              console.log(re);
              owner = re;
            });
          if (
            this.$store.state.currentAddress.toUpperCase() == owner.toUpperCase()
          ) {
            this.isOwnerCheckArray.push(true);
          } else {
            this.isOwnerCheckArray.push(false);
          }
        }
      },
      toMint() {
        this.$router.push("/home/ImitNFT");
      },
      SearchNFT() {
        search(this.SearchVo).then((re) => {
          console.log(re);
          // this.getSetAddressArray(re.data.data);
        });
      },
      CloseMessageBox(opt) {
        this.MessageShow = false;
        this.changeNFT = {};
        this.price = "";
        if (opt == 1) {
          this.$notify({
            title: `用户取消${this.opt == 1 ? "上架" : "下架"}`,
            type: "warning",
            position: "top-left",
            offset: 200,
          });
        }
      },
      OpenMessageBox(NFT, opt) {
        this.opt = opt;
        this.changeNFT = NFT;
        this.MessageShow = true;
      },
    },
  };
</script>

<style lang="scss" scoped>
.userHomeMain {
  width: 100%;
  background-color: var(--White-eee--);
  .MessageMask {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.7);
    display: flex;
    justify-content: center;
    align-items: center;
    .Message {
      transition: all 0.3s ease-in-out;
      overflow: hidden;
      width: 70%;
      height: 70%;
      border-radius: 30px;
      background-color: white;
      display: flex;
      box-shadow: rgba(50, 50, 93, 0.25) 0px 6px 12px -2px,
        rgba(0, 0, 0, 0.3) 0px 3px 7px -3px;
      .MessageLeft {
        flex: 1;
        display: flex;
        justify-content: center;
        align-items: center;
        .imageBox {
          width: 75%;
          height: 75%;
          border-radius: 30px;
          overflow: hidden;
          display: flex;
          justify-content: center;
          align-items: center;
          background-color: transparent;
          img {
            object-fit: contain;
            width: 100%;
            height: 100%;
          }
        }
      }
      .MessageRight {
        position: relative;
        flex: 1;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        /deep/ .select {
          width: 100%;
          text-align: left;
          margin-bottom: 10px;
          .tipsBox {
            width: 100%;
            color: rgb(105, 105, 105);
            text-align: left;
            transition: all 0.3s ease-in-out;
            .tipsTitle2 {
              margin-top: 2%;
              margin-bottom: 1%;
              font-size: 20px;
              font-weight: 800;
              color: black;
              transition: all 0.3s ease-in-out;
              width: 300px;
            }
            .tipsTitle {
              font-size: 25px;
              color: black;
              .tipshelp {
                color: rgb(105, 105, 105);
                font-size: 20px;
              }
            }
          }
          .sumbitBox {
            text-align: left;
            margin-top: 3%;
            .createButton {
              padding: 20px 35px 18px 35px;
              font-size: 17px;
              border-radius: 15px;
              transition: all 0.3s ease-in-out;
              font-family: "Transformers_Movie";
            }
          }
        }
        .colseButton {
          margin-right: 10%;
        }
      }
    }
  }

  .userHomeMainBox {
    width: 100%;
    height: 100vh;
    overflow: hidden;
  }
  .inf {
    width: 100%;
    height: 100%;
    overflow: auto;
    .infBox {
      display: flex;
      justify-content: center;
      align-items: center;
      border-radius: 30px;
      margin-top: 3%;
      width: 100%;
      .content {
        width: 100%;
        height: 100%;
        color: var(--Dark--);
        margin-left: 3%;
        background-color: var(--White-eee--);

        .contentTop {
          width: 100%;
          display: flex;
          justify-content: flex-start;
          align-items: center;

          .userInf {
            background-color: var(--White--);
            border-radius: 50px;
            width: 600px;
            height: 280px;
            display: flex;
            justify-content: center;
            flex-direction: column;
            align-content: flex-start;
            margin-right: 5%;
            overflow: hidden;
            padding: 0;
            transition: all 0.3s ease-in-out;
            &:hover {
              box-shadow: rgba(0, 0, 0, 0.1) 0px 0px 5px 0px,
                rgba(0, 0, 0, 0.1) 0px 0px 1px 0px;
              transition: all 0.3s ease-in-out;
            }
            .userInfTop {
              width: 100%;
              height: 100%;
              display: flex;
              justify-content: flex-start;
              align-items: center;
              margin-top: 20px;
              .userAvatarBox {
                margin-left: 25px;
                margin-right: 25px;
                min-width: 95px;
                min-height: 95px;
                overflow: hidden;
                border-radius: 50%;
                display: flex;
                justify-content: center;
                align-items: center;
                background-color: rgba(230, 230, 230, 1);
                img {
                  object-fit: cover;
                  width: 85px;
                  height: 85px;
                }
              }
              .userName {
                font-weight: 800;
                font-size: 25px;
              }
            }
            .userInfBottom {
              width: 100%;
              height: 100%;
              display: flex;
              justify-content: flex-end;
              align-items: center;
              .UserBalance {
                margin-right: 35px;
                font-size: 1vw;
              }
            }
          }
          .userPet {
            @extend .userInf;
            width: 1000px;
          }
          .TopBackground {
            position: relative;
            width: 100%;
            height: 30%;
            .userName {
              position: absolute;
              width: 726px;
              height: 180px;
              left: 113px;
              top: 185px;
              color: rgb(255, 255, 255);
              font-family: Inter;
              font-size: 5vw;
              font-weight: 800;
              line-height: 90px;
              letter-spacing: -2px;
              text-align: left;
            }
          }
        }
        .contentBottom {
          width: 95%;
          height: 90%;
          margin-top: 5%;
          display: flex;
          justify-content: center;
          .NFTListBox {
            width: 100%;
            background-color: var(--White--);
            border-radius: 50px;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 3%;
            .NFTList {
              flex: 2;
              width: 95%;
              height: 100%;
              /deep/ .el-collapse {
                border: none;
                .el-collapse-item__header {
                  background-color: var(--White--) !important;
                  color: var(--Dark--);
                }
                .el-collapse-item__content {
                  background-color: var(--White--);
                }
              }

              // .NFTListArray {
              //   width: 100%;
              //   display: flex;

              // }
              .collapseInnerBox {
                .NFTInf {
                  margin: 2%;
                  background-color: var(--White--);
                  border-radius: 50px;
                  width: 320px;
                  height: 420px;
                  display: inline-block;
                  overflow: hidden;
                  transition: all 0.7s ease-in-out;
                  box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px,
                    rgba(17, 17, 26, 0.1) 0px 0px 8px;
                  .NFTImage {
                    width: 100%;
                    height: 100%;
                    object-fit: cover;
                    transition: all 0.6s;
                    cursor: pointer;
                    background-color: white;

                    &:hover {
                      transition: all 0.6s;
                    }
                  }
                }

                .NFTInf:hover {
                  transition: all 0.6s ease-in-out;
                  transform: translateY(-5px) perspective(1000px);
                }
                .NFTInf:hover > * {
                  transform: none;
                }

                .Inf {
                  width: 100%;
                  height: 60%;
                  background-color: var(--Dark--);

                  .priceBox {
                    padding-top: 2%;
                    text-align: left;
                    width: 100%;
                    font-size: 0.5vw;
                    color: var(--White--);
                    .price {
                      margin-left: 15%;
                      font-size: 1.5vw;
                    }
                  }
                  .InfTop {
                    width: 100%;
                    height: 40px;
                    position: relative;
                    .NoActive {
                      position: absolute;
                      right: 10%;
                      top: 15px;
                      width: 100px;
                      padding: 4px 0px 4px 0px;
                      color: white;
                      background-color: #d63131b3;
                      border-radius: 10px;
                      transition: all 0.3s ease-in-out;
                    }
                    .Active {
                      @extend .NoActive;
                      background-color: rgba(85, 201, 96, 0.12);
                      color: #55c960;
                    }
                    .NFTName {
                      position: absolute;
                      left: 0%;
                      top: 19px;
                      color: var(--White--);
                      font-weight: 800;
                      text-align: left;
                      font-size: 14px;
                      margin-left: 20%;
                    }
                  }
                  .InfBottom {
                    font-size: 30px;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 20%;
                    padding-bottom: 100px;
                    color: white;
                    background-color: rgba(0, 149, 255, 0.958);
                    cursor: pointer;
                  }
                }
              }

              .ADDNFT {
                margin-left: 1%;
                // background-color: var(--White--);
                // border-radius: 50px;
                // margin-left: 4%;
                // width: 320px;
                // height: 400px;
                // display: inline-block;
                // transition: all 0.7s ease-in-out;
                // box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px,
                //   rgba(17, 17, 26, 0.1) 0px 0px 8px;
                cursor: pointer;
                div {
                  display: flex;
                  justify-content: center;
                  align-items: center;
                  width: 40px;
                  padding: 2% 10% 2% 10%;
                }
                // display: flex;
                // justify-content: center;
                // align-items: center;
                i {
                  font-weight: 800;
                  color: var(--Dark--);
                }
              }
            }
            .NFTMenu {
              margin-right: 3%;
              height: 100%;
              display: flex;
              flex-direction: column;
              flex: 1;
            }
          }
        }
      }
    }
  }
}
</style>
<style lang="scss" scoped>
@import "@/style/index.css";
</style>
<style lang="scss" scoped>
@import "@/style/MarketShop/index.scss";
</style> 
  