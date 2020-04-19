import * as React from "react";
import Badge from "react-bootstrap/Badge";

export default class EndPointStatus extends React.Component {
    constructor(props) {
        super(props);
    }

    getBadge(endPoint) {
        if (endPoint.health.match)
            return <Badge variant="success">Success</Badge>
        return <Badge variant="danger">Fail</Badge>
    }

    render() {
        const endPoint = this.props.endPoint;
        return <div>
            <span>{endPoint.serviceName}</span>
            {this.getBadge(endPoint)}
        </div>

    }
}