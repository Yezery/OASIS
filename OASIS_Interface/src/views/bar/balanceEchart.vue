<template>
  <div
    id="EchartBox"
    ref="EchartBox"
  >
    <div id="EchartsShow" />
  </div>
</template>

<script>
  import { scheduleDailySummary } from "@/api/axios/Sale";
  import * as echarts from "echarts";
  export default {
    data() {
      return {
        title: "",
        isOpen: false,
        subtitle: "",
        chartOptions: {
          title: {
            text: "",
            subtext: "{b}",
            left: "center",
            top: "10%",
            color: "#A1A2AF",
            fontFamily: "Gilroy-Medium",
            fontSize: 20,
            lineHeight: 24,
            fontWeight: "bold",
            subtextStyle: {
              color: "#ffffff",
              fontSize: 28,
              fontFamily: "Gilroy-Bold",
              fontWeight: "bold",
              lineHeight: 36,
            },
          },
          xAxis: {
            type: "category",
            data: ["Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
            axisLine: {
              show: false,
            },
            axisTick: {
              show: false,
            },
            axisLabel: {
              color: "#ffffff", //坐标值得具体的颜色
              fontSize: window.innerHeight * 0.0123 + "",
              fontFamily: "Gilroy-Medium",
            },
            boundaryGap: false, // 坐标轴两侧留白策略
          },
          yAxis: {
            show: false,
          },
          series: [
            {
              type: "line",
              smooth: true,
              data: [],
              itemStyle: {
                color: "#55C960", //改变折线点的颜色
                lineStyle: {
                  width: 6,
                },
              },
              emphasis: {
                itemStyle: {
                  color: "white",
                },
                label: {
                  show: true, // 显示提示标签
                  padding: [15, 25, 15, 25], // 标签内边距
                  formatter: "{b} {c} ETH", // 标签内容格式
                  fontWeight: "bolder",
                  borderRadius: [10, 10, 10, 10], // 设置
                  borderColor: "white",
                  backgroundColor: "#1e1e2c", // 标签背景色
                  position: "bottom",
                  fontSize: 14,
                  color: "white",
                },
              },
              symbol: "emptyCircle", // 空心圆作为转折点标志
              symbolSize: 12, // 转折点标志大小
              lineStyle: {
                type: "solid", // 折线类型为实线
              },
              areaStyle: {
                // 折线图区域填充样式
                color: {
                  type: "linear",
                  x: 0,
                  y: 0,
                  x2: 0,
                  y2: 1,
                  colorStops: [
                    {
                      offset: 0,
                      color: "rgba(0, 255, 0, 0.3)", // 折线图填充色的起始颜色和透明度
                    },
                    {
                      offset: 1,
                      color: "rgba(0, 255, 0, 0)", // 折线图填充色的结束颜色和透明度
                    },
                  ],
                  globalCoord: false, // 缺省为 false
                },
              },
            },
          ],
        },
      };
    },
    created() {
      this.$emit("echart-change", this.echartsChang);
    },
    mounted() {
      this.init();
    },
    computed: {
      Theme() {
        return this.$store.state.Theme;
      },
    },
    watch: {
      Theme() {
        this.echartsChang();
      },
    },
    methods: {
      isDayLight() {
        var currdate = new Date();
        if (currdate.getHours() >= 18 || currdate.getHours() < 6) {
          return true;
        } else {
          return false;
        }
      },
      echartsChang() {
        echarts.dispose(document.getElementById("EchartsShow"));
        if (this.isOpen) {
          var sum = 0;
          for (
            let index = 0;
            index < this.chartOptions.series[0].data.length;
            index++
          ) {
            sum =
              (sum * 100 + this.chartOptions.series[0].data[index] * 100) / 100;
          }
          this.title = `Balance`;
          this.subtitle = `${sum} ETH`;
        }
        let myChart = echarts.init(document.getElementById("EchartsShow"));
        if (this.$store.state.Theme) {
          this.chartOptions.title.text = this.title;
          this.chartOptions.title.subtext = this.subtitle;
          this.chartOptions.title.subtextStyle.color = "#000000";
          this.chartOptions.xAxis.axisLabel.color = "#000000";
        } else {
          this.chartOptions.title.text = this.title;
          this.chartOptions.title.subtext = this.subtitle;
          this.chartOptions.title.subtextStyle.color = "#ffffff";
          this.chartOptions.xAxis.axisLabel.color = "#ffffff";
        }
        myChart.setOption(this.chartOptions);
      },
      async init() {
        await scheduleDailySummary().then(async (re) => {
          if (re.data.data) {
            for (let index = 0; index < re.data.data.length; index++) {
              console.log(re.data.data[index].TotalTurnover);
              this.chartOptions.series[0].data.push(
                re.data.data[index].TotalTurnover
              );
            }
          } else {
            this.chartOptions.series[0].data.push(0);
          }
          this.echartsChang();
            this.$refs.EchartBox.addEventListener("dblclick", () => {
              this.$refs.EchartBox.style.height = "100px";
              setTimeout(() => {
                this.isOpen = false;
                this.subtitle = "";
                this.title = "";
                this.echartsChang();
              }, 300);
            });
            this.$refs.EchartBox.addEventListener("click", () => {
              this.$refs.EchartBox.style.height = "300px";
              setTimeout(() => {
                this.isOpen = true;
                this.echartsChang();
              }, 300);
            });
        });
      },
    },
  };
</script>

<style lang="scss" scoped>
#EchartBox {
  border: $globalBorder;
  display: flex;
  justify-content: flex-start;
  width: 90%;
  height: 100px;
  background-color: var(--White--);
  border-radius: $globalBorderRadius;

  transition: $globalTransitionUI;
  box-shadow: $globalboxshadowfrom;
  &:hover {
    box-shadow: $globalboxshadowto;
    transition: $globalTransitionUI;
  }
}
#EchartsShow {
  width: 100%;
  height: 100%;
}
</style>