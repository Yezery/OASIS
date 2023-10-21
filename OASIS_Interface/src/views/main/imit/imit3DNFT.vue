<template>
  <div class="MainWindow animate__animated animate__fadeInRight" v-loading.fullscreen.lock="isChanging" element-loading-text="交易进行中" element-loading-spinner="el-icon-loading" element-loading-background="rgba(0, 0, 0, 0.8)">
    <div class="imitNFTFrame ">
      <div class="imitNFTFrameTop">
        <div class="ImitTitle">
          IMIT 3DNFT
        </div>
        <div class="imitNFTFrameTop_tipsBox">
          铸造 3DNFT
        </div>
      </div>
      <div class="imitNFTFrameMain">
        <div class="imitNFTFrameMain_form">
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
                <el-upload class="upload-demo" ref="modelUpload" :on-remove="handleRemove" :file-list="fileList" :auto-upload="false" action="#" :limit="1" :on-change="checkFileType">
                  <el-button class="select3DFile" slot="trigger" @click="handleRemove">
                    选取3D文件
                  </el-button>
                </el-upload>
                <div class="show3DBox" v-loading="loading" v-if="isShow3D" element-loading-text="正在上传到IPFS...">
                  <Model :model-path="modelPath" @initModel="seeModel" />
                </div>
              </div>
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Name
                </div>
                <div class="tipsTitle2">
                  NFT系列名
                </div>
              </div>
              <el-input v-model="Name" placeholder="Please enter a series name" />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  <span style="color: red">*</span> Symbol
                </div>
                <div class="tipsTitle2">
                  该系列NFT代币符号 (要求字符长度不超过11个)
                </div>
              </div>
              <el-input v-model="Symbol" placeholder="Please enter the token symbol" maxlength="11" />
            </div>
            <div class="select">
              <div class="tipsBox">
                <div class="tipsTitle">
                  Genesis NFT Name
                </div>
                <div class="tipsTitle2">
                  该系列的一号NFT
                </div>
              </div>
              <el-input v-model="FirstNFTName" placeholder="Please enter the first 3DNFT name" />
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
                <el-input-number v-model="Maxmums" :min="1" />
              </div>
            </div>
            <div class="select">
              <div class="sumbitBox">
                <el-button @click="createNFT" :disabled="!canSubmit" class="createButton">
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
  import { savetoIPFS } from "@/api/axios/contract";
  import Model from "@/components/3DModelShow/3DModel.vue";
  import { MakeNFT } from "@/api/axios/contract";
  export default {
    components: {
      Model,
    },
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
        Maxmums: 1,
        active: 1,
        sumitEnable: true,
        isRepeatClick: true,

        isProcess: false,
        processLoading: false,
        noProcess: false,

        initModel: undefined,
        model: undefined,
        modelPath: "",
        loading: false,
        isShow3D: false,
        isTypeOf3DModel: false,
        isChanging: false,
      };
    },
    mounted() {},
    computed: {
      canSubmit() {
        if (
          this.Name == "" ||
          this.Symbol == "" ||
          this.$refs.modelUpload.uploadFiles.length == 0 ||
          this.FirstNFTName == "" ||
          !this.isTypeOf3DModel
        ) {
          return false;
        } else {
          return true;
        }
      },
    },

    methods: {
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
      async createNFT() {
        if (this.isRepeatClick) {
          this.isRepeatClick = false;
          try {
            this.isChanging = true;
            await MakeNFT(
              this.Name,
              this.Symbol,
              this.$refs.modelUpload.uploadFiles,
              this.Maxmums,
              this.FirstNFTName,
              "3D"
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
            this.Maxmums = 1;
            this.FirstNFTName = "";
            this.handleRemove();
          } catch (error) {
            this.isChanging = false;
            console.log(error);
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
      .ImitTitle {
        font-size: 3.5vw;
        width: 100%;
        color: var(--Dark--);
        margin-top: 1%;
        font-weight: 800;
        transition: all 0.3s ease-in-out;
      }

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
      .imitNFTFrameMain_form {
        width: 100%;
        height: 100%;
        flex: 1;
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        .selectBox {
          flex-direction: column;
          /deep/ .select {
            width: 100%;
            text-align: left;
            margin-bottom: 30px;
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
</style>
