<template>
  <div class="dashboard">
        <Speedometer :speed="currentSpeed"></Speedometer>
  </div>
</template>

<script>
import { setTimeout, setInterval } from "timers";
import { throws } from "assert";
import Speedometer from './Speedometer'
import { Server } from '../config'

export default {
  name: "HelloWorld",
  components: {Speedometer},
  data() {
    return {
      socket: new WebSocket(Server.socket),
      gauge: null,
      currentSpeed: 0,
      currentRPM: 0,
      config: {
        maxSpeed: 180
      },
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
      console.log(response.data);

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

      this.currentSpeed = msg.data.speed;
      this.currentRPM = msg.data.rpm;
    };

    setInterval(() => {
      this.currentSpeed += 2 + (this.currentSpeed * 0.001)

      if (this.currentSpeed <= 100) {
        if (this.currentSpeed < 30) {
          this.currentSpeed += 2 + (this.currentSpeed * 0.001)
        } else {
          if (Math.round(Math.random())) {
            this.currentSpeed -= 2 + (this.currentSpeed * 0.001)
          }
        }
      } else {
        if (this.currentSpeed > 100) {
          this.currentSpeed = 10
        }
      }
    }, 100)
  }
};
</script>


<style scoped>
.dashboard {
    position: absolute;
    width: 100vw;
    height: 100vh;
    background: #13314D;
}
</style>