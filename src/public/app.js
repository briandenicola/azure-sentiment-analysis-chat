new Vue({
    el: '#app',

    data: {
        ws: null, 
        newMsg: '', 
        chatContent: '', 
        username: null,
        joined: false 
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            self.chatContent += '<div class="chip">' + msg.username + '</div> : ' + msg.message + '<br/>';
            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; 
        });
    },

    methods: {
        strip: function(str) {
            return str.replace(/<[^>]+>/g, '');
        },
        send: function () {
            if (this.newMsg) {
                var msg = JSON.stringify({
                    username: this.username,
                    message: this.strip(this.newMsg)
                });
                console.log(msg);
                this.ws.send(msg);
                this.newMsg = ''; 
            }
        },
        join: function () {
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.username = this.strip(this.username);
            this.newMsg = "Welcome " + this.username + " to the chat...";
            this.send();
            this.joined = true;
        }
    }
});