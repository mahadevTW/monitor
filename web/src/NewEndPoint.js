import * as React from "react";
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import {PostEndPoint} from "./constants";

export default class NewEndPoint extends React.Component {
    constructor(props) {
        super(props);
        this.addEndPoint = this.addEndPoint.bind(this)
    }

    addEndPoint(endPoint) {
        fetch(PostEndPoint, {
                method: 'POST',
                mode: 'cors',
                cache: 'no-cache',
                credentials: 'same-origin',
                headers: {
                    'Content-Type': 'application/json'
                },
                redirect: 'follow',
                referrerPolicy: 'no-referrer',
                body: JSON.stringify(endPoint)
            }
        ).then((data) => {
            console.log(data.json());
            this.props.reload()
        });
    }

    render() {
        return (
            <div className="inputForm">
                <Form>
                    <Form.Group>
                        <Form.Label>Service URL</Form.Label>
                        <Form.Control ref="url" placeholder="Enter url"/>
                    </Form.Group>

                    <Form.Group>
                        <Form.Label>Service name</Form.Label>
                        <Form.Control ref="serviceName" placeholder="Enter service name to be visible here"/>
                    </Form.Group>
                    <Form.Group>
                        <Form.Label>Response Codee</Form.Label>
                        <Form.Control ref="responseCode" type="number" placeholder="http response code"/>
                    </Form.Group>
                    <Button variant="primary" size="sm" onClick={() => {
                        this.addEndPoint({
                            url: this.refs.url.value,
                            serviceName: this.refs.serviceName.value,
                            responseCode: parseInt(this.refs.responseCode.value),
                        })
                    }}>
                        Submit
                    </Button>
                </Form>
            </div>
        )
    }
}