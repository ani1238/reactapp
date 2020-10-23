import React, { Component } from 'react';

class Deles extends Component {
    render() {
        function handleClick(e) {
            e.preventDefault();
            console.log('The link was clicked.');
            var AWS = require('aws-sdk');
            let dis = this;
            // you shouldn't hardcode your keys in production! See http://docs.aws.amazon.com/AWSJavaScriptSDK/guide/node-configuring.html
            AWS.config.update({ accessKeyId: 'AKIA6HDERDCENDQIA7EN', secretAccessKey: '3jpDVRyKcBMK+Z1WfIu8evMySk11ILiGyaV/W9+e' });
            AWS.config.update({ region: 'us-east-2' });
            var lambda = new AWS.Lambda();
            var params = {
                FunctionName: 'delete_index_es', /* required */
                Payload: "{}"
            };

            lambda.invoke(params, function (err, data) {
                if (err) console.log(err); // an error occurred
                else {
                    console.log("Index Deleted");
                }
            });

           
            
        }
        return (
            <button onClick={handleClick}> Delete ES Index
            </button>
        );
    }

}
export default Deles;  