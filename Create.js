import React, { Component } from 'react';

class Create extends Component {
    render() {
        function handleClick(e) {
            e.preventDefault();
            console.log('The link was clicked.');
            var AWS = require('aws-sdk');
            let dis = this;
            // you shouldn't hardcode your keys in production! See http://docs.aws.amazon.com/AWSJavaScriptSDK/guide/node-configuring.html
            AWS.config.update({ accessKeyId: 'AKIA6HDERDCELTAOD6GK', secretAccessKey: 'LV8ig7hbgENT1awZI5PqLOVtyHiT+JCYMzQaeNB1' });            AWS.config.update({ region: 'us-east-2' });
            var lambda = new AWS.Lambda();
            var params = {
                FunctionName: 'create_dynamo', /* required */
                Payload: "{}"
            };

            lambda.invoke(params, function (err, data) {
                if (err) console.log(err); // an error occurred
                else {
                    console.log("Put succeed");
                }
            });

           
            
        }
        return (
            <button onClick={handleClick}> Create New Dynamo Table
            </button>
        );
    }

}
export default Create;  