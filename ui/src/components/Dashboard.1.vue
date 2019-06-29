<template>
  <v-app id="dashboard">
    <v-container>
      <v-layout row wrap justify-center>
        <v-flex md4>
          <canvas id="speedometer"></canvas>
        </v-flex>
      </v-layout>
    </v-container>
  </v-app>
</template>

<script>
import { setTimeout } from "timers";
import { throws } from "assert";
import { RadialGauge } from "canvas-gauges";
import Speedometer from './Speedometer'

export default {
  name: "HelloWorld",
  components: {Speedometer},
  data() {
    return {
      socket: new WebSocket("ws://127.0.0.1:5000"),
      gauge: null,
      currentSpeed: 0,
      config: {
        maxSpeed: 180
      },
      test: {
        currentSpeed: 0
      }
    };
  },

  computed: {
    speedometer() {
      return (this.currentSpeed * 100) / this.config.maxSpeed;
    }
  },

  methods: {
    sendMessage(type, data) {
      this.socket.send(
        JSON.stringify({
          type,
          data: JSON.stringify(data)
        })
      );
    },

    readMessage(response) {
      const msg = JSON.parse(response.data);

      if (null !== msg.error) {
        throw new Error(msg.error);
      }

      const data = {
        type: msg.type,
        data: JSON.parse(msg.data)
      };

      return data;
    },

    testUpdateSpeed(value) {
      this.test.currentSpeed = this.test.currentSpeed + value;

      this.sendMessage("speed:update", {
        value: Math.floor(this.test.currentSpeed)
      });

      if (value > 0 && this.test.currentSpeed > 140) {
        value = -(this.test.currentSpeed * 0.1);
      } else if (value < 0 && this.test.currentSpeed < 30) {
        value = this.test.currentSpeed * 0.1;
      }

      setTimeout(() => {
        this.testUpdateSpeed(value);
      }, 200);
    }
  },

  mounted() {
    this.gauge = new RadialGauge({
      renderTo: "speedometer",
      width: 500,
      height: 500,
      units: "Km/h",
      minValue: 0,
      startAngle: 100,
      ticksAngle: 160,
      valueBox: false,
      maxValue: 180,
      majorTicks: [
        "0",
        "20",
        "40",
        "60",
        "80",
        "100",
        "120",
        "140",
        "160",
        "180"
      ],
      minorTicks: 2,
      strokeTicks: true,
      highlights: [
        {
          from: 120,
          to: 180,
          color: "rgba(200, 50, 50, .75)"
        }
      ],
      colorPlate: "#fff",
      borderShadowWidth: 0,
      borders: false,
      needleType: "arrow",
      needleWidth: 2,
      needleCircleSize: 7,
      needleCircleOuter: true,
      needleCircleInner: false,
      animationDuration: 1500,
      animationRule: "linear"
    }).draw();

    this.socket.onopen = () => {
      console.log("SERVER: CONNECTED");
    };

    this.socket.onclose = () => {
      console.log("SERVER: CLOSED");
    };

    this.socket.onmessage = rsp => {
      const msg = this.readMessage(rsp);

      this.currentSpeed = msg.data.Value;

      this.gauge.update({
        value: this.currentSpeed,
        units: `${this.currentSpeed} Km/h`,
      })
    };

    setTimeout(() => {
      this.testUpdateSpeed(4);
    }, 2000);
  }
};
</script>