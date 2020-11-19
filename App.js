import React, { Component } from 'react';

import {
  AreaChart, Area, XAxis, YAxis, CartesianGrid, Tooltip, Legend,
} from 'recharts';


//console.log(data2)
class App extends Component {
  //static jsfiddleUrl = 'https://jsfiddle.net/alidingling/xqjtetw0/';

  constructor(props) {
    super(props);
    this.state = {
      data: []
    };

  }


  esjson() {

    var AWS = require('aws-sdk');
    let dis = this;
    // you shouldn't hardcode your keys in production! See http://docs.aws.amazon.com/AWSJavaScriptSDK/guide/node-configuring.html
    AWS.config.update({ accessKeyId: 'AKIA6HDERDCELTAOD6GK', secretAccessKey: 'LV8ig7hbgENT1awZI5PqLOVtyHiT+JCYMzQaeNB1' });    AWS.config.update({ region: 'us-east-2' });
    var lambda = new AWS.Lambda();
    var params = {
      FunctionName: 'es-to-json', /* required */
      Payload: "{}"
    };


    var datachart = [];
    lambda.invoke(params, function (err, data) {
      if (err) console.log(err); // an error occurred
      else {
        data = JSON.stringify(data)
        data = JSON.parse(data)
        //var js = data["key1"];//.replace(/\"/g,"\'");
        data = JSON.parse(JSON.parse(data.Payload).key2)
        var hits = data.hits.hits
        for (var i in hits) {
          var el = {}
          el.month = hits[i]._source.Month
          el.cupcakes = parseInt(hits[i]._source.Cupcakes)
          datachart.push(el);
        }
        //console.log(datachart);
        datachart.sort(function (a, b) {
          a = a.month.split("-");
          b = b.month.split("-");
          //console.log(new Date(a[0],a[1],1)-new Date(b[0],b[1],1));
          return new Date(a[0], a[1], 1) - new Date(b[0], b[1], 1);
        });

        console.log(datachart.length);
        //debugger;
        dis.setState({ data: datachart })
      }       // successful response

    });



  }

  componentDidMount() {
    this.esjson();
  }
  render() {
    //debugger;
    //console.log(JSON.stringify(this.state.data))
    return (

      <AreaChart width={900} height={500} data={this.state.data}
        margin={{ top: 10, right: 30, left: 0, bottom: 0 }}>
        <defs>
          <linearGradient id="colorUv" x1="0" y1="0" x2="0" y2="1">
            <stop offset="5%" stopColor="#8884d8" stopOpacity={0.8} />
            <stop offset="95%" stopColor="#8884d8" stopOpacity={0} />
          </linearGradient>
          <linearGradient id="colorPv" x1="0" y1="0" x2="0" y2="1">
            <stop offset="5%" stopColor="#82ca9d" stopOpacity={0.8} />
            <stop offset="95%" stopColor="#82ca9d" stopOpacity={0} />
          </linearGradient>
        </defs>
        <XAxis dataKey="month" />
        <YAxis />
        <Tooltip />
        <Area type="monotone" dataKey="cupcakes" stroke="#8884d8" fillOpacity={1} fill="url(#colorUv)" />
      </AreaChart>

    );
  }
}

export default App;  
