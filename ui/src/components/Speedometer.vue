<template>
    <main :style="{width: `${config.Dashboard.size.width}px`, height: `${config.Dashboard.size.height}px`}">
        <svg id="drawing" xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink" :width="config.Dashboard.size.width" :height="config.Dashboard.size.height">
            <g class="circle circle-1" :transform="`translate(${config.position.x}, ${config.position.y}) rotate(180, 0, 0)`">
                <circle :r="config.radius" class="outline" fill="transparent" stroke="#333" stroke-width="10" :stroke-dasharray="bgCircle.start" :stroke-dashoffset="bgCircle.end"></circle>
                <circle ref="speedCircle" :r="config.radius" fill="transparent" stroke="#6495ED" stroke-width="10" stroke-linecap="round"></circle>
            </g>
        </svg>
    </main>
</template>

<script>
    import SVG from 'svg.js'
    import { Car, Dashboard } from '../config'
import { setTimeout } from 'timers';

    export default {
        props: ['speed'],

        data () {
            return {
                draw: null,
                config: {
                    Dashboard,
                    position: {
                        x: Dashboard.size.width / 2,
                        y: Dashboard.size.height / 2,
                    },
                    radius: 120,
                },
                instances: {
                    draw: null,
                    speedCircle: null,
                    speedText: null,
                    speedUnit: null,
                }
            }
        },

        computed: {
            circunference() {
                return 2 * 3.1416 * 120 // 753.984 ( / 2 = 376.992 )
            },
            bgCircle() {
                return {
                    start: this.circunference,
                    end: this.circunference / 2,
                }
            },
            speedCircle() {
                const minCirc = this.circunference / 2

                const newEnd = (this.circunference - minCirc) * this.speed / Car.maxSpeed

                return {
                    start: this.circunference,
                    end: this.circunference - newEnd,
                }
            }
        },

        watch: {
            speed() {
                if (!this.instances.speedCircle) {
                    return
                }

                this.instances.speedText
                    .text(Math.round(this.speed).toString())
                    .x(this.instances.speedUnit.x() - this.instances.speedText.length() - 12)

                this.instances.speedCircle
                    .animate(200)
                    .attr('stroke-dashoffset', this.speedCircle.end)
            }
        },

        mounted() {
            this.instances.draw = SVG('drawing')

            this.instances.speedCircle = SVG.adopt(this.$refs.speedCircle).attr({
                'stroke-dasharray': this.speedCircle.start,
                'stroke-dashoffset': this.speedCircle.end,
            })

            this.instances.speedUnit = this.instances.draw.text(Dashboard.speedUnit).style({
                fill: '#aaa',
                'font-family': '"Comfortaa", cursive',
                'font-size': '16px'
            })

            this.instances.speedUnit
                .x(this.config.position.x + this.instances.speedUnit.length())
                .y(this.config.position.y + 12)

            this.instances.speedText = this.instances.draw.text('').style({
                fill: '#eee',
                'font-family': '"Comfortaa", cursive',
                'font-size': '42px',
            }).y(this.config.position.y)
        },
    }
</script>

<style scoped>
main {
    position: absolute;
    left: 50%;
    top: 50%;
    margin-left: -400px;
    margin-top: -200px;
    width: 800px;
    height: 400px;
}
</style>