<template>
  <div class="dashboard">
        <Speedometer :speed="currentSpeed"></Speedometer>
  </div>
</template>

<script>
import { setTimeout } from "timers";
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