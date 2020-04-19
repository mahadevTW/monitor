import React from 'react';
import logo from './logo.svg';
import './App.css';
import NewEndPoint from './NewEndPoint'
import Badge from "react-bootstrap/Badge";
import EndPoints from "./EndPoints";
import {GetEndPointHealth} from "./constants";

export default class App extends React.Component {
    constructor(props) {
        super(props);
        this.reload = this.reload.bind(this);
        this.state = {
            endPoints: []
        }
    }

    componentDidMount() {
        this.reload()
        const that = this
        window.setInterval(function () {
            that.reload()
        }, 10000);
    }

    reload() {
        fetch(GetEndPointHealth).then(response => {
            console.log(response)
            if (response.status !== 200) {
                alert("failed to load statuses")
            }
            response.json().then(d => {
                this.setState({
                    ...this.state,
                    endPoints: d.data
                })
            })
        })
    }

    render() {
        return (
            <div>
                <NewEndPoint reload={this.reload}/>
                <EndPoints endPoints={this.state.endPoints}/>
            </div>
        );
    }
}
