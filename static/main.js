const ws = new WebSocket(`ws://${location.host}/ws`);

ws.onopen = function() {
  console.log("opened");
};

ws.onclose = function() {
  console.log("closed");
};

ws.onmessage = function({ data }) {
  console.log(data);
};
