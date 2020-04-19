import * as React from "react";
import EndPointStatus from "./EndPointStatus";

export default class EndPoints extends React.Component {
    render() {
        return this.props.endPoints.map(endPoint => {
            return <EndPointStatus endPoint={endPoint}/>
        })
    }
}