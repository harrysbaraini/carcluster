<template>
  <div class="dashboard">
        <Speedometer :speed="currentSpeed"></Speedometer>
  </div>
</template>

<script>
import { setTimeout } from "timers";
import { throws } from "assert";
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
    this.socket.onopen = () => {
      console.log("SERVER: CONNECTED");
    };

    this.socket.onclose = () => {
      console.log("SERVER: CLOSED");
    };

    this.socket.onmessage = rsp => {
      const msg = this.readMessage(rsp);

      this.currentSpeed = msg.data.Value;
    };

    setTimeout(() => {
      this.testUpdateSpeed(4);
    }, 2000);
  }
};
</script>


<style scoped>
.dashboard {
    position: absolute;
    width: 100vw;
    height: 100vh;
    background: #141212;
}
</style>