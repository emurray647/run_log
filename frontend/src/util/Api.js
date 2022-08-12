

async function apiCall(url) {
    return fetch(url)
        .then((response) => response.json())
        .then(response => {
            console.log(response);
            return response;
        })
        .then(checkError)
        .then(function(value) {
            console.log("All done")
            console.log(value)
            return value;
        })
}

function checkError(response) {
    if (response.error != null) {
        return Promise.reject("received an error");
    }
    else {
        return Promise.resolve(response.data);
    }
}

export default apiCall