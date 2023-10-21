<template>
  <div
    class="MainWindow animate__animated animate__fadeInRight"
    v-loading.fullscreen.lock="isChanging"
    element-loading-text="交易进行中"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
    <div class="imitNFTFrame ">
      <div style="width: 100%;text-align: left;margin-left: 5%;margin-top: 3%; ">
        <router-link :to="{ name: 'userhome' }">
          <i
            class="el-icon-arrow-left"
            style="font-size: 2vw;color: var(--Dark--);"
          />
        </router-link>
      </div>
      <div class="imitNFTFrameTop">
        <div class="ImitTitle">
          Add 3DNFT 
        </div>
        <div class="imitNFTFrameTop_tipsBox">
          为你的系列增加 3DNFT
        </div>
      </div>
      <div class="imitNFTFrameMain">
        <div class="imitNFTFrameMain_right">
          <div class="selectBox">
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span>Model file
                </div>
                <div class="tipsTitle2">
                  模型文件<span style="font-weight: 800;font-size: 0.2vw;"> (仅支持 .gltf 嵌入式格式与 .glb非分离格式)</span>
                </div>
              </div>
              <div class="imageUpLoad">
                <el-upload
                  class="upload-demo"
                  ref="modelUpload"
                  :on-remove="handleRemove"
                  :file-list="fileList"
                  :auto-upload="false"
                  action="#"
                  :limit="1"
                  :on-change="checkFileType"
                >
                  <el-button
                    class="select3DFile"
                    slot="trigger"
                    @click="handleRemove"
                  >
                    选取3D文件
                  </el-button>
                </el-upload>
                <div
                  class="show3DBox"
                  v-loading="loading"
                  v-if="isShow3D"
                  element-loading-text="准备模型中..."
                >
                  <Model
                    :model-path="modelPath"
                    @initModel="seeModel"
                  />
                </div>
              </div>
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Name (NFT系列名)
                </div>
                <div class="tipsTitle2">
                  <span style="font-size: 1.5vw;font-weight: 800;margin-top: 2%;"> {{ seriesName }}</span>
                </div>
              </div>
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Symbol (该系列NFT代币符号)
                </div>
                <div class="tipsTitle2">
                  <span style="font-size: 1.5vw;font-weight: 800;margin-top: 2%">{{ symbol }}</span>
                </div>
              </div>
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Add 3DNFT Name
                </div>
                <div class="tipsTitle2">
                  新增加的 3DNFT 名称
                </div>
              </div>
              <el-input
                v-model="NFTName"
                placeholder="Please enter the name of  3DNFT"
              />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  As a gift (optional)
                </div>
                <div class="tipsTitle2">
                  作为礼物赠送（可选）
                </div>
              </div>
              <el-input
                placeholder="请输入赠送者地址"
                v-model="to"
              >
                <template slot="prepend">
                  <img
                    width="25px"
                    height="25px"
                    src="@/assets/SVG/Love.svg"
                    alt="SVG Image"
                    style="color:"
                  >
                </template>
              </el-input>
            </div>
            <div class="select">
              <div class="sumbitBox">
                <el-button
                  @click="addMint"
                  :disabled="!canSubmit"
                  class="createButton"
                >
                  添加
                </el-button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import { getNFTStruct } from "@/api/axios/contract";
  import { addMint } from "@/api/axios/contract";
  import { savetoIPFS } from "@/api/axios/contract";
  import Model from "@/components/3DModelShow/3DModel.vue";
  export default {
    components: {
      Model,
    },
    data() {
      return {
        // imageUrl: "",
        // fileList: [],
        // dialogImageUrl: "",
        // dialogVisible: false,
        // disabled: false,
        // Name: "",
        // Symbol: "",
        // FirstNFTName: "#0",
        // Maxmums: 1,
        // active: 1,
        // sumitEnable: true,
        // isRepeatClick: true,
        imageUrl: "",
        dialogImageUrl: "",
        dialogVisible: false,
        disabled: false,
        fileList: [],
        seriesName: "",
        symbol: "",
        nftContractAddress: "",
        Description: "",
        maxmums: 0,
        NFTName: "",
        to: this.$store.state.currentAddress,
        sumitEnable: true,
        isRepeatClick: true,

        initModel: undefined,
        model: undefined,
        modelPath: "",
        loading: false,
        isShow3D: false,
        isTypeOf3DModel: false,
        isChanging: false,
      };
    },
    async mounted() {
      this.nftContractAddress = this.$route.query.nftContract;
      await this.init();
    },
    computed: {
      canSubmit() {
        if (
          this.seriesName == "" ||
          this.symbol == "" ||
          this.$refs.modelUpload.uploadFiles.length == 0 ||
          this.NFTName == "" ||
          !this.isTypeOf3DModel
        ) {
          return false;
        } else {
          return true;
        }
      },
    },

    methods: {
      async init() {
        let contract = await getNFTStruct(this.nftContractAddress);
        await contract.methods
          .symbol()
          .call()
          .then((re) => {
            this.symbol = re;
          });
        await contract.methods
          .name()
          .call()
          .then((re) => {
            this.seriesName = re;
          });
        await contract.methods
          ._maximums()
          .call()
          .then((re) => {
            this.maxmums = re;
          });
      },
      seeModel(data) {
        this.initModel = data;
      },
      async getModel() {
        this.isShow3D = true;
        this.loading = true;

        let ipfsHash;
        await savetoIPFS(this.$refs.modelUpload.uploadFiles).then((re) => {
          ipfsHash = re;
        });
        this.loading = false;
        this.modelPath = `http://10.39.5.194:8080/ipfs/${ipfsHash}?filename=${this.FirstNFTName}`;
      },
      checkFileType() {
        let prix = this.getFileExtendingName(
          this.$refs.modelUpload.uploadFiles[0].raw.name
        );
        if (prix == ".gltf" || prix == ".glb") {
          this.FirstNFTName = this.$refs.modelUpload.uploadFiles[0].name.replace(
            /(.*\/)*([^.]+).*/gi,
            "$2"
          );
          this.isTypeOf3DModel = true;
          this.getModel();
        } else {
          this.handleRemove();
          this.$notify.error({
            title: "请上传gltf或glb格式的模型",
            position: "top-left",
            offset: 200,
          });
          this.isTypeOf3DModel = false;
        }
      },
      getFileExtendingName(filename) {
        // 文件扩展名匹配正则
        var reg = /.[^.]+$/;
        var matches = reg.exec(filename);
        if (matches) {
          return matches[0];
        }
        return "";
      },
      async addMint() {
        if (this.isRepeatClick) {
          this.isRepeatClick = false;
          try {
            this.isChanging=true
            await addMint(
              this.to,
              this.nftContractAddress,
              this.NFTName,
              "3D",
              this.$refs.modelUpload.uploadFiles,
              this.seriesName,
              this.symbol,
              this.maxmums
            );
            this.isChanging=false
            this.$notify({
              title: "添加成功",
              type: "success",
              position: "top-left",
              offset: 200,
            });
            
            this.fileList = [];
            this.NFTName = "";
            this.maxmums = 1;
            this.Description = "";
            this.handleRemove()
          } catch (error) {
            this.isChanging=false
            this.$notify.error({
              title: "添加失败",
              position: "top-left",
              offset: 200,
            });
            return
          }
        } else {
          this.$notify({
            title: "请勿操作频繁",
            type: "warning",
            position: "top-left",
            offset: 200,
          });
        }

        setTimeout(() => {
          this.isRepeatClick = true;
        }, 5000);
      },
      handleRemove() {
        this.$refs.modelUpload.clearFiles();
        this.isShow3D = false;
        this.isTypeOf3DModel = false;
      },
    },
  };
</script>

<style lang="scss" scoped>
.MainWindow {
  position: relative;
  overflow: hidden;
  font-family: Arial, Helvetica, sans-serif;
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;

  .imitNFTFrame {
    transition: all 0.3s ease-in-out;
    margin: 1.3%;
    border-radius: 40px;
    background-color: var(--White--);
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    width: 80%;
    .imitNFTFrameTop {
      position: relative;
      text-align: left;

      .imitNFTFrameTop_tipsBox {
        width: 100%;
        margin-top: 2%;
        margin-bottom: 2%;
        text-align: center;
        font-size: 1.5vw;
        color: rgb(163, 162, 162);
        font-weight: 800;
      }
      a {
        text-decoration: none;
        color: var(--Dark--);
      }
    }
    .imitNFTFrameMain {
      width: 55%;
      // .imitNFTFrameMain_left {
      //   height: 100%;
      //   text-align: center;
      //   margin-left: 6%;
      //   // /deep/ .imageUpLoad {
      //   //   & div:first-child {
      //   //     position: relative !important;
      //   //     overflow: hidden;
      //   //     height: 500px;
      //   //   }
      //   //   & .el-upload-list--picture-card .el-upload-list__item {
      //   //     width: 500px;
      //   //     height: 500px;
      //   //     border-radius: 20px;
      //   //     position: absolute !important;
      //   //   }
      //   //   & .el-upload--picture-card {
      //   //     object-fit: contain;
      //   //     height: 500px;
      //   //     width: 500px;
      //   //     background-color: rgba(255, 255, 255, 0);
      //   //     border: 3px dashed var(--Dark--);
      //   //     border-radius: 20px;
      //   //     transition: all 0.3s ease-in-out;
      //   //     i {
      //   //       margin-top: 220px;
      //   //     }
      //   //   }
      //   //   & .el-upload--picture-card:hover {
      //   //     background-color: rgba(255, 255, 255, 0);
      //   //   }
      //   // }
      // }
      .imitNFTFrameMain_right {
        padding-top: 3%;
        padding-left: 5%;
        width: 100%;
        height: 100%;
        .selectBox {
          flex-direction: column;

          /deep/ .select {
            width: 100%;
            text-align: left;
            margin-bottom: 30px;

            .el-input__inner:not(.el-input__inner:last-child),
            .el-textarea__inner {
              font-family: Arial, Helvetica, sans-serif;
              height: 50px;
              margin-top: 1%;
              background-color: white;
              border-radius: 12px;
              width: 25vw;
              color: black;
              transition: all 0.3s ease-in-out;
            }
            .input_number {
              margin-top: 1%;
            }
            .tipsBox {
              width: 100%;
              color: rgb(163, 162, 162);
              text-align: left;
              transition: all 0.3s ease-in-out;
              .tipsTitle {
                margin-bottom: 1%;
                width: 76%;
                font-size: 18px;
                font-weight: 800;
                color: var(--Dark--);
                transition: all 0.3s ease-in-out;
              }
              .tipsTitle2 {
                font-size: 12px;
                margin-bottom: 3px;
              }
            }
            .show3DBox {
              width: 30vw;
              height: 30vw;
              display: flex;
              justify-content: center;
              align-items: center;
              border: 1px dashed #ccc;
              border-radius: 15px;
              overflow: hidden;
              margin-top: 2%;
            }
            .imageUpLoad {
              margin-top: 4%;
              .select3DFile {
                border-radius: 8px;
                background-color: var(--Dark--);
                border: 1px dashed var(--Dark--);
                color: var(--White--);
              }
              .seeShow3D {
                border-radius: 8px;
                background-color: var(--Dark--);
                border: 1px dashed var(--Dark--);
                color: var(--White--);
              }
            }
          }
        }
      }
    }
  }
}

.NFTImageSelectBox {
  width: 76%;
  text-align: left;
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  transition: all 0.3s ease-in-out;
}

.ImitTitle {
  font-size: 3.5vw;
  width: 100%;
  color: var(--Dark--);
  margin-top: 1%;
  font-weight: 800;
  transition: all 0.3s ease-in-out;
}

.sumbitBox {
  text-align: left;
  margin-top: 6%;
  margin-bottom: 6%;
  .createButton {
    padding: 20px 35px 18px 35px;
    font-size: 17px;
    background-color: #fff;
    border-radius: 15px;
    transition: all 0.3s ease-in-out;
    font-family: "Transformers_Movie";
    color: black;
    :hover {
      color: black;
    }
  }
  .is-disabled {
    background-color: transparent;
    color: #c0c4cc;
    :hover {
      color: #c0c4cc;
    }
  }
}
</style>
