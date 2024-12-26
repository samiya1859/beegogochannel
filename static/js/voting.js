function sendVote(imageID, isUpvote) {
    const voteValue = isUpvote ? 1 : 2;

    const voteRequest = {
        image_id: imageID,
        sub_id: "my-user-1234", 
        value: voteValue
    };
    console.log(voteRequest)
    fetch(`/vote`, {
        method: "POST",
        headers: {
           "Content-Type": "application/json",
        },
        body: JSON.stringify(voteRequest) 
    })
    .then(response => response.text()) 
    .then(data => {
        alert(data); 
    })
    .catch(error => {
        alert("An error occurred while submitting your vote: " + error);
    });
}
