var ws, connected = false;
var socket_events = {
    "connected": [],
    "disconnected": [],
    "error": []
};

// Crating Websocket service
var tree_connect = function (host) {
    // If host given with Http or https just replacing that, because we have only websocket
    if (host.indexOf("http://") >= 0 || host.indexOf("https://") >= 0)
    {
        host = host.replace("http://", "");
        host = host.replace("https://", "");
    }

    // adding websocket prefix, if it don't have in given host
    if (host.indexOf("ws://") < 0)
    {
        host = "ws://" + host;
    }

    ws = new WebSocket(host);
    ws.onopen = function () {
        connected = true;
        trigger_event("connection", {host: host});
    };
    ws.onclose = function () {
        connected = false;
        trigger_event("disconnect", {host: host});
        setTimeout(function(){
            tree_connect(host);
        }, 1000);
    };
    ws.onmessage = function (message) {
        listen_message(message.data);
    };
    ws.onerror = function(err) {
        console.log(err);
        ws.close();
    }
};


var add_event_callback = function (name, callback) {
    // If we don't have current event we just adding it as an array
    if (!(name in socket_events))
    {
        socket_events[name] = [];
    }
    // Pushing callback to array if we don't have same callback by mistake
    var contains = false;
    for(var i in socket_events[name])
    {
        if(socket_events[name][i] === callback)
        {
            contains = true;
            break;
        }
    }

    if (!contains)
    {
        socket_events[name].push(callback);
    }
};

var remove_event_callback = function (name, callback) {
    if (!(name in socket_events))
    {
        return;
    }

    for(var i in socket_events[name])
    {
        if(socket_events[name][i] === callback)
        {
            socket_events[name] = socket_events[name].splice(i, 1);
            return;
        }
    }
};

var trigger_event = function (event_name, data) {
    if (!(event_name in socket_events))
    {
        return;
    }

    // Calling callbacks
    for(var i in socket_events[event_name])
    {
        socket_events[event_name][i](data);
    }
};

var listen_message = function (message) {
    if(typeof message === 'string')
    {
        try {
            message = JSON.parse(message)
        }
        catch(e) {
            trigger_event("error", e);
            return;
        }
    }

    // If we didn't got object from websocket then just returning from here
    // because we cant compare event with this
    if(typeof message !== 'object' || !("event" in message))
    {
        return;
    }

    trigger_event(message.event, message.data);
};

var send_message = function (event_name, send_data) {
    var message = {
        event: event_name,
        data: send_data
    };
    if(connected)
    {
        try {
            ws.send(JSON.stringify(message));
        }
        catch(e)
        {
            trigger_event("error", e);
        }
        return true;
    }

    return false;
};

module.exports = {
    connect: tree_connect,
    on: add_event_callback,
    off: remove_event_callback,
    emit: send_message
};
