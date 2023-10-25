<template>
  <div
    class="MainWindow animate__animated animate__fadeInRight"
    v-loading.fullscreen.lock="isChanging"
    element-loading-text="交易进行中"
    element-loading-spinner="el-icon-loading"
    element-loading-background="rgba(0, 0, 0, 0.8)"
  >
    <div class="mintNFTFrame ">
      <div class="mintNFTFrameTop">
        <div class="mintTitle">
          Create an NFT
        </div>
        <div class="mintNFTFrameTop_tipsBox">
          铸造NFT
        </div>
      </div>
      <div class="mintNFTFrameMain">
        <div class="mintNFTFrameMain_left">
          <div style="text-align: left;margin-bottom: 2%;">
            <span
              style="color: var(--Dark--);font-weight: 800;font-size: 1vw;"
              v-if="processLoading"
            ><i class="el-icon-loading" />审核中...</span>
            <span
              style="color: rgb(1, 169, 1);font-weight: 800;font-size: 1vw;"
              v-if="isProcess&&!processLoading"
            ><i class="el-icon-success" /> 图片合规</span>
            <span
              style="color: red;font-weight: 800;font-size: 1vw;"
              v-else-if="!isProcess&&!processLoading&&noProcess"
            ><i class="el-icon-error" /> 图片不合规</span>
          </div>
          <div class="imageUpLoad">
            <el-upload
              action="#"
              list-type="picture-card"
              ref="pictureUpload"
              :file-list="fileList"
              :on-change="setPicture"
              :auto-upload="false"
            >
             <div class="innerBox" v-if="!isUpload">
              <i
              
                class="el-icon-upload2"
                style="
              font-size: 3vw;
              color: var(--Dark--);
              transition: all 0.3s ease-in-out;
              margin-bottom:15px;
            "
              />
              <div style="color: var(--Dark--);font-weight: 800;font-size: 1.2vw;">
                将文件拖到此处，或</div>
                <div style="color: #2081E2;font-weight: 800;font-size: 1.2vw;">
                点击浏览文件上传</div>
                <div style="color: #B3B3B3;font-weight: 300;font-size: 1vw;">
                Max size: 50MB</div>
               <div style="color: #B3B3B3;font-weight: 300;font-size: 1vw;">
                JPG,PNG,JEPG,GIF</div>
            </div> 
              <div
                slot="file"
                slot-scope="{ file }"
              >
                <img
                  class="el-upload-list__item-thumbnail"
                  :src="file.url"
                  alt=""
                  ref="realPicture"
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
          </div>
        </div>
        <div class="mintNFTFrameMain_right">
          <div class="selectBox">
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Name *
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
                  Symbol *
                </div>
                <div class="tipsTitle2">
                  该系列NFT代币符号 (要求字符长度不超过11个)
                </div>
              </div>
              <el-input
                v-model="Symbol"
                placeholder="Please enter the token symbol"
                maxlength="11"
              />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Genesis NFT Name *
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
                  描述将包含在商品详情页面上图片下方 <span style="color: red">(不能为关键字 '3D' )</span>
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
                  Supply *
                </div>
                <div class="tipsTitle2">
                  可以铸造的物品数量。
                </div>
              </div>
              <div class="input_number">
                <el-input-number
                  v-model="maximums"
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
    </div>
  </div>
</template>

<script>
  import { process } from "@/utils/nsfwjs";
  import { MakeNFT } from "@/api/axios/contract";
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
        maximums: 1,
        active: 1,
        sumitEnable: true,
        isRepeatClick: true,

        isProcess: false,
        processLoading: false,
        noProcess: false,
        isChanging: false,
        isUpload:false
      };
    },
    mounted() {},
    computed: {
      canSubmit() {
        if (
          this.Name == "" ||
          this.Symbol == "" ||
          this.$refs.pictureUpload.uploadFiles.length == 0 ||
          this.FirstNFTName == "" ||
          !this.isProcess ||
          this.Description == "3D" ||
          this.$refs.pictureUpload.uploadFiles[0].size/1024/1024>50
        ) {
          return false;
        } else {
          return true;
        }
      },
    },
    watch: {
      Description(newData) {
        if (newData == "3D") {
          this.$notify({
            title: "不能为关键字",
            type: "warning",
            position: "top-left",
            offset: 200,
          });
          this.Description = "";
        }
      },
    },
    methods: {
      async setPicture() {
        this.isUpload=true
        this.processLoading = true;
        let prix = this.getFileExtendingName(
          this.$refs.pictureUpload.uploadFiles[0].raw.name
        );
        if (prix == ".jepg" || prix == ".png" || prix == ".jpg"||prix == ".JEPG" || prix == ".PNG" || prix == ".JPG") {
          setTimeout(async () => {
            try {
              await process(this.$refs.realPicture).then((re) => {
                this.processLoading = false;
                if (
                  re[0].className == "Neutral" ||
                  re[0].className == "Drawing"
                ) {
                  this.isProcess = true;
                } else {
                  this.handleRemove();
                  this.isProcess = false;
                  this.noProcess = true;
                  
                  
                }
              });
            } catch (error) {
              this.$notify.error({
                title: "系统异常",
                position: "top-left",
                offset: 200,
              });
              this.isProcess = false;
              this.processLoading = false;
              this.noProcess = false;
              console.error(error);
            }
          }, 100);
        } else {
          if (this.$refs.pictureUpload.uploadFiles[0].size/1024/1024>50) {
            this.$notify.error({
              title:`图片超出大小 ${this.$refs.pictureUpload.uploadFiles[0].size/1024/1024-50}`,
              position: "top-left",
              offset: 200,
            });
            this.isProcess = false;
            this.processLoading = false;
            this.noProcess = false;
            this.handleRemove();
            return
          }
          if (prix == ".gif" || prix == ".GIF") {
            this.isProcess = true;
            this.processLoading = false;
            this.noProcess = false;
          } else {
            this.$notify.error({
              title: "不支持格式",
              position: "top-left",
              offset: 200,
            });
            this.isProcess = false;
            this.processLoading = false;
            this.noProcess = false;
            this.handleRemove();
          }
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
      async createNFT() {
        if (this.isRepeatClick) {
          this.isRepeatClick = false;
          try {
            this.isChanging = true;
            await MakeNFT(
              this.Name,
              this.Symbol,
              this.$refs.pictureUpload.uploadFiles,
              this.maximums,
              this.FirstNFTName,
              this.Description
            );
            this.isChanging = false;
            this.$notify({
              title: "创造成功",
              type: "success",
              position: "top-left",
              offset: 200,
            });
            this.fileList = [];
            this.Name = "";
            this.Symbol = "";
            this.maximums = 1;
            this.FirstNFTName = "";
            this.Description = "";
          } catch (error) {
            this.isChanging = false;
            this.$notify.error({
              title: "创造失败",
              position: "top-left",
              offset: 200,
            });
            return;
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
        // let uploadFiles = this.$refs.pictureUpload.uploadFiles;
        // for (var i = 0; i < uploadFiles.length; i++) {
        //   if (uploadFiles[i]["url"] == file.url) {
        //     uploadFiles.splice(i, 1);
        //   }
        // }
        console.log();
        this.$refs.pictureUpload.clearFiles();
        this.isProcess = false;
        this.noProcess = false;
        this.isUpload=false
      },
    },
  };
</script>

<style lang="scss" scoped>
@import  "@/style/mintPage/mintNFT.scss";

// .NFTImageSelectBox {
//   width: 76%;
//   text-align: left;
//   height: 100%;
//   display: flex;
//   justify-content: center;
//   align-items: center;
//   transition: all 0.3s ease-in-out;
// }
</style>
