<script>
    import {onMount} from "svelte";

    let localAudio = undefined;
    let externalAudio = undefined;


    // Setup RTCP connection. uses google servers
    let yourConn = new RTCPeerConnection({
        iceServers: [{ urls: "stun:stun2.1.google.com:19302" }],
    });

    // TODO refactor to its own js files.
    async function setup() {
        const stream = await window.navigator.mediaDevices.getUserMedia({
            audio: true,
        },);

        // if the user has agreed
        if(stream) {
            // Adds audio tracks to the connection
            stream.getAudioTracks().forEach(track => yourConn.addTrack(
                track,
                stream,
            ));

            localAudio = stream;

            // when a remote user adds stream to the peer connection, we let it hear.
            yourConn.onaddstream = function (e) {
                externalAudio = e.stream;
            };

            // Setup ice handling
            yourConn.onicecandidate = function (event) {
                if (event.candidate) {
                    // TODO handle ice connection handling
                }
            };
        }
    }

    async function sendOffer() {
        // creates a offer to connect
        // const offer = await peerConnection.createOffer((offer) => console.log(offer));
        const offer = await yourConn.createOffer();

        yourConn.setLocalDescription(offer);
        // TODO send
    }

    async function onOfferFromUser (offer) {
        yourConn.setRemoteDescription(new RTCSessionDescription(offer));

        // create an answer to an offer
        yourConn.createAnswer(function (answer) {
            yourConn.setLocalDescription(answer);
            // TODO send to api that we accept it.
        }, function (error) {
            alert("Error when creating an answer");
        });
        // TODO decline or accept implementation.

    }

    async function addClient (answer) {
        yourConn.setRemoteDescription(new RTCSessionDescription(answer));
    }
    // TODO subscribe to a store for displaying the call status.
    onMount(async () => {
        await setup();
        await sendOffer();
    })

</script>

<section>
    <p>TODO</p>
    <audio src={externalAudio}></audio>
    <audio src={localAudio}></audio>
</section>