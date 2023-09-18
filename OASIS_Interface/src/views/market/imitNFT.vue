<template>
  <div class="MainWindow animate__animated animate__fadeInRight">
    <div class="imitNFTFrame ">
      <div class="imitNFTFrameTop">
        <router-link :to="{ name: 'MarketShop' }">
          <i
            class="el-icon-arrow-left"
            style="font-size: 2vw;"
          />
        </router-link>
        <div class="ImitTitle">
          CREATE YOUR NFT
        </div>
        <div class="imitNFTFrameTop_tipsBox">
          创造属于你的NFT
        </div>
      </div>
      <div class="imitNFTFrameMain">
        <div class="imitNFTFrameMain_left">
          <div class="imageUpLoad">
            <el-upload
              action="#"
              list-type="picture-card"
              ref="pictureUpload"
              :file-list="fileList"
              :auto-upload="false"
            >
              <i
                slot="default"
                class="el-icon-picture"
                style="
              font-size: 82px;
              color: var(--Dark--);
              transition: all 0.3s ease-in-out;
            "
              />
              <div
                slot="file"
                slot-scope="{ file }"
              >
                <img
                  class="el-upload-list__item-thumbnail"
                  :src="file.url"
                  alt=""
                >
                <span class="el-upload-list__item-actions">
                  <span
                    v-if="!disabled"
                    class="el-upload-list__item-delete"
                    @click="handleRemove(file)"
                  >
                    <i class="el-icon-delete" />
                  </span>
                </span>
              </div>
            </el-upload>
            <el-dialog :visible.sync="dialogVisible">
              <img
                width="100%"
                :src="dialogImageUrl"
                alt=""
              >
            </el-dialog>
            <span style="color: red">*</span><span style="color: var(--Dark--);font-weight: 800;font-size: 0.2vw;">支持 JPG、PNG、JEPG 格式</span>
          </div>
        </div>
        <div class="imitNFTFrameMain_right">
          <div class="selectBox">
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Name
                </div>
                <div class="tipsTitle2">
                  NFT系列名
                </div>
              </div>
              <el-input
                v-model="Name"
                placeholder="Please enter a series name"
              />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Symbol
                </div>
                <div class="tipsTitle2">
                  该系列NFT代币符号
                </div>
              </div>
              <el-input
                v-model="Symbol"
                placeholder="Please enter the token symbol"
              />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Genesis NFT Name
                </div>
                <div class="tipsTitle2">
                  该系列的一号NFT
                </div>
              </div>
              <el-input
                v-model="FirstNFTName"
                placeholder="Please enter the name of Genesis NFT"
              />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Description
                </div>
                <div class="tipsTitle2">
                  描述将包含在商品详情页面上图片下方。
                </div>
              </div>
              <el-input
                type="textarea"
                v-model="Description"
                placeholder="Please add a description"
              />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Supply
                </div>
                <div class="tipsTitle2">
                  可以铸造的物品数量。
                </div>
              </div>
              <div class="input_number">
                <el-input-number
                  v-model="Maxmums"
                  :min="1"
                />
              </div>
            </div>
            <div class="select">
              <div class="sumbitBox">
                <el-button
                  @click="createNFT"
                  :disabled="!canSubmit"
                  class="createButton"
                >
                  创造
                </el-button>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="imitNFTFrameBottom">
        <div class="createTips">
          <span style="font-size: 60px;color: red;">{{ feePercentage }}%</span> commission for market transactions
        </div>
        <img
          src="@/assets/webAssets/UnLoginPhoto.png"
          alt=""
          width="150px"
          height="150px"
        >
      </div>
    </div>
  </div>
</template>

<script>
  import { MakeNFT, getFeePercentage } from "@/api/axios/contract";
  export default {
    data() {
      return {
        imageUrl: "",
        fileList: [],
        dialogImageUrl: "",
        dialogVisible: false,
        disabled: false,
        Name: "",
        Symbol: "",
        FirstNFTName: "",
        Description: "",
        Maxmums: 1,
        active: 1,
        sumitEnable: true,
        feePercentage: 0,
      };
    },
    mounted() {
      getFeePercentage().then(
        (FeePercentage) => (this.feePercentage = FeePercentage)
      );
    },
    computed: {
      canSubmit() {
        if (
          this.Name == "" ||
          this.Symbol == "" ||
          this.$refs.pictureUpload.uploadFiles.length == 0 ||
          this.FirstNFTName == ""
        ) {
          return false;
        } else {
          return true;
        }
      },
    },
    methods: {
      async createNFT() {
        try {
          if (
            this.$refs.pictureUpload.uploadFiles[0].raw.type == "image/jepg" ||
            this.$refs.pictureUpload.uploadFiles[0].raw.type == "image/png" ||
            this.$refs.pictureUpload.uploadFiles[0].raw.type == "image/jpg"
          ) {
            await MakeNFT(
              this.Name,
              this.Symbol,
              this.$refs.pictureUpload.uploadFiles,
              this.Maxmums,
              this.FirstNFTName,
              this.Description
            );
            this.$notify({
              title: "成功",
              message: "创造成功",
              type: "success",
              position: "top-left",
              offset: 100,
              duration: 2000,
            });
            this.fileList = [];
            this.Name = "";
            this.Symbol = "";
            this.Maxmums = 1;
            this.FirstNFTName = "";
            this.Description = "";
          } else {
            this.$notify({
              title: "警告",
              message: "NFT格式不支持",
              type: "warning",
            });
          }
        } catch (error) {
          this.$notify.error({
            title: "错误",
            message: "创造失败",
            position: "top-left",
            offset: 100,
            duration: 2000,
          });
          return;
        }
      },
      handleRemove(file) {
        let uploadFiles = this.$refs.pictureUpload.uploadFiles;
        for (var i = 0; i < uploadFiles.length; i++) {
          if (uploadFiles[i]["url"] == file.url) {
            uploadFiles.splice(i, 1);
          }
        }
      },
    },
  };
</script>

<style lang="scss" scoped>
.MainWindow {
  position: relative;
  overflow: hidden;
  font-family: Arial, Helvetica, sans-serif;
  width: 97%;
  border-radius: 40px;
  background-color: var(--White--);
  transition: all 0.3s ease-in-out;
  margin: 1.3%;
  padding-bottom: 8%;
  .imitNFTFrame {
    width: 100%;
    .imitNFTFrameTop {
      position: relative;
      text-align: left;
      padding-top: 2%;
      padding-left: 4%;

      .imitNFTFrameTop_tipsBox {
        width: 100%;
        font-size: 20px;
        color: rgb(163, 162, 162);
        font-weight: 800;
      }
      a {
        text-decoration: none;
        color: var(--Dark--);
      }
    }
    .imitNFTFrameMain {
      display: flex;
      justify-content: center;
      align-items: center;
      .imitNFTFrameMain_left {
        height: 100%;
        flex: 1;
        /deep/ .imageUpLoad {
          & div:first-child {
            position: relative !important;
            overflow: hidden;
            height: 330px;
          }
          & .el-upload-list--picture-card .el-upload-list__item {
            width: 305px;
            height: 305px;
            position: absolute !important;
          }
          & .el-upload--picture-card {
            padding: 105px;
            height: 305px;
            width: 305px;
            background-color: rgba(255, 255, 255, 0);
            border: 3px dashed var(--Dark--);
            border-radius: 20px;
            transition: all 0.3s ease-in-out;
          }
          & .el-upload--picture-card:hover {
            background-color: rgba(255, 255, 255, 0);
          }
        }
      }
      .imitNFTFrameMain_right {
        width: 100%;
        height: 100%;
        flex: 1;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
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
        }

        .selectBox {
          flex-direction: column;
        }
      }
    }
    .imitNFTFrameBottom {
      display: flex;
      background-color: blue;
      justify-content: center;
      align-items: center;
      width: 100%;
      .createTips {
        width: 100%;
        font-size: 55px;
        color: var(--Dark--);
        transition: all 0.3s ease-in-out;
        z-index: 2;
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
.tipsBox {
  width: 100%;
  font-size: 15px;
  color: rgb(163, 162, 162);
  text-align: left;
  transition: all 0.3s ease-in-out;
  .tipsTitle {
    margin-bottom: 1%;
    width: 76%;
    font-size: 16px;
    font-weight: 800;
    color: var(--Dark--);
    transition: all 0.3s ease-in-out;
  }
  .tipsTitle2 {
    font-size: 12px;
  }
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
