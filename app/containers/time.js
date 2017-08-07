import React, { Component, PropTypes } from 'react'

export default class Timer extends React.Component {
    constructor(props) {
        super(props);
        /* 注意是 时间戳，单位是毫秒 */
        let end_time = this.props.end_time;
        let now_time = new Date().getTime();
        var result;
        if (end_time < now_time){
            result = 0
        }
        else {
            result = end_time - now_time
        }
        this.state = {
            time_: parseInt(result / 1000)
        };
    }

    tick() {
        let now_time = new Date().getTime();
        var result = 0;
        if (this.props.end_time >= now_time){
            result = this.props.end_time - now_time
        }
        this.setState((prevState) => ({
            time_: parseInt(result / 1000)
        }));
    }

    componentDidMount() {
        this.interval = setInterval(() => this.tick(), 1000);
    }

    componentWillUnmount() {
        clearInterval(this.interval);
    }

    time_trans(d){
        if(d.length == 1){
            return '0' + d;
        }
        return d;
    }
    render() {
        let hour=0, minutes=0, seconds=0;
        if (this.state.time_ != 0) {
            hour = this.time_trans(parseInt(this.state.time_ / 3600).toString());
            let next_time = this.state.time_ % 3600;
            minutes = this.time_trans(parseInt(next_time / 60).toString());
            seconds = this.time_trans((next_time % 60).toString());
        }
        else {
            hour = '00';
            minutes = '00'
            seconds = '00'
        }
        return (
            <div className="sectime">
                <div className="seckill-timer">
                    <span>{hour}</span>
                </div>
                <span>:</span>
                <div className="seckill-timer">
                    <span>{minutes}</span>
                </div>
                <span>:</span>
                <div className="seckill-timer">
                    <span>{seconds}</span>
                </div>
            </div>
        );
    }
}
