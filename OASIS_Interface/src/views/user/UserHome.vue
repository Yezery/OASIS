<template>
  <div class="userHomeMain">
    <el-container class="userHomeMainBox">
      <el-aside
        style="width: auto;"
        class="animate__animated animate__fadeInLeft"
      >
        <ChatMemu ref="ChatMemu" />
      </el-aside>
      <div class="inf">
        <div class="infBox">
          <div class="content ">
            <div class="contentTop animate__animated animate__fadeInDown">
              <div class="userInf ">
                <div class="userInfTop">
                  <div class="userAvatarBox">
                    <img
                      :src="userAvatar"
                      alt=""
                    >
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
                <div
                  class="NFTMenu"
                  v-if="NFTSeriesnameList.length !== 0"
                >
                  <el-input
                    placeholder="请输入内容"
                    v-model="SearchVo.key"
                  >
                    <el-button
                      slot="append"
                      icon="el-icon-search"
                      @click="SearchNFT"
                    />
                  </el-input>
                </div>
                <div class="NFTList">
                  <el-collapse v-if="NFTSeriesnameList.length !== 0">
                    <el-collapse-item
                      v-for="address,i in nftContractAddressList"
                      :key="i"
                      :name="NFTSeriesnameList[i]"
                    >
                      <template slot="title">
                        <span style="margin-left: 2%;">
                          {{ NFTSeriesnameList[i] }}
                        </span>
                        <el-popover
                          title="合约地址"
                          placement="top-start"
                          width="350"
                          trigger="hover"
                          :content="address"
                        >
                          <i
                            class="header-icon el-icon-info"
                            slot="reference"
                          />
                        </el-popover>
                      </template>
                      <div
                        class="ToSellMain"
                        v-for="inf,j in NFTArray"
                        :key="j"
                      >
                        <div
                          class="NFTInfBox"
                          v-for="nft,k in inf"
                          :key="k"
                        >
                          <template v-if="nft.nftAddress == address">
                            <div class="NFTInf">
                              <div style="height:65%;width: 100%;overflow: hidden;">
                                <img
                                  class="NFTImage"
                                  :src="nft.ipfsPath"
                                  alt=""
                                >
                              </div>
                              <div class="Inf">
                                <div class="InfTop">
                                  <div class="NFTName">
                                    {{ nft.nftName }}
                                  </div>
                                  <div
                                    class="Active"
                                    v-if="nft.isActive"
                                  >
                                    <i class="el-icon-sort" />
                                    <span style="font-weight: 800; font-size: 12px;">在售</span>
                                  </div>
                                  <div
                                    class="NoActive"
                                    v-else
                                  >
                                    <i class="el-icon-sort" />
                                    <span style="font-weight: 800; font-size: 12px;"> 未上架 </span>
                                  </div>
                                </div>
                                <div
                                  v-if="nft.isActive"
                                  class="priceBox"
                                >
                                  <span class="price">{{ $store.state.Web3.utils.fromWei(nft.price, 'ether') }}</span> ETH
                                </div>
                                <div
                                  v-else
                                  class="priceBox"
                                >
                                  <span class="price" />
                                </div>
                                <div
                                  class="InfBottom "
                                  style="background-color: #d63131b3;"
                                  @click="downSale(nft)"
                                  v-if="nft.isActive"
                                >
                                  <i class="el-icon-sold-out" />
                                </div>
                                <div
                                  class="InfBottom"
                                  v-else
                                  @click="upSale(nft)"
                                >
                                  <i class="el-icon-sell" />
                                </div>
                              </div>
                            </div>
                          </template>
                        </div>
                      </div>
                      <div class="ADDNFT">
                        <router-link
                          class="addImit"
                          :to="{ name: 'addImit',query:{nftContract:address} }"
                        >
                          <i class="el-icon-plus" />
                        </router-link>
                      </div>
                    </el-collapse-item>
                  </el-collapse>
                  <div v-else>
                    <el-empty>
                      <template slot="description">
                        <div>
                          <span
                            style="font-weight: 800;
                          margin-bottom: 10%;"
                          >未查到相关藏品 </span> 
                          <el-button @click="toMint">
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
        SearchVo: {
          key: "",
          isActive: false,
          minPrice: "",
          maxPrice: "",
          minMaxmums: "",
          maxMaxmums: "",
        },
      };
    },
    async mounted() {
      await this.init();
      await this.getNFTSeriesnameList();
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
      async upSale(NFT) {
        this.$prompt("请输入价格(ETH)", "上架", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          inputPattern: /^(?=.*[1-9])\d*(?:\.\d+)?$/,
          inputErrorMessage: "只能是数字",
        })
          .then(async ({ value }) => {
            NFT.price = value;
            try {
              await UpSale(NFT);
              this.$notify({
              title: "上架成功 🎉",
                type: "success",
                position: "top-left",
            offset: 200,
            });
            } catch (error) {
              this.$notify.error({
                title: "上架异常",
                position: "top-left",
            offset: 200,
              });
            }
          })
          .catch(() => {
            this.$notify({
              title: "用户取消上架",
              type: "warning",
              position: "top-left",
            offset: 200,
            });
          });

        this.UserNFTListInf = await this.$store.state.ownerNFTList;
      },
      async downSale(NFT) {
        try {
          await DownSale(NFT);
          this.$notify({
            title: "下架成功",
            type: "success",
            position: "top-left",
            offset: 200,
          });
          this.UserNFTListInf = this.$store.state.ownerNFTList;
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
            ownerAddress: this.$store.state.currentAddress,
            nftAddress: nftaddress,
          };
          await getOwnerNFTsByAddress(nft).then((re) => {
            NFTInfList.push(re.data.data);
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
        }
      },
      toMint() {
        this.$router.push("/home/ImitNFT")
      },
      SearchNFT() {
        search(this.SearchVo).then((re) => {
          console.log(re);
          // this.getSetAddressArray(re.data.data);
        });
      },
    },
  };
</script>

<style lang="scss" scoped>
.userHomeMain {
  width: 100%;
  background-color: var(--White-eee--);
  .userHomeMainBox{
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
              .ToSellMain {
                height: 100%;
                width: 100%;
                font-family: Arial, Helvetica, sans-serif;
                margin-top: 2%;
                display: flex;
                justify-content: flex-start;
                flex-wrap: wrap;
                margin-left: 2%;
                .NFTInfBox {
                  display: flex;
                  justify-content: flex-start;
                  align-items: center;
                  margin: 1% 2% 2% 2%;
                  .NFTInf {
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
                      object-fit: contain;
                      transition: all 0.6s;
                      cursor: pointer;
                      &:hover {
                        transform: scale(1.1);

                        transition: all 0.6s;
                      }
                    }
                  }

                  .NFTInf:hover {
                    transition: all 0.6s ease-in-out;
                    transform: translateY(-20px) perspective(1000px);
                  }
                  .NFTInf:hover > * {
                    transform: none;
                  }
                }

                .Inf {
                  width: 100%;
                  height: 60%;
                  .priceBox {
                    padding-top: 2%;
                    text-align: left;
                    width: 100%;
                    font-size: 0.5vw;
                    color: var(--Dark--);
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
                      color: var(--Dark--);
                      font-weight: 800;
                      text-align: left;
                      font-size: 14px;
                      margin-left: 20%;
                    }
                  }
                  .InfBottom {
                    font-size: 1.5vw;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 20%;
                    color: white;
                    background-color: rgba(0, 149, 255, 0.958);
                    cursor: pointer;
                  }
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

                .downSale {
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
              .ADDNFT {
                background-color: var(--White--);
                border-radius: 50px;
                margin-left: 4%;
                width: 320px;
                height: 400px;
                display: inline-block;
                transition: all 0.7s ease-in-out;
                box-shadow: rgba(17, 17, 26, 0.05) 0px 1px 0px,
                  rgba(17, 17, 26, 0.1) 0px 0px 8px;
                cursor: pointer;
                display: flex;
                justify-content: center;
                align-items: center;
                i {
                  font-size: 70px;
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
  